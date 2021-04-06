package gigi

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine