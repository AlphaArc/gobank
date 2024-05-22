package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter,status int,toJSON any)error{
	w.Header().Add("Content-type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(toJSON)
}

type apiFunc func(http.ResponseWriter,*http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request)  {
		if err := f(w,r); err!=nil{
			WriteJSON(w,http.StatusBadRequest,ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAdr string
	store storage
}

func NewAPIServer(listenAddr string,store storage) *APIServer {
	return &APIServer{
		listenAdr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account",makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API Server running on port: ",s.listenAdr)

	http.ListenAndServe(s.listenAdr,router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter,r  *http.Request) error {
	switch r.Method{
		case "GET" :  return s.handleGetAccount(w,r)
		case "POST" :  return s.handleCreateAccount(w,r)
		case "DELETE" :  return s.handleDeleteAccount(w,r)
	} 

	return fmt.Errorf("method not allowed %s",r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter,r  *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return WriteJSON(w,http.StatusOK,&Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter,r  *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	createAccountRequest.FirstName,createAccountRequest.LastName = mux.Vars(r)["FirstName"],mux.Vars(r)["LastName"]
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest);err !=nil{
		return err
	}
	account := NewAccount(createAccountRequest.FirstName,createAccountRequest.LastName)
	if err  := s.store.CreateAccount(account);err!=nil{
		return err
	}
	return WriteJSON(w,http.StatusOK,account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter,r  *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter,r  *http.Request) error {
	return nil
}