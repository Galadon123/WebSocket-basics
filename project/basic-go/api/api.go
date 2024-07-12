package api

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type PostData struct {
    Message string `json:"message"`  //When you receive a JSON payload in a POST request with a structure like{"message": "json-data"}
}
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to Homepage")
}
func testroute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to Test Route")
}
func postData(w http.ResponseWriter, r *http.Request){

	//decoder := json.NewDecoder(r.Body)
	var data PostData
	err := json.NewDecoder(r.Body).Decode(&data)  //Decode the JSON payload into the PostData structure
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Message: %s", data.Message)
}
func Api(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/test1",testroute)
	http.HandleFunc("/test2",postData)

    println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil) 
	
}