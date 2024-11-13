package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	//. "github.com/onsi/gomega/matchers"
)

var _ = Describe("HaveKey with iterators", func() {
	When("passed an iter.Seq2", func() {
		It("should do the right thing", func() {
			Expect(universalMapIter2).To(HaveKey("bar"))
			Expect(universalMapIter2).To(HaveKey(HavePrefix("ba")))
			Expect(universalMapIter2).NotTo(HaveKey("barrrrz"))
			Expect(universalMapIter2).NotTo(HaveKey(42))
		})
	})

	/*
		When("passed a correctly typed nil", func() {
			It("should operate succesfully on the passed in value", func() {
				var nilMap map[int]string
				Expect(nilMap).ShouldNot(HaveKey("foo"))
			})
		})

		When("the passed in key is actually a matcher", func() {
			It("should pass each element through the matcher", func() {
				Expect(stringKeys).Should(HaveKey(ContainSubstring("oo")))
				Expect(stringKeys).ShouldNot(HaveKey(ContainSubstring("foobar")))
			})

			It("should fail if the matcher ever fails", func() {
				actual := map[int]string{1: "a", 3: "b", 2: "c"}
				success, err := (&HaveKeyMatcher{Key: ContainSubstring("ar")}).Match(actual)
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())
			})
		})

		When("passed something that is not a map", func() {
			It("should error", func() {
				success, err := (&HaveKeyMatcher{Key: "foo"}).Match([]string{"foo"})
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())

				success, err = (&HaveKeyMatcher{Key: "foo"}).Match(nil)
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())
			})
		})
	*/
})
