package main

import (
	"log"

	"github.com/alextanhongpin/nosetter/foo"
)

func main() {
	var a int
	a = 10
	log.Println(a)
	var f foo.Foo
	f.Name = "name"
}
