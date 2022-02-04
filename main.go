package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/signaes/csv2json/converter"
)

var separator = flag.String("separator", ";", "The csv separator character")

func main() {
	fileName := os.Args[1]

	if fileName == "" {
		log.Fatalln("The file name is required")
	}

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	lines, err := reader.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	var pretty bytes.Buffer

	json.Indent(&pretty, converter.ToJsonBytes(lines, *separator), "", "  ")

	pretty.WriteTo(os.Stdout)
}
