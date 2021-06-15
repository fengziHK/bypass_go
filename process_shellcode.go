package main

import (
	"fmt"
	"encoding/hex"
	"os"
	"strings"
	"strconv"
)

func main() {
	param := os.Args[1]
	isArr := strings.Contains(param, ",")
	if isArr {
		context := strings.Split(param, ",")
		size := len(context)
		dataArr := make([]byte, size)
		for i, v := range context {
			val, _ := strconv.Atoi(v)
			dataArr[i] = byte(val)
		}
		//fmt.Println(dataArr)
		// hexToString
		fmt.Println(hex.EncodeToString([]byte(dataArr)))
	} else {
		val, _ := strconv.Atoi(param)
		data := make([]byte, 1)
		data[0] = byte(val)
		//fmt.Println(data)
		// hexToString
		fmt.Println(hex.EncodeToString([]byte(data)))
	}
}