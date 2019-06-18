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
	Describe("note view", func() {
		Context("Get home page", func() {
			It("should have web title", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				Expect(page.Title()).To(Equal("One Page Note"))
			})
			It("should have default note title", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				text, _ := page.Find("#noteTitle").Text()
				Expect(text).To(Equal("Untitled"))
			})
			It("can change note title by clicking and input", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteTitle := page.Find("#noteTitle")
				Expect(noteTitle.Click()).To(Succeed())
				Expect(noteTitle.Fill("我的note")).To(Succeed())
				noteTitleString, err := noteTitle.Text()
				Expect(err).To(Succeed())
				Expect(noteTitleString).To(Equal("我的note"))
			})
		})

	})
})
