package gohtml

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Título obtem o título de uma página HTML
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			value := r.FindStringSubmatch(string(html))
			if len(value) > 0 {
				c <- value[1]
			}
			c <- "nada"
		}(url)
	}

	return c
}
