package go_ead

import (
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
)

func ContainsInternalAudienceAttr(faid []byte) ([]string, error) {
	elements := []string{}

	doc, err := libxml2.Parse(faid)
	if err != nil {
		return elements, err
	}
	defer doc.Free()
	root, err := doc.DocumentElement()
	ctx, err := xpath.NewContext(root)
	if err != nil {
		return elements, err
	}
	defer ctx.Free()
	if err := ctx.RegisterNS(prefix, nsuri); err != nil {
		return elements, err
	}
	nodelist := xpath.NodeList(ctx.Find(`//ead:*[@audience="internal"]`))

	for i := 0; i < len(nodelist); i++ {
		elements = append(elements, nodelist[i].NodeName())
	}

	return elements, nil
}
