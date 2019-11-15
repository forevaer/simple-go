package trans

import (
	"context"
	"encoding/json"
	"learn_endpoint/param"
	"net/http"
)

func RequestDecode(ctx context.Context, r *http.Request) param.Request {
	var request param.Request
	_ = json.NewDecoder(r.Body).Decode(&request)
	return request
}

func ResponseEncode(ctx context.Context, writer http.ResponseWriter, data interface{}) error {
	_ = json.NewEncoder(writer).Encode(data)
	return nil
}
