package main 


import (
	"net/http"
	"encoding/json"
	"log"
)

type Response struct {
	Message string `json:"message"`
}


func HandleFunc(w http.ResponseWriter, r *http.Request) {


	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", HandleFunc)	

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}