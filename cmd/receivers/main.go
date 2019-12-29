package main

import (
	"alertmanager-receivers/alerts"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)



func main() {
	//if err := alerts.SendMail("subject chinese", "body chinese"); err != nil {
	//	log.Fatalln(err)
	//}
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		alertRecord := alerts.AlertRecord{}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		if err := json.Unmarshal(b, &alertRecord); err != nil {
			panic(err)
		}
		//log.Println(string(b))
		//log.Println(alertRecord)

		if err = alerts.SendAlert(alertRecord); err != nil {
			log.Fatalln(err)
		}
	})))
}