// simple program that defines interface Home, and structs House and Semi, with typed reciver methods

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type roomSz struct {
	width  float32
	length float32
}

type House struct {
	name     string
	rooms    []string
	roomsize []roomSz
}

type Home interface {
	inputSqft()
	printMetric()
}


// methods to create new House objects
func NewHouse() *House {
	r := []string{"kitchen", "living", "dining", "main"}
	rs := make([]roomSz, len(r))
	h := House{"house", r, rs}
	return &h
}

func NewHouseRooms(otherrooms []string) *House {
	r := append([]string{"kitchen", "living", "dining", "main"}, otherrooms...)
	rs := make([]roomSz, len(r))
	h := House{"house", r, rs}
	return &h
}


// methods of interface Home for struct House
func (p *House) inputSqft() {
	for i := 0; i < len(p.roomsize); i++ {
		var inStr string
		fmt.Print(p.rooms[i] + " : width x length: ")
		fmt.Scanf("%s", &inStr)
		s := strings.Split(inStr, "x")
		w, _ := strconv.ParseFloat(s[0], 32)
		l, _ := strconv.ParseFloat(s[1], 32)
		p.roomsize[i] = roomSz{width: float32(w), length: float32(l)}
	}
	fmt.Println()
}

func (p *House) printMetric() {
	fmt.Println(p.name)
	for i := 0; i < len(p.rooms); i++ {
		w := p.roomsize[i].width / 3.2808
		l := p.roomsize[i].length / 3.2808
		fmt.Print(p.rooms[i] + " : ")
		fmt.Printf("%.2f", w)
		fmt.Print("x")
		fmt.Printf("%.2f", l)
		fmt.Println(" m")
	}
	fmt.Println()
}


// subclass of House
type Semi struct {
	House
}

// methods to create new Semi objects
func NewSemi() *Semi {
	r := []string{"kitchen", "living", "dining", "main"}
	rs := make([]roomSz, len(r))
	s := Semi{House{"Semi", r, rs}}
	return &s
}

func NewSemiRooms(otherrooms []string) *Semi {
	r := append([]string{"kitchen", "living", "dining", "main"}, otherrooms...)
	rs := make([]roomSz, len(r))
	s := Semi{House{"Semi", r, rs}}
	return &s
}

// methods of interface Home for struct Semi
func (p *Semi) inputSqft() {
	for i := 0; i < len(p.House.roomsize); i++ {
		var inStr string
		fmt.Print(p.House.rooms[i] + " : width x length: ")
		fmt.Scanf("%s", &inStr)
		s := strings.Split(inStr, "x")
		w, _ := strconv.ParseFloat(s[0], 32)
		l, _ := strconv.ParseFloat(s[1], 32)
		p.House.roomsize[i] = roomSz{width: float32(w), length: float32(l)}
	}
	fmt.Println()
}

func (p *Semi) printMetric() {
	fmt.Println(p.House.name)
	for i := 0; i < len(p.rooms); i++ {
		w := p.House.roomsize[i].width / 3.2808
		l := p.House.roomsize[i].length / 3.2808
		fmt.Print(p.House.rooms[i] + " : ")
		fmt.Printf("%.2f", w)
		fmt.Print("x")
		fmt.Printf("%.2f", l)
		fmt.Println(" m")
	}
	fmt.Println()
}

func main() {

	// creating test objects 
	homes := []Home{NewHouse(), NewSemi(), NewHouseRooms([]string{"bedroom1", "bedroom2"})}

	for i := 0; i < len(homes); i++ {
		homes[i].inputSqft()
	}
	for i := 0; i < len(homes); i++ {
		homes[i].printMetric()
	}

}
