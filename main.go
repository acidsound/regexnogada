package main

import (
	"os"
	"io/ioutil"
	"regexp"
)

func main() {
	src := getSource(os.Args[1])
	str := src
	str = regReplacer(str, `(?im)<script[^>]*>[\s\S]*?</script>`, "")
	str = regReplacer(str, `(?im)<noscript[^>]*>[\s\S]*?</noscript>`, "")
	str = regReplacer(str, `(?im)<!--[\s\S]*?-->`, "")
	str = regReplacer(str, `(?im)<style[^>]*>[\s\S]*?</style>`, "")
	str = regReplacer(str, `(?im)<html[^>]*>`, "")
	str = regReplacer(str, `(?im)<head>[\s\S]*?</head>`, "")
	str = regReplacer(str, `(?im)<article[^>]*>([\s\S]*?)</article>`, "$1")
	str = regReplacer(str, `(?im)(<(head|iframe|body|section|meta|title)[^>]*>)`, "\n$1")
	str = regReplacer(str, `(?im)(<(div|li|ul|h1|h2|h3|h4|p|img)[^>]*>)`, "\n  $1")
	str = regReplacer(str, `(?im)(<section[^>]*>)`, "\n$1")
	str = regReplacer(str, `(?im)<li value="[\d]*" [^>]*>`, `<li class="list">`)
	str = regReplacer(str, `(?im)<section class="slide-content" [^>]*>`, `<section class="slides">`)
	/* font wegiht */
	str = regReplacer(str, `(?im)<p style="font-weight:([\d]*)[^>]*>`, `<p class="fw$1">`)
	println(str)
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
