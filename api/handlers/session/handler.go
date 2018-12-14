package session

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "foo",
		Value: "bar",
	}
	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Hello World!")
}
