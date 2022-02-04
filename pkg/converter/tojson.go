package converter

import (
	"encoding/json"
	"log"
	"strings"
)

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
