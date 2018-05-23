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

// GetPlayerDeaths retrieves player deaths
func GetPlayerDeaths(db *sql.DB, p *Player, limit int) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT most_damage_by, time, level FROM player_deaths WHERE player_id = ? ORDER BY time DESC LIMIT ?", p.ID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse deaths
	for rows.Next() {
		death := &PlayerDeath{}
		if err := rows.Scan(&death.MostDamageBy, &death.Time, &death.Level); err != nil {
			return nil, err
		}
		deaths = append(deaths, death)
	}

	return deaths, nil
}

// GetPlayerDeathsByMonster retrieves deaths caused by a monster
func GetPlayerDeathsByMonster(db *sql.DB, m *xml.Monster, limit int) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT a.name, b.time, b.level FROM players a, player_deaths b WHERE LOWER(b.killed_by) = ? ORDER BY b.time DESC LIMIT ?", strings.ToLower(m.Description), limit)
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
