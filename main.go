package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type LoanInfo struct {
	LoanNumber           string
	LoanBookedDate       string
	LoanType             string
	PrincipalAmount      float64
	InterestRate         float64
	Tenure               int
	OutStandingPrincipal float64
}
type EMIInfo struct {
	PrincipalAmount float64
	InterestAmount  float64
	EMIDate         string
}

const (
	marginX         = 10.0
	marginY         = 20.0
	PaperSize       = "A4"
	LoanNumber      = "Loan Number"
	LoanNumberWidth = 28

	LoanBookedDate      = "Loan Booked Date"
	LoanBookedDateWidth = 25

	LoanType      = "Loan Type"
	LoanTypeWidth = 25

	PrincipalAmt         = "Principal \n Amount (Rs.)"
	PrincipalAmountWidth = 25

	hdfcLogoPath = "//Users//p26345//GoWorkspace//Hdfc-bank-logo.png"

	InterestRate      = "Interest Rate (%)"
	InterestRateWidth = 25

	Tenure      = "Tenure \n (months)"
	TenureWidth = 25

	OutstandingPri      = "Outstanding \nPrincipal (Rs.)"
	OutstandingPriWidth = 30

	FooterMessage = "This is a system generated document and does not require signature."
	FooterMargin  = 25
)

func main4() {

}
func IsEven(n int) bool {
	return n%2 == 0
}

func Concatenatebuffer(s1 string, s2 string) string {
	var buf bytes.Buffer
	buf.WriteString(s1)
	buf.WriteString(s2)
	return buf.String()
}
func main() {

	var ct int

	for t := 0; t < 10; t++ {
		fn := "Payment" + strconv.Itoa(t+1) + ".pdf"
		go CreatePaymentSchedule(fn, 10)
	}

	fmt.Scanln(&ct)

}

func CreatePaymentSchedule(fn string, numEMI int) {

	var pdf *gofpdf.Fpdf
	pdf = PageSetup(pdf)
	pdf.SetFont("Arial", "BU", 12)
	pdf.SetX(80)
	pdf.Cell(100, 25, "Loan EMI Table")

	pdf.Ln(-1)
	AddLoanTableToPage(pdf)
	AddEMIDataToPage(pdf, numEMI)
	AddFooter(pdf, FooterMessage)

	err := pdf.OutputFileAndClose(fn)
	defer RemoveFileFromDisk(fn)
	if err != nil {
		log.Fatal(err)
	}

}

