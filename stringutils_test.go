package utils_test

import (
	. "hillsorcerer.com/utils"

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
})
