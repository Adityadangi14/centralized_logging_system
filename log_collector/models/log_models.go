package models

import "time"

type LogMessage struct {
	Message string `json:"message"`
}

type ParsedLog struct {
	Timestamp     time.Time `json:"timestamp"`
	EventCategory string    `json:"eventCategory"`
	SourceType    string    `json:"sourceType"`
	Username      string    `json:"username"`
	Hostname      string    `json:"hostname"`
	Severity      string    `json:"severity"`
	RawMessage    string    `json:"rawMessage"`
	IsBlacklisted bool      `json:"isBlacklisted"`
}
