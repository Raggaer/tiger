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

// GetPlayerByName retrieves a character by the name
func GetPlayerByName(db *sql.DB, name string) (*Player, error) {
	player := Player{}
	row := db.QueryRow(`
		SELECT 
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
