package composite

import "fmt"

type Parent struct{

}

func (p Parent) Name() string {
	return "parent"
}

func (p Parent) SayHello() {
	fmt.Println("Hello,", p.Name())
}

type Sub struct{
	Parent
}


func (s Sub) Name()  string{
	return "sub"
}	


