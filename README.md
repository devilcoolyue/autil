# AES加、解密工具
## 命令行参数
| 名称                 | 环境变量名称              | 默认值    | 描述                                                                          |
|--------------------|---------------------|--------|-----------------------------------------------------------------------------|
| aesutil.key        | AES_UTIL_KEY        | ""     | 此值作为密钥，默认为空，一般情况不用传参，使用内置的密钥即可，如果设置了动态模式，则加密用途下会生成一个随机的密钥返回，解密用途下必须传递生成的密钥. |
| aesutil.salt       | AES_UTIL_SALT       | ""     | 此值为盐，默认为空，一般情况不用传参，使用内置的盐即可，如果设置了动态模式，则加密用途下会生成一个随机的盐返回，解密用途下必须传递生成的盐.      |
| aesutil.plaintext  | AES_UTIL_PLAINTEXT  | ""     | 此值为明文，此值在加密用途下必须传递，解密用途下不传参.                                                |
| aesutil.ciphertext | AES_UTIL_CIPHERTEXT | ""     | 此值为密文，此值在解密用途下必须传递，加密用途下不传参.                                                |
| aesutil.mode       | AES_UTIL_MODE       | static | 此值为生成密钥和盐的模式，默认为static，即静态模式，该模式下，使用内置的key和salt值，如果需要动态生成，则设置为dynamic.      |

## 查看帮助
### 执行
```sh
./aesutil -h
```
### 返回结果
```text
Usage of ./aesutil:
-aesutil.ciphertext string
Specifies the ciphertext data to be decrypted.
-aesutil.key string
Specifies the AES key used for encryption and decryption. The key length should be 128, 192, or 256 bits
-aesutil.mode string
Generate key and salt mode (static or dynamic). (default "static")
-aesutil.plaintext string
Specifies the plaintext data to be encrypted.
-aesutil.salt string
Specifies the salt used for key derivation to enhance security. The salt helps prevent dictionary attacks.
```
## 静态模式
### 加密
#### 执行
```sh
./aesutil -aesutil.plaintext="hello world"
```
#### 返回结果
```text
INFO[0000] ---------- 加密 开始 ----------                  
INFO[0000] ---------- 模式：static ----------              
INFO[0000] 明文[aesutil.plaintext]：hello world            
INFO[0000] 密文[aesutil.ciphertext]：AeboiwkOF9uKNDxZq3C/F1k7sEm6Vi2J60LbAcbzENKwNjoysxmN6IiLHbj+3KmOZZzB+4HLlF0nRFKnwIHlbwxCGLIkbKtVO6sLFAlmz9M= 
INFO[0000] ---------- 加密 结束 ---------- 
```
### 解密
#### 执行
```sh
./aesutil -aesutil.ciphertext="AeboiwkOF9uKNDxZq3C/F1k7sEm6Vi2J60LbAcbzENKwNjoysxmN6IiLHbj+3KmOZZzB+4HLlF0nRFKnwIHlbwxCGLIkbKtVO6sLFAlmz9M="
```
#### 返回结果
```text
INFO[0000] ---------- 解密 开始 ----------                  
INFO[0000] ---------- 模式：static ----------              
INFO[0000] 明文[aesutil.plaintext]：hello world            
INFO[0000] 密文[aesutil.ciphertext]：AeboiwkOF9uKNDxZq3C/F1k7sEm6Vi2J60LbAcbzENKwNjoysxmN6IiLHbj+3KmOZZzB+4HLlF0nRFKnwIHlbwxCGLIkbKtVO6sLFAlmz9M= 
INFO[0000] ---------- 解密 结束 ---------- 
```
## 动态模式
### 加密
#### 执行
```sh
./aesutil -aesutil.plaintext="hello world" -aesutil.mode=dynamic
```
#### 返回结果
```text
INFO[0000] ---------- 加密 开始 ----------                  
INFO[0000] ---------- 模式：dynamic ----------             
INFO[0000] 明文[aesutil.plaintext]：hello world            
INFO[0000] Key[aesutil.key]：690f27ebaed6028a161feb6938c707d0e30246503f9559b0b82e0d07941fb460 
INFO[0000] 盐[aesutil.salt]：5e84d3e50cb5e4b07d1378d9a415061f54e0ba1a79a31219e9606cbd29e04aa2 
INFO[0000] 密文[aesutil.ciphertext]：hYdx6yNrVllN0S2w7WT1Ig29BjW5GKaRYUbR6o4mruMD/3scKzZOGlhFK+S19RmF1qYovXCnOVjJYWVupqng8qq0o2OCzE3n72Q0NBNEvcs= 
INFO[0000] ---------- 加密 结束 ---------- 
```
### 解密
#### 执行
```sh
./aesutil \
-aesutil.ciphertext="hYdx6yNrVllN0S2w7WT1Ig29BjW5GKaRYUbR6o4mruMD/3scKzZOGlhFK+S19RmF1qYovXCnOVjJYWVupqng8qq0o2OCzE3n72Q0NBNEvcs=" \
-aesutil.mode=dynamic \
-aesutil.key="690f27ebaed6028a161feb6938c707d0e30246503f9559b0b82e0d07941fb460" \
-aesutil.salt="5e84d3e50cb5e4b07d1378d9a415061f54e0ba1a79a31219e9606cbd29e04aa2"
```
#### 返回结果
```text
INFO[0000] ---------- 解密 开始 ----------                  
INFO[0000] ---------- 模式：dynamic ----------             
INFO[0000] 明文[aesutil.plaintext]：hello world            
INFO[0000] Key[aesutil.key]：690f27ebaed6028a161feb6938c707d0e30246503f9559b0b82e0d07941fb460 
INFO[0000] 盐[aesutil.salt]：5e84d3e50cb5e4b07d1378d9a415061f54e0ba1a79a31219e9606cbd29e04aa2 
INFO[0000] 密文[aesutil.ciphertext]：hYdx6yNrVllN0S2w7WT1Ig29BjW5GKaRYUbR6o4mruMD/3scKzZOGlhFK+S19RmF1qYovXCnOVjJYWVupqng8qq0o2OCzE3n72Q0NBNEvcs= 
INFO[0000] ---------- 解密 结束 ---------- 
```
