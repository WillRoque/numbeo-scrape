package numbeo

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type column int

const (
	city column = iota + 1
	country
	netSalary
	monthlyCost
	smallAptCentre
	smallAptOutside
	largeAptCentre
	largeAptOutside
	moneyAfterMonthlyCost
	moneyAfterSmallAptCentre
	moneyAfterSmallAptOutside
	moneyAfterLargeAptCentre
	moneyAfterLargeAptOutside
	moneyAfterSmallAptCentreAndMonthlyCost
	moneyAfterSmallAptOutsideAndMonthlyCost
	moneyAfterLargeAptCentreAndMonthlyCost
	moneyAfterLargeAptOutsideAndMonthlyCost
)

var columnTitleMap = map[column]string{
	city:                                    "City",
	country:                                 "Country",
	netSalary:                               "Net Salary",
	monthlyCost:                             "Monthly Cost (ML)",
	smallAptCentre:                          "1 Bed Apt Centre",
	smallAptOutside:                         "1 Bed Apt Outside",
	largeAptCentre:                          "3 Bed Apt Centre",
	largeAptOutside:                         "3 Bed Apt Outside",
	moneyAfterMonthlyCost:                   "Rem. After ML",
	moneyAfterSmallAptCentre:                "Rem. After 1 Bed Apt Centre",
	moneyAfterSmallAptOutside:               "Rem. After 1 Bed Apt Outside",
	moneyAfterLargeAptCentre:                "Rem. After 3 Bed Apt Centre",
	moneyAfterLargeAptOutside:               "Rem. After 3 Bed Apt Outside",
	moneyAfterSmallAptCentreAndMonthlyCost:  "Rem. After 1 Bed Apt Centre + ML",
	moneyAfterSmallAptOutsideAndMonthlyCost: "Rem. After 1 Bed Apt Outside + ML",
	moneyAfterLargeAptCentreAndMonthlyCost:  "Rem. After 3 Bed Apt Centre + ML",
	moneyAfterLargeAptOutsideAndMonthlyCost: "Rem. After 3 Bed Apt Outside + ML",
}

func (c column) string() string {
	return columnTitleMap[c]
}

var columnNameMap = map[column]string{
	city:                                    "A",
	country:                                 "B",
	netSalary:                               "C",
	monthlyCost:                             "D",
	smallAptCentre:                          "E",
	smallAptOutside:                         "F",
	largeAptCentre:                          "G",
	largeAptOutside:                         "H",
	moneyAfterMonthlyCost:                   "I",
	moneyAfterSmallAptCentre:                "J",
	moneyAfterSmallAptOutside:               "K",
	moneyAfterLargeAptCentre:                "L",
	moneyAfterLargeAptOutside:               "M",
	moneyAfterSmallAptCentreAndMonthlyCost:  "N",
	moneyAfterSmallAptOutsideAndMonthlyCost: "O",
	moneyAfterLargeAptCentreAndMonthlyCost:  "P",
	moneyAfterLargeAptOutsideAndMonthlyCost: "Q",
}

// getColName returns the column name in the spreadsheet.
// Example: column "city" is A, as defined in columnNameMap.
func (c column) getColName() string {
	return columnNameMap[c]
}

// getCellName returns the cell name of a column in the specified row.
// Example: column "city", row 2, will result in A2.
func (c column) getCellName(row int) string {
	return fmt.Sprintf("%s%d", c.getColName(), row)
}

