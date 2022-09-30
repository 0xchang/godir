package common

import (
	"flag"
)

//显示logo信息
func Banner() {
	banner := `
  ____       ____  _
 / ___| ___ |  _ \(_)_ __
| |  _ / _ \| | | | | '__|
| |_| | (_) | |_| | | |
 \____|\___/|____/|_|_|
           godir version: ` + Version + `
           by 0xchang` + "\n"
	Colban.Println(banner)
}

//读取命令行参数
func Flag() {
	Banner()
	flag.StringVar(&Url, "u", "", "Target URL")
	flag.StringVar(&Extention, "e", "php, aspx, jsp, html, js", "Extension list separated by commas (e.g. php,asp)")
	flag.StringVar(&OutFile, "o", "", "Output File(default  url.result.txt)")
	flag.StringVar(&UrlFile, "uf", "", "Target URL File")
	flag.StringVar(&UA, "ua", "", "Set User-Agent(default  random)")
	flag.StringVar(&WordList, "wd", "dict/dict.txt", "Brute Dict(default  dict/dict.txt)")
	flag.StringVar(&ReqMethod, "m", "GET", "Request Method(default GET)")
	flag.IntVar(&ThreadNum, "t", 10, "Number of threads(default 10)")
	flag.IntVar(&Timeout, "to", 3, "Number of timeout(default 3)")
	flag.Parse()

}
