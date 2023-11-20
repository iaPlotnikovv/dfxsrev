package main

import (
	"net/http"
)

func main() {

	//create mux
	mux := http.NewServeMux()

	pHandler := Ilia{}
	mux.Handle("/plotnikov", pHandler)

	//server

	s := &http.Server{
		Addr:    ":1311",
		Handler: mux,
	}
	s.ListenAndServe()

}

type Ilia struct{}

func (p Ilia) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("HELLO WORLD! I'm Ilia!")
	res.WriteHeader(200)
	res.Write(data)
}

//curl -v -X GET http://localhost:1311/plotnikov
