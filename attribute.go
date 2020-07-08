package go_ead

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
)

func ContainsInternalAudienceAttr(faid []byte) (bool, error) {
	default_result := false
	doc, err := libxml2.Parse(faid)
	if err != nil {
		return default_result, err
	}
	defer doc.Free()
	root, err := doc.DocumentElement()
	ctx, err := xpath.NewContext(root)
	if err != nil {
		return default_result, err
	}
	defer ctx.Free()
	if err := ctx.RegisterNS(prefix, nsuri); err != nil {
		return default_result, err
	}
	n := xpath.NodeList(ctx.Find(`//ead:*[@audience="internal"]`))

	if len(n) != 0 {
		return true, nil
	} else {
		return false, nil
	}

}
