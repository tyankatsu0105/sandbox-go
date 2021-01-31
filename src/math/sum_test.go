package math_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tyankatsu0105/sandbox-go/src/math"
)

var _ = Describe("math", func() {
	Describe("Sum", func() {
		It("合計値を返す", func() {
			result := math.Sum(1, 2)
			Expect(result).To(Equal(3))
		})

	})
})
