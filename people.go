package main

import "encoding/xml"

type People struct {
	XMLName xml.Name `xml:"people"`
	Persons []Person `xml:"person"`
}

type Person struct {
	Id           int    `xml:"id" bson:"id,omitempty"`
	First_name   string `xml:"first_name" bson:"first_name,omitempty"`
	Last_name    string `xml:"last_name" bson:"last_name,omitempty"`
	Company      string `xml:"company" bson:"company,omitempty"`
	Email        string `xml:"email" bson:"email,omitempty"`
	Ip_address   string `xml:"ip_address" bson:"ip_address,omitempty"`
	Phone_number string `xml:"phone_number" bson:"phone_number,omitempty"`
}
