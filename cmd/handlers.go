package main

import (
	"net/http"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {
	transactions, err := app.transactionService.ListTransactions()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	app.render(w, "hello.page.html", pageData{"Transactions": transactions})

}
func (app *app) About(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.page.html", nil)
}
func (app *app) Contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.page.html", nil)
}
