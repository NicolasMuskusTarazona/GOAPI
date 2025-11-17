package routes

import(
    "net/http"
    "github.com/gin-gonic/gin"
)

type Usuario struct{
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine){

	r.LoadHTMLGlob("templates/*")

	r.GET("/",func(c *gin.Context) {
		//c.String(200,"Hola :D")
        //c.String(http.StatusOK,"Hola Mundo")
		c.HTML(http.StatusOK,"index.html", gin.H{
			"Title": "Mi aplicacion",
			"Heading": "Nicolas Muskus Tarazona",
			"Message": "Bienvenido a mi aplicacion web con Gin y plantillas HTML.",
		})
	})

	r.Static("/static","./static")
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