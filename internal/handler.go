package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type APIResponse struct {
	Response         string
	IdInstance       string
	ApiTokenInstance string
	ChatId           string
}

var idInstance string = ""
var apiTokenInstance string = ""

func HandleGetSettings(w http.ResponseWriter, r *http.Request) {
	idInstance = r.FormValue("idInstance")
	apiTokenInstance = r.FormValue("apiTokenInstance")
	url := fmt.Sprintf("https://api.green-api.com/waInstance%s/getSettings/%s", idInstance, apiTokenInstance)
	fmt.Println(url)
	response := MakeAPIRequest(url, "GET", nil)
	tmpl, _ := template.ParseFiles("index.html")
	fmt.Println(response, " .. ", apiTokenInstance, "  ..  ", idInstance)
	tmpl.Execute(w, APIResponse{Response: response, IdInstance: idInstance, ApiTokenInstance: apiTokenInstance})
}

func HandleGetStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance = r.FormValue("idInstance")
	apiTokenInstance = r.FormValue("apiTokenInstance")
	url := fmt.Sprintf("https://api.green-api.com/waInstance%s/getStateInstance/%s", idInstance, apiTokenInstance)
	response := MakeAPIRequest(url, "GET", nil)
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, APIResponse{Response: response, IdInstance: idInstance, ApiTokenInstance: apiTokenInstance})
}

func HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	chatId := r.FormValue("chatId")
	message := r.FormValue("message")

	url := fmt.Sprintf("https://api.green-api.com/waInstance%s/sendMessage/%s", idInstance, apiTokenInstance)
	payload := map[string]string{
		"chatId":  chatId,
		"message": message,
	}
	response := MakeAPIRequest(url, "POST", payload)
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, APIResponse{Response: response, IdInstance: idInstance, ApiTokenInstance: apiTokenInstance, ChatId: chatId})
}

func HandleSendFileByUrl(w http.ResponseWriter, r *http.Request) {
	chatId := r.FormValue("chatId")
	urlFile := r.FormValue("fileUrl")
	caption := r.FormValue("caption")

	url := fmt.Sprintf("https://api.green-api.com/waInstance%s/sendFileByUrl/%s", idInstance, apiTokenInstance)
	payload := map[string]string{
		"chatId":   chatId,
		"urlFile":  urlFile,
		"fileName": "File",
		"caption":  caption,
	}
	response := MakeAPIRequest(url, "POST", payload)
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, APIResponse{Response: response, IdInstance: idInstance, ApiTokenInstance: apiTokenInstance, ChatId: chatId})
}

func MakeAPIRequest(url, method string, payload interface{}) string {
	var req *http.Request
	var err error

	if payload != nil {
		jsonData, _ := json.Marshal(payload)
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Reading response body failed: %v", err)
	}

	return string(body)
}
