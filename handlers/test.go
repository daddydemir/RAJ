package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"RAJ/generator"
	"RAJ/models"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Test method is run")
	var model models.Object
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &model)
	query := generator.CreateTable(model)
	_ = json.NewEncoder(w).Encode(query)
}
