package main

import (
	"fmt"
	"github.com/MrBoombastic/VirginMobileClient/src/config"
	"github.com/MrBoombastic/VirginMobileClient/src/http"
	"log"
)

var cfg = config.Get()
var Cookie = ""

func UpdateCookie() error {
	response, err := http.GetCookie(cfg)
	if err != nil {
		return err
	}
	Cookie = response
	return nil
}

func GetData() (config.VMData, error) {
	data, err := http.GetVMData(cfg, Cookie)
	if err != nil {
		return config.VMData{}, err
	}
	return data, nil
}
func main() {
	err := UpdateCookie()
	if err != nil {
		log.Fatal(err)
	}
	data, err := GetData()
	if err != nil {
		log.Fatal(err)
	}
	number := data.MsisdnDetails[0]
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
		number.CustomerBalancesDto.DataBalance.Quantity, number.CustomerBalancesDto.SmsBalance.Quantity, data.UpdateTimestamp))
}
