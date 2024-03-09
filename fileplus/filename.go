package fileplus

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"strings"
)

func UrlToFilename(urlStr string) string {
	filename := strings.Replace(urlStr, "https://", "", 1)
	filename = strings.Replace(filename, "http://", "", 1)

	//log.Println("c:", filename)
	// Remove query parameters
	filename = strings.SplitN(filename, "?", 2)[0]
	//log.Println("c:", filename)
	// Remove special characters
	invalidChars := []string{"/", `\`, ":", "*", "?", "\"", "<", ">", "|"}
	for _, c := range invalidChars {
		filename = strings.ReplaceAll(filename, c, "")
	}

	return filename
}

func UrlEscapeFilename(urlStr string) string {
	filename := url.PathEscape(urlStr)
	filename = strings.ReplaceAll(filename, ":", "")

	return filename
}

func UrlEscapeSha256(urlStr string) string {

	// 计算 URL 的 SHA-256 哈希值
	h := sha256.New()
	h.Write([]byte(urlStr))
	hashInBytes := h.Sum(nil)

	// 将哈希值转换为16进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}

func FileToSha512(content string) string {

	// 计算 URL 的 SHA-512 哈希值
	h := sha512.New()
	h.Write([]byte(content))
	hashInBytes := h.Sum(nil)

	// 将哈希值转换为16进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}
