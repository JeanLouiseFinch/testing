package main

import (
	"fmt"

	"github.com/JeanLouiseFinch/shorter"
)

func main() {
	var s shorter.Shortener
	s = shorter.NewShort()

	fmt.Println(s.Shorten("asdasd.com/asdasda/sdasd/asfdaf/sdfg"))
	//asdasd.com/0

	fmt.Println(s.Shorten(".com/asdasda/sdasd/asfdaf/sdfg"))
	//
	// тут ошибка в пути

	fmt.Println(s.Shorten("asdasd.com/asdasda/sdasd/asfdaf/sdfg"))
	//asdasd.com/0
	//повтор

	fmt.Println(s.Shorten("http://asdasd.com/asdasda/sdasd/asfdaf/sdfg"))
	//http://asdasd.com/1

	fmt.Println(s.Shorten("http:asdasd.com/asdasda/sdasd/asfdaf/sdfg"))
	//
	// тут ошибка в пути

	fmt.Println(s.Resolve("http://asdasd.com/1"))
	//http://asdasd.com/asdasda/sdasd/asfdaf/sdfg

	fmt.Println(s.Resolve("asdasd.com/0"))
	//asdasd.com/asdasda/sdasd/asfdaf/sdfg

	fmt.Println(s.Resolve("asdasd.com/0123"))
	//
	//такого не было
}
