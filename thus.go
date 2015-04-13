package main

import (
       "io"
       "flag"
       "log"
       "net/http"
       "strconv"
)

// Serve the HTML form to upload a file
func FormServer(w http.ResponseWriter, req *http.Request) {
     if "GET" != req.Method {
        log.Fatal("unsupported request method")
     }
     io.WriteString(w, "<html><header></header><body>\n")
     io.WriteString(w, "<form action=\"http://localhost:8080/receive\" method=\"post\" enctype=\"multipart/form-data\">")
     io.WriteString(w, "<label for=\"file\">filename:</label>")
     io.WriteString(w, "<input type=\"file\" name=\"file\" id=\"upfile\">")
     io.WriteString(w, "<input type=\"submit\" name=\"submit\" value=\"submit\">")
     io.WriteString(w, "</form></body></html>\n")
}

func main() {
     var port = flag.Int("port", 8080, "port to listen on (default: 8080)")
     flag.Parse()
     log.Println("listening on port " + strconv.Itoa(*port))
     http.HandleFunc("/", FormServer)
	 err := http.ListenAndServe(":" + strconv.Itoa(*port), nil)
	 if err != nil {
        log.Fatal("server failed: ", err)
     }
}
