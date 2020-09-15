package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// https://gorm.io/docs/
	initialMigration()
	handleRequests2()
}

func initialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

func handleRequests2() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/users", createUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", getUserById).Methods("GET")
	myRouter.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var users []User
	db.Find(&users)
	fmt.Println(users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Valid user body required")
	}
	var newUser User
	json.Unmarshal(reqBody, &newUser)

	db.Create(&newUser)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created")
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "A valid id is required")
	}
	var user User
	db.First(&user, userId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "A valid id is required")
	}
	db.Delete(&User{}, userId)
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintf(w, "Successfully deleted user with id "+strconv.Itoa(userId))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "A valid id is required")
	}
	newUserData, err :=ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "A valid body user data is required")
	}
	var newUser User
	json.Unmarshal(newUserData, &newUser)

	var user User
	db.First(&user, userId)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User id not found")
	} else {
		user.Name = newUser.Name
		user.Email = newUser.Email
		db.Save(&user)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Successfully update user with id "+strconv.Itoa(userId))
	}

}
