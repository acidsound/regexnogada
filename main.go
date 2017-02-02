package main

import (
	"os"
	"io/ioutil"
	"regexp"
	"bufio"
	"strings"
	"fmt"
)

func main() {
	if len(os.Args)<3 {
		showUsage()
		return
	}
	src := getSource(os.Args[1])
	str := src
	file, err := os.Open(os.Args[2])
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != 35 {
			tokens := strings.Split(line, "\t")
			if len(tokens)>1 {
				/* file 에서 읽은 \n 을 newline으로 대체 */
				str = regReplacer(str, tokens[0], strings.Replace(tokens[1], "\\n", "\n", -1 ))
			} else {
				str = regReplacer(str, tokens[0], "")
			}
		}
	}
	fmt.Println(str)
}
func showUsage() {
	println(`
Usage: regexnogada <source file> <rule file>
	`)
}
func regReplacer(source string, regex string, replacement string) string {
	return regexp.MustCompile(regex).ReplaceAllString(source, replacement)
}
func getSource(fn string) string {
	dat, err := ioutil.ReadFile(fn)
	check(err)
	str := string(dat)
	return str
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
