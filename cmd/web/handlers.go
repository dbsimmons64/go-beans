package main

import (
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/dbsimmons64/go-beans/forms"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {

	// The pattern "/" matches all paths not matched by other registered routes.
	// We can use this fact to determine if the request is for an unknown route.
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: handler for %s not found", html.EscapeString(r.URL.Path))
		return
	}

	session := app.sessionStore.GetSession(w, r, "default_session")
	year := session.GetOrSetDefault("year", time.Now().Year())
	month := session.GetOrSetDefault("month", time.Now().Month())

	months := ListMonths()
	years := ListYears()

	transactions, err := app.transactionService.ListTransactions()
	form := forms.New(r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	app.renderPage(w, "hello.page.html", pageData{
		"Year":         year,
		"Month":        month,
		"Months":       months,
		"Years":        years,
		"Form":         form,
		"Transactions": transactions,
	})

}
func (app *app) About(w http.ResponseWriter, r *http.Request) {
	app.renderPage(w, "about.page.html", nil)
}
func (app *app) Contact(w http.ResponseWriter, r *http.Request) {
	app.renderPage(w, "contact.page.html", nil)
}

func (app *app) storeTransaction(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	form := forms.New(r.PostForm)
	form.Required("txn_date", "who", "payee", "amount", "category")
	form.ValidAmount("amount")

	if !form.Valid() {
		transactions, err := app.transactionService.ListTransactions()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		app.renderPage(w, "hello.page.html", pageData{"Form": form, "Transactions": transactions})
		return
	}

	err = app.transactionService.CreateTransaction(r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
