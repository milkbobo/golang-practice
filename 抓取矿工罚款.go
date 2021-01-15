package main

import (
	"encoding/json"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	//"fmt"
	//"os"

	"github.com/tealeg/xlsx/v3"
)

type Transfers struct {
	TotalCount int             `json:totalCount`
	Transfers  []TransfersList `json:transfers`
	Types      []string        `json:types`
}

type TransfersList struct {
	From       string  `json:from`
	Height     int     `json:height`
	Timestamp  int64   `json:timestamp`
	TimeString string  `json:-`
	To         string  `json:to`
	Type       string  `json:type`
	Value      string  `json:value`
	Message    string  `json:message`
	Fil        float64 `json:-`
}

func main() {

	lists := []TransfersList{}

	for pageIndex := 0; pageIndex < 100; pageIndex++ {
		resp, err := http.Get("https://filfox.info/api/v1/address/f07749/transfers?pageSize=100&page=" + strconv.Itoa(pageIndex))
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		result := Transfers{}
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			panic(err)
		}

		for _, transfer := range result.Transfers {
			if transfer.Type == "burn" {
				transfer.TimeString = time.Unix(transfer.Timestamp, 0).Format("2006-01-02 15:04:05")
				de, err := decimal.NewFromString(transfer.Value)
				if err != nil {
					panic(err)
				}
				transfer.Fil, _ = de.Div(decimal.NewFromInt(1000000000000000000)).Float64()
				lists = append(lists, transfer)
			}
		}

		if len(result.Transfers) < 100 {
			break
		}

		time.Sleep(200 * time.Millisecond)
		println(pageIndex)
	}
		s, _ := prettyjson.Marshal(lists)
		fmt.Printf("%+v", string(s))

	//f, err := os.Create("f07749.txt")
	//defer f.Close()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//saveData, err := json.Marshal(lists)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//_, err = f.Write(saveData)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}


	fmt.Printf("%+v",lists)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, transfersList2 := range lists {
		row := sheet.AddRow()
		row.SetHeight(20)
		a := row.WriteStruct(&transfersList2, -1)
		fmt.Println(a)
	}


	file.Save("f07749.xlsx")

}
