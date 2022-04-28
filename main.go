package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "demon slayer",
		Author: "koyoharu gotouge",
		Price:  2.99,
	},
}

func main() {
	server := gin.Default()
	server.GET("/books", getBooks)
	server.GET("/book/:id", getBookByID)
	server.POST("/books", addBook)
	server.PUT("/book/:id", updateBook)
	server.DELETE("/book/:id", deleteBookByID)
	server.Run(":8000")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	paramID := c.Param("id")
	for _, book := range books {
		//ทำการ loop หาหนังสือที่ตรงกับ id ที่ส่งมา
		if book.ID == paramID {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func addBook(c *gin.Context) {
	var newBook book
	// เรียก BindJSON เพื่อผูก JSON ที่รับมากับ newBook
	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
		return
	}
	// เพิ่มรายการหนังสือเล่มใหม่เข้าไปใน slice
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	var editBook book
	if err := c.BindJSON(&editBook); err != nil {
		return
	}
	paramID := c.Param("id")
	for i := 0; i <= len(books)-1; i++ {
		//loop หา id หนังสือตาม id และทำการแก้ไขรายการหนังสือนั่น
		if books[i].ID == paramID {
			books[i].Name = editBook.Name
			books[i].Author = editBook.Author
			books[i].Price = editBook.Price
			c.JSON(http.StatusOK, books[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func deleteBookByID(c *gin.Context) {
	paramID := c.Param("id")
	for i := 0; i <= len(books)-1; i++ {
		if books[i].ID == paramID {
			// ทำการลบรายการหนังสือจาก id ที่ส่งมา
			//โดยตัดข้อมูลแล้วเอามาต่อโดยไม่มีรายการหนังสือที่เราส่งมา
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, "delete success")
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}
