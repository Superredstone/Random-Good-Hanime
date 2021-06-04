package RandomGoodHanime

import (
	"io"
	"log"
	"net/http"
)

//Starts log webserver
func StartWebServer(port string) {
	HandleHttp()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HandleHttp() {
	http.HandleFunc("/", UpdateLog)
}

func UpdateLog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, Log)
}
