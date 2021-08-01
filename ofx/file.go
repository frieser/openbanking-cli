package ofx

import (
	"encoding/xml"
	"io"
)

type TxnList struct {
	XMLName          xml.Name `xml:"OFX"`
	CurDef           string   `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>CURDEF"`
	PayerBank        string   `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>BANKID"`
	PayerAccount     string   `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>ACCTID"`
	PayerAccountType string   `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>ACCTTYPE"`
	Transactions     []Txn    `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKTRANLIST>STMTTRN"`
}

type Txn struct {
	XMLName          xml.Name `xml:"STMTTRN"`
	Type             string   `xml:"TRNTYPE"`
	DatePosted       string   `xml:"DTPOSTED"`
	DateUser         string   `xml:"DTUSER"`
	Amount           float64  `xml:"TRNAMT"`
	ID               string   `xml:"FITID"`
	Payee            string   `xml:"NAME"`
	Memo             string   `xml:"MEMO"`
	PayeeBank        string   `xml:"BANKACCTTO>BANKID"`
	PayeeAccount     string   `xml:"BANKACCTTO>ACCTID"`
	PayeeAccountType string   `xml:"BANKACCTTO>ACCTTYPE"`
}

func (txns *TxnList) Write(w io.Writer) error {
	header := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<?OFX OFXHEADER="200" VERSION="211" SECURITY="NONE" OLDFILEUID="NONE" NEWFILEUID="NONE"?>
`
	_, err := w.Write([]byte(header))

	if err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")

	if err := enc.Encode(txns); err != nil {
		return err
	}

	return nil
}
