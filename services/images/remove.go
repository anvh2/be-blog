package backend

import (
	"encoding/json"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func (s *Server) remove(w http.ResponseWriter, r *http.Request) {
	res := &Response{
		Code: -1,
	}
	encoder := json.NewEncoder(w)

	fileName := r.URL.Query().Get("name")
	if fileName == "" {
		s.logger.Error("[remove] query params is invalid")
		res.Message = "Params is invalid"
		encoder.Encode(res)
		return
	}

	err := os.Remove(s.dirStorage + fileName)
	if err != nil {
		s.logger.Error("[remove] failed to remove file", zap.Error(err))
		res.Message = "Failed to remove file"
		encoder.Encode(res)
		return
	}

	s.logger.Info("[upload] file removed", zap.String("name", fileName))

	res.Code = 1
	res.Message = "OK"
	encoder.Encode(res)
	return
}
