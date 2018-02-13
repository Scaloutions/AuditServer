package data

import (
	"encoding/xml"

	"../utils"
)

/*
	XML structs
*/

// UserCommand infomration necessary for logging user command
type UserCommand struct {
	XMLName        xml.Name `xml:"userCommand"`
	Timestamp      int64    `xml:"timestamp,omitempty"`
	Server         string   `xml:"server,omitempty"`
	TransactionNum int      `xml:"transactionNum,omitempty"`
	Command        string   `xml:"command,omitempty"`
	Username       string   `xml:"username,omitempty"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

// AccountTransaction infomration necessary for logging account transaction
type AccountTransaction struct {
	XMLName        xml.Name `xml:"accountTransaction"`
	Timestamp      int64    `xml:"timestamp,omitempty"`
	Server         string   `xml:"server,omitempty"`
	TransactionNum int      `xml:"transactionNum,omitempty"`
	Action         string   `xml:"action,omitempty"`
	Username       string   `xml:"username,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

// SystemEvent information necessary for logging system event
type SystemEvent struct {
	XMLName        xml.Name `xml:"systemEvent"`
	Timestamp      int64    `xml:"timestamp,omitempty"`
	Server         string   `xml:"server,omitempty"`
	TransactionNum int      `xml:"transactionNum,omitempty"`
	Command        string   `xml:"command,omitempty"`
	Username       string   `xml:"username,omitempty"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

// QuoteServer information necessary for logging quote server hit
type QuoteServer struct {
	XMLName         xml.Name `xml:"quoteServer"`
	Timestamp       int64    `xml:"timestamp,omitempty"`
	Server          string   `xml:"server,omitempty"`
	TransactionNum  int      `xml:"transactionNum,omitempty"`
	QuoteServerTime int64    `xml:"quoteServerTime,omitempty"`
	Command         string   `xml:"command,omitempty"`
	Username        string   `xml:"username,omitempty"`
	StockSymbol     string   `xml:"stockSymbol,omitempty"`
	Price           string   `xml:"price,omitempty"`
	Cryptokey       string   `xml:"cryptokey,omitempty"`
}

type ErrorEvent struct {
	XMLName        xml.Name `xml:"errorEvent"`
	Timestamp      int64    `xml:"timestamp,omitempty"`
	Server         string   `xml:"server,omitempty"`
	TransactionNum int      `xml:"transactionNum,omitempty"`
	Command        string   `xml:"command,omitempty"`
	Username       string   `xml:"username,omitempty"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
	ErrorMessage   string   `xml:"errorMessage,omitempty"`
}

func GetUserCommand(
	server string,
	transactionNum int,
	command string,
	username string,
	stockSymbol string,
	funds float64) UserCommand {

	fundsAsString := utils.GetFundsAsString(funds)

	return UserCommand{
		Timestamp:      utils.GetCurrentTs(),
		Server:         server,
		TransactionNum: transactionNum,
		Command:        command,
		Username:       username,
		StockSymbol:    stockSymbol,
		Funds:          fundsAsString}
}

func GetAccountTransaction(
	server string,
	transactionNum int,
	action string,
	username string,
	funds float64) AccountTransaction {

	fundsAsString := utils.GetFundsAsString(funds)

	return AccountTransaction{
		Timestamp:      utils.GetCurrentTs(),
		Server:         server,
		TransactionNum: transactionNum,
		Action:         action,
		Username:       username,
		Funds:          fundsAsString}
}

func GetSystemEvent(
	server string,
	transactionNum int,
	command string,
	username string,
	stockSymbol string,
	funds float64) SystemEvent {

	fundsAsString := utils.GetFundsAsString(funds)

	return SystemEvent{
		Timestamp:      utils.GetCurrentTs(),
		Server:         server,
		TransactionNum: transactionNum,
		Command:        command,
		Username:       username,
		StockSymbol:    stockSymbol,
		Funds:          fundsAsString}
}

func GetQuoteServer(
	server string,
	transactionNum int,
	quoteServerTime int64,
	command string,
	username string,
	stockSymbol string,
	price float64,
	cryptokey string) QuoteServer {

	priceAsString := utils.GetFundsAsString(price)

	return QuoteServer{
		Timestamp:       utils.GetCurrentTs(),
		Server:          server,
		TransactionNum:  transactionNum,
		QuoteServerTime: quoteServerTime,
		Command:         command,
		Username:        username,
		StockSymbol:     stockSymbol,
		Price:           priceAsString,
		Cryptokey:       cryptokey}
}

func GetErrorEvent(
	server string,
	transactionNum int,
	command string,
	username string,
	stockSymbol string,
	funds float64,
	errorMessage string) ErrorEvent {

	fundsAsString := utils.GetFundsAsString(funds)

	return ErrorEvent{
		Timestamp:      utils.GetCurrentTs(),
		Server:         server,
		TransactionNum: transactionNum,
		Command:        command,
		Username:       username,
		StockSymbol:    stockSymbol,
		Funds:          fundsAsString,
		ErrorMessage:   errorMessage}
}

/*

	Functions

*/

// GetEventMap (to be continued)
func GetEventMap() map[string]int {
	return map[string]int{
		"usercommand":        1,
		"systemevent":        2,
		"errorevent":         3,
		"quoteserver":        4,
		"accounttransaction": 5,
	}
}
