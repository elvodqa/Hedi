package database

import "github.com/elvodqa/hedi/helpers"

type World struct {
	WorldId     uint64
	MapId       uint64
	WorldName   string
	PlayerLimit uint8
	PlayerCount uint8
}

func CreateWorld(worldName string, mapId uint64, owner string) bool {
	duplicateWorldNameQuery, duplicateWorldNameQueryErr := database.Query("SELECT COUNT(*) FROM worlds WHERE worldname = ?", worldName)

	if duplicateWorldNameQueryErr != nil {
		if duplicateWorldNameQuery != nil {
			duplicateWorldNameQuery.Close()
		}

		helpers.Logger.Printf("[Database] Failed to create new world, MySQL query failed.\n")

		return false
	}

	if duplicateWorldNameQuery.Next() {
		var count uint64

		scanErr := duplicateWorldNameQuery.Scan(&count)

		duplicateWorldNameQuery.Close()

		if count != 0 || scanErr != nil {
			return false
		}
	}

	var newWorldId uint64
	var newWorldName string
	var newWorldOwner string

	insertResult, queryErrInsert := database.Query("INSERT INTO worlds (mapid, worldname, owner) VALUES (?, ?, ?)", mapId, worldName, owner)
	queryResult, queryErrGet := database.Query("SELECT worldid, username, owner FROM worlds WHERE worldname = ?", worldName)

	insertResult.Close()

	if queryErrInsert != nil || queryErrGet != nil {
		helpers.Logger.Printf("[Database] Failed to create new world, MySQL query failed.\n")

		return false
	}

	if queryResult.Next() {
		scanErr := queryResult.Scan(&newWorldId, &newWorldName, &newWorldOwner)

		queryResult.Close()

		if scanErr != nil {
			return false
		}

		// TODO: do something usefull with ID AND NAME dunno
		helpers.Logger.Printf("[USER] User %s created a new world named %s(ID: %d).\n", newWorldOwner, newWorldName, newWorldId)

	} else {
		return false
	}

	return true
}
