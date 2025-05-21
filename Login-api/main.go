// Endpoints: Register, Login, Get user profile

// Data: Username, email, password (basic only, no full auth yet)


package main 


import(
	"log"
	"net/http"
	// "encoding/json"
	"fmt"
)

type User struct {
	Username string
	Password string
}

type database map[string]User



func HandleFunc(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func (db database) list(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, user := range db {
		fmt.Fprintln(w, user.Username)
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

	// Store the user in the database
	db[username] = User{
		Username: username,
		Password: password,
	}

	log.Printf("User registered: %s", username)
	fmt.Fprintf(w, "User %s registered successfully\n", username)


}


// Delete user

func (db database) drop(w http.ResponseWriter, req *http.Request){
	username := req.URL.Query().Get("username")

	if _, ok := db[username]; !ok {
		msg := fmt.Sprintf("no such item: %q", username)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, username)
	fmt.Fprintf(w, "deleted %s\n", username)
}

func main() {

		db := database {
		"samartha": {
		Username: "samartha",
		Password: "234",
	},
}
	http.HandleFunc("/register", db.register)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/drop", db.drop)

	log.Fatal(http.ListenAndServe(":8080", nil))

}