package optional

import (
	"log"
	"strconv"
	"testB/pkg/models"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/mdigger/translit"
)

func CreateF(fileDirect string, data []models.DataTcv) error {
	dict := make(map[string][]models.DataTcv)
	for i := 0; i < len(data); i++ {

		dict[data[i].Unit_guid] = append(dict[data[i].Unit_guid], data[i])
	}
	for key := range dict {
		CreatePdf(fileDirect, dict[key])
	}
	return nil
}

func CreatePdf(fileDirect string, datas []models.DataTcv) {
	pdf := newReport()


	pdf = header(pdf, []string{"n", "mqtt", "invid", "msg_id", "text", "context", "class", "level", "area", "addr", "block", "type", "bit", "invert_bit"})


	data := [][]string{}
	for i := 0; i < len(datas); i++ {
		s := []string{
			strconv.Itoa(datas[i].N), datas[i].Mqtt, datas[i].Invid, datas[i].Msg_id, translit.Ru(datas[i].Text), datas[i].Context, datas[i].Class, strconv.Itoa(datas[i].Level), datas[i].Area, datas[i].Addr, datas[i].Block, datas[i].TypE, datas[i].Bit, datas[i].Invert_bit,
		}
		data = append(data, s)
	}

	pdf = table(pdf, data)

	if pdf.Err() {
		log.Fatalf("failed ! %s", pdf.Error())
	}

	err := pdf.OutputFileAndClose(fileDirect + "/" + datas[0].Unit_guid + ".pdf")
	if err != nil {
		log.Fatalf("error saving pdf file: %s", err)
	}
}

func newReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Times", "B", 20)
	pdf.Cell(40, 10, "Test Report")
	pdf.Ln(12)

	pdf.SetFont("Times", "", 14)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)
	return pdf
}

func header(pdf *gofpdf.Fpdf, headerText []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 8)
	pdf.SetFillColor(240, 240, 240)

	for _, str := range headerText {
		pdf.CellFormat(20, 10, str, "1", 0, "", true, 0, "")
	}

	pdf.Ln(-1) 

	return pdf
}

func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 4)
	pdf.SetFillColor(255, 255, 255) 

	align := []string{"L", "C", "L", "L", "L", "L", "L", "L", "L", "L", "L", "L", "L", "L"}
	for _, line := range tbl {
		for i, str := range line {
			pdf.CellFormat(20, 5, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
	pdf.Ln(-1)
	return pdf
}
