package roman

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Roman numerals", func() {
	Context("1", func() {
		It("Returns I", func() {
			Expect(Numeral(1)).To(Equal("I"))
		})
	})
	Context("2", func() {
		It("Returns II", func() {
			Expect(Numeral(2)).To(Equal("II"))
		})
	})
	Context("3", func() {
		It("Returns III", func() {
			Expect(Numeral(3)).To(Equal("III"))
		})
	})
	Context("4", func() {
		It("Returns IV", func() {
			Expect(Numeral(4)).To(Equal("IV"))
		})
	})
	Context("5", func() {
		It("Returns V", func() {
			Expect(Numeral(5)).To(Equal("V"))
		})
	})
	Context("6", func() {
		It("Returns VI", func() {
			Expect(Numeral(6)).To(Equal("VI"))
		})
	})
	Context("8", func() {
		It("Returns VIII", func() {
			Expect(Numeral(8)).To(Equal("VIII"))
		})
	})
	Context("9", func() {
		It("Returns IX", func() {
			Expect(Numeral(9)).To(Equal("IX"))
		})
	})
	Context("10", func() {
		It("Returns X", func() {
			Expect(Numeral(10)).To(Equal("X"))
		})
	})
	Context("13", func() {
		It("Returns XIII", func() {
			Expect(Numeral(13)).To(Equal("XIII"))
		})
	})
	Context("14", func() {
		It("Returns XIV", func() {
			Expect(Numeral(14)).To(Equal("XIV"))
		})
	})
	Context("15", func() {
		It("Returns XV", func() {
			Expect(Numeral(15)).To(Equal("XV"))
		})
	})
	Context("18", func() {
		It("Returns XVIII", func() {
			Expect(Numeral(18)).To(Equal("XVIII"))
		})
	})
	Context("19", func() {
		It("Returns XIX", func() {
			Expect(Numeral(19)).To(Equal("XIX"))
		})
	})
	Context("20", func() {
		It("Returns XX", func() {
			Expect(Numeral(20)).To(Equal("XX"))
		})
	})

	Context("21", func() {
		It("Returns XXI", func() {
			Expect(Numeral(21)).To(Equal("XXI"))
		})
	})

	Context("23", func() {
		It("Returns XXIII", func() {
			Expect(Numeral(23)).To(Equal("XXIII"))
		})
	})

	Context("24", func() {
		It("Returns XXIV", func() {
			Expect(Numeral(24)).To(Equal("XXIV"))
		})
	})

	Context("30", func() {
		It("Returns XXX", func() {
			Expect(Numeral(30)).To(Equal("XXX"))
		})
	})

	Context("35", func() {
		It("Returns XXXV", func() {
			Expect(Numeral(35)).To(Equal("XXXV"))
		})
	})

	Context("40", func() {
		It("Returns XL", func() {
			Expect(Numeral(40)).To(Equal("XL"))
		})
	})

	Context("44", func() {
		It("Returns XLIV", func() {
			Expect(Numeral(44)).To(Equal("XLIV"))
		})
	})

	Context("49", func() {
		It("Returns XLIX", func() {
			Expect(Numeral(49)).To(Equal("XLIX"))
		})
	})

	Context("50", func() {
		It("Returns L", func() {
			Expect(Numeral(50)).To(Equal("L"))
		})
	})

	Context("51", func() {
		It("Returns LI", func() {
			Expect(Numeral(51)).To(Equal("LI"))
		})
	})

	Context("70", func() {
		It("Returns LXX", func() {
			Expect(Numeral(70)).To(Equal("LXX"))
		})
	})

	Context("100", func() {
		It("Returns C", func() {
			Expect(Numeral(100)).To(Equal("C"))
		})
	})

	Context("90", func() {
		It("Returns XC", func() {
			Expect(Numeral(90)).To(Equal("XC"))
		})
	})

	Context("99", func() {
		It("Returns XCIX", func() {
			Expect(Numeral(99)).To(Equal("XCIX"))
		})
	})

	Context("101", func() {
		It("Returns CI", func() {
			Expect(Numeral(101)).To(Equal("CI"))
		})
	})

	Context("300", func() {
		It("Returns CCC", func() {
			Expect(Numeral(300)).To(Equal("CCC"))
		})
	})

	Context("500", func() {
		It("Returns D", func() {
			Expect(Numeral(500)).To(Equal("D"))
		})
	})

	Context("400", func() {
		It("Returns CD", func() {
			Expect(Numeral(400)).To(Equal("CD"))
		})
	})

	Context("441", func() {
		It("Returns CDXL1", func() {
			Expect(Numeral(441)).To(Equal("CDXLI"))
		})
	})

	Context("1974", func() {
		It("Returns MCMLXXIV", func() {
			Expect(Numeral(1974)).To(Equal("MCMLXXIV"))
		})
	})

	Context("102497", func() {
		It("Returns C2MMCDXCVII", func() {
			Expect(Numeral(102497)).To(Equal("C2MMCDXCVII"))
		})
	})
})

