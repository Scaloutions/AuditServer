package data

import (
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID             bson.ObjectId `bson:"_id"`
	Timestamp      int64         `json:"timestamp"`
	Server         string        `json:"server"`
	TransactionNum int           `json:"transactionNum"`
	EventType      string        `json:"eventType"`
	UserName       string        `json:"userName"`
}

type UserCommandEvent struct {
	Event       *Event
	Command     string `json:"command"`
	StockSymbol string `json:"stockSymbol,omitempty"`
	Funds       string `json:"funds,omitempty"`
}

type AccountTransactionEvent struct {
	Event  *Event
	Action string `json:"action"`
	Funds  string `json:"funds"`
}

type SystemEventJ struct {
	Event       *Event
	Command     string `json:"command,omitempty"`
	StockSymbol string `json:"stockSymbol,omitempty"`
	Funds       string `json:"funds,omitempty"`
}

type QuoteServerEvent struct {
	Event           *Event
	QuoteServerTime int64  `json:"quoteServerTime"`
	Command         string `json:"command,omitempty"`
	StockSymbol     string `json:"stockSymbol"`
	Price           string `json:"price"`
	Cryptokey       string `json:"cryptoKey"`
}

type ErrorEventJ struct {
	Event        *Event
	Command      string `json:"command,omitempty"`
	StockSymbol  string `json:"stockSymbol,omitempty"`
	Funds        string `json:"funds,omitempty"`
	ErrorMessage string `json:"errorMessage"`
}
