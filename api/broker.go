package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/arshabbir/brokermod/models"
	"github.com/arshabbir/brokermod/utils"
)

const authUrl = "http://authentication-service:8082/auth"

func (s *server) Broker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"message\" : \"broker invoked\"}"))

}

func (s *server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	authData := models.AuthData{}
	if err := json.NewDecoder(r.Body).Decode(&authData); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalied Request")
		return

	}
	if err := s.autheniticate(w, authData); err != nil {
		log.Println("error in authenitication ", err.Error())
		return
	}
}

func (s *server) autheniticate(w http.ResponseWriter, authdata models.AuthData) error {
	jData, err := json.Marshal(&authdata)
	if err != nil {
		return utils.SendError(w, http.StatusInternalServerError, "Internal server error")

	}
	// For a http request
	aReq, _ := http.NewRequest(http.MethodPost, authUrl, bytes.NewBuffer(jData))
	c := http.Client{}
	resp, err := c.Do(aReq)
	log.Printf("authniticate Response : %v", resp)

	if err != nil {
		return utils.SendError(w, http.StatusInternalServerError, "Auth Service not reachable")
	}
	defer resp.Body.Close()
	// Need fix the return of nil response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.SendError(w, http.StatusInternalServerError, "Auth Service not reachable")
	}
	authResponse := models.UserResponse{}
	if err := json.Unmarshal(b, &authResponse); err != nil {
		return utils.SendError(w, http.StatusInternalServerError, "error parsing auth response")
	}

	log.Println("Status code : ", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return utils.SendError(w, http.StatusUnauthorized, "Bad credientials")
	}
	if err := json.NewEncoder(w).Encode(&authResponse); err != nil {
		log.Fatal("error in sending response ")
	}
	return nil
}
