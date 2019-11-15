package service

import (
	"learn_endpoint/param"
	"time"
)

type Service struct {
}

func (s *Service) Serve(request param.Request) param.Response {
	return param.Response{
		Request: param.Request{
			Question: request.Question,
		},
		Answer: time.Now().String(),
	}
}
