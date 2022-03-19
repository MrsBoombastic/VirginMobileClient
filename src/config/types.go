package config

// Config is config.json structure of all available settings
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// FullDetails is a struct of `user/fullDetails` endpoint
type FullDetails struct {
	Id                    float64     `json:"id"`
	Email                 string      `json:"email"`
	FirstName             string      `json:"firstName"`
	MiddleName            string      `json:"middleName"`
	LastName              string      `json:"lastName"`
	Pesel                 string      `json:"pesel"`
	Gender                interface{} `json:"gender"`
	BirthDate             interface{} `json:"birthDate"`
	InvoiceData           interface{} `json:"invoiceData"`
	CorrespondenceAddress interface{} `json:"correspondenceAddress"`
	ResidenceAddress      interface{} `json:"residenceAddress"`
	CustomerType          string      `json:"customerType"`
	DocumentType          interface{} `json:"documentType"`
	DocumentNumber        interface{} `json:"documentNumber"`
	PayPalAgreement       bool        `json:"payPalAgreement"`
	Registered            bool        `json:"registered"`
	MsisdnDetails         []struct {
		Type                string      `json:"type"`
		Msisdn              string      `json:"msisdn"`
		Name                interface{} `json:"name"`
		Network             string      `json:"network"`
		ContractType        string      `json:"contractType"`
		CustomerBalancesDto struct {
			MsisdnState    string `json:"msisdnState"`
			GeneralBalance struct {
				Quantity  float64 `json:"quantity"`
				Unit      string  `json:"unit"`
				ValidDate string  `json:"validDate"`
				Unlimited bool    `json:"unlimited"`
			} `json:"generalBalance"`
			ValuePackBalance          interface{} `json:"valuePackBalance"`
			VoiceBalance              interface{} `json:"voiceBalance"`
			ComplexBundleVoiceBalance struct {
				Quantity  float64 `json:"quantity"`
				Unit      string  `json:"unit"`
				ValidDate string  `json:"validDate"`
				Unlimited bool    `json:"unlimited"`
			} `json:"complexBundleVoiceBalance"`
			VoiceOnNetBalance     interface{} `json:"voiceOnNetBalance"`
			VoiceToUkraineBalance interface{} `json:"voiceToUkraineBalance"`
			DataBalance           struct {
				Quantity  float64 `json:"quantity"`
				Unit      string  `json:"unit"`
				ValidDate string  `json:"validDate"`
				Unlimited bool    `json:"unlimited"`
			} `json:"dataBalance"`
			RoamingDataBalance struct {
				Quantity  float64 `json:"quantity"`
				Unit      string  `json:"unit"`
				ValidDate string  `json:"validDate"`
				Unlimited bool    `json:"unlimited"`
			} `json:"roamingDataBalance"`
			DataSpecialBalance interface{} `json:"dataSpecialBalance"`
			DataPromoBalance   interface{} `json:"dataPromoBalance"`
			SmsBalance         struct {
				Quantity  float64 `json:"quantity"`
				Unit      string  `json:"unit"`
				ValidDate string  `json:"validDate"`
				Unlimited bool    `json:"unlimited"`
			} `json:"smsBalance"`
			SmsToUkraineBalance interface{} `json:"smsToUkraineBalance"`
		} `json:"customerBalancesDto"`
		OperatorBlockadeComponentStatus           interface{} `json:"operatorBlockadeComponentStatus"`
		OperatorBlockadeComponentInformationError interface{} `json:"operatorBlockadeComponentInformationError"`
		SplitBalancesDto                          struct {
			CreditLimitBalance interface{} `json:"creditLimitBalance"`
			BillshockBalance   struct {
				Quantity        float64 `json:"quantity"`
				InitialQuantity float64 `json:"initialQuantity"`
				Unit            string  `json:"unit"`
				ValidDate       string  `json:"validDate"`
				Unlimited       bool    `json:"unlimited"`
			} `json:"billshockBalance"`
			CyclicVoiceBalance struct {
				Quantity        float64 `json:"quantity"`
				InitialQuantity float64 `json:"initialQuantity"`
				Unit            string  `json:"unit"`
				ValidDate       string  `json:"validDate"`
				Unlimited       bool    `json:"unlimited"`
			} `json:"cyclicVoiceBalance"`
			OneOffVoiceBalance                interface{} `json:"oneOffVoiceBalance"`
			CyclicVoiceOnNetBalance           interface{} `json:"cyclicVoiceOnNetBalance"`
			OneOffVoiceOnNetBalance           interface{} `json:"oneOffVoiceOnNetBalance"`
			CyclicVoiceOnMobileBalance        interface{} `json:"cyclicVoiceOnMobileBalance"`
			OneOffVoiceOnMobileBalance        interface{} `json:"oneOffVoiceOnMobileBalance"`
			CyclicVoiceIncomingRoamingBalance interface{} `json:"cyclicVoiceIncomingRoamingBalance"`
			OneOffVoiceIncomingRoamingBalance interface{} `json:"oneOffVoiceIncomingRoamingBalance"`
			CyclicVoiceToUkraineBalance       interface{} `json:"cyclicVoiceToUkraineBalance"`
			OneOffVoiceToUkraineBalance       interface{} `json:"oneOffVoiceToUkraineBalance"`
			CyclicSmsBalance                  struct {
				Quantity        float64 `json:"quantity"`
				InitialQuantity float64 `json:"initialQuantity"`
				Unit            string  `json:"unit"`
				ValidDate       string  `json:"validDate"`
				Unlimited       bool    `json:"unlimited"`
			} `json:"cyclicSmsBalance"`
			OneOffSmsBalance          interface{} `json:"oneOffSmsBalance"`
			CyclicSmsToUkraineBalance interface{} `json:"cyclicSmsToUkraineBalance"`
			OneOffSmsToUkraineBalance interface{} `json:"oneOffSmsToUkraineBalance"`
			CyclicMmsBalance          interface{} `json:"cyclicMmsBalance"`
			OneOffMmsBalance          interface{} `json:"oneOffMmsBalance"`
			CyclicDataBalance         struct {
				Quantity        float64 `json:"quantity"`
				InitialQuantity float64 `json:"initialQuantity"`
				Unit            string  `json:"unit"`
				ValidDate       string  `json:"validDate"`
				Unlimited       bool    `json:"unlimited"`
			} `json:"cyclicDataBalance"`
			OneOffDataBalance        interface{} `json:"oneOffDataBalance"`
			CyclicDataRoamingBalance struct {
				Quantity        float64 `json:"quantity"`
				InitialQuantity float64 `json:"initialQuantity"`
				Unit            string  `json:"unit"`
				ValidDate       string  `json:"validDate"`
				Unlimited       bool    `json:"unlimited"`
			} `json:"cyclicDataRoamingBalance"`
			OneOffDataRoamingBalance interface{} `json:"oneOffDataRoamingBalance"`
			CyclicDataSpecialBalance interface{} `json:"cyclicDataSpecialBalance"`
			OneOffDataSpecialBalance interface{} `json:"oneOffDataSpecialBalance"`
			CyclicDataPromoBalance   interface{} `json:"cyclicDataPromoBalance"`
			OneOffDataPromoBalance   interface{} `json:"oneOffDataPromoBalance"`
		} `json:"splitBalancesDto"`
		BalanceInformationError               interface{} `json:"balanceInformationError"`
		AccountValidDate                      string      `json:"accountValidDate"`
		ComplexBundleName                     string      `json:"complexBundleName"`
		ComplexBundleValidDate                string      `json:"complexBundleValidDate"`
		ComplexBundleRenewalCount             float64     `json:"complexBundleRenewalCount"`
		NewTariff                             bool        `json:"newTariff"`
		TariffName                            string      `json:"tariffName"`
		MigratedFromOldTariff                 bool        `json:"migratedFromOldTariff"`
		MigrationTargetTariff                 interface{} `json:"migrationTargetTariff"`
		MigrationPending                      bool        `json:"migrationPending"`
		BundleOperationsFlags                 bool        `json:"bundleOperationsFlags"`
		Mnp                                   bool        `json:"mnp"`
		ComplexBundleState                    string      `json:"complexBundleState"`
		BundleFund                            string      `json:"bundleFund"`
		ComplexBundleActivationSubmitted      bool        `json:"complexBundleActivationSubmitted"`
		MgmParentMsisdn                       interface{} `json:"mgmParentMsisdn"`
		Registered                            bool        `json:"registered"`
		JscSubscriberId                       float64     `json:"jscSubscriberId"`
		JscSubscriptionId                     string      `json:"jscSubscriptionId"`
		VoicePostpaidCyclicChangePack         interface{} `json:"voicePostpaidCyclicChangePack"`
		SmsPostpaidCyclicChangePack           interface{} `json:"smsPostpaidCyclicChangePack"`
		DataPostpaidCyclicChangePack          interface{} `json:"dataPostpaidCyclicChangePack"`
		ApplicationPostpaidCyclicChangePack   interface{} `json:"applicationPostpaidCyclicChangePack"`
		ComplexBundlePostpaidCyclicChangePack interface{} `json:"complexBundlePostpaidCyclicChangePack"`
		JscSubscriptionStatusDTO              string      `json:"jscSubscriptionStatusDTO"`
		ContractStartDate                     interface{} `json:"contractStartDate"`
		ContractEndDate                       interface{} `json:"contractEndDate"`
	} `json:"msisdnDetails"`
}
