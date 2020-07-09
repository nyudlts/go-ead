package go_ead

import (
	"io/ioutil"
	"testing"
)

func TestXSDPass(t *testing.T) {

	ead, err := ioutil.ReadFile("static/faids/PASS.xml")
	if err != nil {
		t.Error(err)
	}

	err = ValidateEAD(ead); if err != nil {
		t.Errorf(err.Error())
	}
}

func TestXSDFailWF(t *testing.T) {
	ead, err := ioutil.ReadFile("static/faids/FAIL-not-well-formed.xml")

	err = ValidateEAD(ead); if err == nil {
		t.Error(err)
	}
}

func TestXSDFailSchema(t *testing.T) {
	ead, err := ioutil.ReadFile("static/faids/FAIL-schema.xml")

	err = ValidateEAD(ead); if err == nil {
		t.Error(err.Error())
	}
}
