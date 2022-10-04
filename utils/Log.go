package utils

import (
	"fmt"
	"godir/common"
	"os"
)

func strToFile(res string) {
	var text = []byte(res)
	common.Glock.Lock()
	fl, err := os.OpenFile(common.OutFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", common.OutFile, err)
		return
	}
	_, err = fl.Write(text)
	fl.Close()
	if err != nil {
		fmt.Printf("Write %s error, %v\n", common.OutFile, err)
	}
	common.Glock.Unlock()
}
