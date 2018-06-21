package ical_test

import (
	"bytes"

	"github.com/hrkipp/ical"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContentLines", func() {
	DescribeTable("Degenerate Cases",
		func(doc []byte, expected []string) {
			Expect(
				ical.ContentLines(bytes.NewReader(doc)),
			).To(BeEquivalentTo(expected))
		},
		Entry("nil", nil, nil),
		Entry("empty", []byte{}, nil),
	)

	DescribeTable("Cases",
		func(doc string, expected []string) {
			Expect(
				ical.ContentLines(bytes.NewReader([]byte(doc))),
			).To(BeEquivalentTo(expected))
		},
		Entry("simple single line", "foo\r\n", []string{"foo"}),
		Entry("simple single line no break", "foo", []string{"foo"}),
		Entry("simple single line, only newline", "foo\n", []string{"foo"}),
		Entry("line split once", "fo\r\n o", []string{"foo"}),
		Entry("line split once with tab", "fo\r\n\to", []string{"foo"}),
		Entry("line split once, with final newline", "fo\r\n o\r\n", []string{"foo"}),
		Entry("line split once, with final newline and tab", "fo\r\n\to\r\n", []string{"foo"}),

		Entry("two simple lines", "foo\r\nbar\r\n", []string{"foo", "bar"}),
	)

})
