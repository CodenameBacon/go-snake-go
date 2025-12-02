package game

import "github.com/google/uuid"

func uniqueUsername(base string, players map[uuid.UUID]*Player) string {
	username := base
	for {
		conflict := false
		for _, p := range players {
			if p.username == username {
				conflict = true
				break
			}
		}
		if !conflict {
			return username
		}
		username += "*"
	}
}
