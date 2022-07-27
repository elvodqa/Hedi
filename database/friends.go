package database

type FriendEntry struct {
	UserId   uint32
	FriendID uint32
}

func GetFriendList(userId uint64) (result int, friendList []FriendEntry) {
	var friends = []FriendEntry{}

	queryResult, queryErr := database.Query("SELECT userid, friendid FROM friends WHERE userid = ?", userId)

	if queryErr != nil {
		return -1, friends
	}

	for queryResult.Next() {
		friendEntry := FriendEntry{}

		queryResult.Scan(&friendEntry.UserId, &friendEntry.FriendID)

		friends = append(friends, friendEntry)
	}

	return 0, friends
}

func AddFriend(userId uint64, friendId uint64) bool {
	_, queryErr := database.Exec("INSERT INTO friends (userid, friendid) VALUES (?, ?)", userId, friendId)

	if queryErr != nil {
		return false
	}

	return true
}

func RemoveFriend(userId uint64, friendId uint64) bool {
	_, queryErr := database.Exec("DELETE FROM friends WHERE userid = ? AND friendid = ?", userId, friendId)

	if queryErr != nil {
		return false
	}

	return true
}
