package routes

import (
	"Gin/internal/books"
	"Gin/internal/notebooks"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
)

type Usuario struct{
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

// var usuarios []Usuario

func SetupRoutes(r *gin.Engine){

	r.Static("/static","./static")

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/",func(c *gin.Context) {
		//c.String(200,"Hola :D")
        //c.String(http.StatusOK,"Hola Mundo")
		c.HTML(http.StatusOK,"index.html", nil)
	})

	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")
		if !strings.HasSuffix(page, ".html"){
			page += ".html"
		}
		if _, err := os.Stat("templates/" + page); err == nil{
			c.HTML(http.StatusOK, page, nil)
		}else{
			c.HTML(http.StatusNotFound, "404.html", nil)
		}
	})
	
	// LIBROS / Books
	bookRepo := books.NewInMemoryBookRepository()
	bookService := books.NewBookService(bookRepo)
	// Listado ALL
	r.GET("/books/list", books.BooksListHandler(bookService))
	// Listado Individual
	r.GET("books/list/:id", books.BoksListByIDHandler(bookService))
	//  Crear
	r.GET("/books/new", books.BooksFormHandler(bookService))
	// Guardar
	r.POST("/books/store",books.BooksStoreHandler(bookService))
	//  Editar
	r.GET(  "/books/update/:id",books.BooksFormHandler(bookService))
	r.POST( "/books/update/:id",books.BooksUpdateHandler(bookService))
	// Borrar
	r.GET("/books/delete/:id", books.BooksDeleteHandler(bookService))

	// Cuadernos / NoteBooks
	notebookRepo := notebooks.NewInMemoryNoteBookRepository()
	notebookService := notebooks.NewNoteBookService(notebookRepo)
	// Listado ALL
	r.GET("/notebooks/list", notebooks.NoteBooksListHandler(notebookService))
	// Listado Individual
	r.GET("notebooks/list/:id", notebooks.NoteBooksListByIDHandler(notebookService))
	//  Crear
	r.GET("/notebooks/new", notebooks.NoteBooksFormHandler(notebookService))
	// Guardar
	r.POST("/notebooks/store",notebooks.NoteBooksStoreHandler(notebookService))
	//  Editar
	r.GET(  "/notebooks/update/:id",notebooks.NoteBooksFormHandler(notebookService))
	r.POST( "/notebooks/update/:id",notebooks.NoteBooksUpdateHandler(notebookService))
	// Borrar
	r.GET("/notebooks/delete/:id", notebooks.NoteBooksDeleteHandler(notebookService))
}