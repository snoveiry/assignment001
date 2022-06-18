// Package config defines the main assignment001 configuration
package config

import (
	//authservice "github.com/snoveiry/assignment001/v1/services/auth"
)

type Assignment001 struct {
	//AuthService                  *authservice.Authenticator
	Port                         string
	PostgresHost                 string
	PostgresDatabase             string
	PostgresUsername             string
	PostgresPassword             string
	TokenSecret                    string
	DBInitArgs                   string
	Environment                  string
	BaseURL                      string
	SMTPHost                     string
	SMTPPort                     string
	SMTPUsername                 string
	SMTPPassword                 string
	SMTPSender                   string
	CorsDomainWhitelist          string
	EmailDelivery                bool
	SMSDelivery                  bool
	TwilioAuthToken              string
	TwilioSender                 string
	SMSPerHour                   string
	EmailsPerHour                string
	GoogleCloudStorageBucket     string
	GoogleApplicationCredentials string
	AzureStorageAccountName      string
	AzureStorageAccountKey       string
	AzureStorageServiceURL       string
	SlackStatusURL               string
	MapboxPublicToken            string
	MapboxRequestsPerHour        string
}
