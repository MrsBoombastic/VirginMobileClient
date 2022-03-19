package main

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"github.com/MrBoombastic/VirginMobileClient/src/cookie"
	"github.com/MrBoombastic/VirginMobileClient/src/http"
	"github.com/MrBoombastic/VirginMobileClient/src/tools"
	"log"
	"math"
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
		numberValidDate, err := tools.FormatVMDate(number.CustomerBalancesDto.GeneralBalance.ValidDate)
		if err != nil {
			log.Fatal(err)
		}
		parsedComplexBundleValidDate, err := tools.FormatVMDate(number.ComplexBundleValidDate)
		if err != nil {
			log.Fatal(err)
		}
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

`, number.Msisdn, numberValidDate.Format("2 Jan 2006 15:04:05"), number.CustomerBalancesDto.GeneralBalance.Quantity, number.TariffName, number.ComplexBundleName,
			parsedComplexBundleValidDate.Format("2 Jan 2006 15:04:05"), math.Round(number.CustomerBalancesDto.ComplexBundleVoiceBalance.Quantity/60),
			number.CustomerBalancesDto.ComplexBundleVoiceBalance.Quantity, tools.ByteCountIEC(int64(number.CustomerBalancesDto.DataBalance.Quantity)*1024), number.CustomerBalancesDto.SmsBalance.Quantity,
			time.Now().Format("2006.01.02 15:04:05"))
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

	log.Println("Done. Output saved to info.txt and info.json files.")
}
