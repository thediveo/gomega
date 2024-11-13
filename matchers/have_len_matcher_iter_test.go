package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HaveLen with iterators", func() {
	When("passed an iterator type", func() {
		It("should do the right thing", func() {
			Expect(emptyIter).To(HaveLen(0))
			Expect(emptyIter2).To(HaveLen(0))

			Expect(universalIter).To(HaveLen(len(universalElements)))
			Expect(universalIter2).To(HaveLen(len(universalElements)))
		})
	})

	When("passed a correctly typed nil", func() {
		It("should operate succesfully on the passed in value", func() {
			var nilIter func(func(string) bool)
			Expect(nilIter).Should(HaveLen(0))

			var nilIter2 func(func(int, string) bool)
			Expect(nilIter2).Should(HaveLen(0))
		})
	})
})
