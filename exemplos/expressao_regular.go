package exemplos

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	example := "#GoLangCode!$!"

	// Expressão para pegar apenas alfanuméricos
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(example, "")

	fmt.Printf("A string of %s becomes %s \n", example, processedString)

}
