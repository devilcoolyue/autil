package main

import (
	"flag"
	"github.com/devilcoolyue/autil/aes"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	var (
		aesutilKey        = flag.String("aesutil.key", getEnv("AES_UTIL_KEY", ""), "Specifies the AES key used for encryption and decryption. The key length should be 128, 192, or 256 bits")
		aesutilSalt       = flag.String("aesutil.salt", getEnv("AES_UTIL_SALT", ""), "Specifies the salt used for key derivation to enhance security. The salt helps prevent dictionary attacks.")
		aesutilPlaintext  = flag.String("aesutil.plaintext", getEnv("AES_UTIL_PLAINTEXT", ""), "Specifies the plaintext data to be encrypted.")
		aesutilCiphertext = flag.String("aesutil.ciphertext", getEnv("AES_UTIL_CIPHERTEXT", ""), "Specifies the ciphertext data to be decrypted.")
		aesutilMode       = flag.String("aesutil.mode", getEnv("AES_UTIL_MODE", "static"), "Generate key and salt mode (static or dynamic).")
	)
	flag.Parse()

	// 校验密钥、盐生成模式
	if *aesutilMode != "static" && *aesutilMode != "dynamic" {
		log.Errorf("Unknown mode: %s, please use `static` or `dynamic` mode", *aesutilMode)
		return
	}

	// 校验 明文 和 密文 参数，有且仅有一个有值
	if (*aesutilPlaintext == "" && *aesutilCiphertext == "") || (*aesutilPlaintext != "" && *aesutilCiphertext != "") {
		log.Errorf("Either aesutilPlaintext or aesutilCiphertext must be provided, but not both or neither.")
		return
	}

	// 加密
	if *aesutilPlaintext != "" {
		// 进行加密操作，根据模式生成或使用静态的 key 和 salt
		if *aesutilMode == "static" {
			*aesutilKey = aes.StaticKey
			*aesutilSalt = aes.StaticSalt
		} else if *aesutilMode == "dynamic" {
			*aesutilKey, _ = aes.Generate256BitString()
			*aesutilSalt, _ = aes.Generate256BitString()
		}
		eUtil, err := aes.NewAesUtil(*aesutilPlaintext, *aesutilKey, "", *aesutilSalt).Encrypt()
		if err != nil {
			log.Errorf("Error during encryption: %v", err)
		}
		// 返回结果
		printResult(*aesutilMode, eUtil, true)
	} else if *aesutilCiphertext != "" { // 解密
		// 进行解密操作，如果是动态模式必须提供 key 和 salt
		if (*aesutilKey == "" || *aesutilSalt == "") && *aesutilMode == "dynamic" {
			log.Errorf("Key and salt must be provided for decryption on dynamic mode.")
			return
		}
		if *aesutilMode == "static" {
			*aesutilKey = aes.StaticKey
			*aesutilSalt = aes.StaticSalt
		}
		dUtil, err := aes.NewAesUtil("", *aesutilKey, *aesutilCiphertext, *aesutilSalt).Decrypt()
		if err != nil {
			log.Errorf("Error during decryption: %v", err)
			return
		}
		// 返回结果
		printResult(*aesutilMode, dUtil, false)
	}
}

// 打印返回体
func printResult(mode string, aesUtil *aes.Aes, isEncrypt bool) {
	des := "解"
	if isEncrypt {
		des = "加"
	}
	log.Printf("---------- %s密 开始 ----------", des)
	log.Printf("---------- 模式：%s ----------", mode)
	log.Printf("明文[aesutil.plaintext]：%s", aesUtil.Plaintext)
	// 只有动态模式下才输出 key 和 salt
	if mode == "dynamic" {
		log.Printf("Key[aesutil.key]：%s", aesUtil.Key)
		log.Printf("盐[aesutil.salt]：%s", aesUtil.Salt)
	}
	log.Printf("密文[aesutil.ciphertext]：%s", aesUtil.Ciphertext)
	log.Printf("---------- %s密 结束 ----------", des)
}

// 获取环境变量
func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}
