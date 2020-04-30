package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Public data query methods

func (db *DataStorage) GetPlayerSummary(steamID string) (PlayerSummary, error) {
	//TODO implement
	// rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	// var id int
	// var firstname string
	// var lastname string
	// for rows.Next() {
	// 	rows.Scan(&id, &firstname, &lastname)
	// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	// }
	return PlayerSummary{}, nil
}

func (db *DataStorage) GetUserStatsForGame(steamID string) (UserStatsForGame, error) {
	//TODO implement
	return UserStatsForGame{}, nil
}

func (db *DataStorage) GetRecentlyPlayedGames(steamID string) (RecentlyPlayedGames, error) {
	//TODO implement
	return RecentlyPlayedGames{}, nil
}

func (db *DataStorage) GetPlayerHistory(steamID string, numPoints int) (PlayerHistory, error) {
	//TODO implement
	return PlayerHistory{}, nil
}

// Private data retrieval methods

func (db *DataStorage) updatePlayerSummary(steamID string) {
	// statement.Exec("Nic", "Raboy")
}

func (db *DataStorage) updateUserStatsForGame(steamID string) {
}

func (db *DataStorage) updateRecentlyPlayedGames(steamID string) {
}

// Initialization
type DataStorage struct {
	db         *sql.DB
	statements map[string]sql.Stmt
}

