package models

import (
	"database/sql"
	"strings"

	"github.com/raggaer/tiger/app/xml"
)

// PlayerDeath defines a death of a character
type PlayerDeath struct {
	PlayerID              int64
	Player                *Player
	Time                  int64
	Level                 int
	KilledBy              string
	IsPlayer              bool
	MostDamageBy          string
	MostDamageIsPlayer    bool
	Unjustified           bool
	MostDamageUnjustified bool
}

// GetPlayerDeathsByMonster retrieves deaths caused by a monster
func GetPlayerDeathsByMonster(db *sql.DB, m *xml.Monster) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT a.name, b.time, b.level FROM players a, player_deaths b WHERE LOWER(b.killed_by) = ? ORDER BY time DESC LIMIT 10", strings.ToLower(m.Description))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse deaths
	for rows.Next() {
		death := &PlayerDeath{
			Player: &Player{},
		}
		if err := rows.Scan(&death.Player.Name, &death.Time, &death.Level); err != nil {
			return nil, err
		}
		deaths = append(deaths, death)
	}

	return deaths, nil
}
