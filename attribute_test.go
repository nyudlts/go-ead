package go_ead

import (
	"testing"
)

func TestInternalAudiencePass(t *testing.T) {

	pass := b.Get("/faids/PASS.xml")
	result, err := ContainsInternalAudienceAttr(pass)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result == true {
		t.Error("EAD contained internal audience content")
	}
}

func TestInternalAudienceFail(t *testing.T) {

	pass := b.Get("/faids/FAIL-internal-audience.xml")
	result, err := ContainsInternalAudienceAttr(pass)
	if err != nil {
		t.Errorf(err.Error())
	}
	if result != true {
		t.Error("EAD did not contain internal audience content")
	}
}