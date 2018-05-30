package routing

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"text/template"
)

func universalRouter(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(
		"./templates/"+page+".html",
		"./templates/header.html",
		"./templates/footer.html")

	if err != nil {
		fmt.Println(err.Error())
	}

	t.ExecuteTemplate(w, page, nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "index") }
func BlogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "blog") }
func ShopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "shop") }
func BankHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "bank") }
func KriptHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "kript") }
func ErrorHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "error") }
