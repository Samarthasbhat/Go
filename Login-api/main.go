// Endpoints: Register, Login, Get user profile

// Data: Username, email, password (basic only, no full auth yet)


package main 


import(
	"log"
	"net/http"
	// "encoding/json"
	"fmt"
)

type User string 

type database map[string]User


func HandleFunc(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func (db database) register(w http.ResponseWriter, r *http.Request) {
	username  := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if _, ok := db[username]; ok {
		msg := "duplicate username"
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
		// Basic validation (optional)
	if username == "" || password == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}

	log.Printf("User registered: %s", username)
	fmt.Fprintf(w, "User %s registered successfully\n", username)


}

func main() {

		db := database {
		"Username": "Samartha",
		"Password":"@123#23",
	}
	http.HandleFunc("/register", db.register)

	log.Fatal(http.ListenAndServe(":8080", nil))

}