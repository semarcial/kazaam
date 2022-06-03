package converter

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/semarcial/kazaam/v4/transform"
)

type Trim struct {
	ConverterBase
}

func (c *Trim) Convert(jsonData []byte, value []byte, args []byte) (newValue []byte, err error) {
	newValue = value

	var jsonValue *transform.JSONValue
	jsonValue, err = c.NewJSONValue(value)
	if err != nil {
		return
	}

	if jsonValue.IsString() == false {
		err = errors.New("value must be string for trim converter")
		return
	}

	origValue := jsonValue.GetStringValue()

	newValue = []byte(strconv.Quote(strings.Trim(origValue, " \t")))

	return
}
