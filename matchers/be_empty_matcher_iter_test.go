package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BeEmpty with iterators", func() {
	When("passed an iterator type", func() {
		It("should do the right thing", func() {
			Expect(emptyIter).To(BeEmpty())
			Expect(emptyIter2).To(BeEmpty())

			Expect(universalIter).NotTo(BeEmpty())
			Expect(universalIter2).NotTo(BeEmpty())
		})
	})

	When("passed a correctly typed nil", func() {
		It("should be true", func() {
			var nilIter func(func(string) bool)
			Expect(nilIter).Should(BeEmpty())

			var nilIter2 func(func(int, string) bool)
			Expect(nilIter2).Should(BeEmpty())
		})
	})
})
