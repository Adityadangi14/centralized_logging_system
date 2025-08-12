package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type ParsedLog struct {
	ID            int32     `json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	EventCategory string    `json:"eventCategory"`
	SourceType    string    `json:"sourceType"`
	Username      string    `json:"username"`
	Hostname      string    `json:"hostname"`
	Severity      string    `json:"severity"`
	RawMessage    string    `json:"rawMessage"`
	IsBlacklisted bool      `json:"isBlacklisted"`
}

func main() {

	serverAddr := "log_collector:3000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to log collector at", serverAddr)

	rand.Seed(time.Now().UnixNano())
	hostnames := []string{"aiops9242", "appnode01", "srv-backup", "dev-machine"}
	users := []string{"root", "admin", "john", "serviceuser"}
	severities := []string{"info", "warn", "error"}
	eventCategories := []string{"syslog", "login_audit", "logout_audit"}
	sourceTypes := []string{"linux", "windows"}

	var idCounter int32 = 1

	for {
		logEntry := ParsedLog{
			ID:            idCounter,
			Timestamp:     time.Now().UTC(),
			EventCategory: eventCategories[rand.Intn(len(eventCategories))],
			SourceType:    sourceTypes[rand.Intn(len(sourceTypes))],
			Username:      users[rand.Intn(len(users))],
			Hostname:      hostnames[rand.Intn(len(hostnames))],
			Severity:      severities[rand.Intn(len(severities))],
			RawMessage:    "<86> sudo: pam_unix(sudo:session): session opened for user",
			IsBlacklisted: rand.Intn(10) == 0, // 10% chance true
		}
		idCounter++

		data, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshalling log:", err)
			continue
		}

		_, err = conn.Write(append(data, '\n'))
		if err != nil {
			fmt.Println("Error sending log:", err)
			break
		}

		fmt.Println("Sent:", string(data))
		time.Sleep(2 * time.Second)
	}
}
