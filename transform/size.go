package transform

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Size(spec *Config, data []byte) ([]byte, error) {
	var outData []byte
	if spec.InPlace {
		outData = data
	} else {
		outData = []byte(`{}`)
	}

	nameData, err := getJSONRaw(data, "name", true)
	if err != nil {
		return nil, err
	}
	sizeData, err := getJSONRaw(data, "size", true)
	if err != nil {
		return nil, err
	}

	sizeDataClean := strings.ReplaceAll(string(sizeData), "\"", "")

	fmt.Println(sizeDataClean)
	if string(sizeDataClean) == "null" || string(sizeDataClean) == "" {

		fmt.Println("Here")
		newName := strings.ReplaceAll(strings.ReplaceAll(string(nameData), "'", "ft"), "-", " ")

		sizePattern := `\s*[-,(\[]*\s*((?:(?:\d+|\d*\.\d+)\s*)(?:[-/_]+\s*(?:\d+|\d*\.\d+)\s*)?(?:pack|pk|in|inches|count|cnt|ct|ctn|ounce|oz|gallon|gal|pound|lb|ml|liter|l|(?:fl(?:uid)?\.?\s+(?:oz|ounce))|quart|qt|piece|pc|pint|pt|g|mg|serving|(?:sq\ )?ft|tablet|softgel|capsule|lozenge|can|bottle|sheet|roll)s?\b\.?)\s*(?:bag|bars?|box(?:es)?|cans?)?(?:\W*\beach\b\W*$)?[)\]]?`

		r := regexp.MustCompile(sizePattern)
		lowerName := strings.ToLower(newName)
		fmt.Println(lowerName)
		matches := r.FindAllString(lowerName, -1)

		sizeData = []byte(strconv.Quote(strings.TrimSpace(strings.Join(matches, ","))))
	}

	outData, err = setJSONRaw(outData, sizeData, "size")
	if err != nil {
		return nil, err
	}

	return outData, nil
}
