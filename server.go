package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"./utils"
)

func checkAndHandleError(err error) {
	if err != nil {
		utils.ERROR.Println(err)
	}
}

func getBody(req *http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	checkAndHandleError(err)
	return body
}

func getObject(i interface{}, body []byte) interface{} {
	err := json.Unmarshal(body, &i)
	checkAndHandleError(err)
	return i
}

func logXMLObj(i interface{}) {
	xmlEvent := utils.GetXMLEventString(i)
	utils.XMLLOGGER.Println(xmlEvent)
}

func usercommand(rw http.ResponseWriter, req *http.Request) {
	processing("usercommand", rw, req)
}

func systemevent(rw http.ResponseWriter, req *http.Request) {
	processing("systemevent", rw, req)
}

func quoteserver(rw http.ResponseWriter, req *http.Request) {
	processing("quoteserver", rw, req)
}

func accounttransaction(rw http.ResponseWriter, req *http.Request) {
	processing("accounttransaction", rw, req)
}

func errorevent(rw http.ResponseWriter, req *http.Request) {
	processing("errorevent", rw, req)
}

func processing(
	commandType string,
	rw http.ResponseWriter,
	req *http.Request) {

	body := getBody(req)
	utils.INFO.Println(string(body))

	switch commandType {
	case "usercommand":
		var userCommand utils.UserCommand
		getObject(&userCommand, body)
		logXMLObj(userCommand)
	case "systemevent":
		var systemEvent utils.SystemEvent
		getObject(&systemEvent, body)
		logXMLObj(systemEvent)
	case "accounttransaction":
		var accountTransaction utils.AccountTransaction
		getObject(&accountTransaction, body)
		logXMLObj(accountTransaction)
	case "quoteserver":
		var quoteServer utils.QuoteServer
		getObject(&quoteServer, body)
		logXMLObj(quoteServer)
	case "errorevent":
		var errorEvent utils.ErrorEvent
		getObject(&errorEvent, body)
		logXMLObj(errorEvent)
	}

}

func main() {

	utils.Init() // initialize loggers
	http.HandleFunc("/systemevent", systemevent)
	http.HandleFunc("/usercommand", usercommand)
	http.HandleFunc("/quoteserver", quoteserver)
	http.HandleFunc("/errorevent", errorevent)
	http.HandleFunc("/accounttransaction", accounttransaction)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
