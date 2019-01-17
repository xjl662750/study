package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var a float64 = 3000
	var b float64 = 1800
	var c float64 = 500
	dist := strconv.FormatFloat((a-b)/c, 'f', 1, 64)
	distint := float64(int64((a-b)/c) + 2)
	distfloat64, _ := strconv.ParseFloat("0", 64)
	fmt.Println(dist)
	fmt.Println(distint)
	fmt.Println(distfloat64)

	xlsx := excelize.NewFile()
	// Create a new sheet.
	// index := xlsx.NewSheet("Sheet2")
	// Set active sheet of the workbook.
	// xlsx.SetActiveSheet(index)
	// Set value of a cell.
	xlsx.SetCellValue("Sheet1", "A1", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B1", "100")
	for i := 2; i < 10; i++ {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), i)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), i)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), i+i)
	}

	//读
	// rows := xlsx.GetRows("Sheet1")
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		// colCell = i
	// 		fmt.Println(colCell)
	// 	}
	// }

	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
