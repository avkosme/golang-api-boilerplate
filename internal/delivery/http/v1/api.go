package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/avkosme/golang-api-boilerplate/internal/config"
	pages "github.com/avkosme/golang-api-boilerplate/internal/service"
	"github.com/avkosme/golang-api-boilerplate/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type Http struct {
	req *http.Request
	res http.ResponseWriter
}

type Message struct {
	Update_id int64 `json:"update_id"`
	Message   struct {
		Text string `json:"text"`
	}
}

// Parse http body
func (h *Http) parse() (message *Message) {

	body, err := ioutil.ReadAll(h.req.Body)

	if err != nil {
		logger.ForError(err)
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		logger.ForError(err)
	}

	logger.LogFile.Println(fmt.Sprintf("%s", body))

	return
}

// Run root http route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpObj := Http{r, w}

	message := httpObj.parse()

	if true == pages.Manage(message.Message.Text) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Run http handler
func Run() {
	router := httprouter.New()
	router.HandleMethodNotAllowed = false
	router.POST("/", Index)

	err := http.ListenAndServeTLS(
		fmt.Sprintf(
			"%s:%s", config.BotBindAddress, config.BotPort), config.CertPath, config.KeyPath, router)
	if err != nil {
		logger.ForError(err)
	}
}
