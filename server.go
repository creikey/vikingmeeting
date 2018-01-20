package main

import (
	"log"
)

const (
	LOGFILE = "log.txt"
)

func main() {
	log.Println("Starting server")
	test := NewChat("Hello!")
	log.Printf("Chat to write|%v|filename of the log: %v", test, LOGFILE)
	err := test.LogToFile(LOGFILE)
	if err != nil {
		log.Fatal(err)
	}
	/*logf, err := os.Open(LOGFILE)
	if err != nil {
		log.Fatal(err)
	}
	logf.WriteString("Test")
	logf.Close()*/
}

/*func mainHandler(w http.ResponseWriter, r *http.Request) {
	title := "Viking Games Meeting"
	p, _ := loadPa
}*/
