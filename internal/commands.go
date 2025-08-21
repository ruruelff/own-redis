package internal

import (
	"strings"
	"time"
)

func (db *DB) ProcessCommand(cmd string) string {
	parts := strings.Fields(strings.ToLower(cmd))
	if len(parts) == 0 {
		return "(error) empty command\n"
	}

	switch parts[0] {
	case "ping":
		return "PONG\n"

	case "set":
		if len(parts) < 3 {
			return "(error) ERR wrong number of arguments for 'set' command\n"
		}

		key := parts[1]
		value := parts[2]

		var ttl time.Duration
		if len(parts) > 4 && parts[3] == "px" {
			ms, err := time.ParseDuration(parts[4] + "ms")
			if err != nil {
				return "(error) ERR invalid expire time\n"
			}
			ttl = ms
		}

		db.Set(key, value, ttl)
		return "OK\n"

	case "get":
		if len(parts) != 2 {
			return "(error) ERR wrong number of arguments for 'get' command\n"
		}
		return db.Get(parts[1])

	default:
		return "(error) ERR unknown command\n"
	}
}
