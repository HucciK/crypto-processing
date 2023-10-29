package config

type CryptoProcessingConfig struct {
	PaymentsConfig PaymentsConfig
	ObserverConfig ObserverConfig
}

type PaymentsConfig struct {
	PaymentTTL int `json:"payment_ttl"`
}

type ObserverConfig struct {
	ObserveInterval int `json:"observ_interval"`
}
