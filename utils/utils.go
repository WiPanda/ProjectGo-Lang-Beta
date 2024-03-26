package utils

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"Project3/config"
	"Project3/model"

	"gorm.io/gorm/clause"
)

func AddBook(book model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := config.Mysql.DB.WithContext(ctx).Create(&book).Error
	if err != nil {
		return fmt.Errorf("Gagal menambahkan buku: %w", err)

	}

	return nil
}

// UpdateBook memperbarui buku yang sudah ada di database
func UpdateBook(id uint, book model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the book by ID first
	existingBook, err := GetBookByID(id)
	if err != nil {
		return fmt.Errorf("failed to get book by ID: %w", err)
	}

	// Update the fields of the existing book
	existingBook.ISBN = book.ISBN
	existingBook.Penulis = book.Penulis
	existingBook.Tahun = book.Tahun
	existingBook.Judul = book.Judul
	existingBook.Gambar = book.Gambar
	existingBook.Stok = book.Stok

	// Save the updated book
	err = config.Mysql.DB.WithContext(ctx).Save(&existingBook).Error

	// err := config.Mysql.DB.WithContext(ctx).
	// 	Model(&model.Book{}).
	// 	Where("id = ?", id).
	// 	Updates(book).
	// 	Error
	if err != nil {
		return fmt.Errorf("gagal mengubah data buku : %w", err)
	}

	return nil
}

// GetBookByID
func GetBookByID(id uint) (model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book model.Book
	err := config.Mysql.DB.WithContext(ctx).First(&book, id).Error
	if err != nil {
		return model.Book{}, fmt.Errorf("failed to get book by ID: %w", err)
	}

	return book, nil
}

// ListBooks mengambil data buku dari database.
func ListBooks() ([]model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var books []model.Book
	err := config.Mysql.DB.WithContext(ctx).Find(&books).Error
	if err != nil {
		return nil, fmt.Errorf("Gagal untuk mendapat Daftar Buku: %w", err)
	}

	return books, nil
}

// DeleteBook, soft-deletes buku dari database.
func DeleteBook(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := config.Mysql.DB.WithContext(ctx).Delete(&model.Book{}, id).Error
	if err != nil {
		return fmt.Errorf("Gagal untuk menghapus buku: %w", err)
	}

	return nil
}

// ImportCSV, memasukan data bukku fari CSV file.
func ImportCSV(csvPath string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	csvFile, err := os.Open(csvPath)
	if err != nil {
		return fmt.Errorf("Gagal untuk membuka file CSV: %w", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // mengizinkan jumlah bidang yang bervariasi

	// Melewati Row HEADER
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("Gagal untuk Membaca ROW HEADER: %w", err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Gagal untuk membaca ROW CSV: %w", err)
		}

		tahun, err := strconv.Atoi(row[2])
		if err != nil {
			return fmt.Errorf("Tahun tidak Valid : %w", err)
		}

		stok, err := strconv.Atoi(row[5])
		if err != nil {
			return fmt.Errorf("Stok tidak valid : %w", err)
		}

		book := model.Book{
			ISBN:    row[0],
			Penulis: row[1],
			Tahun:   uint(tahun),
			Judul:   row[3],
			Gambar:  row[4],
			Stok:    uint(stok),
		}

		err = config.Mysql.DB.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "isbn"}},
			DoUpdates: clause.AssignmentColumns([]string{"penulis", "tahun", "judul", "gambar", "stok"}),
		}).Create(&book).Error
		if err != nil {
			return fmt.Errorf("Gagal untuk memasukan Buku: %w", err)
		}
	}

	return nil
}
