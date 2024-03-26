package controllers

import (
	"Project3/model"
	"Project3/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	var book model.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := utils.AddBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk menambahkan buku",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Buku berhasil di tambahkan",
		"book":    book,
	})
}

// UpdateBook menangani permintaan PUT untuk memperbarui buku yang sudah ada.
func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID Buku tidak benar",
		})
	}

	var book model.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Body request tidak benar",
		})
	}

	if err := utils.UpdateBook(uint(id), book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk mengupdate book",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book update Berhasil",
		"book":    book,
	})
}

// ListBooks menangani permintaan GET untuk mengambil daftar buku.
func ListBooks(c *fiber.Ctx) error {
	books, err := utils.ListBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk menglist buku",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Buku berhasil di dapat",
		"books":   books,
	})
}

// DeleteBook menangani permintaan untuk soft-delete.
func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID Buku tidak valid",
		})
	}

	if err := utils.DeleteBook(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus buku",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Buku berhasil di hapus",
	})
}

// ImportBooks menangani permintaan POST untuk mengimpor buku dari file CSV.
func ImportBooks(c *fiber.Ctx) error {
	csvPath := c.FormValue("csv_path")
	if csvPath == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "csv_path parameter tidak ada / hilang",
		})
	}

	if err := utils.ImportCSV(csvPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal untuk mengimport buku",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Buku berhasil di import",
	})
}
