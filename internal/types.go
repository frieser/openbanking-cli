package internal

import (
	"fmt"
	"github.com/frieser/nordigen-go-lib"
)

type Country struct {
	Emoji string
	Name  string
	Code  string
}

type Bank struct {
	Id                   string
	Name                 string
	TransactionTotalDays int
}

type Account struct {
	Name         string
	Id           string
	Iban         string
	Owner        string
	Currency     string
	Status       string
	CreateAt     string
	LastAccessed string
	Balances     nordigen.AccountBalances
}

var Countries = []Country{
	{Emoji: "🇬🇧", Name: "United Kingdom", Code: "GB"},
	{Emoji: "🇩🇪", Name: "Germany", Code: "DE"},
	{Emoji: "🇲🇫", Name: "France", Code: "FR"},
	{Emoji: "🇮🇹", Name: "Italy", Code: "IT"},
	{Emoji: "🇪🇸", Name: "Spain", Code: "ES"},
	{Emoji: "🇵🇹", Name: "Portugal", Code: "PT"},
}

func (c Country) String() string {
	return fmt.Sprintf("%s %s (%s)", c.Emoji, c.Name, c.Code)
}

func (b Bank) String() string {
	return fmt.Sprintf("%s (%s)", b.Name, b.Id)
}

func (a Account) String() string {
	return fmt.Sprintf("%s (%s)", a.Name, a.Iban)
}
