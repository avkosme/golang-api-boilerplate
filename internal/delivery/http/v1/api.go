package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/avkosme/golang-api-boilerplate/internal/config"
	"github.com/avkosme/golang-api-boilerplate/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type Http struct {
	req *http.Request
	res http.ResponseWriter
}

type Message struct {
	Message_id int64 `json:"message_id"`
	Message    struct {
		Text string `json:"text"`
	}
}

// Parse http body
func (h *Http) parse() (message *Message, err error) {

	defer func() (message, err error) {
		if c := recover(); c != nil {
			return message, err
		}
		return message, err
	}()

	body, err := ioutil.ReadAll(h.req.Body)

	if err != nil {
		logger.ForError(err)
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		panic(err)
	}

	logger.LogFile.Println(fmt.Sprintf("%s", body))

	return message, err

}

// Run root http route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpObj := Http{r, w}
	message, err := httpObj.parse()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("Hello %s", message.Message.Text)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger.ForError(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write(jsonResp)
	}
}

// Run http handler
func Run() {
	router := httprouter.New()
	router.HandleMethodNotAllowed = false
	router.POST("/", Index)

	err := http.ListenAndServeTLS(
		fmt.Sprintf(
			"%s:%s",
			config.BotBindAddress,
			config.BotPort),
		config.CertPath,
		config.KeyPath,
		router)
	if err != nil {
		logger.ForError(err)
	}
}
