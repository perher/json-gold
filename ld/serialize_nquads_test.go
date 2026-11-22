package ld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNQuads_RejectsRelativeIRIs(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
	}{
		{"subject", `<relative> <http://example.org/p> "x" .`},
		{"predicate", `_:b0 <relative> "x" .`},
		{"object", `<http://example.org/s> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <CredentialStatusList2017> .`},
		{"literal datatype", `<http://example.org/s> <http://example.org/p> "x"^^<relative> .`},
		{"graph name", `<http://example.org/s> <http://example.org/p> "x" <relative> .`},
		{"empty scheme", `<:relative> <http://example.org/p> "x" .`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseNQuads(tc.input)
			assert.ErrorContains(t, err, "error while parsing N-Quads; invalid quad. line: 1")
			assert.ErrorContains(t, err, "is not an absolute IRI")
		})
	}
}

func TestParseNQuads_AcceptsAbsoluteIRIs(t *testing.T) {
	input := `<http://example.org/s> <http://example.org/p> <did:example:123> <urn:uuid:42> .
_:b0 <http://example.org/p> "x"^^<http://www.w3.org/2001/XMLSchema#string> .
_:b0 <http://example.org/p> "x"@en _:g0 .
`

	ds, err := ParseNQuads(input)
	assert.NoError(t, err)
	assert.Len(t, ds.Graphs["urn:uuid:42"], 1)
	assert.Len(t, ds.Graphs["@default"], 1)
	assert.Len(t, ds.Graphs["_:g0"], 1)
}
