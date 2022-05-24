package helpers

import (
	"fmt"
	"github.com/essentialkaos/translit/v2"
	"io"
	"log"
	"os"
)

func TranslitRuToEng() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(translit.EncodeToISO9B(string(in)))
}
