package model

const groups_state_yes = 1
const groups_state_no = 0

func GetGroupsAll(appId string) (error ,[]Groups) {
	var groups []Groups
	result :=Db.Where("appid =?",appId).Where("state =?",groups_state_yes).Find(&groups)
	return result.Error,groups
}
