package main

import (
       "io"
       "flag"
       "log"
       "net/http"
       "strconv"
)

// upload page
func HelloServer(w http.ResponseWriter, req *http.Request) {
     if "GET" != req.Method {
        log.Fatal("unsupported request method")
     }
     io.WriteString(w, "<html><header></header><body>hi there</body></html>\n")
}

func main() {
     var port = flag.Int("port", 8080, "port to listen on (default: 8080)")
     flag.Parse()
     log.Println("listening on port " + strconv.Itoa(*port))
     http.HandleFunc("/", HelloServer)
	 err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
	 if err != nil {
        log.Fatal("server failed: ", err)
     }
}
