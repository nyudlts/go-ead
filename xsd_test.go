package go_ead

import (
	"testing"
)

func TestXSDPass(t *testing.T) {

	pass := b.Get("/faids/PASS.xml")
	err := ValidateEAD(pass); if err != nil {
		t.Errorf(err.Error())
	}
}

func TestXSDFailWF(t *testing.T) {
	failwf := b.Get("/faids/FAIL-not-well-formed.xml")
	err := ValidateEAD(failwf); if err == nil {
		t.Error(err.Error())
	}
}

func TestXSDFailSchema(t *testing.T) {
	failwf := b.Get("/faids/FAIL-schema.xml")
	err := ValidateEAD(failwf); if err == nil {
		t.Error(err.Error())
	}
}
