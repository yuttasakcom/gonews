package app

import (
	"net/http"

	"github.com/yuttasakcom/gonews/pkg/view"
)

func newsViews(w http.ResponseWriter, r *http.Request) {

}

func newsID(w http.ResponseWriter, r *http.Request) {
	view.NewsID(w, nil)
}
