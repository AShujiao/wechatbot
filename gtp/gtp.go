
 package gtp

 import (
	 "bytes"
	 "encoding/json"
	 "io/ioutil"
	 "log"
	 "net/http"
 
	 "github.com/869413421/wechatbot/config"
 )
 
 const BASEURL = "https://gpt.baixing.com"
 
 // ChatGPTResponseBody 请求体
 type ChatGPTResponseBody struct {
	 Data    string `json:"data"`
	 Type    string `json:"type"`
	 Code    int    `json:"code"`
	 Message string `json:"message"`
 }
 
 type ChoiceItem struct {
 }
 
 // ChatGPTRequestBody 响应体
 type ChatGPTRequestBody struct {
	 Prompt string `json:"p"`
	 K      string `json:"k"`
 }
 
 // Completions gtp文本模型回复
 // curl https://api.openai.com/v1/completions
 // -H "Content-Type: application/json"
 // -H "Authorization: Bearer your chatGPT key"
 // -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
 func Completions(msg string) (string, error) {
	 apiKey := config.LoadConfig().ApiKey
	 requestBody := ChatGPTRequestBody{
		 K:      apiKey,
		 Prompt: msg,
	 }
	 requestData, err := json.Marshal(requestBody)
 
	 if err != nil {
		 return "", err
	 }
	 //log.Printf("request gtp json string : %v", string(requestData))
	 req, err := http.NewRequest("POST", BASEURL, bytes.NewBuffer(requestData))
	 if err != nil {
		 return "", err
	 }
	 req.Header.Set("Content-Type", "application/json")
	 //req.Header.Set("Authorization", "Bearer "+apiKey)
	 client := &http.Client{}
	 response, err := client.Do(req)
	 if err != nil {
		 return "", err
	 }
	 defer response.Body.Close()
 
	 body, err := ioutil.ReadAll(response.Body)
	 if err != nil {
		 return "", err
	 }
 
	 gptResponseBody := &ChatGPTResponseBody{}
	 log.Println(string(body))
	 err = json.Unmarshal(body, gptResponseBody)
	 if err != nil {
		 return "", err
	 }
	 var reply string
	 if gptResponseBody.Code == 0 {
		 reply = gptResponseBody.Data
	 }
	 //log.Printf("gpt response text: %s \n", reply)
	 return reply, nil
 }
 