package hello

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateToken() (int64, string) {
	callTime := time.Now().UnixNano() / 1000000
	block, _ := des.NewTripleDESCipher([]byte("XVlB>gba*CMRAj&whTHc}cuA"))
	blockSize := block.BlockSize()
	message := fmt.Sprintf("%d#%s", callTime, GenerateRandomStr(16))
	origData, _ := PKCS7Padding([]byte(message), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte("X`l$zg<a"))
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return callTime, base64.StdEncoding.EncodeToString(cryted)
}
