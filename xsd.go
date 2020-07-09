package go_ead

import (
	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xsd"
	"io/ioutil"
)

var p = parser.New()

func ValidateEAD(fa []byte) error {
	schema, err := ioutil.ReadFile("static/schema/ead.xsd")
	if err != nil {
		return err
	}
	eadxsd, err := xsd.Parse(schema)
	if err != nil {
		return err
	}
	doc, err := p.Parse(fa)
	if err != nil {
		return err
	}

	if err := eadxsd.Validate(doc); err != nil {
		return err
	}

	return nil
}
