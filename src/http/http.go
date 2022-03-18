package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func setHeaders(req *http.Request) {
	req.Header.Set("Accept-Language", "pl-PL")
	req.Header.Set("Authorization", "Basic dmlyZ2luOnZtMTIzNA==")
	req.Header.Set("x-app-id", "U2FsdGVkX1/yC2X2icYybogWyqspG/4fEzuF9FwS8Tw=")
	req.Header.Set("User-Agent", "VirginMobile 2.6.11, Android 5.1.1, unknown")
	req.Header.Set("Connection", "Keep-Alive")
}
func GetCookie(cfg config.Config) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer([]byte(fmt.Sprintf("username=%s&password=%s", cfg.Username, cfg.Password)))
	req, _ := http.NewRequest("POST", "https://virginmobile.pl/spitfire-web-api/api/v1/authentication/login", body)
	setHeaders(req)
	req.Header.Set("Content-Length", string(rune(body.Len())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		return strings.Split(res.Header.Get("set-cookie"), ";")[0], nil
	} else {
		return "", errors.New(res.Status)
	}
}

func GetVMData(cfg config.Config, cookie string) (config.VMData, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://virginmobile.pl/spitfire-web-api/api/v1/user/fullDetails", nil)
	setHeaders(req)
	req.Header.Set("msisdn", cfg.Username)
	req.Header.Set("Host", "virginmobile.pl")
	req.Header.Set("Cookie", cookie)

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return config.VMData{}, err
		}
		var VMData config.VMData
		err = json.Unmarshal(body, &VMData)
		if err != nil {
			return config.VMData{}, err
		}
		return VMData, nil
	} else {
		return config.VMData{}, errors.New(res.Status)
	}
}
