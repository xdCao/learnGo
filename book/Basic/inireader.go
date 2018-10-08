package Basic

import (
	"bufio"
	"os"
	"strings"
)

func main() {

}

func getValue(filename, section, expectKey string) string {

	file, err := os.Open(filename)
	if err != nil {
		return ""
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var sectionName string

	for {

		linestr, e := reader.ReadString('\n')
		if e != nil {
			break
		}

		linestr = strings.TrimSpace(linestr)

		if linestr == "" {
			continue
		}

		if linestr[0] == ';' {
			continue
		}

		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			sectionName = linestr[1 : len(linestr)-1]
		} else if sectionName == section {
			pair := strings.Split(linestr, "=")
			if len(pair) == 2 {
				key := strings.TrimSpace(pair[0])
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}
		}

	}

}
