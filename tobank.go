package tobibank

import (
	"github.com/Oloruntobi1/payload"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type TobiBank struct {
	ApiKey string
}

func WireUpTobiBank() *TobiBank {
	return &TobiBank{}
}

func(w *TobiBank) Create() string {
	return "Created by Wallet africa"
}

type Result struct {
	UserId int `json:"userId"`
}

func(m *TobiBank) Payout() payload.WalletPayoutResponse{

	var res Result

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		log.Fatalln(err)
	}

	return payload.WalletPayoutResponse{
		Balance: res.UserId,
	}
}

// func(m *TobiBank) Payout() string{
// 	return "my head oo"
// }

// git commit -a -m "my new version"
// git push
// git tag v1.1.5
// git push -q origin v1.1.5


// creating v2
// git checkout -b v2
// git push --set-upstream origin v2
// git commit -a -m "version 2"
// git tag v2.0.0
// git push --tags origin v2


//git commit -a -m "v2.1.0"
// git push
// git tag v2.1.0
//git push -q origin v.2.1.0

