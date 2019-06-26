package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_filesystemStore_can_get_note(t *testing.T) {

	tempFile, clean := createTempFile(t, []byte(`[{"Id":1, "Title":"title1"}]`))

	defer clean()

	store := NewFileSystemStore(tempFile)
	note := store.GetNote(1)
	assert.Equal(t, 1, note.Id)
	assert.Equal(t, "title1", note.Title)
}

func createTempFile(t *testing.T, initialData []byte) (*os.File, func()) {
	tempFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("cannot create temp file %s %v", tempFile.Name(), err)
	}
	tempFile.Write(initialData)
	removeFile := func() {
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}
