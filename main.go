package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
	"strconv"
)

const DefaultPort = "8000"

type GetResponse struct {
	Result string 	`json:"result"`
}


func main() {
	port := GetPort()

	engine := gin.Default()

	engine.GET("/get", func(c *gin.Context) {
		rs := &GetResponse{
			Result: "ok",
		}
		header, _ := json.Marshal(c.Request.Header)

		fmt.Println(strings.Repeat("-", 15), "GET")
		fmt.Printf("url:    %s\n", c.Request.RequestURI)
		fmt.Printf("header: %s\n", header)
		fmt.Println(strings.Repeat("-", 20))

		c.JSON(http.StatusOK, rs)
	})


	engine.POST("/post", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func(){
			err = c.Request.Body.Close()
			if err!=nil {
				fmt.Println(err)
			}
		}()
		header, _ := json.Marshal(c.Request.Header)

		fmt.Println(strings.Repeat("-", 15), "POST")
		fmt.Printf("url:    %s\n", c.Request.RequestURI)
		fmt.Printf("header: %s\n", header)
		fmt.Printf("body:   %s\n", body)
		fmt.Println(strings.Repeat("-", 20))

		_, err = c.Writer.Write(header)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.PUT("/put", func(c *gin.Context) {

		body, err := ioutil.ReadAll(c.Request.Body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func(){
			err = c.Request.Body.Close()
			if err!=nil {
				fmt.Println(err)
			}
		}()
		header, _ := json.Marshal(c.Request.Header)

		fmt.Println(strings.Repeat("-", 15), "PUT")
		fmt.Printf("url:    %s\n", c.Request.RequestURI)
		fmt.Printf("header: %s\n", header)
		fmt.Printf("body:   %s\n", body)
		fmt.Println(strings.Repeat("-", 20))

		_, err = c.Writer.Write(header)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.DELETE("/delete", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func(){
			err = c.Request.Body.Close()
			if err!=nil {
				fmt.Println(err)
			}
		}()
		header, _ := json.Marshal(c.Request.Header)

		fmt.Println(strings.Repeat("-", 15), "DELETE")
		fmt.Printf("url:    %s\n", c.Request.RequestURI)
		fmt.Printf("header: %s\n", header)
		fmt.Printf("body:   %s\n", body)
		fmt.Println(strings.Repeat("-", 20))

		_, err = c.Writer.Write(header)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.POST("/upload", func(c *gin.Context){
		file, err := c.FormFile("file")
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		usr, err := user.Current()
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		dir := path.Join(usr.HomeDir, "test_download")
		if IsExists(dir) {
			_ = os.RemoveAll(dir)
		}
		_ = os.MkdirAll(dir, os.ModePerm)

		filePath := path.Join(dir, strings.ReplaceAll(uuid.NewV4().String(), "-", ""))
		err = c.SaveUploadedFile(file, filePath)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		//c.String(http.StatusOK, filePath)
		var result = make(map[string]interface{})
		result["filePath"] = filePath

		var bodyResult = make(map[string][]string)
		for k, v := range c.Request.PostForm {
			bodyResult[k] = v
		}
		result["body"] = bodyResult
		result["headers"] = c.Request.Header

		data, err := json.Marshal(result)
		if err!=nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(strings.Repeat("-", 15), "UPLOAD")
		fmt.Println(string(data))
		fmt.Println(strings.Repeat("-", 20))

		c.JSON(http.StatusOK, result)
	})

	_ = engine.Run(":" + port)
}


func IsExists(p string) bool{
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetPort() string {
	args := os.Args
	if args != nil && len(args) > 1 {
		_, err := strconv.ParseInt(args[1], 10, 32)
		if err == nil {
			return args[1]
		}
	}
	return DefaultPort
}