//Add footer message at the end of the page.
//Default margin set to 25
func AddFooter(pdf *gofpdf.Fpdf, msg string) {
	_, h := pdf.GetPageSize()
	var Y = pdf.GetY()
	Y = h - FooterMargin
	pdf.SetY(Y)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFooterFunc(func() {
		pdf.CellFormat(200, 25, msg, "0", 0, "L", false, 0, "")
	})

}
func RemoveFileFromDisk(fn string) bool {
	//remove the file.
	fmt.Println("FN ", fn)
	time.Sleep(time.Second * 60)
	_, err := os.Stat(fn)
	if err != nil {
		log.Fatal("File ", fn, " path does not exist.")
		return false
	} else {
		err := os.Remove(fn)
		if err != nil {
			log.Fatal("File ", fn, " could not be removed.", err)
			return false
		} else {
			fmt.Println("File Removed")
		}
	}
	return true

}
func AddEMIDataToPage(pdf *gofpdf.Fpdf, numEMI int) {
	//set header
	pdf.SetFont("Arial", "B", 7)
	for _, val := range GetEMIHeader() {
		pdf.CellFormat(60, 6, val, "1", 0, "CM", false, 0, "")

	}

	pdf.Ln(-1)
	for _, val := range MockEMIData(numEMI) {
		amount, interest := strconv.FormatFloat(val.PrincipalAmount, 'f', 2, 64), strconv.FormatFloat(val.InterestAmount, 'f', 2, 64)
		pdf.SetFont("Arial", "", 8)

		pdf.CellFormat(60, 5, amount, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(60, 5, interest, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(60, 5, val.EMIDate, "1", 0, "CM", false, 0, "")
		pdf.Ln(-1)

	}

}
func AddLoanTableToPage(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 7)
	pdf.CellFormat(LoanNumberWidth, 8, LoanNumber, "1", 0, "C", false, 0, "")
	pdf.CellFormat(LoanBookedDateWidth, 8, LoanBookedDate, "1", 0, "C", false, 0, "")
	pdf.CellFormat(LoanTypeWidth, 8, LoanType, "1", 0, "C", false, 0, "")
	X := pdf.GetX()
	Y := pdf.GetY()
	pdf.MultiCell(PrincipalAmountWidth, 4, PrincipalAmt, "1", "C", false)
	X = X + PrincipalAmountWidth
	pdf.SetXY(X, Y)
	pdf.CellFormat(InterestRateWidth, 8, InterestRate, "1", 0, "C", false, 0, "")
	X = X + InterestRateWidth
	pdf.SetXY(X, Y)
	pdf.MultiCell(TenureWidth, 4, Tenure, "1", "C", false)
	X = X + TenureWidth
	pdf.SetXY(X, Y)
	pdf.MultiCell(OutstandingPriWidth, 4, OutstandingPri, "1", "C", false)
	loanInfo := LoanInfo{LoanNumber: "000000000000003193", LoanBookedDate: "31 Aug 2020",
		LoanType: "JUMBOLOAN", PrincipalAmount: 4626.3,
		InterestRate:         17.3,
		Tenure:               18,
		OutStandingPrincipal: 3882.93,
	}

	pdf.SetFont("Arial", "", 7)
	pdf.CellFormat(LoanNumberWidth, 6, loanInfo.LoanNumber, "1", 0, "C", false, 0, "")
	pdf.CellFormat(LoanBookedDateWidth, 6, loanInfo.LoanBookedDate, "1", 0, "C", false, 0, "")
	pdf.CellFormat(LoanTypeWidth, 6, loanInfo.LoanType, "1", 0, "C", false, 0, "")
	pdf.CellFormat(PrincipalAmountWidth, 6, fmt.Sprintf("%v", loanInfo.PrincipalAmount), "1", 0, "C", false, 0, "")
	pdf.CellFormat(InterestRateWidth, 6, fmt.Sprintf("%v", loanInfo.InterestRate), "1", 0, "C", false, 0, "")
	pdf.CellFormat(TenureWidth, 6, fmt.Sprintf("%v", loanInfo.Tenure), "1", 0, "C", false, 0, "")
	pdf.CellFormat(OutstandingPriWidth, 6, fmt.Sprintf("%v", loanInfo.OutStandingPrincipal), "1", 0, "C", false, 0, "")
	//fill EMI data
	pdf.Ln(-1)
	pdf.Ln(-1)

}
func PageSetup(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf = gofpdf.New("P", "mm", PaperSize, "")
	pdf.SetMargins(marginX, marginY, marginX)

	pdf.AddPage()
	pdf.SetAuthor("HDFC Credit Card Division", true)
	//HDFC Logo image.. (Common images path/resources path)

	pdf.ImageOptions(hdfcLogoPath, 10, 10, 40, 10, false, gofpdf.ImageOptions{}, 0, "")
	return pdf
}

func GetHeader() []string {

	return []string{LoanNumber, LoanBookedDate, LoanType, PrincipalAmt, InterestRate, Tenure, OutstandingPri}

}
func GetEMIHeader() []string {

	return []string{"Principal Amount (Rs)", "Interest Amount (Rs)", "EMI Date"}
}

func MockEMIData(numEMI int) []EMIInfo {
	//generate random float amounts.
	rand.Seed(time.Now().UnixNano())

	info := make([]EMIInfo, 0)
	for t := 0; t < numEMI; t++ {
		principal := fmt.Sprintf("%0.2f", rand.Float64()*100)
		principalf, _ := strconv.ParseFloat(principal, 64)
		interest := fmt.Sprintf("%0.2f", rand.Float64()*100)
		interestf, _ := strconv.ParseFloat(interest, 64)
		info = append(info, EMIInfo{PrincipalAmount: principalf, InterestAmount: interestf,
			EMIDate: GetRandomDate()})
	}
	return info

}
func GetRandomDate() string {
	months := []string{"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	dates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28}
	years := []int{2019, 2020, 2021, 2022, 2023}
	rdate := dates[rand.Intn(len(dates)-1)]
	rmonth := months[rand.Intn(len(months)-1)]
	ryear := years[rand.Intn(len(years)-1)]
	val := fmt.Sprintf("%d %s %d", rdate, rmonth, ryear)
	return val
}

//test

func EncodePaymentSchedule(payment gofpdf.Pdf) {
	var buf bytes.Buffer
	err := payment.Output(&buf)

	if err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println(str)
	fn, err := os.Create("sm.pdf")
	dec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	num, err2 := fn.Write(dec)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(num)
	defer fn.Close()

}

/*
	func DecodePdf(arr []byte) {
		fn, err := os.Create("sm.pdf")
		if err != nil {
			log.Fatal(err)
		}
		defer fn.Close()
		buf := make([]byte, len(arr))
		fmt.Println("BUF Length :: ", len(buf))
		_, err1 := base64.StdEncoding.Decode(buf, arr)
		if err1 != nil {
			log.Fatal(err1)
		}
		fn.WriteString(string(buf))
}
*/
