package model

type Groups struct {
	ID int64 `gorm:"primaryKey;column:id;type:bigint(20);not null"` // 主键
}

func GetGroupsAll(appId string) (error ,[]Groups) {
    var groupsArr []Groups
    groupsArr=append(groupsArr,Groups{ID:1})
	return nil,groupsArr
}
