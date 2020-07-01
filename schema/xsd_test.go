package schema

import (
	"github.com/nyudlts/go-ead/box"
	"testing"
)

var bx = box.Box

func TestXSDPass(t *testing.T) {

	pass := bx.Get("/faids/PASS.xml")
	err := ValidateEAD(pass); if err != nil {
		t.Errorf(err.Error())
	}
}

func TestXSDFailWF(t *testing.T) {
	failwf := bx.Get("/faids/FAIL-not-well-formed.xml")
	err := ValidateEAD(failwf); if err == nil {
		t.Error(err.Error())
	}
}

func TestXSDFailSchema(t *testing.T) {
	failwf := bx.Get("/faids/FAIL-schema.xml")
	err := ValidateEAD(failwf); if err == nil {
		t.Error(err.Error())
	}
}
