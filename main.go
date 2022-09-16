package main

import (
	"io/ioutil"
	"os"
)

func main() {
	path, _ := os.LookupEnv("LOCALAPPDATA")
	path = path + "\\Packages\\MicrosoftWindows.Client.CBS_cw5n1h2txyewy\\TempState\\ScreenClip\\"
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var recentTime int64 = 0
	fileName := ""

	for _, entry := range dirs {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		time_ := info.ModTime().Unix()
		if recentTime < time_ {
			recentTime = time_
			fileName = info.Name()
		}
	}
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	CopyFile(path+fileName, currentPath+"\\"+fileName)
	return
}

func CopyFile(filePath string, targetPath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(targetPath, bytes, 0777)
	if err != nil {
		panic(err)
	}
	return
}