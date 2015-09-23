package utils_test

import (
	. "github.com/bkroeze/go.utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stringutils", func() {
	Context("Parsing a CSV line", func() {
		It("Should parse a simple line into fields", func() {
			line := "one,two,three"
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "two", "three"}))
		})

		It("Should parse a line with a quote", func() {
			line := "one,\"two\",three"
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "two", "three"}))
		})

		It("Should automatically trim whitespace between fields", func() {
			line := "one, two, three"
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "two", "three"}))
		})

		It("Should automatically trim whitespace between quoted fields", func() {
			line := "one, \"two\",\"three\" "
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "two", "three"}))
		})

		It("Should handle internal quotes for quoted fields", func() {
			line := "one, \"I have an \"Internal\" quote\",\"three\" "
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "I have an \"Internal\" quote", "three"}))
		})

		It("Should handle internal commas for quoted fields", func() {
			line := "one, \"I have, a \"comma\" inside\",\"three\" "
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "I have, a \"comma\" inside", "three"}))
		})

		It("Should allow fields to start with internal quotes", func() {
			line := "\"\"one\"\", \"\"two\", \"three\"\""
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"\"one\"", "\"two", "three\""}))
		})
		It("Should allow empty records", func() {
			line := "one,,three"
			Expect(SplitCSVStringIntoFields(line)).To(HaveLen(3))
			Expect(SplitCSVStringIntoFields(line)).To(ConsistOf([]string{"one", "", "three"}))
		})
	})

	Context("Parsing multiple lines", func() {
		It("Should parse multiple lines", func() {
			lines := "one,two\nthree,four"
			parsed, err := SplitMultilineCSV(lines, false)
			Expect(err).To(BeNil())
			Expect(parsed).To(HaveLen(2))
			Expect(parsed[0]).To(ConsistOf([]string{"one", "two"}))
			Expect(parsed[1]).To(ConsistOf([]string{"three", "four"}))
		})

		It("Should parse multiple lines with a header", func() {
			lines := "f1,f2\none,two\nthree,four"
			parsed, err := SplitMultilineCSV(lines, true)
			Expect(err).To(BeNil())
			Expect(parsed).To(HaveLen(2))
			Expect(parsed[0]).To(ConsistOf([]string{"one", "two"}))
			Expect(parsed[1]).To(ConsistOf([]string{"three", "four"}))
		})
	})

	Context("Inserting text", func() {
		It("Should find the start and end tokens of a string", func() {
			template := "<a></a>"
			start, end, ok := GetTokenPositions("<a>", "</a>", template)
			Expect(ok).To(BeTrue())
			Expect(start).To(Equal(0))
			Expect(end).To(Equal(3))
		})

		It("Should find the start and end tokens of a multiline string", func() {
			template := `<a>
test
</a>`
			start, end, ok := GetTokenPositions("<a>", "</a>", template)
			Expect(ok).To(BeTrue())
			Expect(start).To(Equal(0))
			Expect(end).To(Equal(9))
		})

		It("Should remove text between tokens", func() {
			template := "<a>test</a>"
			expected := "<a></a>"
			result, changed := RemoveTextBetween("<a>", "</a>", template)
			Expect(changed).To(BeTrue())
			Expect(result).To(Equal(expected))
		})

		It("Should remove text between tokens in a multiline string", func() {
			template := `<a>
test
</a>`
			expected := "<a></a>"
			result, changed := RemoveTextBetween("<a>", "</a>", template)
			Expect(changed).To(BeTrue())
			Expect(result).To(Equal(expected))
		})

		It("Should not remove anything if token not found", func() {
			template := "<a>test"
			expected := "<a>test"
			result, changed := RemoveTextBetween("<a>", "</a>", template)
			Expect(changed).To(BeFalse())
			Expect(result).To(Equal(expected))
		})

		It("Should insert text between HTML comment tags", func() {
			template := "<!-- test --><!-- /test -->"
			expected := "<!-- test -->inserted<!-- /test -->"
			Expect(InsertTextBetween("<!-- test -->", "<!-- /test -->", template, "inserted")).To(Equal(expected))
		})

		It("Should work on multiline tags", func() {
			template := `
Test
<!-- runetable -->
Inside
<!-- /runetable -->
`
			expected := `
Test
<!-- runetable -->Inserted<!-- /runetable -->
`
			Expect(InsertTextBetween("<!-- runetable -->", "<!-- /runetable -->", template, "Inserted")).To(Equal(expected))
		})

		It("Should work on multiline tags 2", func() {
			template := "test\n <a>\nfoo\n</a>"
			expected := "test\n <a>Inserted</a>"
			Expect(InsertTextBetween("<a>", "</a>", template, "Inserted")).To(Equal(expected))
		})

		It("Should insert text between tags, removing the prior text", func() {
			template := "<!-- test -->\nfoo\n<!-- /test -->"
			expected := "<!-- test -->inserted<!-- /test -->"
			Expect(InsertTextBetween("<!-- test -->", "<!-- /test -->", template, "inserted")).To(Equal(expected))
		})

		It("Should not insert or remove anything if both tags are not found", func() {
			template := "<!-- test -->foo<!-- /xxx -->"
			Expect(InsertTextBetween("<!-- test -->", "<!-- /test -->", template, "inserted")).To(Equal(template))
		})

	})
})
