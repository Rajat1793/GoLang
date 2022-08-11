package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course -file
type Course struct {
	CourseId    string  `json:courseid`
	CourseName  string  `json:coursename`
	CoursePrice string  `json:price`
	Author      *Author `json:author` //data type is pointer
}

type Author struct {
	Fullname string `json:fullname`
	Website  string `json:website`
}

// fake DB
var courses []Course

// middleware, helper-file
func (c *Course) IsEmpty() bool { //pointer of struct
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //if we want to manullly create a courseid
}

func main() {
	fmt.Println("Welcome to go API")
	r := mux.NewRouter()

	// enter the value in fake db
	courses = append(courses, Course{CourseId: "2", CourseName: "React", CoursePrice: "299", Author: &Author{Fullname: "Rajat", Website: "LCO"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Go lang", CoursePrice: "199", Author: &Author{Fullname: "Ram", Website: "LCO"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") //with the route we have to pass the {id} which is same as params id what we have passed in other function (both should be same)
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controller files goes in seperate file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API </h1>")) //slice of byte data
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //syntax for encoder it will throw back the json value to who ever is calling
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)

	// loop though courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//if its {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// todo: check only if title is duplicate
	// loop, title matches with course name, json

	// generate a unique id, string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000)) //formatting int into string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with myid(updated)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// send a response when id is not found
	for _, course := range courses {
		if course.CourseId != params["id"] {
			json.NewEncoder(w).Encode("No course found with given id")
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove(index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("one course deleted")
			break
		}
	}
}
