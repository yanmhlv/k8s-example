package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"io/ioutil"
)

func main() {
	n := rand.Intn(100)

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "service.1 %d", n)
	})

	router.GET("/:name", func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		urlStr := fmt.Sprintf("http://%s:3000/", p.ByName("name"))
		resp, err := http.Get(urlStr)
		if err != nil {
			fmt.Fprintf(w, "get error: %+v", err)
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "read resp error: %+v", err)
			return
		}

		fmt.Fprintf(w, "service-%d: response from %s is [%s]", n, p.ByName("name"), string(data))
	})

	(&http.Server{
		Handler: router,
		Addr:    ":3000",
	}).ListenAndServe()
}
