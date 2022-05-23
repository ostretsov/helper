package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

	placeholders := placeholders(os.Args[3])
	if len(placeholders) == 0 {
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

		cmd := os.Args[3]
		for _, p := range placeholders {
			cmd = strings.Replace(cmd, "${"+strconv.Itoa(p)+"}", line[p-1], -1)
		}
		fmt.Println(cmd)
	}
}

func placeholders(s string) []int {
	re := regexp.MustCompile(`\${(\d+)}`)

	var pl []int
	found := re.FindAllStringSubmatch(s, -1)
	for _, f := range found {
		colNum, err := strconv.Atoi(f[1])
		if err != nil {
			panic(fmt.Sprintf("placeholder must be a number: %s", err))
		}
		pl = append(pl, colNum)
	}
	return pl
}
