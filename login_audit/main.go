package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type ParsedLog struct {
	ID            int32     `json:"id,omitempty"`
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

	hostnames := []string{"aiops9242", "appnode01", "srv-auth", "dev-machine"}
	users := []string{"root", "admin", "alice", "bob", "serviceuser"}
	outcomes := []string{"SUCCESS", "FAILURE"}
	severities := map[string]string{"SUCCESS": "info", "FAILURE": "warn"}

	var idCounter int32 = 1

	for {

		sleepDur := time.Duration(1+rand.Intn(2)) * time.Second

		username := users[rand.Intn(len(users))]
		hostname := hostnames[rand.Intn(len(hostnames))]
		outcome := outcomes[rand.Intn(len(outcomes))]

		raw := fmt.Sprintf(
			`<86> %s sshd[%d]: %s: pam_unix(sshd:session): session %s for user %s`,
			hostname,
			1000+rand.Intn(9000),
			hostname,
			lowerFirst(outcome),
			username,
		)

		entry := ParsedLog{
			ID:            idCounter,
			Timestamp:     time.Now().UTC(),
			EventCategory: "login.audit",
			SourceType:    "linux",
			Username:      username,
			Hostname:      hostname,
			Severity:      severities[outcome],
			RawMessage:    raw,
			IsBlacklisted: rand.Intn(10) == 0,
		}
		idCounter++

		data, err := json.Marshal(entry)
		if err != nil {
			fmt.Println("marshal error:", err)
			time.Sleep(sleepDur)
			continue
		}

		_, err = conn.Write(append(data, '\n'))
		if err != nil {
			fmt.Println("send error:", err)
			return
		}

		fmt.Println("Sent:", string(data))
		time.Sleep(sleepDur)
	}
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("%s%s", string(s[0]+32), s[1:])
}
