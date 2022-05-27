package transform

import (
	"fmt"
)

type TransformFunc func(spec *Config, data []byte) ([]byte, error)

// Shift moves values from one provided json path to another in raw []byte.
func Conditional(spec *Config, data []byte) ([]byte, error) {
	var outData []byte
	if spec.InPlace {
		outData = data
	} else {
		outData = []byte(`{}`)
	}

	defaultSpecTypes := map[string]TransformFunc{
		"pass":        Pass,
		"shift":       Shift,
		"extract":     Extract,
		"default":     Default,
		"delete":      Delete,
		"concat":      Concat,
		"coalesce":    Coalesce,
		"timestamp":   Timestamp,
		"uuid":        UUID,
		"steps":       Steps,
		"merge":       Merge,
		"conditional": Conditional,
	}

	if conditional, ok := (*spec.Spec)["conditional"]; ok {
		if !ok {
			return nil, ParseError(fmt.Sprintln("Warn: Missing conditional"))
		}

		conditionals := conditional.(map[string]interface{})

		if len(conditionals) >= 2 {
			if condition, then := conditionals["condition"], conditionals["then"]; condition == nil && then == nil {
				return nil, ParseError(fmt.Sprintln("Warn: Condition should have at least 2 sections, condition and then"))
			}
		}

		condition := fmt.Sprintf("%v", conditionals["condition"])

		be, err := NewBasicExpr(data, condition)
		if err != nil {
			return nil, ParseError(fmt.Sprintf("Warn: Condition %s failed to compile %v", condition, err))
		}

		evaluation, _ := be.Eval()
		var operation map[string]interface{}
		if evaluation {
			operation = conditionals["then"].(map[string]interface{})
		} else {
			operation = conditionals["else"].(map[string]interface{})
		}

		var dataOutput []byte = data
		for key, value := range operation {
			function := defaultSpecTypes[key]
			specFunction := value.(map[string]interface{})
			config := Config{Spec: &specFunction, Require: spec.Require, InPlace: true}
			dataOutput, err = function(&config, dataOutput)
			if err != nil {
				return nil, ParseError(fmt.Sprintf("Warn: Error executing %s operation %v", key, err))
			}
		}
		outData = dataOutput

	}

	return outData, nil
}
