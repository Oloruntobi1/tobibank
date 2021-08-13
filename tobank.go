package tobibank

import "github.com/Oloruntobi1/payload"

type TobiBank struct {
	ApiKey string
}

func WireUpTobiBank() *TobiBank {
	return &TobiBank{}
}

func(w *TobiBank) Create() string {
	return "Created by Wallet africa"
}

func(m *TobiBank) Payout() payload.WalletPayoutResponse{
	return payload.WalletPayoutResponse{}
}

// func(m *TobiBank) Payout() string{
// 	return "my head oo"
// }

// git commit -a -m "my new version"
// git push
// git tag v1.1.3
// git push -q origin v1.1.3


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

