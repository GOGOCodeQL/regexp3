package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	x := r.FindStringIndex("peach punch")
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println("[+]", x)

	y := r.FindStringSubmatch("peach punch")
	fmt.Println(y)

	z := r.FindStringSubmatchIndex("peach punch")
	fmt.Println("[!]", z)
	index := len(z)
	i := len(z)
	if i <= index {
		fmt.Println(z[i])
	}
	if i <= len(z) {
		fmt.Println(z[i])
	}
	if i < len(z)+1 {
		fmt.Println(z[i])
	}
	for i := 0; i <= len(z); i++ {
		fmt.Println("[-]", z[i])
	}

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
