package books

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
// LIST
func BooksListHandler(s*BookService) gin.HandlerFunc{
	return func(c *gin.Context) {
		list, _ := s.GetAll()
		c.HTML(http.StatusOK, "books_list.html", gin.H{
			"Books": list,
		})
	}
}

// Listar ID
func BoksListByIDHandler(s*BookService) gin.HandlerFunc{
	return  func (c *gin.Context)  {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		book, err := s.GetByID(id)
		if err != nil{
			c.String(http.StatusNotFound, "book not found")
			return 
		}
		// Obtener todos para seguir mostrando la tabla completa
		allBooks, _ := s.GetAll()
		c.HTML(http.StatusOK, "books_list.html", gin.H{
			"Books": allBooks,
			"Book": book,
		})
	}
}
// FORM ( CREAR + EDITAR)
func BooksFormHandler(s *BookService) gin.HandlerFunc{
	return  func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr != ""{
			id, err := strconv.Atoi(idStr)
			if err != nil{
				c.String(http.StatusBadRequest, "invalid id")
				return 
			}
			book, err := s.GetByID(id)
			if err != nil{
				c.String(http.StatusNotFound, "book not found")
				return 
			}
			c.HTML(http.StatusOK, "books_list.html", gin.H{
				"Book": book,
			})
			return 
		}
		// Crear LIBRO
		c.HTML(http.StatusOK, "books_list.html", gin.H{
			"Book": nil,
		})
	}
}

// STORE (guardar nuevo)
func BooksStoreHandler(s *BookService) gin.HandlerFunc  {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		author := c.PostForm("author")
		priceStr := c.PostForm("price")
		stockStr := c.PostForm("stock")
		description := c.PostForm("description")
		
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil{
			c.String(http.StatusBadRequest, "price must be a number")
			return
		}

		stock, err := strconv.Atoi(stockStr)
		if err != nil{
			c.String(http.StatusBadRequest, "stock must be a number")
			return
		}
		book :=Book{
			Title: title,
			Author: author,
			Price: price,
			Stock: stock,
			Description: description,
		}
		if err := s.CreateBook(book); err != nil{
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/books/list")
	}
}

// UPDATE
	func BooksUpdateHandler(s*BookService) gin.HandlerFunc{
		return func(c *gin.Context) {
			idStr := c.Param("id")
			title := c.PostForm("title")
			author := c.PostForm("author")
			priceStr := c.PostForm("price")
			stockStr := c.PostForm("stock")
			description := c.PostForm("description")

			id, err := strconv.Atoi(idStr)
			if err != nil{
				c.String(http.StatusBadRequest, "invalid id")
				return
			}

			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil{
				c.String(http.StatusBadRequest, "price must be a number")
				return
			}

			stock, err := strconv.Atoi(stockStr)
			if err != nil{
				c.String(http.StatusBadRequest, "stock must be a number")
				return
			}
			book :=Book{
				ID: id,
				Title: title,
				Author: author,
				Price: price,
				Stock: stock,
				Description: description,
			}
			if err := s.UpdateBook(book); err != nil{
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.Redirect(http.StatusSeeOther, "/books/list")
		}
	}
// DELETE
func BooksDeleteHandler(s *BookService) gin.HandlerFunc{
	return func (c*gin.Context)  {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		if err := s.DeleteBook(id); err != nil{
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/books/list")
	}
}
