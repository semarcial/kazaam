package transform

import "testing"

func TestThenConditional(t *testing.T) {
	jsonOut := `{
		"sku": "810483427",
		"gtin": "00481048342701",
		"brand": "",
		"item_name": "Denim Blue Diamond Textured Wash Cloth",
		"size": "",
		"cost_unit": "each",
		"category": "Big Lots Catch All",
		"corporate_price": "2.99",
		"shipping_cost": "5.99",
		"shipping_weight": "0.38",
		"weight_uom": "lbs",
		"sold_length": "7.6",
		"sold_height": "1.85",
		"sold_width": "8.3",
		"parentSku": "svc720003",
		"variantsize": "Wash"
	   ,"variation_size":"Wash","variation_parent":""}`
	spec := `{
		"conditional": {
			"condition": "parentSku != \"\"",
			"then": {
				"shift": {"variation_size": "variantsize"},
				"default": {"variation_parent": ""}
			},
			"else": "{}"
		}
	}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Conditional, cfg, `{
		"sku": "810483427",
		"gtin": "00481048342701",
		"brand": "",
		"item_name": "Denim Blue Diamond Textured Wash Cloth",
		"size": "",
		"cost_unit": "each",
		"category": "Big Lots Catch All",
		"corporate_price": "2.99",
		"shipping_cost": "5.99",
		"shipping_weight": "0.38",
		"weight_uom": "lbs",
		"sold_length": "7.6",
		"sold_height": "1.85",
		"sold_width": "8.3",
		"parentSku": "svc720003",
		"variantsize": "Wash"
	}`)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}

func TestElseConditional(t *testing.T) {
	jsonOut := `{
		"sku": "810483427",
		"gtin": "00481048342701",
		"brand": "",
		"item_name": "Denim Blue Diamond Textured Wash Cloth",
		"size": "",
		"cost_unit": "each",
		"category": "Big Lots Catch All",
		"corporate_price": "2.99",
		"shipping_cost": "5.99",
		"shipping_weight": "0.38",
		"weight_uom": "lbs",
		"sold_length": "7.6",
		"sold_height": "1.85",
		"sold_width": "8.3",
		"parentSku": "svc720003",
		"variantsize": "Wash"
	,"variation_size":"","weight_uom_test":""}`
	spec := `{
		"conditional": {
			"condition": "parentSku == \"\"",
			"then": {
				"shift": {"variation_size": "variantsize"}
			},
			"else": {
				"default": {"variation_size": "",
				"weight_uom_test": ""}
			}
		}
	}`

	cfg := getConfig(spec, false)
	cfg.InPlace = false
	kazaamOut, _ := getTransformTestWrapper(Conditional, cfg, `{
		"sku": "810483427",
		"gtin": "00481048342701",
		"brand": "",
		"item_name": "Denim Blue Diamond Textured Wash Cloth",
		"size": "",
		"cost_unit": "each",
		"category": "Big Lots Catch All",
		"corporate_price": "2.99",
		"shipping_cost": "5.99",
		"shipping_weight": "0.38",
		"weight_uom": "lbs",
		"sold_length": "7.6",
		"sold_height": "1.85",
		"sold_width": "8.3",
		"parentSku": "svc720003",
		"variantsize": "Wash"
	}`)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}
