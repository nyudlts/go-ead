package schema

import (
	"github.com/lestrrat/go-libxml2/parser"
	"github.com/lestrrat/go-libxml2/xsd"
	"github.com/nyudlts/go-ead/box"
)

var p = parser.New()
var b = box.Box

func ValidateEAD(fa []byte) error {
	eadxsd, err := xsd.Parse(b.Get("/schema/ead.xsd"))
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
