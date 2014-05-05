package main

import(
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	proxyUrl, err := url.Parse("http://127.0.0.1:8081/")
	if err != nil {
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:     true,
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	req, err := http.NewRequest("GET", "http://115.28.240.199/test2.php", nil)
	if err != nil || req == nil || req.Header == nil {
		return
	}
	req.Header.Set("hello", "world")
	resp, err := client.Do(req)

	fmt.Println(resp.Header)
	fmt.Println("Content-Length", resp.ContentLength)
	if err != nil || resp == nil || resp.Body == nil {
		return
	} else {
		html, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(html))
	}
}