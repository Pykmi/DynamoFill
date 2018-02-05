package main

type Movie struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Genre 	string 	`json:"genre"`
	Year 	int 	`json:"year"`
}

type List []Movie
