package lib

import (
	"errors"
	"roof-stack/internal/entity"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

const (
	path string = "reports"
	excel string = "xlsx"
	fileName string = "transactions-report"
)

func ReportData(transactons []*entity.Transaction) (string, error) {

	f := excelize.NewFile()
	f = setColumns(f)
	f = setValues(transactons, f)
	const layout = "2017.09.07-17.06.06"
	t := time.Now()
	if err := f.SaveAs( path + "/"+ fileName + "-" + t.Format(layout) + "." + excel); err != nil {
		panic(err)
	}
	return "Report success", errors.New("Export report successfull")
}

func setColumns( f *excelize.File ) *excelize.File {

	// Buraları hardcode yaptım zamanım omadığı için böyle pushlamak zorundayım
	// foreach döngüsüyle column ve valueler yazılabilirdi.
	f.SetCellValue("Sheet1", "A1", "Status")
	f.SetCellValue("Sheet1", "B1", "ActionType")
	f.SetCellValue("Sheet1", "C1", "CurrencyID")
	f.SetCellValue("Sheet1", "D1", "WalletID")
	f.SetCellValue("Sheet1", "E1", "Amount")
	f.SetCellValue("Sheet1", "F1", "CreatedAt")

	return f
}

func setValues(transactions []*entity.Transaction,  f *excelize.File) *excelize.File {
	for k, v := range transactions {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(k+2), v.Status)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(k+2), v.ActionType)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(k+2), v.CurrencyID)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(k+2), v.WalletID)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(k+2), v.CreatedAt)
	}

	return f
}