package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Country struct {

	// defining struct variables
	Name      string
	Capital   string
	Continent string
}

type Organizacion struct {
	Name string `json:"organizacion"`
	User User
	Rol  Rol
}

type User struct {
	UserName string `json:"username"`
}
type Rol struct {
	RolName string `json:"rolname"`
}

type Userss struct {
	Namess string
}

func main() {
	records, err := readData("data.csv")
	org := Organizacion{}
	var result []Organizacion

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		if record[0] == "org1" {
			org = Organizacion{
				Name: record[0],
				User: User{
					record[1],
				},
				Rol: Rol{
					record[2],
				},
			}
			result = addOgr(&org)

		}
		if record[0] == "org2" {
			org = Organizacion{
				Name: record[0],
				User: User{
					record[1],
				},
				Rol: Rol{
					record[2],
				},
			}
			result = addOgr(&org)
		}

		e, _ := json.Marshal(result)

		fmt.Println(string(e))
	}

}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func addOgr(org *Organizacion) []Organizacion {
	var results []Organizacion

	results = append(results, *org)

	return results
}
