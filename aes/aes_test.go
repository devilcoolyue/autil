package aes

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	// 定义密钥、加盐值和加密的密文
	key := "91f122e9bcd674a4a6b28fc6322a9f4ce0b28635532eba406828bf5610db1a3b"
	salt := "1c35a61272f4a0a67af35ea765b7cbb6dad10194618b87f605af79661ca57d19"
	plaintext := "你好，这是一条秘密消息！"

	aesUtil, err := NewAesUtil(plaintext, key, "", salt).Encrypt()
	if err != nil {
		panic(err)
	}
	fmt.Println(aesUtil.Ciphertext)
}

func TestAesDecrypt(t *testing.T) {
	// 定义密钥、加盐值和加密的密文
	key := "91f122e9bcd674a4a6b28fc6322a9f4ce0b28635532eba406828bf5610db1a3b"
	salt := "1c35a61272f4a0a67af35ea765b7cbb6dad10194618b87f605af79661ca57d19"
	ciphertext := "zVzTCnTbzIrQF9yTJOYJ7in97PcFaIf7bVk+WNI1NSJkFE7GQWAaU0yJ1F6W8fzVhkzl4zPL9pqCZXiCXHHLe3NMEl3xtYqSVSb3c9wneaNLReCc1Ufz5sdT8AEl3nctk8IoExXzBZKGyFDOIfdEOg=="

	aesUtil, err := NewAesUtil("", key, ciphertext, salt).Decrypt()
	if err != nil {
		panic(err)
	}
	fmt.Println(aesUtil.Plaintext)
}

func TestStaticAllFlow(t *testing.T) {
	eUtil, err := NewAesUtil("你好，这是一条秘密消息！", StaticKey, "", StaticSalt).Encrypt()
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------- 加密 开始 -----------------")
	fmt.Printf("明文：%s\n", eUtil.Plaintext)
	fmt.Printf("Key：%s\n", eUtil.Key)
	fmt.Printf("盐：%s\n", eUtil.Salt)
	fmt.Printf("密文：%s\n", eUtil.Ciphertext)
	fmt.Println("--------------- 加密 结束 -----------------")

	dUtil, err := NewAesUtil("", StaticKey, eUtil.Ciphertext, StaticSalt).Decrypt()

	if err != nil {
		panic(err)
	}
	fmt.Println("--------------- 解密 开始 -----------------")
	fmt.Printf("明文：%s\n", dUtil.Plaintext)
	fmt.Printf("Key：%s\n", dUtil.Key)
	fmt.Printf("盐：%s\n", dUtil.Salt)
	fmt.Printf("密文：%s\n", dUtil.Ciphertext)
	fmt.Println("--------------- 解密 结束 -----------------")
}

func TestDynamicAllFlow(t *testing.T) {
	key, _ := Generate256BitString()
	salt, _ := Generate256BitString()
	eUtil, err := NewAesUtil("你好，这是一条秘密消息！", key, "", salt).Encrypt()
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------- 加密 开始 -----------------")
	fmt.Printf("明文：%s\n", eUtil.Plaintext)
	fmt.Printf("Key：%s\n", eUtil.Key)
	fmt.Printf("盐：%s\n", eUtil.Salt)
	fmt.Printf("密文：%s\n", eUtil.Ciphertext)
	fmt.Println("--------------- 加密 结束 -----------------")

	dUtil, err := NewAesUtil("", key, eUtil.Ciphertext, salt).Decrypt()

	if err != nil {
		panic(err)
	}
	fmt.Println("--------------- 解密 开始 -----------------")
	fmt.Printf("明文：%s\n", dUtil.Plaintext)
	fmt.Printf("Key：%s\n", dUtil.Key)
	fmt.Printf("盐：%s\n", dUtil.Salt)
	fmt.Printf("密文：%s\n", dUtil.Ciphertext)
	fmt.Println("--------------- 解密 结束 -----------------")
}
