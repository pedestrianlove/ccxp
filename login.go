package ccxp

import (
	"io"
	"net/http"
	"strings"
)

func (c *client) Login() error {
	return nil
}

func (c *client) getpwdstr(imageUrl string) (string, error) {
	result := strings.Split(imageUrl, "pwdstr=")

	return result[1], nil
}

func (c *client) requestOCR(imageUrl string) (string, error) {
	response, err := http.Get("https://ocr.nthumods.com/?url=https://www.ccxp.nthu.edu.tw/ccxp/INQUIRE/" + imageUrl)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
