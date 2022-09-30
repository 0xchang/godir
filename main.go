package main

import (
	"godir/common"
	"godir/utils"

	"github.com/fatih/color"
)

func main() {
	common.Flag()
	common.Parse()

	bar := utils.NewBarOption(0, common.FileLine) //任务条
	lim := utils.NewLimit(common.ThreadNum)       //线程数

	go utils.ReadFromFile()
	for i := 0; i < int(bar.GetTotal()); i++ {
		lim.Add()
		go utils.ScanTask(lim, bar)
	}
	lim.Wait()
	bar.Finish()
	color.New(color.FgYellow).Printf("\nTask Completed")
}
