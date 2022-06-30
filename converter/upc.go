package converter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/semarcial/kazaam/v5/transform"
)

var errorInvalidArgs = errors.New("invalid value or arguments for upc converter")

type Upc struct {
	ConverterBase
}

func (c *Upc) Convert(jsonData []byte, value []byte, args []byte) (newValue []byte, err error) {

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

	if isNumeric(strings.ReplaceAll(jsonValue.GetStringValue(), "\"", "")) && isNumeric(argsValue.GetStringValue()) {
		limit, intErr := strconv.ParseInt(argsValue.GetStringValue(), 10, 0)
		if intErr != nil {
			err = intErr
			return
		}

		value := jsonValue.GetStringValue()
		if utf8.RuneCountInString(value) > int(limit) {
			inputFmt := leftTruncateString(value, int(limit))
			newValue = []byte(strconv.Quote(inputFmt))
		} else if utf8.RuneCountInString(value) < int(limit) && utf8.RuneCountInString(value) > 0 {
			result := fmt.Sprintf("%012s", value)
			newValue = []byte(strconv.Quote(result))
		} else if utf8.RuneCountInString(value) == int(limit) {
			newValue = []byte(strconv.Quote(value))
		}
	}

	return
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
