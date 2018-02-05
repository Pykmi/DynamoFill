# DynamoFill

Small helper to fill dynamo tables with mock data. Just change the Movie and List structs in types.go, as well as AttributeValues in Add() function to match your data.

*types.go*
```go
type Movie struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Genre 	string 	`json:"genre"`
	Year 	int 	`json:"year"`
}

type List []Movie
```

*main.go*
```go
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
```