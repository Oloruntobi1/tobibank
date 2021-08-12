package TobiBank


type Response struct {
	Balance int
}

type TobiBank struct {
	ApiKey string
}

func WireUpTobiBank() *TobiBank {
	return &TobiBank{}
}

func(w *TobiBank) Create() string {
	return "Created by Wallet africa"
}

func(m *TobiBank) Payout() Response {
	return Response{}
}

