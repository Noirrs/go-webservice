package main

type Word struct {
	Value  string `json:"value" bson:"value"`
	Trued  int    `json:"trued" bson:"trued"`
	Falsed int    `json:"falsed" bson:"falsed"`
}

type Category struct {
	Name  string `json:"name" bson:"name"`
	Words []Word `json:"words" bson:"words"`
}
