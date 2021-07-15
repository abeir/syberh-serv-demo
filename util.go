package main

import (
	"encoding/json"
	"mime/multipart"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const SAVED_DIR = "test_download"

func SaveFiles(files map[string][]*multipart.FileHeader, context *gin.Context) (names []string, pathes []string, err error) {
	usr, err := user.Current()
	if err != nil {
		return nil, nil, err
	}
	dir := path.Join(usr.HomeDir, SAVED_DIR)
	if IsExists(dir) {
		_ = os.RemoveAll(dir)
	}
	_ = os.MkdirAll(dir, os.ModePerm)

	fields := make([]string, 0)
	saved := make([]string, 0)
	for n, v := range files {
		filePath := path.Join(dir, n+"_"+strings.ReplaceAll(uuid.NewV4().String(), "-", ""))
		for _, f := range v {
			if err = context.SaveUploadedFile(f, filePath); err != nil {
				return nil, nil, err
			}
			saved = append(saved, filePath)
		}
		fields = append(fields, n)
	}
	return fields, saved, nil
}

func IsExists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Contians(arr []string, str string) bool {
	if arr == nil {
		return false
	}
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func PrintMap(m map[string][]string) string {
	if m == nil || len(m) == 0 {
		return "{}"
	}
	data, err := json.Marshal(m)
	if err != nil {
		return "{}"
	}
	return string(data)
}

func PrintArray(arr []string) string {
	if arr == nil || len(arr) == 0 {
		return "[]"
	}
	data, err := json.Marshal(arr)
	if err != nil {
		return "[]"
	}
	return string(data)
}
