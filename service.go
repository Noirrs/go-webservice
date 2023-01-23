package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

func AddWord(word Word, cc *mongo.Collection, category string) error {

	var decodedCategory Category
	filter := bson.D{{Key: "name", Value: category}}

	err := cc.FindOne(context.TODO(), filter).Decode(&decodedCategory)

	if err != nil && err == mongo.ErrNoDocuments {

		var category Category = Category{Name: category, Words: []Word{word}}
		cc.InsertOne(context.TODO(), category)

	} else if err == nil {

		var WordNames []string

		for _, ch := range decodedCategory.Words {
			WordNames = append(WordNames, ch.Value)
		}

		if slices.Contains(WordNames, word.Value) {
			return errors.New("word already exists")
		}

		newWords := append(decodedCategory.Words, word)

		res := cc.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "words": newWords})

		if res.Err() != nil {
			log.Fatal("unknown err occurred ", res.Err())
			return res.Err()
		}

	}

	return nil
}

func Corrector(word Word, cc *mongo.Collection, category string) error {

	var decodedCategory Category

	ctg := checkErr(cc, category)
	fmt.Println(ctg)
	typer := reflect.TypeOf(ctg).String()

	if typer == "main.Category" {
		decodedCategory = ctg.(Category)
		var foundedWord Word
		fmt.Println(decodedCategory.Words)
		for i, ch := range decodedCategory.Words {
			if ch.Value == word.Value {
				foundedWord = ch
				decodedCategory.Words = append(decodedCategory.Words[:i], decodedCategory.Words[i+1:]...)
			}
		}

		if (foundedWord == Word{}) {
			return errors.New("couldn't find the word")
		}

		foundedWord.Trued += word.Trued
		foundedWord.Falsed += word.Falsed

		decodedCategory.Words = append(decodedCategory.Words, foundedWord)

		cc.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "words": decodedCategory.Words})
	} else {
		return errors.New("couldn't find the category")
	}
	return nil
}

func Edit(word string, newWordValue string, cc *mongo.Collection, category string) error {

	var decodedCategory Category

	ctg := checkErr(cc, category)
	typer := reflect.TypeOf(ctg).String()

	if typer == "main.Category" {
		decodedCategory = ctg.(Category)

		var foundedWord Word

		for i, ch := range decodedCategory.Words {
			if ch.Value == word {
				foundedWord = ch
				decodedCategory.Words = append(decodedCategory.Words[:i], decodedCategory.Words[i+1:]...)
			}
		}

		if (foundedWord == Word{}) {
			return errors.New("couldn't find the word")
		}

		foundedWord.Value = newWordValue

		decodedCategory.Words = append(decodedCategory.Words, foundedWord)

		cc.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "words": decodedCategory.Words})
	} else {
		return errors.New("couldn't find the category")
	}
	return nil
}

func Delete(word string, cc *mongo.Collection, category string) error {

	var decodedCategory Category

	ctg := checkErr(cc, category)
	typer := reflect.TypeOf(ctg).String()
	if typer == "main.Category" {
		decodedCategory = ctg.(Category)
		var foundedWord Word

		for i, ch := range decodedCategory.Words {
			if ch.Value == word {
				foundedWord = ch
				decodedCategory.Words = append(decodedCategory.Words[:i], decodedCategory.Words[i+1:]...)
			}
		}

		if (foundedWord == Word{}) {
			return errors.New("couldn't find the word")
		}
		if len(decodedCategory.Words) == 0 {
			cc.DeleteOne(context.TODO(), bson.M{"name": category})
		} else {
			cc.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "words": decodedCategory.Words})
		}
	} else {
		return errors.New("couldn't find the category")
	}
	return nil
}

func checkErr(cc *mongo.Collection, category string) any {
	var decodedCategory Category

	fmt.Println(category)
	filter := bson.D{{Key: "name", Value: category}}

	err := cc.FindOne(context.TODO(), filter).Decode(&decodedCategory)

	if err != nil && err == mongo.ErrNoDocuments {
		return errors.New("couldn't find the category")
	}

	return decodedCategory
}
