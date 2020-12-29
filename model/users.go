package model

type Users struct {
	ID    int64     `gorm:"primaryKey;column:id;type:bigint(20);not null"`
}

func GetUsersAll(groupsId string)  []Users{
	var userArr []Users
	userArr=append(userArr,Users{ID:1})
	userArr=append(userArr,Users{ID:2})
	return userArr
}
