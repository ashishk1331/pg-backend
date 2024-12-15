package generate

import (
	"github.com/gin-gonic/gin"
	"pg-backend/template"
)

func Get(c *gin.Context) {
	var id string = c.Param("id")

	var template string = template.Generate(id)

	c.String(200, template)
}