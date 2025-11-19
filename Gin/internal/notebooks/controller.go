package notebooks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteBookController struct{
	service *NoteBookService
}

func NewNoteBookControllers(s *NoteBookService) *NoteBookController {
	return  &NoteBookController{service: s,}
}
// Rutas
func (c *NoteBookController)RegisterRoutes(r * gin.RouterGroup){
	{
		r.GET("/notebooks", c.GetAll)
		r.GET("/notebooks/:id", c.GetByID)
		r.POST("/notebooks", c.Create)
		r.PUT("/notebooks/:id",c.Put)
		r.DELETE("/notebooks/:id", c.Delete)
	}
}
// Obtener todos
func (c*NoteBookController)GetAll(ctx *gin.Context)  {
	list, err := c.service.repo.FindAll()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,list)
}
// Obtener por ID
func (c*NoteBookController) GetByID(ctx *gin.Context){
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	NoteBook, err := c.service.repo.FindByID(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"error": "NoteBook not found"})
		return
	}

	ctx.JSON(http.StatusOK, NoteBook)
}
// Crear
func (c *NoteBookController)Create(ctx *gin.Context)  {
	var notebook NoteBook

	if err := ctx.ShouldBindJSON(&notebook); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.service.CreateNoteBook(notebook)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "NoteBook created successfully!"})
}
// Eliminar
func (c*NoteBookController)Delete(ctx*gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = c.service.repo.Delete(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error":"NoteBook not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "NoteBook deleted"})
}
// Actualizar
func (c*NoteBookController)Put(ctx*gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var notebook NoteBook
	if err := ctx.ShouldBindJSON(&notebook); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	notebook.ID = id

	err = c.service.UpdateNoteBook(notebook)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "NoteBook not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "NoteBook UPDATE! :D"})
}