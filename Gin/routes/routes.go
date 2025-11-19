package routes

import (
	"Gin/internal/books"
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
	// CONFIG API (JSON)
	repo := books.NewInMemoryBookRepository()
	service := books.NewBookService(repo)
	controller := books.NewBookControllers(service)

	// API: /api/books
	api := r.Group("/api")
	{
		controller.RegisterRoutes(api)
	}
	// LIBROS
	// Listado ALL
	r.GET("/books/list", books.BooksListHandler(service))
	// Listado Individual
	r.GET("books/list/:id", books.BoksListByIDHandler(service))
	//  Crear
	r.GET("/books/new", books.BooksFormHandler(service))
	// Guardar
	r.POST("/books/store",books.BooksStoreHandler(service))
	//  Editar
	r.GET(  "/books/update/:id",books.BooksFormHandler(service))
	r.POST( "/books/update/:id",books.BooksUpdateHandler(service))
	// Borrar
	r.GET("/books/delete/:id", books.BooksDeleteHandler(service))
}