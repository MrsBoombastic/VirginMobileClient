package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"github.com/MrBoombastic/VirginMobileClient/src/cookie"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var API = "https://virginmobile.pl/spitfire-web-api/api/v1/"

// SetHeaders sets proper headers for API requests
func SetHeaders(req *http.Request) {
	req.Header.Set("Accept-Language", "pl-PL")
	req.Header.Set("Authorization", "Basic dmlyZ2luOnZtMTIzNA==") //LOL
	req.Header.Set("x-app-id", "U2FsdGVkX1/yC2X2icYybogWyqspG/4fEzuF9FwS8Tw=")
	req.Header.Set("User-Agent", "VirginMobile 2.6.11, Android 5.1.1, unknown")
	req.Header.Set("Connection", "Keep-Alive")
}

// Login returns login cookie as a string
func Login(cfg config.Config) (string, error) {
	client := &http.Client{}
	// Creating login body and request
	body := bytes.NewBuffer([]byte(fmt.Sprintf("username=%s&password=%s", cfg.Username, cfg.Password)))
	req, _ := http.NewRequest("POST", API+"authentication/login", body)
	SetHeaders(req)
	req.Header.Set("Content-Length", string(rune(body.Len())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		// Returning cookie, if success
		return strings.Split(res.Header.Get("set-cookie"), ";")[0], nil
	} else {
		return "", errors.New(res.Status)
	}
}

// GetFullDetails returns all fetched data from VM. If 403 error is present, it will try to regenerate cookie and try again.
func GetFullDetails(cfg config.Config, savedCookie string) (config.FullDetails, error) {
	client := &http.Client{}
	// Creating request
	req, _ := http.NewRequest("GET", API+"user/fullDetails", nil)
	SetHeaders(req)
	req.Header.Set("msisdn", cfg.Username)
	req.Header.Set("Host", "virginmobile.pl")
	req.Header.Set("Cookie", savedCookie)

	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		// If everything is OK, return the data
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return config.FullDetails{}, err
		}
		var VMData config.FullDetails
		err = json.Unmarshal(body, &VMData)
		if err != nil {
			return config.FullDetails{}, err
		}
		return VMData, nil
	} else if res.StatusCode == 403 {
		// If cookie is obsolete, regenerate it!
		log.Println("Error 403! Fetching new cookie and trying again...")
		newCookie, err := Login(cfg)
		if err != nil {
			return config.FullDetails{}, err
		}
		err = cookie.SaveCookie(newCookie)
		if err != nil {
			return config.FullDetails{}, err
		}
		time.Sleep(3 * time.Second)
		newVMData, err := GetFullDetails(cfg, newCookie)
		if err != nil {
			return config.FullDetails{}, err
		}
		return newVMData, nil
	} else {
		// If other error is present, fail!
		return config.FullDetails{}, errors.New(res.Status)
	}
}
