package runtime

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mrzack99s/mrz-captive-portal-api/configs"
)

func allowNet(message map[string]interface{}) {

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(configs.SystemConfig.ZAuth.Operator.HostURL+"/allowNet", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()

}

func logout(message map[string]interface{}) {

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(configs.SystemConfig.ZAuth.Operator.HostURL+"/logout", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close()

}

func Run(todo string, data map[string]interface{}) {
	switch todo {
	case "allowNet":
		go allowNet(data)

	case "logout":
		go logout(data)

	}

}
