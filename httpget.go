package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var log string
	fmt.Println("Логин")
	fmt.Scanf("%s\n", &log)

	var pass int
	fmt.Println("Пароль")
	fmt.Scanf("%d\n", &pass)

	router := mux.NewRouter()

	router.HandleFunc("/avatar", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		login := r.FormValue("login")
		password, err := strconv.Atoi(r.FormValue("password"))
		if err != nil {
			password = 200
		}

		data := r.FormValue("data")
		dataanswer := data + "1"

		logintrue := bool(log == login)
		passtrue := bool(pass == password)

		autorizationok := logintrue && passtrue

		if autorizationok {
			fmt.Fprint(w, dataanswer, login, password, t)

		} else {

			fmt.Fprint(w, " access denied ")
		}

	})
	http.ListenAndServe(":3000", router)

}
