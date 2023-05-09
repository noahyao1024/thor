package server
 import ( "encoding/base64"
 "golib/fakehtml"
 "net/http"
 "github.com/gin-gonic/gin"
 )
 func stage2(e *gin.Engine) *gin.Engine { 
	e.GET("/fakehtml/bootstrap.min.css", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Bootstrapmincss())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/bootstrap.min.css.map", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Bootstrapmincssmap())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/bootstrap.min.js", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Bootstrapminjs())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/bootstrap.min.js.map", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Bootstrapminjsmap())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/detail.html", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Detailhtml())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/jquery.min.js", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Jqueryminjs())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/popper.min.js", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Popperminjs())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/popper.min.js.map", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Popperminjsmap())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	e.GET("/fakehtml/reports.html", func(c *gin.Context) {
		decodedBytes, _ := base64.StdEncoding.DecodeString(fakehtml.Reportshtml())
		c.Data(http.StatusOK, "text/html; charset=utf-8", decodedBytes)
	})
	return e
}
