package routes

import (
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
)

type Usuario struct{
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

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
/*	r.GET("/saludo/:nombre",func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola, %s :D",nombre)
	})

	r.POST("/usuarios", func (c *gin.Context)  {
		var nuevoUsuario Usuario 

		if err := c.BindJSON(&nuevoUsuario); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H {"error": "Error al decodificar el JSON"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == ""{
			c.JSON(http.StatusBadRequest, gin.H {"eror": "Nombre y correo electronico son campos requeridos"})
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"mensaje":"Usuario registrado","datos": usuarios})
	})
		*/
}