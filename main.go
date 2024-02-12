package main

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"flag"
	"fmt"
)

var eCsv string
var ePort int

func init() {
	flag.StringVar(&eCsv, "csv", "", "Path to CSV file")
	flag.IntVar(&ePort, "port", 8080, "Port for WebServer running")
}

func main() {
	flag.Parse()
	http.HandleFunc("/grandstream.xml", grandstreamHandler)
	http.HandleFunc("/gigaset.xml", gigasetHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ePort), nil))
}

func grandstreamHandler(w http.ResponseWriter, r *http.Request) {
	generateXML(w, eCsv, "Grandstream")
}

func gigasetHandler(w http.ResponseWriter, r *http.Request) {
	generateXML(w, eCsv, "Gigaset")
}

func generateXML(w http.ResponseWriter, filename string, phoneType string) {
	// Otevřít CSV soubor
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, "Chyba při otevírání souboru", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Vytvořit čteč pro CSV soubor
	reader := csv.NewReader(file)

	// Načíst záznamy z CSV
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Chyba při čtení dat ze souboru", http.StatusInternalServerError)
		return
	}

	var xmlData []byte
	if phoneType == "Gigaset" {
		xmlData, err = xml.MarshalIndent(prepareGigasetXML(records), "", "    ")
	} else {
		xmlData, err = xml.MarshalIndent(prepareGrandstreamXML(records), "", "    ")
	}

	// Převést telefonní seznam na XML
	
	if err != nil {
		http.Error(w, "Chyba při generování XML", http.StatusInternalServerError)
		return
	}

	// Nastavit hlavičky HTTP odpovědi
	w.Header().Set("Content-Type", "application/xml")

	// Odeslat XML jako odpověď
	_, err = w.Write(xmlData)
	if err != nil {
		http.Error(w, "Chyba při odesílání XML", http.StatusInternalServerError)
		return
	}
}
