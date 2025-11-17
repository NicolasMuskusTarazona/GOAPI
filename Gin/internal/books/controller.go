package books

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct{
	service *BookService
}

func NewBookControllers(s *BookService) *BookController {
	return  &BookController{service: s,}
}

func (c *BookController)RegisterRoutes(r * gin.Engine){
	books := r.Group("/books")
	{
		books.GET("/", c.GetAll)
		books.POST("/", c.Create)
		books.GET("/:id", c.GetByID)
		books.DELETE("/:id", c.Delete)
		books.PUT("/:id",c.Put)
	}
}

func (c*BookController)GetAll(ctx *gin.Context)  {
	list, err := c.service.repo.FindAll()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,list)
}

func (c *BookController)Create(ctx *gin.Context)  {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.CreateBook(book)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book created successfully!"})
}

func (c*BookController) GetByID(ctx *gin.Context){
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	Book, err := c.service.repo.FindByID(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"error": "book not found"})
		return
	}

	ctx.JSON(http.StatusOK, Book)
}

func (c*BookController)Delete(ctx*gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = c.service.repo.Delete(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error":"book not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func (c*BookController)Put(ctx*gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id

	err = c.service.UpdateBook(book)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "book UPDATE! :D"})
}