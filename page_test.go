package onePage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var _ = Describe("Page", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	It("should show page view", func() {
		By("redirecting the user to the page view", func() {
			Expect(page.Navigate("http://localhost:7000")).To(Succeed())
			Expect(page.Title()).To(Equal("One Pics Note"))
		})
	})
})
