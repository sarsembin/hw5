package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

var rawXml = []byte(`<root>
  <user>
    <id>1</id>
    <address>
      <city_id>5</city_id>
      <street>Satbayev</street>
    </address>
    <age>20</age>
  </user>
  <user>
    <id>1</id>
    <address>
      <city_id>6</city_id>
      <street>Al-Farabi</street>
    </address>
    <age>32</age>
  </user>
</root>`)


func main() {
	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		panic(err)
	}


	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	var xmlUsers Users
	if err := xml.Unmarshal(rawXml, &xmlUsers); err != nil {
		panic(err)
	}
	fmt.Println("XML users")
	for _, xmlUsers := range xmlUsers.Users {
		fmt.Printf("%#v\n", xmlUsers)
	}
}

type Users struct {
	Users []User	`xml:"user"`
}

type User struct {
	ID      CustomFloat64   `json:"id" xml:"id"`
	Address Address 		`json:"address" xml:"address"`
	Age     CustomFloat64	`json:"age" xml:"age"`
}

type Address struct {
	CityID CustomFloat64 	`json:"city_id" xml:"city_id"`
	Street string			`json:"street" xml:"street"`
}

type CustomFloat64 struct{
	Float64 float64
}

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil

}

func (cf *CustomFloat64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	(*cf).Float64, err = strconv.ParseFloat(s, 64)
	if err != nil{
		return err
	}
	return nil
}
