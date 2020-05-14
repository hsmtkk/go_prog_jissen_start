package chitchat

import "net/http"

func Handler() http.Handler {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	return mux
}
