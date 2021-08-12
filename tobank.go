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

