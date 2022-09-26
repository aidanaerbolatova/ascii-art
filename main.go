package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		os.Exit(1)
	} else {
		if args[0] == "" {
			return
		}
		if isEmpty(args[0]) {
			count := len(args[0]) / 2
			for i := 0; i < count; i++ {
				fmt.Println()
			}
			return
		}
		data, err := ioutil.ReadFile("standard.txt")
		if err != nil {
			os.Exit(1)
		}
		x := makeMap(string(data))
		args[0] = strings.ReplaceAll(args[0], "\n", "\\n")
		split_data := strings.Split(args[0], "\\n")
		for _, value := range split_data {
			fmt.Print(getToStr(value, x))
		}

	}
}

func isEmpty(s string) bool {
	return strings.ReplaceAll(s, "\\n", "") == ""
}

func makeMap(s string) map[rune]string {
	simvol := make(map[rune]string)
	temp := ""
	count := 0
	j := rune(32)

	for _, data := range s {
		temp += string(data)
		if data == '\n' {
			count++
		}
		if count == 9 {
			count = 0
			simvol[j] = temp[1 : len(temp)-1]
			temp = ""
			j++
		}
	}
	return simvol
}

func getToStr(s string, x map[rune]string) string {
	temp := [8]string{}
	if s == "" {
		return "\n"
	}
	for _, char := range s {
		for index, value := range strings.Split(x[char], "\n") {
			temp[index] += value
		}
	}
	result := ""
	for _, value := range temp {
		result += value + "\n"
	}
	return result
}
