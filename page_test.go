package onePage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"net/http/httptest"
	"onePage"
)

var _ = Describe("Page", func() {
	var page *agouti.Page
	var rootURL string

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		onePageNoteServer := &onePage.OnePageNoteServer{}

		server := httptest.NewServer(onePageNoteServer)
		rootURL = server.URL
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	It("should show page view", func() {
		By("redirecting the user to the page view", func() {
			Expect(page.Navigate(rootURL)).To(Succeed())
			Expect(page.Title()).To(Equal("One Page Note"))
			text, _ := page.Find("#noteTitle").Text()
			Expect(text).To(Equal("Untitled"))
		})
	})
})
