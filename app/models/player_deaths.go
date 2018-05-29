package models

import (
	"database/sql"
	"strings"
	"time"

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

// GetTimeServerDeaths retrieves server deaths by the given time
func GetTimeServerDeaths(db *sql.DB, limit int, t time.Time) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT a.name, b.time, b.level, b.killed_by FROM players a, player_deaths b WHERE a.id = b.player_id AND b.time >= ? ORDER BY b.time DESC LIMIT ?", t.Unix(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse deaths
	for rows.Next() {
		death := &PlayerDeath{
			Player: &Player{},
		}
		if err := rows.Scan(&death.Player.Name, &death.Time, &death.Level, &death.KilledBy); err != nil {
			return nil, err
		}
		deaths = append(deaths, death)
	}

	return deaths, nil
}

// GetServerDeaths retrieves server deaths
func GetServerDeaths(db *sql.DB, limit int) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT a.name, b.time, b.level, b.killed_by FROM players a, player_deaths b WHERE a.id = b.player_id ORDER BY b.time DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse deaths
	for rows.Next() {
		death := &PlayerDeath{
			Player: &Player{},
		}
		if err := rows.Scan(&death.Player.Name, &death.Time, &death.Level, &death.KilledBy); err != nil {
			return nil, err
		}
		deaths = append(deaths, death)
	}

	return deaths, nil
}

// GetPlayerDeaths retrieves player deaths
func GetPlayerDeaths(db *sql.DB, p *Player, limit int) ([]*PlayerDeath, error) {
	deaths := []*PlayerDeath{}

	// Retrieve deaths using monster description
	rows, err := db.Query("SELECT killed_by, mostdamage_by, time, level FROM player_deaths WHERE player_id = ? ORDER BY time DESC LIMIT ?", p.ID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse deaths
	for rows.Next() {
		death := &PlayerDeath{}
		if err := rows.Scan(&death.KilledBy, &death.MostDamageBy, &death.Time, &death.Level); err != nil {
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
	rows, err := db.Query("SELECT a.name, b.time, b.level FROM players a, player_deaths b WHERE LOWER(b.killed_by) = ? AND a.id = b.player_id ORDER BY b.time DESC LIMIT ?", strings.ToLower(m.Description), limit)
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
