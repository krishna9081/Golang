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

func (p person) sayHello() {
	fmt.Println("Namaste")
}

func (sa secretAgent) sayHello() {
	fmt.Println("I'm in spain,Let me say Hola")
}

type greeting interface {
	sayHello()
}

func vomit(g greeting) {
	switch v := g.(type) {
	case person:
		fmt.Println("I'm", v.fname+" "+v.lname)
	case secretAgent:
		v.sayHello()
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	p := person{"Shahurukh", "Khan"}
	sa := secretAgent{person{"James", "Bond"}, true}

	vomit(p)
	vomit(sa)

}
