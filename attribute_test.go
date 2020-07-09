package go_ead

import (
	"io/ioutil"
	"testing"
)

func TestInternalAudiencePass(t *testing.T) {

	ead, err := ioutil.ReadFile("static/faids/PASS.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := ContainsInternalAudienceAttr(ead)
	if err != nil {
		t.Error(err)
	}

	if len(result) != 0 {
		t.Error("EAD contained internal audience content")
	}
}

func TestInternalAudienceFail(t *testing.T) {

	ead, err := ioutil.ReadFile("static/faids/FAIL-internal-audience.xml")

	result, err := ContainsInternalAudienceAttr(ead)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("EAD did not contain internal audience content")
	}

	if result[0] != "editionstmt" {
		t.Errorf("editionstmt element was not returned, received %s", result[0])
	}
}
