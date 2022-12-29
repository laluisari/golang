package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"
)

func dataJson(status int, w http.ResponseWriter, msg string) {
	w.WriteHeader(status)
	jsonData, _ := json.Marshal(model.ErrorResponse{
		Error: msg,
	})
	w.Write(jsonData)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	card := model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		dataJson(400, w, "Internal Server Error")
		return
	}
	if card.Username == "" || card.Password == "" {
		dataJson(400, w, "Username or Password empty")
		return
	}
	_, ok := db.Users[card.Username]
	if ok {
		dataJson(409, w, "Username already exist")
		return
	}

	db.Users[card.Username] = card.Password
	w.WriteHeader(200)
	jsonData, _ := json.Marshal(model.SuccessResponse{
		Username: card.Username,
		Message:  "Register Success",
	})
	w.Write(jsonData)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	card := model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&card)

	if err != nil {
		dataJson(400, w, "Internal Server Error")
		return
	}

	if card.Username == "" || card.Password == "" {
		dataJson(400, w, "Username or Password empty")
		return
	}

	cek := db.Users[card.Username]
	if cek != card.Password {
		dataJson(401, w, "Wrong User or Password!")
		return
	}

	stoken := uuid.NewString()
	expired := time.Now().Add(5 * time.Hour)
	db.Sessions[stoken] = model.Session{
		Username: card.Username,
		Expiry:   expired,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   stoken,
		Expires: expired,
	})

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(model.SuccessResponse{
		Username: card.Username,
		Message:  "Login Success",
	})
	w.Write(jsonData)
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		dataJson(400, w, "Internal Server Error")
		return
	}
	todo := model.Todo{}
	json.Unmarshal(body, &todo)

	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(401)
		jsonData, _ := json.Marshal(model.ErrorResponse{
			Error: "http: named cookie not present",
		})
		w.Write(jsonData)
		return
	}

	sesUname := db.Sessions[cookie.Value]

	db.Task[sesUname.Username] = []model.Todo{}
	db.Task[sesUname.Username] = append(db.Task[sesUname.Username], model.Todo{
		Id:   cookie.Value,
		Task: todo.Task,
		Done: todo.Done,
	})

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(model.SuccessResponse{
		Username: sesUname.Username,
		Message:  "Task " + todo.Task + " added!",
	})
	w.Write(jsonData)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(401)
		jsonData, _ := json.Marshal(model.ErrorResponse{
			Error: "http: named cookie not present",
		})
		w.Write(jsonData)
		return
	}

	sesUname := db.Sessions[cookie.Value]

	getTodos, ok := db.Task[sesUname.Username]

	if !ok {
		w.WriteHeader(404)
		jsonData, _ := json.Marshal(model.ErrorResponse{
			Error: "Todolist not found!",
		})
		w.Write(jsonData)
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(getTodos)
	w.Write(jsonData)

	// TODO: answer here
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(401)
		jsonData, _ := json.Marshal(model.ErrorResponse{
			Error: "http: named cookie not present",
		})
		w.Write(jsonData)
		return
	}

	sesUname := db.Sessions[cookie.Value]

	db.Task[sesUname.Username] = []model.Todo{}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(model.SuccessResponse{
		Username: sesUname.Username,
		Message:  "Clear ToDo Success",
	})
	w.Write(jsonData)
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	c, err := r.Cookie("session_token")
	if err != nil {
		w.WriteHeader(401)
		jsonData, _ := json.Marshal(model.ErrorResponse{
			Error: "http: named cookie not present",
		})
		w.Write(jsonData)
		return
	}
	stoken := c.Value
	delete(db.Sessions, stoken)

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(model.SuccessResponse{
		Username: username,
		Message:  "Logout Success",
	})
	w.Write(jsonData)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})
	// TODO: answer here
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	// TODO: answer here
	mux.Handle("/todo/read", middleware.Get(http.HandlerFunc(ListToDo)))
	mux.Handle("/todo/create", middleware.Post(http.HandlerFunc(AddToDo)))
	mux.Handle("/todo/clear", middleware.Delete(http.HandlerFunc(ClearToDo)))
	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
