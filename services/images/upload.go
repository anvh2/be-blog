package backend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func (s *Server) upload(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Code: -1,
	}
	encoder := json.NewEncoder(w)

	r.ParseMultipartForm(10 << 20)

	// fmt.Println(r.MultipartForm.File["image"])

	file, handler, err := r.FormFile(viper.GetString("images.form_file"))
	if err != nil {
		s.logger.Error("[upload] failed to retrieve file", zap.Error(err))
		res.Message = "Failed to retrieve file"
		encoder.Encode(res)
		return
	}
	defer file.Close()

	s.logger.Info("[upload] file info", zap.String("name", handler.Filename),
		zap.Int64("size", handler.Size), zap.Any("header", handler.Header))

	dir, err := s.getDir()
	if err != nil {
		s.logger.Error("[upload] failed to get directory", zap.Error(err))
		return
	}

	pattern := "image-*." + getExtension(handler.Filename)
	save, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		s.logger.Error("[upload] failed to create file", zap.Error(err))
		res.Message = "Failed to create file"
		encoder.Encode(res)
		return
	}
	defer save.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		s.logger.Error("[upload] failed to read file", zap.Error(err))
		res.Message = "Failed to read file"
		encoder.Encode(res)
		return
	}

	save.Write(fileBytes)

	name := save.Name()
	name = strings.Replace(name, "public/images", "static", -1)

	s.logger.Info("[upload] file saved", zap.String("name", name))

	link := "http://localhost:55201/" + name

	res.Code = 1
	res.Message = "OK"
	res.Data = link
	encoder.Encode(res)
	return
}

func getExtension(name string) string {
	arr := strings.Split(name, ".")
	if len(arr) <= 0 {
		return ""
	}
	return arr[len(arr)-1]
}

func (s *Server) getDir() (string, error) {
	dateFormat := time.Now().Format("060102")
	dir := s.dirStorage + dateFormat
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
	}
	return dir, err
}
