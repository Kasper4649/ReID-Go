package controller

import (
	pb "ReID-Go/src/message"
	"ReID-Go/src/util"
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

func Search(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if strings.HasPrefix(fileHeader.Header["Content-Type"][0], "video/") {

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		client(fileContent, fileHeader.Filename)

	} else {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "非视频格式:)",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("search successfully"),
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
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "非法上传:)",
		})
		return
	}

	for _, file := range files {
		if strings.HasPrefix(file.Header["Content-Type"][0], "image/") {
			filePath := path.Join(config.QueryDirectory, "0001_c1s1_0_"+file.Filename)
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

func client(file []byte, filename string) {
	// Set up a connection to the server.
	config.GetConf()
	conn, err := grpc.Dial(config.GRPCServerAddress, grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(60*1024*1024),
			grpc.MaxCallSendMsgSize(60*1024*1024)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewSearchServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	r, err := c.Search(ctx, &pb.SearchRequest{File: file, Name: filename})
	if err != nil {
		log.Fatalf("could not receive output: %v", err)
	}
	err = ioutil.WriteFile("output.mp4", r.GetFile(), 0644)
	if err != nil {
		panic(err)
	}
}
