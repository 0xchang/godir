package utils

import (
	"fmt"
	"godir/common"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var format_str = "\r[%02d:%02d:%02d] %d -  %s  -  %s               \n"

func GenRequest(url string) *http.Request {

	req, err := http.NewRequest(common.ReqMethod, url, nil)
	if err != nil {
		fmt.Println("\r url is ", url)
		panic("here")
	}
	req.Header.Add("User-Agent", Random_UA())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	return req
}

func DoRequest(url string, uri string) {
	client := http.Client{
		Timeout: time.Duration(common.Timeout) * time.Second,
	}
	req := GenRequest(url + uri)
	res, err := client.Do(req)
	if err != nil {
		common.Glock.Lock()
		common.Colerr.Printf("\r%-50s %s %s\n", req.URL.RequestURI(), "网络错误,并发过高或网络不佳", strings.Repeat(" ", 10))
		common.Glock.Unlock()
		time.Sleep(time.Millisecond * 200)
		return
	}
	defer res.Body.Close()
	now := time.Now()
	status_code := res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		common.Colerr.Printf("Read body error!")
		return
	}
	bodylen := len(body)
	result_fmt := fmt.Sprintf(format_str, now.Hour(), now.Minute(), now.Second(), status_code, Lenout(bodylen), req.URL.RequestURI())
	common.Glock.Lock()
	if status_code < 300 {
		common.Col200.Printf(result_fmt)
	} else if status_code < 400 {
		common.Col30x.Printf(result_fmt)
	} else if status_code == 403 {
		common.Col40x.Printf(result_fmt)
	} else if status_code < 500 {
		//common.Col40x.Printf(result_fmt)
	} else {
		common.Col50x.Printf(result_fmt)
	}
	common.Glock.Unlock()
}

func Lenout(size int) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%dKB", size/1024)
	} else {
		return fmt.Sprintf("%dMB", size/(1024*1024))
	}
}
