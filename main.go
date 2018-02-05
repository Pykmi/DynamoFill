package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// globals
var db *dynamodb.DynamoDB
var table string

func main() {
	// read cli arguments
	tableName := flag.String("table-name", "", "Table name")
	file := flag.String("file", "MOCK_DATA.json", "JSON file name")
	region := flag.String("region", "", "AWS server region")

	flag.Parse()

	table = deref(tableName)

	// read data from json file
	items, err := readJson(deref(file))
	if err != nil {
		stop(err)
	}

	// connect to database
	err = Connect(region)
	if err != nil {
		stop(err)
	}

	// add items to the database
	var n = 0

	for i := range items {
		err = Add(items[i])
		if err != nil {
			stop(err)
		}

		n++
	}

	fmt.Printf("\nAdded %v item total\n", n)
}

// add one item to the database
func Add(item Movie) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(strconv.Itoa(item.ID)),
			},
			"name": {
				S: aws.String(item.Name),
			},
			"genre": {
				S: aws.String(item.Genre),
			},
			"year": {
				S: aws.String(strconv.Itoa(item.Year)),
			},
		},
		ReturnConsumedCapacity: aws.String("NONE"),
		TableName: aws.String(table),
	}

	_, err := db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

// connect to the database
func Connect(region *string) error {
	s, err := session.NewSession(&aws.Config{
		Region: region},
	)

	if err != nil {
		return err
	}

	db = dynamodb.New(s)
	return nil
}

// read json data file
func readJson(file string) (List, error) {
	fh, err := ioutil.ReadFile( file )
	if err != nil {
		return nil, err
	}

	var items List

	err = json.Unmarshal(fh, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// dereference string pointer
func deref(p *string) string {
	return *p
}

// stop execution and report error message
func stop(err error) {
	fmt.Println(err)
	os.Exit(1)
}

