package main

import (
	"fmt"

	"github.com/wolfeidau/shorten"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	length = kingpin.Arg("length", "Length of random string").Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Println("Random String:", shorten.RandSeq(*length))
}
