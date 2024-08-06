package ccxp

import (
	"errors"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Client interface {
	Login() error
	GetCaptcha() string
}

var captcha string

type client struct {
	crawler  *colly.Collector
	username string
	password string
}

func New(username, password string) Client {
	return &client{
		crawler:  colly.NewCollector(colly.AllowedDomains("www.ccxp.nthu.edu.tw")),
		username: username,
		password: password,
	}
}

func (c *client) Login() error {
	loginClient := c.crawler.Clone()

	loginClient.OnHTML("form[action='pre_select_entry.php']", func(e *colly.HTMLElement) {
		go func() {
			e.DOM.ChildrenFiltered("img").Each(func(_ int, img *goquery.Selection) {
				imageUrl, exists := img.Attr("src")
				if exists {
					response, err := http.Get("https://ocr.nthumods.com/?url=https://www.ccxp.nthu.edu.tw/ccxp/INQUIRE/" + imageUrl)
					if err != nil {
						errors.New(err.Error())
					}

					result, err := io.ReadAll(response.Body)
					if err != nil {
						errors.New(err.Error())
					}

					captcha = string(result)
				} else {
					errors.New("Captcha url does not exist.")
					return
				}
			})

		}()
	})

	loginClient.Visit("https://www.ccxp.nthu.edu.tw/ccxp/INQUIRE/")

	// loginClient.OnScraped(func(response *colly.Response) {
	// 	err := loginClient.Post("https://www.ccxp.nthu.edu.tw/ccxp/INQUIRE/pre_select_entry.php", map[string]string{
	// 		"account": c.username,
	// 		"passwd":  c.password,
	// 	})
	//
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// })

	return nil
}

func (c *client) GetCaptcha() string {
	return captcha
}
