package model

const groups_state_yes = 1
const groups_state_no = 0

func GetGroupsAll(appId string) (groups []Groups) {
	//var groups = []Groups{}

	Db.Where("appid =?",appId).Where("state =?",groups_state_yes).Find(&groups)
	return
}
