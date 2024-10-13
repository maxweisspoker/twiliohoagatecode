package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", doRespond)
	_ = r.Run(":8080")
}

func doRespond(c *gin.Context) {
	resp := `
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play digits="wwww9wwwwwwww"/>
    <Hangup/>
</Response>
	`
	c.Data(200, "text/xml", []byte(strings.TrimSpace(resp)+"\n"))
}
