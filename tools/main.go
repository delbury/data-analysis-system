package main

import (
	"data-analysis-system/pkg"
	"fmt"
)

// 请求 host
const HOST = "http://127.0.0.1:4000"
const URL_PATH = "/api/env/backup"
const DOWNLOAD_DIR = "./download/"

var HEADERS = map[string]string{"Authorization": "Basic YWRtaW46MTIzNDU2YQ=="}

func main() {
	fmt.Println("start daily task")
	pkg.DailyTask(func() {
		pkg.GetDownload(HOST+URL_PATH, HEADERS, DOWNLOAD_DIR)
	}, 2, 0, 0)
}