// GenerateSpreadsheet generates a spreadsheet with the data contained in CityInfo.
func GenerateSpreadsheet(cis []CityInfo) excelize.File {
	xlsx := excelize.NewFile()
	sheet := "Remaining Money"
	xlsx.SetSheetName("Sheet1", sheet)

	// Set columns width
	xlsx.SetColWidth(sheet, city.getColName(), netSalary.getColName(), 15)
	xlsx.SetColWidth(sheet, monthlyCost.getColName(), monthlyCost.getColName(), 17)
	xlsx.SetColWidth(sheet, smallAptCentre.getColName(), largeAptOutside.getColName(), 18)
	xlsx.SetColWidth(sheet, moneyAfterMonthlyCost.getColName(), moneyAfterMonthlyCost.getColName(), 15)
	xlsx.SetColWidth(sheet, moneyAfterSmallAptCentre.getColName(), moneyAfterLargeAptOutside.getColName(), 28)
	xlsx.SetColWidth(sheet, moneyAfterSmallAptCentreAndMonthlyCost.getColName(), moneyAfterLargeAptOutsideAndMonthlyCost.getColName(), 32)

	// Insert columns titles
	xlsx.SetCellValue(sheet, city.getCellName(1), city.string())
	xlsx.SetCellValue(sheet, country.getCellName(1), country.string())
	xlsx.SetCellValue(sheet, netSalary.getCellName(1), netSalary.string())
	xlsx.SetCellValue(sheet, monthlyCost.getCellName(1), monthlyCost.string())
	xlsx.SetCellValue(sheet, smallAptCentre.getCellName(1), smallAptCentre.string())
	xlsx.SetCellValue(sheet, smallAptOutside.getCellName(1), smallAptOutside.string())
	xlsx.SetCellValue(sheet, largeAptCentre.getCellName(1), largeAptCentre.string())
	xlsx.SetCellValue(sheet, largeAptOutside.getCellName(1), largeAptOutside.string())
	xlsx.SetCellValue(sheet, moneyAfterMonthlyCost.getCellName(1), moneyAfterMonthlyCost.string())
	xlsx.SetCellValue(sheet, moneyAfterSmallAptCentre.getCellName(1), moneyAfterSmallAptCentre.string())
	xlsx.SetCellValue(sheet, moneyAfterSmallAptOutside.getCellName(1), moneyAfterSmallAptOutside.string())
	xlsx.SetCellValue(sheet, moneyAfterLargeAptCentre.getCellName(1), moneyAfterLargeAptCentre.string())
	xlsx.SetCellValue(sheet, moneyAfterLargeAptOutside.getCellName(1), moneyAfterLargeAptOutside.string())
	xlsx.SetCellValue(sheet, moneyAfterSmallAptCentreAndMonthlyCost.getCellName(1), moneyAfterSmallAptCentreAndMonthlyCost.string())
	xlsx.SetCellValue(sheet, moneyAfterSmallAptOutsideAndMonthlyCost.getCellName(1), moneyAfterSmallAptOutsideAndMonthlyCost.string())
	xlsx.SetCellValue(sheet, moneyAfterLargeAptCentreAndMonthlyCost.getCellName(1), moneyAfterLargeAptCentreAndMonthlyCost.string())
	xlsx.SetCellValue(sheet, moneyAfterLargeAptOutsideAndMonthlyCost.getCellName(1), moneyAfterLargeAptOutsideAndMonthlyCost.string())

	// Insert values
	for i, ci := range cis {
		xlsx.SetCellValue(sheet, city.getCellName(i+2), ci.City)
		xlsx.SetCellValue(sheet, country.getCellName(i+2), ci.Country)
		xlsx.SetCellValue(sheet, netSalary.getCellName(i+2), ci.NetSalary)
		xlsx.SetCellValue(sheet, monthlyCost.getCellName(i+2), ci.MonthlyCost)
		xlsx.SetCellValue(sheet, smallAptCentre.getCellName(i+2), ci.SmallAptCentre)
		xlsx.SetCellValue(sheet, smallAptOutside.getCellName(i+2), ci.SmallAptOutside)
		xlsx.SetCellValue(sheet, largeAptCentre.getCellName(i+2), ci.LargeAptCentre)
		xlsx.SetCellValue(sheet, largeAptOutside.getCellName(i+2), ci.LargeAptOutside)
		xlsx.SetCellValue(sheet, moneyAfterMonthlyCost.getCellName(i+2), ci.MoneyAfterMonthlyCost)
		xlsx.SetCellValue(sheet, moneyAfterSmallAptCentre.getCellName(i+2), ci.MoneyAfterSmallAptCentre)
		xlsx.SetCellValue(sheet, moneyAfterSmallAptOutside.getCellName(i+2), ci.MoneyAfterSmallAptOutside)
		xlsx.SetCellValue(sheet, moneyAfterLargeAptCentre.getCellName(i+2), ci.MoneyAfterLargeAptCentre)
		xlsx.SetCellValue(sheet, moneyAfterLargeAptOutside.getCellName(i+2), ci.MoneyAfterLargeAptOutside)
		xlsx.SetCellValue(sheet, moneyAfterSmallAptCentreAndMonthlyCost.getCellName(i+2), ci.MoneyAfterSmallAptCentreAndMonthlyCost)
		xlsx.SetCellValue(sheet, moneyAfterSmallAptOutsideAndMonthlyCost.getCellName(i+2), ci.MoneyAfterSmallAptOutsideAndMonthlyCost)
		xlsx.SetCellValue(sheet, moneyAfterLargeAptCentreAndMonthlyCost.getCellName(i+2), ci.MoneyAfterLargeAptCentreAndMonthlyCost)
		xlsx.SetCellValue(sheet, moneyAfterLargeAptOutsideAndMonthlyCost.getCellName(i+2), ci.MoneyAfterLargeAptOutsideAndMonthlyCost)
	}

	return *xlsx
}
