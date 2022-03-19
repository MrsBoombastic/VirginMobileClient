package main

import (
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"github.com/MrBoombastic/VirginMobileClient/src/cookie"
	"github.com/MrBoombastic/VirginMobileClient/src/http"
	"log"
	"time"
)

var cfg = config.Get()

func main() {
	savedCookie, err := cookie.ReadCookie()
	if err != nil {
		log.Fatal(err)
	}
	data, err := http.GetFullDetails(cfg, savedCookie)
	if err != nil {
		log.Fatal(err)
	}
	for _, number := range data.MsisdnDetails {
		fmt.Println(fmt.Sprintf(`Phone number: %v
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
			number.CustomerBalancesDto.DataBalance.Quantity, number.CustomerBalancesDto.SmsBalance.Quantity, time.Now().Format("2006.01.02 15:04:05")))
	}
}
