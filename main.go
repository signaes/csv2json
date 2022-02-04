package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
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

	json.Indent(&pretty, ToJsonBytes(lines, *separator), "", "  ")

	pretty.WriteTo(os.Stdout)
}

func ToJsonBytes(lines [][]string, separator string) []byte {
	header := strings.Split(lines[0][0], separator)

	var data []interface{}

	for _, line := range lines[1:] {
		var item map[string]interface{} = make(map[string]interface{})

		body := strings.Split(line[0], separator)

		for i, head := range header {
			if i < len(body) {
				item[head] = body[i]
			} else {
				item[head] = ""
			}
		}

		data = append(data, item)
	}

	bytes, err := json.Marshal(data)

	if err != nil {
		log.Fatalln(err)
	}

	return bytes
}
