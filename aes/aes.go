package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

const (
	StaticKey  = "91f122e9bcd674a4a6b28fc6322a9f4ce0b28635532eba406828bf5610db1a3b" // 内置的静态key
	StaticSalt = "1c35a61272f4a0a67af35ea765b7cbb6dad10194618b87f605af79661ca57d19" // 内置的静态salt
)

type Aes struct {
	Plaintext  string // 明文
	Key        string // 密钥
	Ciphertext string // 密文
	Salt       string // 盐
}

func NewAesUtil(plaintext, key, ciphertext, salt string) *Aes {
	return &Aes{
		Plaintext:  plaintext,
		Key:        key,
		Ciphertext: ciphertext,
		Salt:       salt,
	}
}

// Encrypt 加密
func (aesUtil *Aes) Encrypt() (*Aes, error) {
	if len(aesUtil.Plaintext) == 0 || len(aesUtil.Key) == 0 || len(aesUtil.Salt) == 0 {
		return nil, errors.New("plaintext or key or salt is empty")
	}

	// 在明文末尾添加盐值
	plaintextWithSalt := aesUtil.Plaintext + aesUtil.Salt

	// 将明文转换为字节数组
	plaintextBytes := []byte(plaintextWithSalt)

	// 将十六进制密钥转换为字节数组
	keyByte, err := hex.DecodeString(aesUtil.Key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to decode hex key: %v", err))
	}

	// 使用 AES 算法创建 Cipher
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	// 对明文进行 PKCS#7 填充
	blockSize := block.BlockSize()
	plaintextBytes = pkcs7Pad(plaintextBytes, blockSize)

	// 创建 ECB 加密器
	encrypted := make([]byte, len(plaintextBytes))
	for i := 0; i < len(plaintextBytes); i += block.BlockSize() {
		block.Encrypt(encrypted[i:], plaintextBytes[i:])
	}

	// 将加密后的字节数组进行 Base64 编码
	ciphertext := base64.StdEncoding.EncodeToString(encrypted)

	aesUtil.Ciphertext = ciphertext

	// 返回加密后的字符串
	return aesUtil, nil
}

// Decrypt 解密
func (aesUtil *Aes) Decrypt() (*Aes, error) {
	if len(aesUtil.Ciphertext) == 0 || len(aesUtil.Key) == 0 || len(aesUtil.Salt) == 0 {
		return nil, errors.New("encrypted or key or salt is empty")
	}
	// 解析 Base64 编码的密文
	ciphertext, err := base64.StdEncoding.DecodeString(aesUtil.Ciphertext)
	if err != nil {
		return nil, err
	}

	// 将十六进制密钥转换为字节数组
	keyByte, err := hex.DecodeString(aesUtil.Key)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to decode hex key: %v", err))
	}

	// 使用 AES 算法创建 Cipher
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	// 创建 ECB 解密器
	decrypted := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(decrypted[i:], ciphertext[i:])
	}

	// 去除 PKCS#7 填充
	plaintextWithSalt, err := pkcs7Unpad(decrypted)
	if err != nil {
		return nil, err
	}
	plaintext := string(plaintextWithSalt)

	// 去除加盐值
	if strings.HasSuffix(plaintext, aesUtil.Salt) {
		plaintext = plaintext[:len(plaintext)-len(aesUtil.Salt)]
	}

	aesUtil.Plaintext = plaintext

	// 返回解密后的字符串
	return aesUtil, nil
}

// PKCS#7 填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// 去除 PKCS#7 填充
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("pkcs7Unpad: input data is empty")
	}
	unpadding := int(data[length-1])

	// 校验填充长度是否合理
	if unpadding <= 0 || unpadding > length {
		return nil, fmt.Errorf("pkcs7Unpad: invalid padding size")
	}

	return data[:(length - unpadding)], nil
}

// Generate256BitString 随机生成 256 位密钥或盐的字符串
func Generate256BitString() (string, error) {
	key := make([]byte, 32) // 32 bytes for 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
