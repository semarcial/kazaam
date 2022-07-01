package transform

import "testing"

func TestExtractFeetSize(t *testing.T) {
	spec := `{
		"source": "name",
		"targetPath": "size",
		"pattern": null
	}`
	jsonIn := `{"name":"My Magic Carpet Ramage Maroon Washable Area Rug, (5' x 7')", "size": null}`
	jsonOut := `{"size":"(5ft , 7ft)"}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Size, cfg, jsonIn)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}

func TestExtractPiecesSizeWhenNull(t *testing.T) {
	spec := `{
		"source": "name",
		"targetPath": "size",
		"pattern": null
	}`
	jsonIn := `{"name":"Tasha Purple & Gray Medallion Queen 6-Piece Sheet Set Living Colors", "size": null}`
	jsonOut := `{"size":"6 piece"}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Size, cfg, jsonIn)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}

func TestExtractPiecesSizeWhenEmpty(t *testing.T) {
	spec := `{
		"source": "name",
		"targetPath": "size",
		"pattern": null
	}`
	jsonIn := `{"name":"Tasha Purple & Gray Medallion Queen 6-Piece Sheet Set Living Colors", "size": ""}`
	jsonOut := `{"size":"6 piece"}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Size, cfg, jsonIn)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}

func TestNotExtractPiecesSize(t *testing.T) {
	spec := `{
		"source": "name",
		"targetPath": "size",
		"pattern": null
	}`
	jsonIn := `{"name":"Tasha Purple & Gray Medallion Queen 6-Piece Sheet Set Living Colors", "size": "6 pieces"}`
	jsonOut := `{"size":"6 pieces"}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Size, cfg, jsonIn)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}

func TestNotExtractCustomPattern(t *testing.T) {
	spec := `{
		"source": "name",
		"targetPath": "size",
		"pattern": "([a-z])\\w+"
	}`
	jsonIn := `{"name":"Tasha Purple & Gray Medallion Queen 6-Piece Sheet Set Living Colors", "size": ""}`
	jsonOut := `{"size":"tasha,purple,gray,medallion,queen,piece,sheet,set,living,colors"}`

	cfg := getConfig(spec, false)
	kazaamOut, _ := getTransformTestWrapper(Size, cfg, jsonIn)
	areEqual, _ := checkJSONBytesEqual(kazaamOut, []byte(jsonOut))

	if !areEqual {
		t.Error("Transformed data does not match expectation.")
		t.Log("Expected:   ", jsonOut)
		t.Log("Actual:     ", string(kazaamOut))
		t.FailNow()
	}
}
