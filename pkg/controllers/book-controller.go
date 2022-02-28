package routes

import (
	"github.com/gorilla/mux"
	"github.com/shem/go-book/pkg/controllers"
)


var RegisterBookStoreRoutes = func(router *mux.Router)  {
	router.Handlefunc("/book/", controllers.CreateBook).methods("POST")
	router.Handlefunc("/book/", controllers.GetBook).methods("GET")
	router.Handlefunc("/book/bookId", controllers.GetBookById).methods("GET")
	router.Handlefunc("/book/bookId", controllers.updateBook).methods("PUT")
	router.Handlefunc("/book/bookId", controllers.DeleteBook).methods("DELETE")
}