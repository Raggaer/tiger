package models

import (
	"database/sql"
	"strings"
)

// Player defines a server character
type Player struct {
	ID                   int64
	Name                 string
	GroupID              int
	Account              *Account
	AccountID            int
	Level                int
	Vocation             int
	Health               int
	HealthMax            int
	Experience           int64
	Lookbody             int
	LookFeet             int
	LookHead             int
	LookLegs             int
	LookType             int
	LookAddons           int
	MagicLevel           int
	Mana                 int
	ManaMax              int
	ManaSpent            int
	Soul                 int
	TownID               int
	Posx                 int64
	Posy                 int64
	Posz                 int64
	Cap                  int64
	Sex                  int
	LastLogin            int64
	LastIP               int
	Save                 bool
	Skull                int
	SkullTime            int64
	LastLogout           int64
	Blessings            int
	OnlineTime           int64
	Deletion             int
	Balance              int
	OfflineTrainingTime  int64
	OfflineTrainingSkill int
	Stamina              int64
	SkillFist            int
	SkillFistTries       int
	SkillClub            int
	SkillClubTries       int
	SkillSword           int
	SkillSwordTries      int
	SkillAxe             int
	SkillAxeTries        int
	SkillDist            int
	SkillDistTries       int
	SkillShielding       int
	SkillShieldingTries  int
	SkillFishing         int
	SkillFishingTries    int
	Conditions           []byte
}

// GetTopPlayersBySkillFishing retrieves top players by skill fishing field
func GetTopPlayersBySkillFishing(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_fishing FROM players ORDER BY skill_fishing DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillFishing); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillShield retrieves top players by skill shield field
func GetTopPlayersBySkillShield(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_shielding FROM players ORDER BY skill_shielding DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillShielding); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillDist retrieves top players by skill dist field
func GetTopPlayersBySkillDist(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_dist FROM players ORDER BY skill_dist DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillDist); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillAxe retrieves top players by skill axe field
func GetTopPlayersBySkillAxe(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_axe FROM players ORDER BY skill_axe DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillAxe); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillSword retrieves top players by skill sword field
func GetTopPlayersBySkillSword(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_sword FROM players ORDER BY skill_sword DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillSword); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillClub retrieves top players by skill club field
func GetTopPlayersBySkillClub(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_club FROM players ORDER BY skill_club DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillClub); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersBySkillFist retrieves top players by skill fist field
func GetTopPlayersBySkillFist(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, skill_fist FROM players ORDER BY skill_fist DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.SkillFist); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersByExperience retrieves top players by experience field
func GetTopPlayersByExperience(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, level, experience FROM players ORDER BY experience DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.Level, &p.Experience); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetTopPlayersByMagicLevel retrieves top players by maglevel field
func GetTopPlayersByMagicLevel(db *sql.DB, limit int) ([]*Player, error) {
	players := []*Player{}

	rows, err := db.Query("SELECT name, maglevel FROM players ORDER BY maglevel DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Player{}
		if err := rows.Scan(&p.Name, &p.MagicLevel); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}

	return players, nil
}

// GetPlayerByName retrieves a character by the name
func GetPlayerByName(db *sql.DB, name string) (*Player, error) {
	player := Player{}
	row := db.QueryRow(`
		SELECT 
			id,
			name,
			group_id,
			account_id,
			level,
			vocation,
			health,
			healthmax,
			experience,
			lookbody,
			lookfeet,
			lookhead,
			looklegs,
			looktype,
			lookaddons,
			maglevel,
			stamina,
			skill_fist,
			skill_club,
			skill_sword,
			skill_axe,
			skill_dist,
			skill_shielding,
			skill_fishing
		FROM players WHERE LOWER(name) = ?
	`, strings.ToLower(name))
	if err := row.Scan(
		&player.ID,
		&player.Name,
		&player.GroupID,
		&player.AccountID,
		&player.Level,
		&player.Vocation,
		&player.Health,
		&player.HealthMax,
		&player.Experience,
		&player.Lookbody,
		&player.LookFeet,
		&player.LookHead,
		&player.LookLegs,
		&player.LookType,
		&player.LookAddons,
		&player.MagicLevel,
		&player.Stamina,
		&player.SkillFist,
		&player.SkillClub,
		&player.SkillSword,
		&player.SkillAxe,
		&player.SkillDist,
		&player.SkillShielding,
		&player.SkillFishing,
	); err != nil {
		return nil, err
	}
	return &player, nil
}
