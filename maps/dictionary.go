// Package main
package main

import (
	"fmt"
)

const (
	ErrInvalidKey       = DictionaryErr("invalid key")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]

	if !ok {
		return "", ErrInvalidKey
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrInvalidKey:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrInvalidKey:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrInvalidKey:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func main() {
	dictionary := Dictionary{
		"Hana":   "Fire",
		"Akira":  "Wind",
		"Hisoka": "Lightning",
		"Mori":   "Water",
	}

	fmt.Println(dictionary.Search("Hana"))   // Fire
	fmt.Println(dictionary.Search("Hisoka")) // Lightning
	fmt.Println(dictionary.Search("Masaru")) // Key doesn't exist

	// length
	fmt.Println(len(dictionary)) // 4

	// delete method
	delete(dictionary, "Mori")
	fmt.Printf("After deleting Mori: %v\n", dictionary)

	dictionary["Shingen"] = "Wind, Earth"

	// Looping
	for key, value := range dictionary {
		fmt.Printf("key: %s, Value: %s\n", key, value)
	}
}
