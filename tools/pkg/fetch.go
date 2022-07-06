package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func GetDownload(path string, headers map[string]string, downloadDir string) {
	// 需要自定义 headers
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	req.Header.Set("content-type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	// 获取文件名
	fileInfo := strings.Split(resp.Header.Get("Content-Disposition"), ";")
	var fileName string
	if len(fileInfo) == 2 {
		fileInfo = strings.Split(fileInfo[1], "=")
		if len(fileInfo) == 2 {
			fileName = fileInfo[1]
		}
	}
	if fileName == "" {
		curTimeString := time.Now().String()
		fileName = curTimeString
	}
	// 生成文件
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = ioutil.WriteFile(downloadDir+fileName, data, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("file: (%v) was downloaded\n", fileName)

	ManageFile(downloadDir, 5)
}

// 已下载文件管里
func ManageFile(dir string, limit int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 按修改时间排序，降序
	sort.Slice(files, func(i, j int) bool {
		diff := files[i].ModTime().Sub(files[j].ModTime())
		return diff > 1
	})

	// 处理超出容量限制的文件
	for ind, file := range files {
		if ind > limit-1 {
			fileName := file.Name()
			err = os.Remove(dir + fileName)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Printf("file: (%v) was removed\n", fileName)
		}
	}
}
