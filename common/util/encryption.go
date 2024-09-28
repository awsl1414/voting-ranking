package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptonMd5 使用 MD5 对字符串进行加密，返回加密后的十六进制字符串
func EncryptionMd5(s string) string {
	// 创建一个新的 MD5 哈希对象
	hash := md5.New()
	// 将字符串转换为字节数组并写入哈希对象
	hash.Write([]byte(s))
	// 计算哈希值，并将结果转换为十六进制字符串返回
	return hex.EncodeToString(hash.Sum(nil))
}
