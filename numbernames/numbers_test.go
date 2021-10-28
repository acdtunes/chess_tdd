package numbernames
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Number names", func() {
	Context("0 to 20", func() {
		It("Returns zero to twenty", func() {
			Expect(Name(0)).To(Equal("zero"))
			Expect(Name(1)).To(Equal("one"))
			Expect(Name(2)).To(Equal("two"))
			Expect(Name(3)).To(Equal("three"))
			Expect(Name(4)).To(Equal("four"))
			Expect(Name(5)).To(Equal("five"))
			Expect(Name(6)).To(Equal("six"))
			Expect(Name(7)).To(Equal("seven"))
			Expect(Name(8)).To(Equal("eight"))
			Expect(Name(9)).To(Equal("nine"))
			Expect(Name(10)).To(Equal("ten"))
			Expect(Name(11)).To(Equal("eleven"))
			Expect(Name(12)).To(Equal("twelve"))
			Expect(Name(13)).To(Equal("thirteen"))
			Expect(Name(14)).To(Equal("fourteen"))
			Expect(Name(15)).To(Equal("fifteen"))
			Expect(Name(16)).To(Equal("sixteen"))
			Expect(Name(17)).To(Equal("seventeen"))
			Expect(Name(18)).To(Equal("eighteen"))
			Expect(Name(19)).To(Equal("nineteen"))
			Expect(Name(20)).To(Equal("twenty"))
		})
	})

	Context("21 to 30", func() {
		It("Returns twenty one to thirty", func() {
			Expect(Name(21)).To(Equal("twenty one"))
			Expect(Name(22)).To(Equal("twenty two"))
			Expect(Name(23)).To(Equal("twenty three"))
			Expect(Name(24)).To(Equal("twenty four"))
			Expect(Name(25)).To(Equal("twenty five"))
			Expect(Name(26)).To(Equal("twenty six"))
			Expect(Name(27)).To(Equal("twenty seven"))
			Expect(Name(28)).To(Equal("twenty eight"))
			Expect(Name(29)).To(Equal("twenty nine"))
			Expect(Name(30)).To(Equal("thirty"))
		})
	})

	Context("37", func() {
		It("Returns thirty seven", func() {
			Expect(Name(37)).To(Equal("thirty seven"))
		})
	})

	Context("40, 50, 60, 70, 80, 90", func() {
		It("Returns forty, fifty, sixty, seventy, eighty, ninety", func() {
			Expect(Name(40)).To(Equal("forty"))
			Expect(Name(50)).To(Equal("fifty"))
			Expect(Name(60)).To(Equal("sixty"))
			Expect(Name(70)).To(Equal("seventy"))
			Expect(Name(80)).To(Equal("eighty"))
			Expect(Name(90)).To(Equal("ninety"))
		})
	})

	Context("79", func() {
		It("Returns seventy nine", func() {
			Expect(Name(79)).To(Equal("seventy nine"))
		})
	})

	Context("100", func() {
		It("Returns one hundred", func() {
			Expect(Name(100)).To(Equal("one hundred"))
		})
	})

	Context("110", func() {
		It("Returns one hundred and ten", func() {
			Expect(Name(110)).To(Equal("one hundred and ten"))
		})
	})

	Context("463", func() {
		It("Returns four hundred and sixty three", func() {
			Expect(Name(463)).To(Equal("four hundred and sixty three"))
		})
	})

	Context("1000", func() {
		It("Returns one thousand", func() {
			Expect(Name(1000)).To(Equal("one thousand"))
		})
	})

	Context("1200", func() {
		It("Returns one thousand, two hundred", func() {
			Expect(Name(1200)).To(Equal("one thousand, two hundred"))
		})
	})

	Context("9950", func() {
		It("Returns nine thousand, nine hundred and fifty", func() {
			Expect(Name(9950)).To(Equal("nine thousand, nine hundred and fifty"))
		})
	})

	Context("1000000", func() {
		It("Returns one million", func() {
			Expect(Name(1000000)).To(Equal("one million"))
		})
	})

	Context("2000000", func() {
		It("Returns two millions", func() {
			Expect(Name(2000000)).To(Equal("two millions"))
		})
	})

	Context("960248390", func() {
		FIt("Returns nine hundred and sixty millions, two hundred and forty eight thousand, three hundred and ninety", func() {
			Expect(Name(960248390)).To(Equal("nine hundred and sixty millions, two hundred and forty eight thousand, three hundred and ninety"))
		})
	})
})



