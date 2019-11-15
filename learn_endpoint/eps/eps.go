package eps

import (
	"context"
	"learn_endpoint/param"
	"learn_endpoint/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndPoint(srv service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.Serve(request.(param.Request)), nil
	}
}
