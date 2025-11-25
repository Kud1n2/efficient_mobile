package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func CheckAvailable() URLresponse {
	dict := make(map[string]string)
	if len(URLrequests) > 0 {
		used := URLrequests[len(URLrequests)-1]
		for i := 0; i < len(used.Links); i++ {
			_, err := net.Dial("tcp", used.Links[i]+":http")
			if err != nil {
				dict[used.Links[i]] = "not available"
			} else {
				dict[used.Links[i]] = "available"
			}

		}
		URLresponses = append(URLresponses, URLresponse{Links: dict, Links_num: len(URLresponses) + 1})
	}
	return URLresponse{Links: dict, Links_num: len(URLrequests)}
}

func makePDF(list []int) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	for i := 0; i < len(URLresponses); i++ {
		for _, j := range list {
			if URLresponses[i].Links_num == j {
				for key, val := range URLresponses[i].Links {
					str := key + ":" + val
					pdf.Cell(0, 0, str)
					pdf.Ln(10)
				}
			}
		}
	}
	return pdf.OutputFileAndClose("file.pdf")
}

func writeToLog(resp URLresponse) {
	file, err := os.ReadFile("log.json")
	var resps []URLresponse
	if err != nil {
		resps = []URLresponse{}
	} else {
		err = json.Unmarshal(file, &resps)
		if err != nil {
			fmt.Println("Ошибка чтения")
		}
	}

	resps = append(resps, resp)

	jsonData, err := json.Marshal(resps)
	if err != nil {
		log.Fatal("Ошибка маршалинга")
	}

	err = os.WriteFile("log.json", jsonData, 0644)
	if err != nil {
		log.Fatal("Ошибка записи")
	}
}
