package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"net/http/httptest"
	"testing"
	"time"
)

var _ = Describe("Note", func() {
	var page *agouti.Page
	var rootURL string
	var store *FileSystemStore
	var server *httptest.Server
	var cleanDatabase func()

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		//store = NewInMemoryStore()
		//database, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
		var t *testing.T
		database, clean := createTempFile(t, []byte("[]"))
		cleanDatabase = clean
		store = NewFileSystemStore(database)
		onePageNoteServer := NewOnePageNoteServer(store)

		server = httptest.NewServer(onePageNoteServer)
		rootURL = server.URL
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
		cleanDatabase()
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
			It("should have default Date", func() {
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
			It("can change note Date by clicking and input", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteDate := page.Find("#noteDate")
				Expect(noteDate.Click()).To(Succeed())
				Expect(noteDate.Fill("2018-01-01")).To(Succeed())
				noteTitleString, err := noteDate.Text()
				Expect(err).To(Succeed())
				Expect(noteTitleString).To(Equal("2018-01-01"))

			})
			It("should change grids number", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				By("clicking button 16", func() {
					Expect(page.Find("#button16").Click()).To(Succeed())
					count, err := page.All(".baseGrid").Count()
					Expect(err).To(Succeed())
					Expect(count).To(Equal(15))
				})
				By("clicking button 32", func() {
					Expect(page.Find("#button32").Click()).To(Succeed())
					count, err := page.All(".baseGrid").Count()
					Expect(err).To(Succeed())
					Expect(count).To(Equal(31))
				})
				By("clicking button 64", func() {
					Expect(page.Find("#button64").Click()).To(Succeed())
					count, err := page.All(".baseGrid").Count()
					Expect(err).To(Succeed())
					Expect(count).To(Equal(63))
				})
				By("clicking button 8", func() {
					Expect(page.Find("#button8").Click()).To(Succeed())
					count, err := page.All(".baseGrid").Count()
					Expect(err).To(Succeed())
					Expect(count).To(Equal(7))
				})
			})

			It("will remember stored noteDate", func() {
				date, _ := time.Parse("2006-01-02", "2018-02-02")
				store.notes = []Note{{Id: 1, Date: &date}}
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteDate := page.Find("#noteDate")
				noteDateString, err := noteDate.Text()
				Expect(err).To(Succeed())
				Expect(noteDateString).To(Equal("2018-02-02"))
			})
			It("will remember last edited noteDate", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteDate := page.Find("#noteDate")
				Expect(noteDate.Click()).To(Succeed())
				Expect(noteDate.Fill("2018-02-02")).To(Succeed())

				noteDateString, err := noteDate.Text()
				Expect(err).To(Succeed())
				Expect(noteDateString).To(Equal("2018-02-02"))

				time.Sleep(2 * time.Second)

				Expect(page.Navigate(rootURL)).To(Succeed())
				noteDateString, err = noteDate.Text()
				Expect(err).To(Succeed())
				Expect(noteDateString).To(Equal("2018-02-02"))
			})
			It("will remember last edited noteTitle", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteTitle := page.Find("#noteTitle")
				Expect(noteTitle.Click()).To(Succeed())
				Expect(noteTitle.Fill("筆記主題")).To(Succeed())

				noteTitleString, err := noteTitle.Text()
				Expect(err).To(Succeed())
				Expect(noteTitleString).To(Equal("筆記主題"))

				time.Sleep(2 * time.Second)

				Expect(page.Navigate(rootURL)).To(Succeed())
				noteTitleString, err = noteTitle.Text()
				Expect(err).To(Succeed())
				Expect(noteTitleString).To(Equal("筆記主題"))
			})
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

				Expect(page.Navigate(rootURL)).To(Succeed())
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
			It("could create note", func() {
				Expect(page.Navigate(rootURL)).To(Succeed())
				By("click new Note button", func() {
					Expect(page.Find("#noteTitle").Fill("note 1"))
					time.Sleep(1 * time.Second)
					time.Sleep(3 * time.Second)
					Expect(page.Find("#newNoteBtn").Click()).To(Succeed())
					time.Sleep(3 * time.Second)
					noteItem1 := page.All(".noteItem").At(0)
					text, err := noteItem1.Text()
					Expect(err).To(Succeed())
					Expect(text).To(Equal("Untitled"))
					noteTitle, err := page.Find("#noteTitle").Text()
					Expect(err).To(Succeed())
					Expect(noteTitle).To(Equal("Untitled"))
				})
			})
			It("could change note ", func() {
				By("click on note list", func() {
					Expect(page.Navigate(rootURL)).To(Succeed())
					Expect(page.Find("#noteTitle").Fill("note1"))
					Expect(page.Find("#noteDate").Fill("2011-01-01"))
					keywords := page.All(".keyword")
					Expect(keywords.At(0).Fill("keyword1"))
					Expect(keywords.At(1).Fill("keyword2"))
					Expect(keywords.At(2).Fill("keyword3"))
					Expect(keywords.At(3).Fill("keyword4"))
					Expect(keywords.At(4).Fill("keyword5"))
					Expect(keywords.At(5).Fill("keyword6"))

					time.Sleep(2 * time.Second)

					Expect(page.Find("#newNoteBtn").Click()).To(Succeed())
					noteItems := page.All(".noteItem")
					text, err := noteItems.At(0).Text()
					Expect(err).To(Succeed())
					Expect(text).To(Equal("Untitled"))

					time.Sleep(2 * time.Second)
					Expect(noteItems.At(1).Click()).To(Succeed())
					time.Sleep(2 * time.Second)
					keyword1, err := keywords.At(0).Text()
					Expect(err).To(Succeed())
					Expect(keyword1).To(Equal("keyword1"))
					keyword2, err := keywords.At(1).Text()
					Expect(err).To(Succeed())
					Expect(keyword2).To(Equal("keyword2"))
					keyword3, err := keywords.At(2).Text()
					Expect(err).To(Succeed())
					Expect(keyword3).To(Equal("keyword3"))
				})
			})
		})
		Context("note list", func() {
			It("should show note title.", func() {
				store.notes = []Note{{Title: "title1"}}
				Expect(page.Navigate(rootURL)).To(Succeed())
				noteItems := page.All(".noteItem")
				note1Title, err := noteItems.At(0).Text()
				Expect(err).To(Succeed())
				Expect(note1Title).To(Equal("title1"))
			})
			It("should show the latest note title", func() {
				store.notes = []Note{{Id: 1, Title: "title1"}}
				Expect(page.Navigate(rootURL)).To(Succeed())
				title := page.Find("#noteTitle")
				Expect(title.Click()).To(Succeed())
				Expect(title.Fill("my title2")).To(Succeed())

				time.Sleep(5 * time.Second)

				noteItems := page.All(".noteItem")
				note1Title, err := noteItems.At(0).Text()
				Expect(err).To(Succeed())
				Expect(note1Title).To(Equal("my title2"))
			})
			It("should remember edited notes title", func() {
				store.notes = []Note{{Id: 1, Title: "title1"}}
				Expect(page.Navigate(rootURL)).To(Succeed())
				title := page.Find("#noteTitle")
				Expect(title.Click()).To(Succeed())
				Expect(title.Fill("my title2")).To(Succeed())

				time.Sleep(3 * time.Second)

				noteItems := page.All(".noteItem")
				note1Title, err := noteItems.At(0).Text()
				Expect(err).To(Succeed())
				Expect(note1Title).To(Equal("my title2"))
			})
		})
		Context("others", func() {

			//Test_Server_after_open_will_focus_on_latest_created_note_and_can_continue_edit_it
			//func Test_Server_after_open_will_focus_on_last_edited_note_and_can_continue_edit_it(t *testing.T) {
		})
	})
})
