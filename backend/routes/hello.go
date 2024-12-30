package routes

import (
	"encoding/json"
	"net/http"
)

func Print(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode("encoded json")
}
