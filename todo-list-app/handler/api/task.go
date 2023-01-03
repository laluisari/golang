package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	idCtx := r.Context().Value("id")
	id := idCtx.(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	idLogin, _ := strconv.Atoi(idCtx.(string))
	idTask := r.URL.Query().Get("task_id")

	if idTask == "" {
		listTask, err := t.taskService.GetTasks(r.Context(), idLogin)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&listTask)
		return
	}

	idTaskInt, _ := strconv.Atoi(idTask)

	dataTaskId, err := t.taskService.GetTaskByID(r.Context(), idTaskInt)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(&dataTaskId)

}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}
	categoryId := strconv.Itoa(task.CategoryID)

	if task.Title == "" || task.Description == "" || categoryId == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}
	idCtx := r.Context().Value("id")
	idLogin, err := strconv.Atoi(idCtx.(string))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	tempTask := entity.Task{
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      idLogin,
	}
	dataTask, err := t.taskService.StoreTask(r.Context(), &tempTask)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "error internal server"})
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": idLogin,
		"task_id": dataTask.ID,
		"message": "success create new task",
	})

	// TODO: answer here
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idCtx := r.Context().Value("id")
	idLogin, err := strconv.Atoi(idCtx.(string))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	taskID := r.URL.Query().Get("task_id")
	int_taskID, _ := strconv.Atoi(taskID)

	err = t.taskService.DeleteTask(r.Context(), int_taskID)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": idLogin,
		"task_id": int_taskID,
		"message": "success delete task",
	})
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}
	userId := r.Context().Value("id")
	idLogin, err := strconv.Atoi(userId.(string))

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid user id"})
		return
	}
	tempTask := entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
	}
	dataTask, err := t.taskService.UpdateTask(r.Context(), &tempTask)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "error internal server"})
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": idLogin,
		"task_id": dataTask.ID,
		"message": "success update task",
	})

	// TODO: answer here
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
