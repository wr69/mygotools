package notice

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Post(sendUrl string, sendKey string, channel string, mode string, name string, postData string) {
	//base64.StdEncoding.EncodeToString
	//base64.URLEncoding.EncodeToString
	encoder := base64.URLEncoding
	formData := url.Values{}
	formData.Set("password", sendKey)
	formData.Set("channel", channel) //
	formData.Set("mode", mode)
	formData.Set("name", name)
	formData.Set("json", encoder.EncodeToString([]byte(postData)))

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	req, err := http.NewRequest("POST", sendUrl, strings.NewReader(formData.Encode())) //http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(name, "Failed to post HTTP URL :", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "github action")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(name, "Failed to post resp :", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(name, "Failed to post body :", err)
		return
	}

	log.Println(name, string(body))
}
