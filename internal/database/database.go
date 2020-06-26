package database

import (
	"database/sql"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	// Use sqlite backend
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pinpox/megaclan3000/internal/steamclient"
)

// DataStorage is the main interface to the saved data. It provides methods for
// retrieval as well as methods to ingress new data from the API or update
// existing values
type DataStorage struct {
	db         *sqlx.DB
	dbm        *modl.DbMap
	statements map[string]*sql.Stmt
}

func (ds *DataStorage) getPlayerSummary(steamID string) (steamclient.PlayerSummary, error) {

	summary := steamclient.PlayerSummary{}
	var err error

	if err = ds.db.Get(&summary, "SELECT * FROM player_summary WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warnf("Error retrieving player_summary for steamID %v: %v", steamID, err)
	}

	return summary, err
}

func (ds *DataStorage) getRecentlyPlayedGames(steamID string) (steamclient.RecentlyPlayedGames, error) {

	recentGames := steamclient.RecentlyPlayedGames{}
	var err error

	if err = ds.db.Get(&recentGames, "SELECT * FROM recently_played WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warnf("Error retrieving recently_played for steamID %v: %v", steamID, err)
	}

	return recentGames, err
}

func (ds *DataStorage) getUserStatsForGameStats(steamID string) (steamclient.GameStats, error) {
	var err error
	gamestats := steamclient.GameStats{}
	if err = ds.db.Get(&gamestats, "SELECT * FROM player_stats WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warnf("Error retrieving player_stats for steamID %v: %v", steamID, err)
		return gamestats, err
	}
	gamestats.SteamID = steamID
	return gamestats, err
}

func (ds *DataStorage) getUserStatsForGameExtra(steamID string) (steamclient.GameExtras, error) {

	var err error
	extra := steamclient.GameExtras{}

	if err = ds.db.Get(&extra, "SELECT * FROM player_extra WHERE steamid=? LIMIT 1", steamID); err != nil {
		log.Warnf("Error retrieving player_extra for steamID %v: %v", steamID, err)
		return extra, err
	}
	extra.SteamID = steamID
	return extra, err
}

func (ds *DataStorage) getUserStatsForGame(steamID string) (steamclient.UserStatsForGame, error) {

	var err error
	usfg := steamclient.UserStatsForGame{}

	if usfg.Stats, err = ds.getUserStatsForGameStats(steamID); err != nil {
		return steamclient.UserStatsForGame{}, err
	}

	if usfg.Extra, err = ds.getUserStatsForGameExtra(steamID); err != nil {
		return steamclient.UserStatsForGame{}, err
	}

	usfg.SteamID = steamID

	return usfg, err

}

func (ds *DataStorage) getPlayerHistory(steamID string) (steamclient.PlayerHistory, error) {

	entries := []steamclient.PlayerHistoryEntry{}
	history := steamclient.PlayerHistory{}
	var err error

	if err = ds.db.Select(&entries, "SELECT * FROM player_history WHERE steamid=? ORDER BY time LIMIT 10", steamID); err != nil {
		log.Warnf("Error retrieving player_history for steamID %v: %v", steamID, err)
	}

	history.SteamID = steamID
	history.Data = entries

	return history, err
}

// GetPlayerInfoBySteamID returns a PlayerInfo from a steamID. It will try to
// get the needed values from the database and return an error if steamID
// cannot be found in it.
func (ds *DataStorage) GetPlayerInfoBySteamID(steamID string) (steamclient.PlayerInfo, error) {

	if len(steamID) <= 1 {
		panic("Tried to get info for empty ID")
	}

	info := steamclient.PlayerInfo{}
	var err error

	if info.PlayerSummary, err = ds.getPlayerSummary(steamID); err != nil {
		panic(err)
	}

	if info.RecentlyPlayedGames, err = ds.getRecentlyPlayedGames(steamID); err != nil {
		panic(err)
	}

	if info.UserStatsForGame, err = ds.getUserStatsForGame(steamID); err != nil {
		panic(err)
	}

	if info.PlayerHistory, err = ds.getPlayerHistory(steamID); err != nil {
		panic(err)
	}

	return info, nil
}

// NewDataStorage creates a new DataStorage for a given sqlite database filepath
func NewDataStorage(pathStorage, pathSchema string) (*DataStorage, error) {
	var err error

	// Initialize database
	storage := new(DataStorage)
	storage.statements = make(map[string]*sql.Stmt)

	// Connect to database
	log.Debugf("Reading %v", pathStorage)
	if storage.db, err = sqlx.Open("sqlite3", pathStorage); err != nil {
		log.Fatal("Failed to open sqlite file", err)
	}

	// Read and execute schema from schema.sql
	schema, err := ioutil.ReadFile(pathSchema)
	if err != nil {
		log.Fatal(err)
	}

	dbm := modl.NewDbMap(storage.db.DB, modl.SqliteDialect{})

	dbm.AddTableWithName(steamclient.PlayerSummary{}, "player_summary").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.PlayerHistoryEntry{}, "player_history")
	dbm.AddTableWithName(steamclient.GameStats{}, "player_stats").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.GameExtras{}, "player_extra").SetKeys(false, "steamid")
	dbm.AddTableWithName(steamclient.RecentlyPlayedGames{}, "recently_played").SetKeys(false, "steamid")

	log.Debugf("Mapping for PlayerSummary{}: %v", dbm.TableFor(steamclient.PlayerSummary{}).TableName)

	if err := dbm.CreateTablesIfNotExists(); err != nil {
		log.Warn("Database not creatable: ", err)
		panic(err)
	}

	storage.dbm = dbm

	storage.db.MustExec(string(schema))

	return storage, nil
}

// GetAllPlayers returns a PlayerInfo object for all players known to the
// database
func (ds *DataStorage) GetAllPlayers(steamIDs []string) ([]steamclient.PlayerInfo, error) {

	var players []steamclient.PlayerInfo

	for _, v := range steamIDs {
		log.Debugf("Getting PlayerInfo from DB for ID: %v", v)
		if pi, err := ds.GetPlayerInfoBySteamID(v); err == nil {
			players = append(players, pi)
		} else {
			log.Fatal(err)
		}
	}

	return players, nil
}

// UpdatePlayerInfo receives a PlayerInfo object and updates the database entry
// for it's steamID
func (ds *DataStorage) UpdatePlayerInfo(pi steamclient.PlayerInfo) error {
	var err error

	if err = ds.UpdatePlayerSummary(pi.PlayerSummary); err != nil {
		log.Fatalf("Error saving PlayerSummary for %v (%v): %v", pi.PlayerSummary.SteamID, pi.PlayerSummary.Personaname, err)
		return err
	}
	if err = ds.UpdateRecentlyPlayedGames(pi.RecentlyPlayedGames); err != nil {
		log.Fatalf("Error saving RecentlyPlayedGames for %v (%v)", pi.RecentlyPlayedGames.SteamID, pi.PlayerSummary.Personaname)
		return err
	}
	if err = ds.UpdateUserStatsForGame(pi.UserStatsForGame); err != nil {
		log.Fatalf("Error saving UserStatsForGame for %v (%v)", pi.UserStatsForGame.SteamID, pi.PlayerSummary.Personaname)
		return err
	}

	return nil
}
