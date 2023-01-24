package enums

import (
	"bufio"
	"js-strings/config"
	"log"
	"os"
	"regexp"
	"strings"
)

func New(c *config.Config) Enums {

	file, err := os.Open(c.StringsEnumPath)
	if err != nil {
		log.Fatalf("can't open strings enum file with path %s with error: %v", c.StringsEnumPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), c.StringsEnumName) {
			break
		}
	}

	stringEnums := Enums{}

	re := regexp.MustCompile(`[A-Za-z]+`)

	for scanner.Scan() {
		t := scanner.Text()

		if strings.Contains(t, "}") {
			break
		}

		f := re.FindStringSubmatch(t)

		stringEnums = append(stringEnums, f...)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error with reading the strings enum file: ", err)
	}

	return stringEnums
}
