package serialization

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Plant struct {
	XmlName xml.Name `xml:"xml_name,omitempty"`
	Id      int      `xml:"id,omitempty"`
	Name    string   `xml:"name,omitempty"`
	Origin   []string `xml:"orgin,omitempty"`
}

func (p Plant) String() string {
	return fmt.Sprintf("plant id=%d, name=%s, origin=%v", p.Id, p.Name, p.Origin)
}

func Test_xml(t *testing.T) {
	coffee:= &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin =[]string{"Ethiopia", "Brazil"}

	out, _ :=xml.Marshal(coffee)
	fmt.Printf("out: %s\n", out)

	fmt.Printf("xml header: %s , out: %s\n", xml.Header, out)


	var p Plant
	if err:=xml.Unmarshal(out , &p) ;err!= nil {
		panic(err)
	}
	fmt.Printf("p: %+v\n", p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin =[]string{"Mexico", "California"}

	type Nesting struct {
		XMLlName xml.Name `xml:"nesting" `
		Plants   []*Plant `xml:"plants"`
	}

	nesting :=&Nesting{}
	nesting.Plants = []*Plant{coffee ,tomato}
	out, _= xml.Marshal(nesting)
	fmt.Printf("out: %s\n", out)
}
