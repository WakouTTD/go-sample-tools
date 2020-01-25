package lock

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

var lockFileDirectory = "./"

func addLockFile(fileName string, beginTime time.Time) error {
	content := []byte(beginTime.Format("2006/01/02 15:04:05") + "\n")
	err := ioutil.WriteFile(fileName, content, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Lock(mediaName string, beginTime time.Time) {
	fileName := lockFileDirectory + mediaName
	if exists(fileName) {
		fmt.Printf("異常終了: The file \"%s\" already exists . Check the directory \"%s\".\n", fileName, lockFileDirectory)
		os.Exit(1)
	} else {
		err := addLockFile(fileName, beginTime)
		if err != nil {
			fmt.Printf("異常終了: The file \"%s\". Check the directory \"%s\". %s\n", fileName, lockFileDirectory, err)
			os.Exit(1)
		}
	}
}

func removeLockFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func UnLock(mediaName string) {
	fileName := lockFileDirectory + mediaName
	if exists(fileName) {
		err := removeLockFile(fileName)
		if err != nil {
			fmt.Printf("異常終了: The file \"%s\". Check the directory \"%s\". %s\n", fileName, lockFileDirectory, err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("異常終了: The file \"%s\" already does not exist . Check the directory \"%s\".\n", fileName, lockFileDirectory)
		os.Exit(2)
	}
}
