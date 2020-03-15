package backend

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func (s *Server) move(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Code: -1,
	}
	encoder := json.NewEncoder(w)

	src := r.URL.Query().Get("src")
	if src == "" {
		s.logger.Error("[move] query params src is invalid")
		res.Message = "Src is invalid"
		encoder.Encode(res)
		return
	}
	des := r.URL.Query().Get("des")
	if des == "" {
		s.logger.Error("[move] query params des is invalid")
		res.Message = "Des is invalid"
		encoder.Encode(res)
		return
	}

	fileSrc, err := os.Open(s.dirStorage + src)
	if err != nil {
		s.logger.Error("[move] failed to open source file", zap.Error(err))
		res.Message = "Move file failed"
		encoder.Encode(res)
		return
	}

	fileDes, err := os.Create(s.dirStorage + des)
	if err != nil {
		s.logger.Error("[move] failed to create destination file", zap.Error(err))
		res.Message = "Move file failed"
		encoder.Encode(res)
		return
	}
	defer fileDes.Close()

	_, err = io.Copy(fileDes, fileSrc)
	if err != nil {
		s.logger.Error("[move] failed to copy from source to destination file", zap.Error(err))
		res.Message = "Move file failed"
		encoder.Encode(res)
		return
	}

	err = os.Remove(s.dirStorage + src)
	if err != nil {
		s.logger.Error("[move] failed to remove source file", zap.Error(err))

		fileSrc.Close()
		res.Message = "Move file failed"
		encoder.Encode(res)
		return
	}

	s.logger.Info("[upload] file moved", zap.String("src", src), zap.String("des", des))

	res.Code = 1
	res.Message = "OK"
	encoder.Encode(res)
	return
}
