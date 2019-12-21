package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	txt, err := ioutil.ReadFile(`students.txt`)

	if err != nil {
		panic(err)
	}

	content := string(txt)

	reg := regexp.MustCompile(`"xh":"2017211573",(?s:(.*?))","xb"`)
	result1 := reg.FindAllString(content, -1)

	var str string
	for _, b := range result1 {
		str = str + b
	}
	reg2 := regexp.MustCompile(`"xmEn":"(.*?)"`)
	result2 := reg2.FindAllString(str, -1)

	fmt.Println(result2)

}
