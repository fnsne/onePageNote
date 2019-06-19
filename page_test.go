package main

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe("Page", func() {
	var page *agouti.Page
	var rootURL string

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		err = os.Remove("db.notes.json")
		if err != nil {
			fmt.Printf("err= %v", err)
		}

		onePageNoteServer := &OnePageNoteServer{}

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
			It("should have default date", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				expectDate := time.Now().Format("2006-01-02")
				text, _ := page.Find("#noteDate").Text()
				Expect(text).To(Equal(expectDate))
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
			It("can change note date by clicking and input", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteDate := page.Find("#noteDate")
				Expect(noteDate.Click()).To(Succeed())
				Expect(noteDate.Fill("2018-01-01")).To(Succeed())
				noteTitleString, err := noteDate.Text()
				Expect(err).To(Succeed())
				Expect(noteTitleString).To(Equal("2018-01-01"))

			})
			//It("will remember last edited noteDate", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteDate := page.Find("%noteDate")
			//})
		})

	})
})
