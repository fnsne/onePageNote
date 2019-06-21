package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"net/http/httptest"
)

var _ = Describe("NoteList", func() {
	var page *agouti.Page
	var store *InMemoryStore
	var rootURL string

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		store = NewInMemoryStore()
		onePageNoteServer := NewOnePageNoteServer(store)

		server := httptest.NewServer(onePageNoteServer)
		rootURL = server.URL
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})
	FDescribe("note list", func() {
		It("should show all existed notes", func() {
			store.notes = map[int]Note{1: {Title: "note1"}, 2: {Title: "note2"}, 3: {Title: "note3"}}
			Expect(page.Navigate(rootURL + "/note/")).To(Succeed())
			count, err := page.Find("#noteList").All(".note").Count()
			Expect(err).To(Succeed())
			Expect(count).To(Equal(3))
		})
	})
})
