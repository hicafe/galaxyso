package service

import (
	"bufio"
	"fmt"
	"io"
	"models"
	"os"
	"strconv"
	"strings"
)

var SourceUrls = make([]models.SourceUrl, 0)

func InitSourceUrls() {
	fi, err := os.Open("sourceurl.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	reader := bufio.NewReader(fi)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		array := strings.Split(line, ",")
		url := array[0]
		deep := array[1]
		deep = strings.Replace(deep, "\r\n", "", -1)
		sourceUrl := models.SourceUrl{}
		sourceUrl.Url = url
		sourceUrl.Deep, _ = strconv.Atoi(deep)
		SourceUrls = append(SourceUrls, sourceUrl)
	}

	fmt.Println("SourceUrls init success!", SourceUrls)
}
