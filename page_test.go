package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"net/http/httptest"
	"time"
)

var _ = Describe("Page", func() {
	var page *agouti.Page
	var rootURL string
	var store *InMemoryStore

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		store = &InMemoryStore{}
		onePageNoteServer := NewOnePageNoteServer(store)

		server := httptest.NewServer(onePageNoteServer)
		rootURL = server.URL
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	Describe("note view", func() {
		Context("Get home page", func() {
			//It("should have web title", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	Expect(page.Title()).To(Equal("One Page Note"))
			//})
			//It("should have default note title", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	text, _ := page.Find("#noteTitle").Text()
			//	Expect(text).To(Equal("Untitled"))
			//})
			//It("should have default Date", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	expectDate := time.Now().Format("2006-01-02")
			//	text, _ := page.Find("#noteDate").Text()
			//	Expect(text).To(Equal(expectDate))
			//})
			//It("can change note title by clicking and input", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteTitle := page.Find("#noteTitle")
			//	Expect(noteTitle.Click()).To(Succeed())
			//	Expect(noteTitle.Fill("我的note")).To(Succeed())
			//	noteTitleString, err := noteTitle.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteTitleString).To(Equal("我的note"))
			//})
			//It("can change note Date by clicking and input", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteDate := page.Find("#noteDate")
			//	Expect(noteDate.Click()).To(Succeed())
			//	Expect(noteDate.Fill("2018-01-01")).To(Succeed())
			//	noteTitleString, err := noteDate.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteTitleString).To(Equal("2018-01-01"))
			//
			//})
			//It("should change grids number", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	By("clicking button 16", func() {
			//		Expect(page.Find("#button16").Click()).To(Succeed())
			//		count, err := page.All(".baseGrid").Count()
			//		Expect(err).To(Succeed())
			//		Expect(count).To(Equal(15))
			//	})
			//	By("clicking button 32", func() {
			//		Expect(page.Find("#button32").Click()).To(Succeed())
			//		count, err := page.All(".baseGrid").Count()
			//		Expect(err).To(Succeed())
			//		Expect(count).To(Equal(31))
			//	})
			//	By("clicking button 64", func() {
			//		Expect(page.Find("#button64").Click()).To(Succeed())
			//		count, err := page.All(".baseGrid").Count()
			//		Expect(err).To(Succeed())
			//		Expect(count).To(Equal(63))
			//	})
			//	By("clicking button 8", func() {
			//		Expect(page.Find("#button8").Click()).To(Succeed())
			//		count, err := page.All(".baseGrid").Count()
			//		Expect(err).To(Succeed())
			//		Expect(count).To(Equal(7))
			//	})
			//})
			//
			//It("will remember stored noteDate", func() {
			//	date, _ := time.Parse("2006-01-02", "2018-02-02")
			//	store.note = Note{&date}
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	Expect(page.Find("#testBtn").Click()).To(Succeed())
			//	noteDate := page.Find("#noteDate")
			//	noteDateString, err := noteDate.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteDateString).To(Equal("2018-02-02"))
			//})
			//It("will remember last edited noteDate", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteDate := page.Find("#noteDate")
			//	Expect(noteDate.Click()).To(Succeed())
			//	Expect(noteDate.Fill("2018-02-02")).To(Succeed())
			//
			//	noteDateString, err := noteDate.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteDateString).To(Equal("2018-02-02"))
			//
			//	time.Sleep(2 * time.Second)
			//
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteDateString, err = noteDate.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteDateString).To(Equal("2018-02-02"))
			//})
			//It("will remember last edited noteTitle", func() {
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteTitle := page.Find("#noteTitle")
			//	Expect(noteTitle.Click()).To(Succeed())
			//	Expect(noteTitle.Fill("筆記主題")).To(Succeed())
			//
			//	noteTitleString, err := noteTitle.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteTitleString).To(Equal("筆記主題"))
			//
			//	time.Sleep(2 * time.Second)
			//
			//	Expect(page.Navigate(rootURL)).To(Succeed())
			//	noteTitleString, err = noteTitle.Text()
			//	Expect(err).To(Succeed())
			//	Expect(noteTitleString).To(Equal("筆記主題"))
			//})
			It("will remember input keywords and comments", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				keywords := page.All(".keyword")
				comments := page.All(".comment")
				Expect(keywords.At(0).Click()).To(Succeed())
				Expect(keywords.At(0).Fill("關鍵字1")).To(Succeed())
				Expect(comments.At(0).Click()).To(Succeed())
				Expect(comments.At(0).Fill("評論1")).To(Succeed())

				Expect(keywords.At(1).Click()).To(Succeed())
				Expect(keywords.At(1).Fill("關鍵字2")).To(Succeed())
				Expect(comments.At(1).Click()).To(Succeed())
				Expect(comments.At(1).Fill("評論2")).To(Succeed())

				Expect(keywords.At(2).Click()).To(Succeed())
				Expect(keywords.At(2).Fill("關鍵字3")).To(Succeed())
				Expect(comments.At(2).Click()).To(Succeed())
				Expect(comments.At(2).Fill("評論3")).To(Succeed())

				keyword1, err := keywords.At(0).Text()
				Expect(err).To(Succeed())
				Expect(keyword1).To(Equal("關鍵字1"))

				comment1, err := comments.At(0).Text()
				Expect(err).To(Succeed())
				Expect(comment1).To(Equal("評論1"))

				keyword2, err := keywords.At(1).Text()
				Expect(err).To(Succeed())
				Expect(keyword2).To(Equal("關鍵字2"))

				comment2, err := comments.At(1).Text()
				Expect(err).To(Succeed())
				Expect(comment2).To(Equal("評論2"))

				keyword3, err := keywords.At(2).Text()
				Expect(err).To(Succeed())
				Expect(keyword3).To(Equal("關鍵字3"))

				comment3, err := comments.At(2).Text()
				Expect(err).To(Succeed())
				Expect(comment3).To(Equal("評論3"))

				time.Sleep(2 * time.Second)

				keyword1, err = keywords.At(0).Text()
				Expect(err).To(Succeed())
				Expect(keyword1).To(Equal("關鍵字1"))

				comment1, err = comments.At(0).Text()
				Expect(err).To(Succeed())
				Expect(comment1).To(Equal("評論1"))

				keyword2, err = keywords.At(1).Text()
				Expect(err).To(Succeed())
				Expect(keyword2).To(Equal("關鍵字2"))

				comment2, err = comments.At(1).Text()
				Expect(err).To(Succeed())
				Expect(comment2).To(Equal("評論2"))

				keyword3, err = keywords.At(2).Text()
				Expect(err).To(Succeed())
				Expect(keyword3).To(Equal("關鍵字3"))

				keyword3, err = keywords.At(2).Text()
				Expect(err).To(Succeed())
				Expect(keyword3).To(Equal("關鍵字3"))

				comment3, err = comments.At(2).Text()
				Expect(err).To(Succeed())
				Expect(comment3).To(Equal("評論3"))
			})
		})

	})
})

func shouldEqual(page *agouti.Page, selector, expect string) {
}
