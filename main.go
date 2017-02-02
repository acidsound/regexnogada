package main

import (
	"os"
	"io/ioutil"
	"regexp"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args)<3 {
		showUsage()
		return
	}
	/* load source file */
	src := getSource(os.Args[1])
	str := src
	/* open rule file */
	file, err := os.Open(os.Args[2])
	check(err)
	defer file.Close()

	/* convert */
	result := convertRegEx(file, str)

	/* output */
	if len(os.Args) > 3 {
		check(saveResult(os.Args[3], result))
	} else {
		fmt.Println(result)
	}
}
func convertRegEx(file *os.File, str string) string {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != 35 {
			tokens := strings.Split(line, "\t")
			replacement := ""
			if len(tokens) > 1 {
				/* \n, \b 같은 escape 문자들 변환 : https://play.golang.org/p/AZ82pxX64b */
				replacement, _ = strconv.Unquote(`"` + tokens[1] + `"`)
			}
			str = regReplacer(str, tokens[0], replacement)
		}
	}
	return str
}
func saveResult(fileName string, result string) error {
	return ioutil.WriteFile(fileName, []byte(result), 0644)
}
func showUsage() {
	println(`
Usage: regexnogada [source file] [rule file]
  or
  regexnogada [source file] [rule file] [target file]

If called without a target file, 'regexnogada' will run a rule file at the terminal.
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
