package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {
	res := showFunc2("test.txt")
	fmt.Printf("%s", res)
}

func showFunc2(file string) []string {
	root, _ := os.Getwd()
	filePath := path.Join(root, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("failed to read file %s", file)
		return nil
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	res := make([]string, 0, 16)
	var tempRes string
	var first bool
	for _, line := range lines {
		line = strings.TrimSpace(line)
		fmt.Printf("%s\n", line)
		if "" == line {
			continue
		}
		if strings.HasPrefix(line, "Thread") {
			if 0 != len(tempRes) {
				res = append(res, tempRes)
			}
			tempRes = ""
			first = true
		}
		parts := strings.Split(line, " ")
		if 5 != len(parts) {
			continue
		}
		if first {
			first = false
			tempRes += fmt.Sprintf("%s", parts[4])
		} else {
			tempRes += fmt.Sprintf(" <- %s", parts[4])
		}
	}
	if 0 != len(tempRes) {
		res = append(res, tempRes)
	}
	return res
}
