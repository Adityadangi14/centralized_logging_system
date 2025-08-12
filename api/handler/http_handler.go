package handler

import (
	"log"

	db "github.com/Adityadangi14/centralized_logging_system/api/db/gen"
	"github.com/Adityadangi14/centralized_logging_system/api/initilizers"
	"github.com/gofiber/fiber/v2"
)

func GetLogsHandler(c *fiber.Ctx) error {
	service := c.Query("service")
	level := c.Query("level")
	username := c.Query("username")
	isBlacklistedStr := c.Query("is.blacklisted")

	var logs []db.ParsedLog
	var err error

	switch {

	case service != "" && level == "" && username == "" && isBlacklistedStr == "":
		logs, err = initilizers.Q.ListLogsByService(c.Context(), service)

	case service == "" && level != "" && username == "" && isBlacklistedStr == "":
		logs, err = initilizers.Q.ListLogsBySeverity(c.Context(), level)

	case service != "" && level != "" && username == "" && isBlacklistedStr == "":
		logs, err = initilizers.Q.ListLogsByServiceAndSeverity(c.Context(), db.ListLogsByServiceAndSeverityParams{
			EventCategory: service,
			Severity:      level,
		})

	case service == "" && level == "" && username != "" && isBlacklistedStr != "":
		isBlacklisted := isBlacklistedStr == "true" || isBlacklistedStr == "1"
		logs, err = initilizers.Q.ListLogsByUsernameAndBlacklisted(c.Context(), db.ListLogsByUsernameAndBlacklistedParams{
			Username:      username,
			IsBlacklisted: isBlacklisted,
		})

	default:
		return c.Status(400).JSON(fiber.Map{"error": "invalid query parameters"})
	}

	if err != nil {
		log.Printf("Failed to fetch logs: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}

	return c.JSON(logs)
}
