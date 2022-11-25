package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("My name is", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Println(`I don't want to tell openly my LastName is `, sa.lname, `and my fistname is  `, `sa.fname`)
}

func main() {
	p := person{"Shahurukh", "Khan"}
	sa := secretAgent{person{"James", "Bond"}, true}

	fmt.Println(p.fname)
	p.pSpeak()
	fmt.Println(sa.lname)
	sa.saSpeak()
	sa.pSpeak()
}
