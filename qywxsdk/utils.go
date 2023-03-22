package qywxsdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math/rand"
	"sort"
	"strings"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//验签
func MakeSignature(token, timestamp, nonce, msg string) string {
	strs := sort.StringSlice{token, timestamp, nonce, msg}
	sort.Strings(strs)
	str := strings.Join(strs, "")
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

//随机获取一个数值内的string
func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

//pKCS7Padding
func pKCS7Padding(plaintext string, block_size int) []byte {
	padding := block_size - (len(plaintext) % block_size)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	var buffer bytes.Buffer
	buffer.WriteString(plaintext)
	buffer.Write(padtext)
	return buffer.Bytes()
}

//cbc加密
func cbcEncrypter(encoding_aeskey, plaintext string) ([]byte, error) {
	if !strings.HasSuffix(encoding_aeskey, "=") {
		encoding_aeskey = encoding_aeskey + "="
	}
	aeskey, err := base64.StdEncoding.DecodeString(encoding_aeskey)
	if nil != err {
		return nil, err
	}
	const block_size = 32
	pad_msg := pKCS7Padding(plaintext, block_size)

	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(pad_msg))
	iv := aeskey[:aes.BlockSize]

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext, pad_msg)
	base64_msg := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(base64_msg, ciphertext)

	return base64_msg, nil
}

func pKCS7Unpadding(plaintext []byte, block_size int) ([]byte, error) {
	plaintext_len := len(plaintext)
	if nil == plaintext || plaintext_len == 0 {
		return nil, errors.New("pKCS7Unpadding error nil or zero")
	}
	if plaintext_len%block_size != 0 {
		return nil, errors.New("pKCS7Unpadding text not a multiple of the block size")
	}
	padding_len := int(plaintext[plaintext_len-1])
	return plaintext[:plaintext_len-padding_len], nil
}
func Encrypt(receiver_id, encoding_aeskey, reply_msg string) (string, error) {
	rand_str := randString(16)
	var buffer bytes.Buffer
	buffer.WriteString(rand_str)

	msg_len_buf := make([]byte, 4)
	binary.BigEndian.PutUint32(msg_len_buf, uint32(len(reply_msg)))
	buffer.Write(msg_len_buf)
	buffer.WriteString(reply_msg)
	buffer.WriteString(receiver_id)

	tmp_ciphertext, err := cbcEncrypter(encoding_aeskey, buffer.String())
	if nil != err {
		return "", err
	}
	ciphertext := string(tmp_ciphertext)
	return ciphertext, nil
}

func ParsePlainText(plaintext []byte) ([]byte, uint32, []byte, []byte, error) {
	const block_size = 32
	plaintext, err := pKCS7Unpadding(plaintext, block_size)
	if nil != err {
		return nil, 0, nil, nil, err
	}

	text_len := uint32(len(plaintext))
	if text_len < 20 {
		return nil, 0, nil, nil, errors.New("plain is to small 1")
	}
	random := plaintext[:16]
	msg_len := binary.BigEndian.Uint32(plaintext[16:20])
	if text_len < (20 + msg_len) {
		return nil, 0, nil, nil, errors.New("plain is to small 2")
	}

	msg := plaintext[20 : 20+msg_len]
	receiver_id := plaintext[20+msg_len:]

	return random, msg_len, msg, receiver_id, nil
}

//cbc解密
func cbcDecrypter(encoding_aeskey, base64_encrypt_msg string) ([]byte, error) {
	if !strings.HasSuffix(encoding_aeskey, "=") {
		encoding_aeskey = encoding_aeskey + "="
	}
	aeskey, err := base64.StdEncoding.DecodeString(encoding_aeskey)
	if nil != err {
		return nil, err
	}

	encrypt_msg, err := base64.StdEncoding.DecodeString(base64_encrypt_msg)
	if nil != err {
		return nil, err
	}

	block, err := aes.NewCipher(aeskey)
	if err != nil {
		return nil, err
	}

	if len(encrypt_msg) < aes.BlockSize {
		return nil, errors.New("encrypt_msg size is not valid")
	}

	iv := aeskey[:aes.BlockSize]

	if len(encrypt_msg)%aes.BlockSize != 0 {
		return nil, errors.New("encrypt_msg not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encrypt_msg, encrypt_msg)

	return encrypt_msg, nil
}
func Decrypt(receiver_id, encoding_aeskey, encrypt string) ([]byte, error) {
	plaintext, crypt_err := cbcDecrypter(encoding_aeskey, encrypt)
	if nil != crypt_err {
		return nil, crypt_err
	}
	_, _, msg, plan_receiver_id, crypt_err := ParsePlainText(plaintext)
	if nil != crypt_err {
		return nil, crypt_err
	}

	if len(receiver_id) > 0 && strings.Compare(string(plan_receiver_id), receiver_id) != 0 {
		return nil, errors.New("receiver_id is not equil")
	}

	return msg, nil
}
