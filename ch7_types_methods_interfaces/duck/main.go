package main

type LogicProvider struct{}

func (l LogicProvider) Process(data string) string {
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.Process("...data..")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
