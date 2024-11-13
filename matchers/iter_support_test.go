package matchers_test

var universalElements = []string{"foo", "bar", "baz"}

func universalIter(yield func(string) bool) {
	for _, element := range universalElements {
		if !yield(element) {
			return
		}
	}
}

func universalIter2(yield func(int, string) bool) {
	for idx, element := range universalElements {
		if !yield(idx, element) {
			return
		}
	}
}

func emptyIter(yield func(string) bool) {}

func emptyIter2(yield func(int, string) bool) {}
