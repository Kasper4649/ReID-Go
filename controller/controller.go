package controller

import (
	pb "ReID-Go/message"
	"ReID-Go/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

var config util.Conf

// @ID Search
// @Summary 查找特定的行人
// @Description 查找特定的行人
// @Accept mpfd
// @Produce json
// @Param file formData file true "待查找视频"
// @Success 200 {string} string "{"message": "searched successfully"}"
// @Failure 400 {string} string "{"error": {}}"
// @Failure 415 {string} string "{"error": {}}"
// @Failure 500 {string} string "{"error": {}}"
// @Router /search [post]
func Search(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var outputUrl string

	if strings.HasPrefix(fileHeader.Header["Content-Type"][0], "video/") {

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 查找特定行人
		if outputUrl, err = client(fileContent, fileHeader.Filename); err != nil {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{
				"error": err,
			})
			return
		}

	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "非视频格式:)",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("searched successfully"),
		"url":  outputUrl,
	})
}

// @ID Query
// @Summary 选取特定的行人
// @Description 选取特定的行人
// @Accept mpfd
// @Produce json
// @Param files formData file true "行人图片"
// @Success 200 {string} string "{"message": "uploaded successfully"}"
// @Failure 400 {string} string "{"error": {}}"
// @Failure 415 {string} string "{"error": {}}"
// @Failure 500 {string} string "{"error": {}}"
// @Router /query [post]
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
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "非法上传:)",
		})
		return
	}

	for _, file := range files {
		if strings.HasPrefix(file.Header["Content-Type"][0], "image/") {
			filePath := path.Join(config.QueryDirectory, "0001_c1s1_000"+file.Filename+"_00.jpg")
			if err = c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{
				"error": "非图片格式:)",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("uploaded successfully %d files", len(files)),
	})
}

func client(file []byte, filename string) (string, error) {
	// Set up a connection to the server.
	config.GetConf()
	conn, err := grpc.Dial(config.GRPCServerAddress, grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(60*1024*1024),
			grpc.MaxCallSendMsgSize(60*1024*1024)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	c := pb.NewSearchServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	r, err := c.Search(ctx, &pb.SearchRequest{ File: file, Name: filename })
	if err != nil {
		log.Fatalf("could not receive output: %v", err)
		return "", err
	}

	return r.GetUrl(), nil
}
