package parser

import "testing"

func TestTranslation(t *testing.T) {
	result := Translate("Text string")
	if result != "Text string -> Parsed" {
		t.Error("Invalid translation")
	}
}
