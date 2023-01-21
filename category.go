package main



type Char struct {
	Value   string  `json:"value" bson:"value"`
	Trued   int    `json:"trued" bson:"trued"`
	Falsed 	int    `json:"falsed" bson:"falsed"`
}

type Category struct {
	Name string `json:"name" bson:"name"`
	Chars []Char `json:"chars" bson:"chars"`
}