package types

// Money представляет собой денежную единицу в минимальных единицах
// (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

// Category представляет собой категорю платежа
type Category string

// Status представляет собой статус платежа
type Status string


// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Card предствляет информацию о платежной карте
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money // Used Money type
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

const (
	StatusOK Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID 		int
	Amount Money
	Category Category
	Status Status
}