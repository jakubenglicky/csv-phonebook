package main

import "encoding/xml"

type GrandstreamPhonebook struct {
	XMLName xml.Name  `xml:"AddressBook"`
	Contacts []GrandstreamContact `xml:"Contact"`
}

type GrandstreamContact struct {
	LastName  string `xml:"LastName"`
	Phone1    string `xml:"Phone1"`
}

func prepareGrandstreamXML(records [][]string) GrandstreamPhonebook{
	
	// Vytvořit strukturu pro telefonní seznam
	phonebook := GrandstreamPhonebook{}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		contact := GrandstreamContact{
			LastName:  record[0],
			Phone1:    record[1],
			}
			phonebook.Contacts = append(phonebook.Contacts, contact)
	}

	return phonebook
}
