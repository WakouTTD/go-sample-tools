package main

import (
	"fmt"
	"os"
	"time"

	"lockfile/lock"
)

func main() {

	p, _ := os.Getwd()
	fmt.Println(p)

	programName := "samplemedia"

	beginTime := time.Now()
	fmt.Printf("----- begin %s list -----\n", programName)
	fmt.Println("開始 日時 : ", beginTime.Format("2006/01/02 15:04:05"))

	lock.Lock(programName, beginTime)

	time.Sleep(5 * time.Second)

	//	lock.UnLock(programName)

	endTime := time.Now()
	fmt.Println("終了 日時 : ", endTime.Format("2006/01/02 15:04:05"))
	fmt.Println("稼働時間約: ", time.Since(beginTime))
	fmt.Printf("----- end %s list -----\n", programName)
}
