package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pmutua/health-insurance-api/app/config"
	"github.com/pmutua/health-insurance-api/app/handler"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/persons", a.GetAllPersons)
	a.Post("/persons", a.CreatePerson)
	a.Get("/persons", a.GetPerson)
	a.Put("/persons/{title}", a.UpdatePerson)
	a.Delete("/persons/{title}", a.DeletePerson)
	a.Put("/persons/{title}/disable", a.DisablePerson)
	a.Put("/persons/{title}/enable", a.EnablePerson)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Person Data
func (a *App) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPersons(a.DB, w, r)
}

func (a *App) CreatePerson(w http.ResponseWriter, r *http.Request) {
	handler.CreatePerson(a.DB, w, r)
}

func (a *App) GetPerson(w http.ResponseWriter, r *http.Request) {
	handler.GetPerson(a.DB, w, r)
}

func (a *App) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePerson(a.DB, w, r)
}

func (a *App) DeletePerson(w http.ResponseWriter, r *http.Request) {
	handler.DeletePerson(a.DB, w, r)
}

func (a *App) DisablePerson(w http.ResponseWriter, r *http.Request) {
	handler.DisablePerson(a.DB, w, r)
}

func (a *App) EnablePerson(w http.ResponseWriter, r *http.Request) {
	handler.EnablePerson(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
