package utils

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

func GetStaticResourceURL(c *gin.Context, resource_path string) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	basePath := fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	return path.Join(basePath, "static", resource_path)
}
