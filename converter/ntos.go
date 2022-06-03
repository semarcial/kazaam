package converter

import (
	"errors"
	"strconv"

	"github.com/semarcial/kazaam/v1/transform"
)

type Ntos struct {
	ConverterBase
}

func (c *Ntos) Convert(jsonData []byte, value []byte, args []byte) (newValue []byte, err error) {
	var jsonValue *transform.JSONValue
	jsonValue, err = c.NewJSONValue(value)
	if err != nil {
		return
	}

	if jsonValue.IsString() {
		newValue = value
	} else if jsonValue.IsNumber() {
		num := jsonValue.GetNumber()
		newValue = []byte(strconv.Quote(num.String()))
	} else {
		err = errors.New("unexpected type")
	}

	return
}
