package controller

import (
	"ReID-Go/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

var config util.Conf

func Search(c *gin.Context) {

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

func Query(c *gin.Context) {

	config.GetConf()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		filePath := path.Join(config.QueryDirectory, "0001_c1s1_0_"+file.Filename)
		if err = c.SaveUploadedFile(file, filePath); err != nil {
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
