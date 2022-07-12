package handlers

import (
	"github.com/SaiKumarThota9/awesomeProject/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
