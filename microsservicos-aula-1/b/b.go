package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Result - Retorno do resultado da chamada http
type Result struct {
	Status     string
	PcDesconto int
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9091", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")

	resultCoupon := makeHTTPCall("http://localhost:9092", coupon)

	result := Result{Status: "declined", PcDesconto: resultCoupon.PcDesconto}

	if ccNumber == "1" {
		result.Status = "approved"
	}

	if resultCoupon.Status == "invalid" {
		result.Status = "invalid coupon"
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))

}

func makeHTTPCall(urlMicroservice string, coupon string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Servidor fora do ar!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}
