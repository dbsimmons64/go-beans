package main

import (
	"net/http"

	"github.com/dbsimmons64/go-beans/forms"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {
	transactions, err := app.transactionService.ListTransactions()
	form := forms.New(r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	app.render(w, "hello.page.html", pageData{"Form": form, "Transactions": transactions})

}
func (app *app) About(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.page.html", nil)
}
func (app *app) Contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.page.html", nil)
}

func (app *app) storeTransaction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	form := forms.New(r.PostForm)
	form.Required("txn_date", "who", "payee", "amount", "category")

	if !form.Valid() {
		transactions, err := app.transactionService.ListTransactions()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		app.render(w, "hello.page.html", pageData{"Form": form, "Transactions": transactions})
		return
	}

	err = app.transactionService.CreateTransaction(r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
