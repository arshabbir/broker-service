package api

import "net/http"

func (s *server) Broker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"message\" : \"broker invoked\"}"))

}
