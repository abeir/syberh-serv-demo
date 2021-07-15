package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const DefaultPort = "8000"

type GetResponse struct {
	Result string `json:"result"`
}

func launch(port string) {
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
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func() {
			err = c.Request.Body.Close()
			if err != nil {
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
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.PUT("/put", func(c *gin.Context) {

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func() {
			err = c.Request.Body.Close()
			if err != nil {
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
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.DELETE("/delete", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer func() {
			err = c.Request.Body.Close()
			if err != nil {
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
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		_, err = c.Writer.Write(body)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	})

	engine.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		var saved []string
		files := form.File
		if files == nil || len(files) == 0 {
			fmt.Println("not found any upload files")
		} else {
			if _, saved, err = SaveFiles(files, c); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
		}

		formData := make(map[string][]string)
		for n, v := range form.Value {
			formData[n] = v
		}

		cookies := make([]string, 0)
		for _, cookie := range c.Request.Cookies() {
			cookies = append(cookies, cookie.String())
		}

		paths := strings.Join(saved, ";")

		fmt.Println(strings.Repeat("-", 15), "UPLOAD")
		fmt.Printf("file: %s\n", paths)
		fmt.Printf("form: %s\n", PrintMap(formData))
		fmt.Printf("header: %s\n", PrintMap(c.Request.Header))
		fmt.Printf("cookie: %s\n", PrintArray(cookies))
		fmt.Println(strings.Repeat("-", 20))

		result := make(map[string]interface{})
		result["file"] = paths
		c.JSON(http.StatusOK, result)
	})

	_ = engine.Run(":" + port)
}

func getPort() string {
	args := os.Args
	if args != nil && len(args) > 1 {
		_, err := strconv.ParseInt(args[1], 10, 64)
		if err == nil {
			return strings.TrimSpace(args[1])
		}
	}
	return DefaultPort
}

func main() {
	launch(getPort())
}
