package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Result - Estrutura de retorno da funcao
type Result struct {
	Status     string
	PcDesconto int
}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")

	var pcDesc = 0
	if coupon == "sub" {
		pcDesc = 15
	} else {
		pcDesc = 12

	}

	result := Result{Status: "valid", PcDesconto: pcDesc}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
