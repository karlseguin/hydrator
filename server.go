package hydrator

import (
	"log"
	"net/http"
	"strconv"
)

func Start() {
	http.HandleFunc("/", handler)
	log.Println("Server running at http://127.0.0.1:4006/")
	log.Fatal(http.ListenAndServe("127.0.0.1:4006", nil))
}

func handler(output http.ResponseWriter, req *http.Request) {
	response, err := Proxy(req)
	if err != nil {
		log.Println(err.Error())
		output.WriteHeader(500)
		return
	}
	for k, v := range response.Header() {
		response.Header()[k] = v
	}
	body := response.Body()
	output.Header()["Content-Length"] = []string{strconv.Itoa(len(body))}
	output.WriteHeader(response.Status())
	output.Write(body)
}
