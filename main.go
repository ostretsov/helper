package main

import (
	"helper/helpers"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("No arguments provided")
	}
	switch os.Args[1] {
	case "csv-to-commands", "csvcmd":
		helpers.CSVToCommands()
	case "translit-ru-to-eng", "tr":
		helpers.TranslitRuToEng()
	}
}
