package transform

import (
	"html"
	"regexp"
	"strconv"
	"strings"
)

const sizePattern = `\s*[-,(\[]*\s*((?:(?:\d+|\d*\.\d+)\s*)(?:[-/_]+\s*(?:\d+|\d*\.\d+)\s*)?(?:pack|pk|in|inches|count|cnt|ct|ctn|ounce|oz|gallon|gal|pound|lb|ml|liter|l|(?:fl(?:uid)?\.?\s+(?:oz|ounce))|quart|qt|piece|pc|pint|pt|g|mg|serving|(?:sq\ )?ft|tablet|softgel|capsule|lozenge|can|bottle|sheet|roll)s?\b\.?)\s*(?:bag|bars?|box(?:es)?|cans?)?(?:\W*\beach\b\W*$)?[)\]]?`
const oneCount = "1 ct"

func Size(spec *Config, data []byte) ([]byte, error) {
	var outData []byte
	if spec.InPlace {
		outData = data
	} else {
		outData = []byte(`{}`)
	}

	nameData, err := getJSONRaw(data, (*spec.Spec)["source"].(string), true)
	if err != nil {
		return nil, err
	}
	sizeData, err := getJSONRaw(data, (*spec.Spec)["targetPath"].(string), true)
	if err != nil {
		return nil, err
	}

	replacer := strings.NewReplacer("'", "ft", "-", " ", "\"", "in")

	unscaped := html.UnescapeString(strings.ReplaceAll(string(sizeData), "\"", ""))
	sizeDataClean := replacer.Replace(unscaped)
	sizeData = []byte(strconv.Quote(sizeDataClean))

	if string(sizeDataClean) == "null" || string(sizeDataClean) == "" {

		output := replacer.Replace(html.UnescapeString(string(nameData)))
		pattern := (*spec.Spec)["pattern"]

		var re *regexp.Regexp
		if pattern == nil {
			re = regexp.MustCompile(sizePattern)
		} else {
			re = regexp.MustCompile(pattern.(string))
		}

		lowerName := strings.ToLower(output)
		matches := re.FindAllString(lowerName, -1)

		defaultValue := (*spec.Spec)["default"]
		var defaultSize string
		if defaultValue == nil {
			defaultSize = oneCount
		} else {
			defaultSize = defaultValue.(string)
		}

		if len(matches) != 0 {
			sizeData = []byte(strconv.Quote(strings.TrimSpace(strings.Join(matches, ","))))
		} else {
			sizeData = []byte(strconv.Quote(defaultSize))
		}

	}

	outData, err = setJSONRaw(outData, sizeData, "size")
	if err != nil {
		return nil, err
	}

	return outData, nil
}
