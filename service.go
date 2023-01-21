package main

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

func AddChar(char Char, c *mongo.Collection, category string) error {
	// findOne and findOneAndUpdate methods are not working as expected so fix them
	filter := bson.D{{Key: "name", Value: category}}
	var decodedCategory Category
	err := c.FindOne(context.TODO(), filter).Decode(&decodedCategory)

	if err != nil && err == mongo.ErrNoDocuments {

		var category Category = Category{Name: category, Chars: []Char{char}}
		c.InsertOne(context.TODO(), category)
	} else if err == nil {

		if slices.Contains(decodedCategory.Chars, char) {

			return errors.New("char already exists")
		}

		newChars := append(decodedCategory.Chars, char)

		res := c.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "chars": newChars})

		if res.Err() != nil {
			log.Fatal("unknown err occurred ", res.Err())
			return res.Err()
		}

	}

	return nil
}

func AddCorrectorFalse(char Char, c *mongo.Collection, category string) error {

	filter := bson.D{{Key: "name", Value: category}}
	var decodedCategory Category

	err := c.FindOne(context.TODO(), filter).Decode(&decodedCategory)

	if err != nil && err == mongo.ErrNoDocuments {

		return errors.New("couldn't find the category")

	} else if err == nil {
		var foundedChar Char

		for i, ch := range decodedCategory.Chars {
			if ch.Value == char.Value {
				foundedChar = ch
				decodedCategory.Chars = append(decodedCategory.Chars[:i], decodedCategory.Chars[i+1:]...)
			}
		}
		if (foundedChar == Char{}) {
			return errors.New("couldn't find the char")
		}
		foundedChar.Trued += char.Trued
		foundedChar.Falsed += char.Falsed

		decodedCategory.Chars = append(decodedCategory.Chars, foundedChar)

		c.FindOneAndReplace(context.TODO(), bson.M{"name": category}, bson.M{"name": category, "chars": decodedCategory.Chars})

	}

	return nil
}

type TypedChar struct {
	collection *mongo.Collection
}
