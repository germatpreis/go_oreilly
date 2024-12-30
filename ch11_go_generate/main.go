package main

import (
	"ch11_go_generate/data"
	"fmt"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=ch11_go_generate --go_opt=Mperson.proto=ch11_go_generate/data person.proto

type Direction int

const (
	_ Direction = iota
	North
	South
	East
	West
)

//go:generate stringer -type=Direction

func main() {
	// proto
	p := &data.Person{
		Name:  "Bob Bobson",
		Id:    20,
		Email: "bob@bobson.com",
	}
	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)

	// stringer
	fmt.Println(North.String())
}
