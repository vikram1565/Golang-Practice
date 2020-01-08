package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"username": []string{"required", "alpha"},
		"email":    []string{"required", "email"},
		"age":      []string{"digits:2", "digits:2"},
	}
	data := map[string]interface{}{
		"username": "vikram",
		"email":    "invalid",
		"age":      12,
	}
	opts := govalidator.Options{
		// Request: r,
		Data:  &data,
		Rules: rules,
		// RequiredDefault: true,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()
	err := map[string]interface{}{"validationError": e}
	fmt.Println(err)
	// w.Header().Set("Content-type", "application/json")
	// json.NewEncoder(w).Encode(err)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}
