package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", doRespond)
	r.GET("/recordvoice", doRespond2)
	_ = r.Run(":8080")
}

func doRespond(c *gin.Context) {
	resp := `
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play digits="wwwwwwwwwwww9ww9ww9wwwwwwwwwwwwwwwwwwwwwwww"/>
    <Hangup/>
</Response>
	`
	c.Data(200, "text/xml", []byte(strings.TrimSpace(resp)+"\n"))
}

func doRespond2(c *gin.Context) {
	resp := `
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Record timeout="30" transcribe="false" />
    <Hangup/>
</Response>
	`
	c.Data(200, "text/xml", []byte(strings.TrimSpace(resp)+"\n"))
}

