package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(n string) string

}

type Italian struct {
	languageName string
}

type Portuguese struct {
	languageName string

}

func (i Italian) LanguageName() string{
	i.languageName = "Italian"
	return i.languageName
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", i.LanguageName(), name)
}

func (p Portuguese) LanguageName() string{
	p.languageName = "Portuguese"
	return p.languageName
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", p.LanguageName(), name)
}

func SayHello(n string, g Greeter) string{
	return g.Greet(n)

}


