package numbeo

import "testing"

func TestGenerateSpreadsheet(t *testing.T) {
	basel := CityInfo{City: "Basel", Country: "Switzerland"}
	munich := CityInfo{City: "Munich", Country: "Germany"}
	cases := make([]CityInfo, 2)
	cases[0] = basel
	cases[1] = munich

	spreadSheet := GenerateSpreadsheet(cases)
	sheetName := "Remaining Money"

	// Check the spreadsheet header (first row)
	for i := city; i <= moneyAfterLargeAptOutsideAndMonthlyCost; i++ {
		cellName := column(i).getCellName(1)
		got := spreadSheet.GetCellValue(sheetName, cellName)
		want := columnTitleMap[column(i)]
		if got != want {
			t.Errorf("Spreadsheet header is incorrect. At cell %s got %s, want %s", cellName, got, want)
		}
	}

	for i, row := range spreadSheet.GetRows(sheetName) {
		if i == 0 {
			continue // Header was already checked
		}
		c := cases[i-1]
		if row[0] != c.City {
			t.Errorf("Spreadsheet city got %s, want %s", row[0], c.City)
		}
		if row[1] != c.Country {
			t.Errorf("Spreadsheet country got %s, want %s", row[1], c.Country)
		}
	}
}
