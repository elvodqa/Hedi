package database

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elvodqa/hedi/helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       int64
	Username     string
	Password     string
	Country      uint64
	Banned       int8
	BannedReason string
	Privileges   int32
	JoinedAt     string
}

type Session struct {
	UserID      uint64
	Online      uint8
	WorldID     uint64
	PrevWorldId uint64

	PositionX int64
	PositionY int64

	ItemHead  uint64
	ItemTorso uint64
	ItemLegs  uint64
	ItemShoes uint64

	ItemHeadSecondary  uint64
	ItemTorsoSecondary uint64
	ItemLegsSecondary  uint64
	ItemShoesSecondary uint64

	WeaponPrimary   uint64
	WeaponSecondary uint64
	Pet             uint64
}

func UserFromDatabaseById(id uint64) (int8, User) {
	returnUser := User{}

	queryResult, queryErr := database.Query("SELECT userid, username, password, country, banned, banned_reason, privileges, joinedat FROM userinfo WHERE userid = ?", id)

	if queryErr != nil {
		helpers.Logger.Printf("[Database] Failed to Fetch User from Database, MySQL query failed.\n")

		if queryResult != nil {
			queryResult.Close()
		}

		return -2, returnUser
	}

	if queryResult.Next() {
		scanErr := queryResult.Scan(&returnUser.UserID, &returnUser.Username, &returnUser.Password, &returnUser.Country, &returnUser.Banned, &returnUser.BannedReason, &returnUser.JoinedAt)

		queryResult.Close()

		if scanErr != nil {
			// log

			return -2, returnUser
		}

		return 0, returnUser
	}

	queryResult.Close()
	// User not found
	return -1, returnUser
}

func UserFromDatabaseByUsername(username string) (int8, User) {
	returnUser := User{}

	queryResult, queryErr := database.Query("SELECT userid, username, password, country, banned, banned_reason, privileges, joined_at FROM userinfo WHERE username = ?", username)

	if queryErr != nil {
		// log

		if queryResult != nil {
			queryResult.Close()
		}

		return -2, returnUser
	}

	//If there is a result
	if queryResult.Next() {
		scanErr := queryResult.Scan(&returnUser.UserID, &returnUser.Username, &returnUser.Password, &returnUser.Country, &returnUser.Banned, &returnUser.BannedReason, &returnUser.Privileges, &returnUser.JoinedAt)

		queryResult.Close()

		if scanErr != nil {
			// log

			return -2, returnUser
		}

		return 0, returnUser
	}

	queryResult.Close()
	//User not found
	return -1, returnUser
}

func CreateNewUser(username string, rawPassword string) bool {
	duplicateUsernameQuery, duplicateUsernameQueryErr := database.Query("SELECT COUNT(*) FROM userinfo WHERE username = ?", username)

	if duplicateUsernameQueryErr != nil {
		if duplicateUsernameQuery != nil {
			duplicateUsernameQuery.Close()
		}

		helpers.Logger.Printf("[Database] Failed to create new user, MySQL query failed.\n")

		return false
	}

	if duplicateUsernameQuery.Next() {
		var count uint64

		scanErr := duplicateUsernameQuery.Scan(&count)

		duplicateUsernameQuery.Close()

		if count != 0 || scanErr != nil {
			return false
		}
	}

	passwordHashed := md5.Sum([]byte(rawPassword))
	passwordHashedString := hex.EncodeToString(passwordHashed[:])
	bcryptPassword, bcryptErr := bcrypt.GenerateFromPassword([]byte(passwordHashedString), bcrypt.DefaultCost)

	if bcryptErr != nil {
		return false
	}

	var newUserId uint64
	var newUsername string

	insertResult, queryErrInsert := database.Query("INSERT INTO userinfo (username, password) VALUES (?, ?)", username, bcryptPassword)
	queryResult, queryErrGet := database.Query("SELECT user_id, username FROM userinfo WHERE username = ?", username)

	insertResult.Close()

	if queryErrInsert != nil || queryErrGet != nil {
		helpers.Logger.Printf("[Database] Failed to create new user, MySQL query failed.\n")

		return false
	}

	if queryResult.Next() {
		scanErr := queryResult.Scan(&newUserId, &newUsername)

		queryResult.Close()

		if scanErr != nil {
			return false
		}

		// TODO: DO SOMETHING USEFUL HERE

	} else {
		return false
	}

	return true
}

func AuthenticateUser(username string, password string) (userId int32, authSuccess bool) {
	query, queryErr := database.Query("SELECT user_id, username, password FROM userinfo WHERE username = ?", username)

	var scanUsername, scanPassword string
	var scanUserId int32

	if queryErr != nil {
		if query != nil {
			query.Close()
		}

		return -2, false
	}

	if query.Next() {
		scanErr := query.Scan(&scanUserId, &scanUsername, &scanPassword)

		query.Close()

		if scanErr != nil {
			return -2, false
		}

		if bcrypt.CompareHashAndPassword([]byte(scanPassword), []byte(password)) == nil {
			return scanUserId, true
		} else {
			return scanUserId, false
		}
	} else {
		return -1, false
	}
}
