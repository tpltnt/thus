package thus

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Serve the HTML form to upload a file
func formServer(w http.ResponseWriter, req *http.Request) {
	if "GET" != req.Method {
		log.Fatal("unsupported request method")
	}
	io.WriteString(w, "<html><header></header><body>\n")
	io.WriteString(w, "<form action=\"http://localhost:8080/receive\" method=\"post\" enctype=\"multipart/form-data\">")
	io.WriteString(w, "<label for=\"file\">filename:</label>")
	io.WriteString(w, "<input type=\"file\" name=\"file\" id=\"file\">")
	io.WriteString(w, "<input type=\"submit\" name=\"submit\" value=\"submit\">")
	io.WriteString(w, "</form></body></html>\n")
}

// Handle the file upload
func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if "POST" != req.Method {
		log.Fatal("unsupported request method for file upload")
	}

	// get file by ID from POST request
	file, header, err := req.FormFile("file")
	if err != nil {
		log.Fatal("error getting file from form data: ", err)
	}
	defer file.Close()

	// try to create file handler
	out, err := os.Create("/tmp/uploaded_file")
	if err != nil {
		log.Fatal("Could not create file. Maybe wrong privileges?")
	}
	defer out.Close()

	// write content of file to disk
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// indicate success
	io.WriteString(w, "<html><header></header><body>sucessfully uploaded ")
	io.WriteString(w, header.Filename)
	io.WriteString(w, "</body></html>")
}

func main() {
	var port = flag.Int("port", 8080, "port to listen on (default: 8080)")
	var ip = flag.String("ip", "", "IP to bind to")
	flag.Parse()
	log.Println("listening on " + *ip + ":" + strconv.Itoa(*port))
	http.HandleFunc("/", formServer)
	http.HandleFunc("/receive", uploadHandler)
	err := http.ListenAndServe(*ip+":"+strconv.Itoa(*port), nil)
	if err != nil {
		log.Fatal("server failed: ", err)
	}
}
