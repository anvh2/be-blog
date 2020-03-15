package backend

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"go.uber.org/zap"
)

func (s *Server) download(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Code: -1,
	}
	encoder := json.NewEncoder(w)

	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		s.logger.Error("[download] query params is invalid")
		res.Message = "Params is invalid"
		encoder.Encode(res)
		return
	}

	file, err := os.Open(s.dirStorage + fileName)
	if err != nil {
		s.logger.Error("[download] failed to open file", zap.Error(err))
		res.Message = "File not found"
		encoder.Encode(res)
		return
	}
	defer file.Close()

	header := make([]byte, 512)
	file.Read(header)
	contentType := http.DetectContentType(header)
	stat, _ := file.Stat()
	size := strconv.FormatInt(stat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", size)

	file.Seek(0, 0)
	io.Copy(w, file)

	s.logger.Info("[upload] file downloaded", zap.String("name", fileName))

	res.Code = 1
	res.Message = "OK"
	encoder.Encode(res)
	return
}
