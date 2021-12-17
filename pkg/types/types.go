package types

type Money int64

type Category string

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "Fail"
	StatusInProgress Status = "INPROGRESS"
)
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}