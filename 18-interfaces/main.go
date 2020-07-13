package main

import "fmt"

// Entertainer represents any kind of enterainer
type Entertainer interface {
	Play()
}

// Beatle represents music grıup
type Beatle struct {
	Name       string
	Instrument string
}

// AnadoluAtesi a dance group
type AnadoluAtesi struct {
	Name      string
	DanceType string
}

// Play for Beatle tpye
func (b Beatle) Play() {
	fmt.Printf("%s plays %s\n", b.Name, b.Instrument)
}

// Play for AnadoluAtesi type
func (a AnadoluAtesi) Play() {
	fmt.Printf("%s plays %s\n", a.Name, a.DanceType)
}

func main() {
	b := Beatle{Name: "Ringo", Instrument: "Drums"}
	a := AnadoluAtesi{Name: "Anadolu Atesi", DanceType: "Halk Oyunları"}
	entertain(b)
	entertain(a)
}

func entertain(e Entertainer) {
	e.Play()
}
