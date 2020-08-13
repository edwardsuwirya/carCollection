package entity

type Car struct {
	CarDetail CarDetail
}
type CarDetail struct {
	Id       int
	Car      string
	CarModel string
	CarColor string
}
