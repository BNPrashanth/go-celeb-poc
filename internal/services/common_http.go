package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/BNPrashanth/go-celeb-poc/internal/logger"
	"github.com/spf13/viper"
)

func requestWithBasic(req *http.Request, obj interface{}) error {
	client := &http.Client{}
	req.SetBasicAuth(viper.GetString("pp_client"), viper.GetString("pp_secret"))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	return json.Unmarshal(bodyText, &PayPalClient)
}
