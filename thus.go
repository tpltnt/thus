package main

import (
       "io"
       "log"
       "net/http"
)

// upload page
func HelloServer(w http.ResponseWriter, req *http.Request) {
     if "GET" != req.Method {
        log.Fatal("unsupported request method")
     }
     io.WriteString(w, "<html><header></header><body>hi there</body></html>\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server failed: ", err)
	}
}