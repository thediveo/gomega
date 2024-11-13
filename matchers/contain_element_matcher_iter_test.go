//go:build go1.23

package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/matchers"
)

var _ = Describe("ContainElement on iterators", func() {

	elements := []string{"foo", "bar", "baz"}
	elementIter := func(yield func(string) bool) {
		for _, element := range elements {
			if !yield(element) {
				return
			}
		}
	}
	elementIter2 := func(yield func(int, string) bool) {
		for idx, element := range elements {
			if !yield(idx, element) {
				return
			}
		}
	}

	When("only matching", func() {
		Context("and expecting a non-matcher", func() {
			It("should do the right thing", func() {
				Expect(elementIter).To(ContainElement("baz"))
				Expect(elementIter).NotTo(ContainElement("barrrrz"))

				Expect(elementIter2).To(ContainElement("baz"))
				Expect(elementIter2).NotTo(ContainElement("barrrrz"))
			})
		})

		Context("and expecting a matcher", func() {
			It("should pass each element through the matcher", func() {
				Expect(elementIter).To(ContainElement(HaveLen(3)))
				Expect(elementIter).NotTo(ContainElement(HaveLen(5)))

				Expect(elementIter2).To(ContainElement(HaveLen(3)))
				Expect(elementIter2).NotTo(ContainElement(HaveLen(5)))
			})

			It("should power through even if the matcher ever fails", func() {
				elements := []any{1, 2, "3", 4}
				it := func(yield func(any) bool) {
					for _, element := range elements {
						if !yield(element) {
							return
						}
					}
				}
				Expect(it).Should(ContainElement(BeNumerically(">=", 3)))

				it2 := func(yield func(int, any) bool) {
					for idx, element := range elements {
						if !yield(idx, element) {
							return
						}
					}
				}
				Expect(it2).Should(ContainElement(BeNumerically(">=", 3)))
			})

			It("should fail if the matcher fails", func() {
				elements := []interface{}{1, 2, "3", "4"}
				it := func(yield func(any) bool) {
					for _, element := range elements {
						if !yield(element) {
							return
						}
					}
				}
				success, err := (&ContainElementMatcher{Element: BeNumerically(">=", 3)}).Match(it)
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())

				it2 := func(yield func(int, any) bool) {
					for idx, element := range elements {
						if !yield(idx, element) {
							return
						}
					}
				}
				success, err = (&ContainElementMatcher{Element: BeNumerically(">=", 3)}).Match(it2)
				Expect(success).Should(BeFalse())
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	Describe("returning findings", func() {
		Context("with match(es)", func() {
			It("should assing a single finding to a scalar result reference", func() {
				var stash string
				Expect(elementIter).To(ContainElement("bar", &stash))
				Expect(stash).To(Equal("bar"))

				Expect(elementIter2).To(ContainElement("baz", &stash))
				Expect(stash).To(Equal("baz"))
			})

			It("should assign a single finding to a slice return reference", func() {
				var stash []string
				Expect(elementIter).To(ContainElement("bar", &stash))
				Expect(stash).To(HaveLen(1))
				Expect(stash).To(ContainElement("bar"))

				Expect(elementIter2).To(ContainElement("baz", &stash))
				Expect(stash).To(HaveLen(1))
				Expect(stash).To(ContainElement("baz"))
			})
		})
	})

})
