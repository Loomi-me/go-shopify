package goshopify

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed fixtures/abandonedCheckouts.json
var data []byte

func TestCanParseAbandonedCheckoutResponse(t *testing.T) {
	res := new(AbandonedCartResource)
	err := json.Unmarshal(data, &res)
	if err != nil {
		t.Errorf("failed parsing checkout response %v",err)
	}
}
