package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func CSVToCommands() {
	if len(os.Args) != 4 {
		log.Fatal(`csv-to-commands /file.csv '/cmd ${1} ${2}'`)
	}

	csvFile := os.Args[2]
	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("failed to open file %s: %s", csvFile, err)
	}
	defer func() { _ = f.Close() }()

	placeholders := countPlaceholders(os.Args[3])
	if placeholders == 0 {
		log.Fatal("no placeholders found in command")
	}

	r := csv.NewReader(f)
	for {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read line: %s", err)
		}

		if len(line) != placeholders {
			log.Fatalf("line has %d columns, but %d placeholders", len(line), placeholders)
		}

		cmd := os.Args[3]
		for i := 1; i <= placeholders; i++ {
			cmd = strings.Replace(cmd, "${"+strconv.Itoa(i)+"}", line[i-1], -1)
		}
		fmt.Println(cmd)
	}
}

func countPlaceholders(s string) int {
	count := 0
	for i := 1; ; i++ {
		if strings.Contains(s, "${"+strconv.Itoa(i)+"}") {
			count++
		} else {
			break
		}
	}
	return count
}
