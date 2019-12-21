// 加解密相关方法
package crypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// SHA1 SHA1哈希加密
func SHA1(plainText []byte) string {
	sha := sha1.New()
	sha.Write(plainText)
	return hex.EncodeToString(sha.Sum(nil))
}

// MD5 MD5哈希加密， 返回32位字符串
func MD5(plainText []byte) string {
	m := md5.New()
	m.Write(plainText)
	return hex.EncodeToString(m.Sum(nil))
}

// SHA256 加密
func SHA256(plainText []byte) string {
	sha := sha256.New()
	sha.Write([]byte(plainText))
	return hex.EncodeToString(sha.Sum(nil))
}

