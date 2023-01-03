package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	idCtx := r.Context().Value("id")
	idLogin, err := strconv.Atoi(idCtx.(string))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid user id"})
		return
	}

	listCategory, err := c.categoryService.GetCategories(r.Context(), int(idLogin))

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&listCategory)

	// TODO: answer here
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}
	if category.Type == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}
	idCtx := r.Context().Value("id")
	idLogin, err := strconv.Atoi(idCtx.(string))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid user id"})
		return
	}

	var temp = entity.Category{
		UserID: idLogin,
		Type:   category.Type,
	}

	data, err := c.categoryService.StoreCategory(r.Context(), &temp)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "error internal server"})
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":     int(idLogin),
		"category_id": data.ID,
		"message":     "success create new category",
	})

}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idCtx := r.Context().Value("id")
	idLogin, err := strconv.Atoi(idCtx.(string))
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "invalid user id"})
		return
	}

	idCtgr := r.URL.Query().Get("category_id")
	idCtgrInt, _ := strconv.Atoi(idCtgr)

	err = c.categoryService.DeleteCategory(r.Context(), idCtgrInt)

	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "error internal server"})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":     int(idLogin),
		"category_id": idCtgrInt,
		"message":     "success delete category",
	})
	// TODO: answer here
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
