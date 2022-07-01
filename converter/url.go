package converter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/semarcial/kazaam/v5/transform"
)

type Url struct {
	ConverterBase
}

func (c *Url) Convert(jsonData []byte, value []byte, args []byte) (newValue []byte, err error) {

	newValue = []byte(value)

	var jsonValue, argsValue *transform.JSONValue
	jsonValue, err = c.NewJSONValue(value)
	if err != nil {
		return nil, err
	}

	argsValue, err = transform.NewJSONValue(args)
	if err != nil {
		return
	}

	if !jsonValue.IsString() || !argsValue.IsString() {
		err = errorInvalidArgs
		return
	}

	url := jsonValue.GetStringValue()
	protocol := argsValue.GetStringValue()

	if !strings.HasPrefix(url, "http:") || !strings.HasPrefix(url, "https:") || !strings.HasPrefix(url, "http") || !strings.HasPrefix(url, "https") {
		var result string
		if strings.Contains(protocol, ":") {
			result = fmt.Sprintf("%s%s", protocol, url)
		} else {
			result = fmt.Sprintf("%s:%s", protocol, url)
		}
		newValue = []byte(strconv.Quote(strings.TrimSpace(result)))
	}

	if strings.Contains(string(newValue), "set=") && !strings.Contains(string(newValue), "?set=") {
		result := strings.ReplaceAll(string(newValue), "set=", "?set=")
		newValue = []byte(result)
	}

	return
}
