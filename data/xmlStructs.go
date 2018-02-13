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
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum int      `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

// AccountTransaction infomration necessary for logging account transaction
type AccountTransaction struct {
	XMLName        xml.Name `xml:"accountTransaction"`
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum int      `xml:"transactionNum"`
	Action         string   `xml:"action"`
	Username       string   `xml:"username"`
	Funds          string   `xml:"funds"`
}

// SystemEvent information necessary for logging system event
type SystemEvent struct {
	XMLName        xml.Name `xml:"systemEvent"`
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum int      `xml:"transactionNum"`
	Command        string   `xml:"command,omitempty"`
	Username       string   `xml:"username"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

// QuoteServer information necessary for logging quote server hit
type QuoteServer struct {
	XMLName         xml.Name `xml:"quoteServer"`
	Timestamp       int64    `xml:"timestamp"`
	Server          string   `xml:"server"`
	TransactionNum  int      `xml:"transactionNum"`
	QuoteServerTime int64    `xml:"quoteServerTime"`
	Command         string   `xml:"command,omitempty"`
	Username        string   `xml:"username"`
	StockSymbol     string   `xml:"stockSymbol"`
	Price           string   `xml:"price"`
	Cryptokey       string   `xml:"cryptokey"`
}

type ErrorEvent struct {
	XMLName        xml.Name `xml:"errorEvent"`
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum int      `xml:"transactionNum"`
	Command        string   `xml:"command,omitempty"`
	Username       string   `xml:"username"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
	ErrorMessage   string   `xml:"errorMessage"`
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
