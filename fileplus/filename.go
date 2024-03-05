package fileplus

import (
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
