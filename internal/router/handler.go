package router

import (
	"context"
	"gRPC/internal/proto/pb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type Router struct {
	grpcClient pb.ServiceClient
}

func NewRouter(grpcClient pb.ServiceClient) *Router {
	return &Router{
		grpcClient: grpcClient,
	}
}

func (a *Router) PrepareRouter() http.Handler {
	rout := mux.NewRouter()
	v2Rout := rout.PathPrefix("/v2").Subrouter()
	v2Rout.Path("/test/get-test").
		HandlerFunc(a.TestHandler).
		Methods(http.MethodGet, http.MethodOptions)
	rout.Use(mux.CORSMethodMiddleware(rout))
	return rout
}

func (a *Router) TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	txtParam := r.URL.Query().Get("text")
	if len(txtParam) == 0 {
		sendPlainText(w, "add query parameter `text` to method", http.StatusBadRequest)
		return
	}
	log.Printf("get from client txt parameter: %s", txtParam)
	ctx := context.Background()
	req := pb.GetRequestTest{Text: txtParam}
	res, err := a.grpcClient.Test(ctx, &req)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			sendPlainText(w, grpcErr.Message(), http.StatusBadRequest)
			return
		}
		sendPlainText(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("get from grpc error response: %s", res.Result)
	sendPlainText(w, res.Result, http.StatusOK)
}

func sendPlainText(w http.ResponseWriter, text string, code int) {
	w.WriteHeader(code)
	_, err := w.Write([]byte(text))
	if err != nil {
		log.Printf("Failed to send response: %v\n", err)
	}
}
