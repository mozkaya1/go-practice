package main
import (
	"log"
	"os"
)
func main() {
	file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "My Logger ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Application started...")
	logger.Println("My first custom logger")
	logger.Println("Application stopped...")
}

