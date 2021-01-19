package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kheadjr-rv/golang"
	"github.com/kheadjr-rv/golang/language"
)

func main() {
	lang := flag.String("language", "english", "Language to speak")

	flag.Parse()

	var greeting string

	switch *lang {
	case "english":
		greeting = golang.Greet(language.English())
	case "spanish":
		greeting = golang.Greet(language.Spanish())
	default:
		fmt.Fprintf(os.Stderr, "sorry unknown language: %v\n", *lang)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
