package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.POST("/upload", upload)
	r.POST("/search", search)
	err := r.Run()
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

func search(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	// 提交给 Python 模型
	file.Open()

	c.JSON(http.StatusOK, gin.H{

	})
}


func upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		if err = c.SaveUploadedFile(file, path.Join("./", file.Filename)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("uploaded successfully %d files", len(files)),
	})
}
