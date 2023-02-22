package main

import (
	"curso_golang/Faydiamond/basesweb/internal/course"
	"curso_golang/Faydiamond/basesweb/internal/enrollment"
	"curso_golang/Faydiamond/basesweb/internal/pkg/bootstrap"
	"curso_golang/Faydiamond/basesweb/internal/user"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	_ = godotenv.Load()
	l := bootstrap.InitLogger()

	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	userRepo := user.NewRepo(l, db)
	userSrv := user.NewService(l, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	courseRepo := course.NewRepo(db, l)
	courseSrv := course.NewService(l, courseRepo)
	courseEnd := course.MakeEndpoints(courseSrv)

	enrollRepo := enrollment.NewRepo(db, l)
	enrollSrv := enrollment.NewService(l, userSrv, courseSrv, enrollRepo)
	enrollEnd := enrollment.MakeEndpoints(enrollSrv)

	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Delete).Methods("DELETE")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")

	router.HandleFunc("/enrollments", enrollEnd.Create).Methods("POST")
	server := &http.Server{
		Addr:         "127.0.0.1:3333",
		Handler:      router,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
	}
	fmt.Println(" serve in 127.0.0.1:3333 ")

	log.Fatal(server.ListenAndServe())

}
