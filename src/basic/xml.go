package basic

import (
	"encoding/xml"
	"fmt"
)

type MyStruct struct {
	Field1 string `xml:"element1"`
	Field2 int    `xml:"element2"`
}

func main18() {
	data := MyStruct{Field1: "value1", Field2: 42}
	xmlData, err := xml.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(xmlData))

	var newData MyStruct
	err = xml.Unmarshal(xmlData, &newData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newData)
}
