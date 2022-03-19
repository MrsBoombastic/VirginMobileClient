package main

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"github.com/MrBoombastic/VirginMobileClient/src/cookie"
	"github.com/MrBoombastic/VirginMobileClient/src/http"
	"log"
	"os"
	"time"
)

var cfg = config.Get()

func main() {
	// Getting saved cookie
	savedCookie, err := cookie.ReadCookie()
	if err != nil {
		log.Fatal(err)
	}
	// Getting latest data using saved cookie
	data, err := http.GetFullDetails(cfg, savedCookie)
	if err != nil {
		log.Fatal(err)
	}

	// Saving most important data to info.txt file
	var text string
	for _, number := range data.MsisdnDetails {
		text += fmt.Sprintf(`Phone number: %v
Valid until: %v
Balance: %v
Tariff: %v
Bundle name: %v
Bundle valid until: %v
Minutes: %v mins (%v secs)
Mobile data: %v
SMS: %v
Updated: %v

`, number.Msisdn, number.CustomerBalancesDto.GeneralBalance.ValidDate,
			number.CustomerBalancesDto.GeneralBalance.Quantity, number.TariffName, number.ComplexBundleName,
			number.ComplexBundleValidDate, number.CustomerBalancesDto.ComplexBundleVoiceBalance.Quantity/60, number.CustomerBalancesDto.ComplexBundleVoiceBalance.Quantity,
			number.CustomerBalancesDto.DataBalance.Quantity, number.CustomerBalancesDto.SmsBalance.Quantity, time.Now().Format("2006.01.02 15:04:05"))
	}

	err = os.WriteFile("info.txt", []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Saving all data to info.json file
	file, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("info.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
