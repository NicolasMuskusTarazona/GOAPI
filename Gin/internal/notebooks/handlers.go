package notebooks

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Lista
func NoteBooksListHandler(s*NoteBookService) gin.HandlerFunc{
	return func(c *gin.Context) {
		list, _ := s.GetAll()
		c.HTML(http.StatusOK, "notebooks_list.html", gin.H{
			"NoteBooks": list,
		})
	}
}
// Listar por ID
func NoteBooksListByIDHandler(s*NoteBookService) gin.HandlerFunc{
	return  func (c *gin.Context)  {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		notebook, err := s.GetByID(id)
		if err != nil{
			c.String(http.StatusNotFound, "notebook not found")
			return 
		}
		// Obtener todos para seguir mostrando la tabla completa
		allNoteBooks, _ := s.GetAll()
		c.HTML(http.StatusOK, "notebooks_list.html", gin.H{
			"NoteBooks": allNoteBooks,
			"NoteBook": notebook,
		})
	}
}

// FORM ( CREAR + EDITAR)
func NoteBooksFormHandler(s *NoteBookService) gin.HandlerFunc{
	return  func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr != ""{
			id, err := strconv.Atoi(idStr)
			if err != nil{
				c.String(http.StatusBadRequest, "invalid id")
				return 
			}
			notebook, err := s.GetByID(id)
			if err != nil{
				c.String(http.StatusNotFound, "notebook not found")
				return 
			}
			c.HTML(http.StatusOK, "books_list.html", gin.H{
				"NoteBook": notebook,
			})
			return 
		}
		// Crear LIBRO
		c.HTML(http.StatusOK, "notebooks_list.html", gin.H{
			"NoteBook": nil,
		})
	}
}


// STORE (guardar nuevo)
func NoteBooksStoreHandler(s *NoteBookService) gin.HandlerFunc  {
	return func(c *gin.Context) {
		priceStr := c.PostForm("price")
		stockStr := c.PostForm("stock")
		
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
		notebook :=NoteBook{
			Price: price,
			Stock: stock,
		}
		if err := s.CreateNoteBook(notebook); err != nil{
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/notebooks/list")
	}
}
// UPDATE
func NoteBooksUpdateHandler(s*NoteBookService) gin.HandlerFunc{
	return func(c *gin.Context) {
		idStr := c.Param("id")
		priceStr := c.PostForm("price")
		stockStr := c.PostForm("stock")
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
		notebook :=NoteBook{
			ID: id,
			Price: price,
			Stock: stock,
		}
		if err := s.UpdateNoteBook(notebook); err != nil{
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/notebooks/list")
	}
}

// DELETE
func NoteBooksDeleteHandler(s *NoteBookService) gin.HandlerFunc{
	return func (c*gin.Context)  {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.String(http.StatusBadRequest, "invalid id")
			return
		}
		if err := s.DeleteNoteBook(id); err != nil{
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Redirect(http.StatusSeeOther, "/notebooks/list")
	}
}


