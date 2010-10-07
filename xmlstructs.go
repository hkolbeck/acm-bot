//Copyright 2010 Cory Kolbeck <ckolbeck@gmail.com>.
//So long as this notice remains in place, you are welcome 
//to do whatever you like to or with this code.  This code is 
//provided 'As-Is' with no warrenty expressed or implied. 
//If you like it, and we happen to meet, buy me a beer sometime

package main

//Calendar feed structs
type calEntry struct {
	Title string
	Content string
}

type calFeed struct {
	TotalResults int
	Entry []calEntry
}

//Dictionary structs
type dictWordnet struct {
	Pos []dictPos
}

type dictPos struct {
	Name string "attr"
	Category dictCategory
}

type dictCategory struct {
	Sense []dictSense
}

type dictSense struct{
	Synset struct {Word []string; Definition, Sample string}
}

//Wiki struct
type wiki struct {
	Page struct {
		Title string;
		Body struct{Text string}
	}
}
