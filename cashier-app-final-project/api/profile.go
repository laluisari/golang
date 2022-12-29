package api

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("assets/images/", "img-avatar.png"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "directory tidak ada", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("assets/images/", "img-avatar.png"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "directory tidak ada", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

	api.dashboardView(w, r)
}
