package api

import "net/http"

func (s *server) Broker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("broker invoked"))

}