func NewDB(path string) (*DataStorage, error) {
	//TODO implement
	var err error
	storage := new(DataStorage)
	database, _ := sql.Open("sqlite3", path)

	// Prepare all statements
	statements := make(map[string]*sql.Stmt)

	if statements["create_player_summary"], err = database.Prepare(
		`CREATE TABLE IF NOT EXISTS player_summary (
			steamid INTEGER PRIMARY KEY,
			communityvisibilitystate INTEGER,
			profilestate INTEGER,
			personaname TEXT,
			profileurl TEXT,
			avatar TEXT,
			avatarmedium TEXT,
			avatarfull TEXT,
			lastlogoff INTEGER,
			personastate INTEGER,
			primaryclanid INTEGER,
			timecreated INTEGER,
			personastateflags INTEGER)`); err != nil {
		log.Fatal(err)
	}

	if statements["create_player_stats"], err = database.Prepare(
		`CREATE TABLE IF NOT EXISTS player_stats (
			steamid INTEGER PRIMARY KEY,
			total_kills INTEGER,
			total_deaths INTEGER,
			total_time_played INTEGER,
			total_planted_bombs INTEGER,
			total_defused_bombs INTEGER,
			total_wins INTEGER,
			total_damage_done INTEGER,
			total_money_earned INTEGER,
			total_kills_knife INTEGER,
			total_kills_hegrenade INTEGER,
			total_kills_glock INTEGER,
			total_kills_deagle INTEGER,
			total_kills_elite INTEGER,
			total_kills_fiveseven INTEGER,
			total_kills_xm1014 INTEGER,
			total_kills_mac10 INTEGER,
			total_kills_ump45 INTEGER,
			total_kills_p90 INTEGER,
			total_kills_awp INTEGER,
			total_kills_ak47 INTEGER,
			total_kills_aug INTEGER,
			total_kills_famas INTEGER,
			total_kills_g3sg1 INTEGER,
			total_kills_m249 INTEGER,
			total_kills_headshot INTEGER,
			total_kills_enemy_weapon INTEGER,
			total_wins_pistolround INTEGER,
			total_wins_map_cs_assault INTEGER,
			total_wins_map_de_dust2 INTEGER,
			total_wins_map_de_inferno INTEGER,
			total_wins_map_de_train INTEGER,
			total_weapons_donated INTEGER,
			total_kills_enemy_blinded INTEGER,
			total_kills_knife_fight INTEGER,
			total_kills_against_zoomed_sniper INTEGER,
			total_dominations INTEGER,
			total_domination_overkills INTEGER,
			total_revenges INTEGER,
			total_shots_hit INTEGER,
			total_shots_fired INTEGER,
			total_rounds_played INTEGER,
			total_shots_deagle INTEGER,
			total_shots_glock INTEGER,
			total_shots_elite INTEGER,
			total_shots_fiveseven INTEGER,
			total_shots_awp INTEGER,
			total_shots_ak47 INTEGER,
			total_shots_aug INTEGER,
			total_shots_famas INTEGER,
			total_shots_g3sg1 INTEGER,
			total_shots_p90 INTEGER,
			total_shots_mac10 INTEGER,
			total_shots_ump45 INTEGER,
			total_shots_xm1014 INTEGER,
			total_shots_m249 INTEGER,
			total_hits_deagle INTEGER,
			total_hits_glock INTEGER,
			total_hits_elite INTEGER,
			total_hits_fiveseven INTEGER,
			total_hits_awp INTEGER,
			total_hits_ak47 INTEGER,
			total_hits_aug INTEGER,
			total_hits_famas INTEGER,
			total_hits_g3sg1 INTEGER,
			total_hits_p90 INTEGER,
			total_hits_mac10 INTEGER,
			total_hits_ump45 INTEGER,
			total_hits_xm1014 INTEGER,
			total_hits_m249 INTEGER,
			total_rounds_map_cs_assault INTEGER,
			total_rounds_map_de_dust2 INTEGER,
			total_rounds_map_de_inferno INTEGER,
			total_rounds_map_de_train INTEGER,
			last_match_t_wins INTEGER,
			last_match_ct_wins INTEGER,
			last_match_wins INTEGER,
			last_match_max_players INTEGER,
			last_match_kills INTEGER,
			last_match_deaths INTEGER,
			last_match_mvps INTEGER,
			last_match_favweapon_id INTEGER,
			last_match_favweapon_shots INTEGER,
			last_match_favweapon_hits INTEGER,
			last_match_favweapon_kills INTEGER,
			last_match_damage INTEGER,
			last_match_money_spent INTEGER,
			last_match_dominations INTEGER,
			last_match_revenges INTEGER,
			total_mvps INTEGER,
			total_rounds_map_de_lake INTEGER,
			total_rounds_map_de_safehouse INTEGER,
			total_rounds_map_de_bank INTEGER,
			total_TR_planted_bombs INTEGER,
			total_gun_game_rounds_won INTEGER,
			total_gun_game_rounds_played INTEGER,
			total_wins_map_de_bank INTEGER,
			total_wins_map_de_lake INTEGER,
			total_matches_won_bank INTEGER,
			total_matches_won INTEGER,
			total_matches_played INTEGER,
			total_gg_matches_won INTEGER,
			total_gg_matches_played INTEGER,
			total_progressive_matches_won INTEGER,
			total_trbomb_matches_won INTEGER,
			total_contribution_score INTEGER,
			last_match_contribution_score INTEGER,
			last_match_rounds INTEGER,
			total_kills_hkp2000 INTEGER,
			total_shots_hkp2000 INTEGER,
			total_hits_hkp2000 INTEGER,
			total_hits_p250 INTEGER,
			total_kills_p250 INTEGER,
			total_shots_p250 INTEGER,
			total_kills_sg556 INTEGER,
			total_shots_sg556 INTEGER,
			total_hits_sg556 INTEGER,
			total_hits_scar20 INTEGER,
			total_kills_scar20 INTEGER,
			total_shots_scar20 INTEGER,
			total_shots_ssg08 INTEGER,
			total_hits_ssg08 INTEGER,
			total_kills_ssg08 INTEGER,
			total_shots_mp7 INTEGER,
			total_hits_mp7 INTEGER,
			total_kills_mp7 INTEGER,
			total_kills_mp9 INTEGER,
			total_shots_mp9 INTEGER,
			total_hits_mp9 INTEGER,
			total_hits_nova INTEGER,
			total_kills_nova INTEGER,
			total_shots_nova INTEGER,
			total_hits_negev INTEGER,
			total_kills_negev INTEGER,
			total_shots_negev INTEGER,
			total_shots_sawedoff INTEGER,
			total_hits_sawedoff INTEGER,
			total_kills_sawedoff INTEGER,
			total_shots_bizon INTEGER,
			total_hits_bizon INTEGER,
			total_kills_bizon INTEGER,
			total_kills_tec9 INTEGER,
			total_shots_tec9 INTEGER,
			total_hits_tec9 INTEGER,
			total_shots_mag7 INTEGER,
			total_hits_mag7 INTEGER,
			total_kills_mag7 INTEGER,
			total_gun_game_contribution_score INTEGER,
			last_match_gg_contribution_score INTEGER,
			total_kills_m4a1 INTEGER,
			total_kills_galilar INTEGER,
			total_kills_molotov INTEGER,
			total_kills_taser INTEGER,
			total_shots_m4a1 INTEGER,
			total_shots_galilar INTEGER,
			total_shots_taser INTEGER,
			total_hits_m4a1 INTEGER,
			total_hits_galilar INTEGER,
			total_matches_won_train INTEGER,
			total_matches_won_lake INTEGER,
			GI_lesson_csgo_instr_explain_buymenu INTEGER,
			GI_lesson_csgo_instr_explain_buyarmor INTEGER,
			GI_lesson_csgo_instr_explain_plant_bomb INTEGER,
			GI_lesson_csgo_instr_explain_bomb_carrier INTEGER,
			GI_lesson_bomb_sites_A INTEGER,
			GI_lesson_defuse_planted_bomb INTEGER,
			GI_lesson_csgo_instr_explain_follow_bomber INTEGER,
			GI_lesson_csgo_instr_explain_pickup_bomb INTEGER,
			GI_lesson_csgo_instr_explain_prevent_bomb_pickup INTEGER,
			GI_lesson_Csgo_cycle_weapons_kb INTEGER,
			GI_lesson_csgo_instr_explain_zoom INTEGER,
			GI_lesson_csgo_instr_explain_reload INTEGER,
			GI_lesson_tr_explain_plant_bomb INTEGER,
			GI_lesson_bomb_sites_B INTEGER,
			GI_lesson_version_number INTEGER,
			GI_lesson_find_planted_bomb INTEGER,
			GI_lesson_csgo_hostage_lead_to_hrz INTEGER,
			GI_lesson_csgo_instr_rescue_zone INTEGER,
			GI_lesson_csgo_instr_explain_inspect INTEGER,
			steam_stat_xpearnedgames INTEGER)`); err != nil {
		log.Fatal(err)
	}

	if statements["create_recently_played"], err = database.Prepare(
		`CREATE TABLE IF NOT EXISTS recently_played (
			steamid INTEGER PRIMARY KEY,
			playtime_2weeks INTEGER,
			playtime_forever INTEGER,
			playtime_windows_forever INTEGER,
			playtime_mac_forever INTEGER,
			playtime_linux_forever INTEGER)`); err != nil {
		log.Fatal(err)
	}

	// TODO add all fields for which we want historical info
	if statements["create_player_history"], err = database.Prepare(
		`CREATE TABLE IF NOT EXISTS player_history (
			steamid INTEGER,
			time INTEGER,
			total_kills INTEGER)`); err != nil {
		log.Fatal(err)
	}

	// - update recently_played
	if statements["update_recently_played"], err = database.Prepare(
		`UPDATE recently_played SET
			playtime_2weeks = ?,
			playtime_forever= ?,
			playtime_windows_forever= ?,
			playtime_mac_forever = ?,
			playtime_linux_forever = ?
			WHERE steamid = ?`); err != nil {
		log.Fatal(err)
	}

	// - insert history
	if statements["insert_history"], err = database.Prepare(
		`INSERT INTO player_history (
			steamid,
			time,
			total_kills
		) VALUES (?, ?, ?)`); err != nil {
		log.Fatal(err)
	}

	// - update player_summary
	if statements["update_player_summary"], err = database.Prepare(
		`UPDATE player_summary SET
			communityvisibilitystate = ?,
			profilestate = ?,
			personaname = ?,
			profileurl = ?,
			avatar = ?,
			avatarmedium = ?,
			avatarfull = ?,
			lastlogoff = ?,
			personastate = ?,
			primaryclanid = ?,
			timecreated = ?,
			personastateflags = ?
			WHERE steamid = ?`); err != nil {
		log.Fatal(err)
	}
	// - TODO update player_stats
	// - TODO update history time (if everything else is unchanged)
	// - TODO query player_summary
	// - TODO query player_stats
	// - TODO query recently_played
	// - TODO query history

	// Create tables, if necessary
	// statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")

	storage.db = database
	return storage, nil
}