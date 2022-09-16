package database

import (
	"day-4/config"
	"day-4/models"
)

func GetBooks() (*[]models.Book, error) {
	var books *[]models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(id int) (*models.Book, error) {
	var book *models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(book *models.Book) (*models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(id int, book *models.Book) (*models.Book, error) {

	if err := config.DB.Model(&book).Where("id = ?", id).Updates(models.Book{Title: book.Title, Page: book.Page, Author: book.Author}).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int, book *models.Book) (*models.Book, error) {

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}
