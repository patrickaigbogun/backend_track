package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// var Name string = "patrick"

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", Name, Name)
// }

func formattedDateTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func main() {

	date := formattedDateTime()

	message := map[string]string{"email": "patrickaigbogunoti@gmail.com", "current_datetime": date, "github_url": "https://github.com/patrickaigbogun/backend_track"}

	messageInJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling JSON", err)
		return
	}

	http.HandleFunc("/hng_12/v0/1/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(messageInJson)
	})

	http.ListenAndServe(":3000", nil)
}
