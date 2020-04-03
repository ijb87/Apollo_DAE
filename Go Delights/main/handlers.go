package main

type IHandler struct {
  type SHandler struct {}
}

func (ih IHandler) ServeHTTP(w http.ResponseWrite, r *http.Request){
  hostdom := strings.Splite(r.Host, ":")[0]
  http.Redirect(w, r, "https://" + hostdom + ":8181" + r.URL.Path, 302)
}

func (sh SHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "I hope you feel secure now you are here")
}