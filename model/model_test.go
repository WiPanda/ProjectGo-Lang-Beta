package model_test

import (
	"Project3/config"
	"Project3/model"
	"Project3/utils"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestAddBook(t *testing.T) {
	Init()

	book := model.Book{
		ISBN:    "123-0987654321",
		Penulis: "Martin Tin Martin",
		Tahun:   2020,
		Judul:   "Coding, ga error ga ganteng",
		Gambar:  "https://contoh.com/gambar.jpg",
		Stok:    10,
	}

	err := utils.AddBook(book)
	assert.Nil(t, err)

	err = utils.AddBook(book)
	assert.NotNil(t, err)
}

func TestUpdateBook(t *testing.T) {
	Init()

	book := model.Book{
		ISBN:    "321-0987654321",
		Penulis: "Update Martin Tin Martin",
		Tahun:   2025,
		Judul:   "Oprek, ga bootloop ga ganteng",
		Gambar:  "https://contoh.com/oprek_gambar.jpg",
		Stok:    20,
	}

	err := utils.UpdateBook(1, book)
	assert.Nil(t, err)

	err = utils.UpdateBook(999, book)
	assert.NotNil(t, err)
}

func TestListBooks(t *testing.T) {
	// Test List semua buku
	books, err := utils.ListBooks()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(books), 0)
}

func TestDeleteBook(t *testing.T) {
	err := utils.DeleteBook(1) // anggep bukunya id 1
	assert.Nil(t, err)

	err = utils.DeleteBook(999)
	assert.NotNil(t, err)
}
