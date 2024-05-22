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
	router.HandleFunc("/account/{id}",makeHTTPHandleFunc(s.handleAccountByID))

	log.Println("JSON API Server running on port: ",s.listenAdr)

	http.ListenAndServe(s.listenAdr,router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter,r  *http.Request) error {
	switch r.Method{
		case "GET" :  return s.handleGetAllAccount(w,r)
		case "POST" :  return s.handleCreateAccount(w,r)
		// case "DELETE" :  return s.handleDeleteAccount(w,r)
	} 

	return fmt.Errorf("method not allowed %s",r.Method)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter,r  *http.Request) error {
	switch r.Method{
		case "GET" :  return s.handleGetAccountByID(w,r)
		case "POST" :  return s.handleCreateAccount(w,r)
		case "DELETE" :  return s.handleDeleteAccount(w,r)
	} 

	return fmt.Errorf("method not allowed %s",r.Method)
}

func (s *APIServer) handleGetAllAccount(w http.ResponseWriter,r  *http.Request) error {
	fmt.Println(r.Method+" Need Auth")
	accounts,  err := s.store.GetAllAccounts();if err!=nil{
		return nil
	}
	return WriteJSON(w,http.StatusOK,accounts)
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

func (s *APIServer) handleGetAccountByID(w  http.ResponseWriter,r *http.Request) error{
	// var id string = mux.Vars(r)["id"]
	// account,err :=  s.store.GetAccountByID(int(id))
	// if err!=nil{
	// 	return err
	// }
	fmt.Println(r.Body)
	acc := `"AVC":"123"`
	return WriteJSON(w,http.StatusOK,acc)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter,r  *http.Request) error {
	fmt.Println(r.Body)
	acc := `"AVC":"123"`
	return WriteJSON(w,http.StatusOK,acc)
}

// func (s *APIServer) handleTransfer(w http.ResponseWriter,r  *http.Request) error {
// 	fmt.Println(r.Body)
// 	acc := `"AVC":"123"`
// 	return WriteJSON(w,http.StatusOK,acc)
// }