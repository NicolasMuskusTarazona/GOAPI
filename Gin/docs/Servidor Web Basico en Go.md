# Servidor Web Basico 

```go
package main

import"github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/",func(c *gin.Context) {
		//c.String(200,"Hola :D")
		c.JSON(200,gin.H{
			"message":"Hola :D",
		})
	})

	r.Run(":8080")
}
```