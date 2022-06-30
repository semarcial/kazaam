package converter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/semarcial/kazaam/v5/transform"
)

type UpcBiglots struct {
	ConverterBase
}

func (c *UpcBiglots) Convert(jsonData []byte, value []byte, args []byte) (newValue []byte, err error) {

	var jsonValue *transform.JSONValue
	jsonValue, err = c.NewJSONValue(value)
	if err != nil {
		return nil, err
	}

	if !isNumeric(strings.ReplaceAll(jsonValue.GetStringValue(), "\"", "")) {
		return []byte(strconv.Quote(jsonValue.GetStringValue())), nil
	} else {
		value := jsonValue.GetStringValue()
		if utf8.RuneCountInString(value) > 12 {
			inputFmt := leftTruncateString(value, 12)
			newValue = []byte(strconv.Quote(inputFmt))
		} else if utf8.RuneCountInString(value) < 12 && utf8.RuneCountInString(value) > 0 {
			result := fmt.Sprintf("%012s", value)
			newValue = []byte(strconv.Quote(result))
		} else if utf8.RuneCountInString(value) == 12 {
			newValue = []byte(strconv.Quote(value))
		}
	}
	return newValue, nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 0)
	return err == nil
}

func leftTruncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}

	stringSize := utf8.RuneCountInString(str)
	if stringSize < length {
		return str
	}
	return string([]rune(str)[stringSize-length:])
}
