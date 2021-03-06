package db

import (
	"github.com/number571/hiddenlake/models"
	"github.com/number571/hiddenlake/settings"
)

func GetAllFriends(user *models.User) []string {
	var (
		friends  []string
		hashname string
	)
	id := GetUserId(user.Username)
	if id < 0 {
		return nil
	}
	rows, err := settings.DB.Query(
		"SELECT Hashname FROM Friend WHERE IdUser=$1",
		id,
	)
	if err != nil {
		panic("query 'getallclients' failed")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&hashname)
		if err != nil {
			return nil
		}
		friends = append(friends, hashname)
	}
	return friends
}
