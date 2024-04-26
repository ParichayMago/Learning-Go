package main

import (

	"database/sql"

		"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {

	Router *mux.Router 
	DB *sql.DB

}

type Product struct {
	ID int `json:"id"`
	Name string `json:name`
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(addr string) { }