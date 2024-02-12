package main

import (
	"encoding/xml"
)

type GigasetPhonebook struct {
	XMLName xml.Name  `xml:"phonebooks"`
	Phonebook []GigasetContact `xml:"phonebook"`
}

type GigasetContact struct {
	XMLName xml.Name `xml:"contact"`
	Person  GigasetPerson   `xml:"person"`
	Telephony GigasetTelephony `xml:"telephony"`
}

type GigasetPerson struct {
	RealName string `xml:"realName"`
}

type GigasetTelephony struct {
	Number string `xml:"number,attr"`
}


func prepareGigasetXML(records [][]string) GigasetPhonebook{
	
	phonebook := GigasetPhonebook{}
	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		contact := GigasetContact{
			Person: GigasetPerson{
				RealName: record[0],
				},
				Telephony: GigasetTelephony{
				Number: record[1],
				},
				}
				phonebook.Phonebook = append(phonebook.Phonebook, contact)
	}
	
	return phonebook
}
