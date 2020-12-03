package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"context"
)

type _MessageMgr struct {
	*_BaseMgr
}

// MessageMgr open func
func MessageMgr(db *gorm.DB) *_MessageMgr {
	if db == nil {
		panic(fmt.Errorf("MessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MessageMgr) GetTableName() string {
	return "message"
}

// Get 获取
func (obj *_MessageMgr) Get() (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MessageMgr) Gets() (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_MessageMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMsgSerID msg_ser_id获取 服务端消息id
func (obj *_MessageMgr) WithMsgSerID(msgSerID uint64) Option {
	return optionFunc(func(o *options) { o.query["msg_ser_id"] = msgSerID })
}

// WithGroupsID groups_id获取 群id
func (obj *_MessageMgr) WithGroupsID(groupsID int) Option {
	return optionFunc(func(o *options) { o.query["groups_id"] = groupsID })
}

// WithFromUId from_uid获取 发送方UID
func (obj *_MessageMgr) WithFromUId(fromUId uint64) Option {
	return optionFunc(func(o *options) { o.query["from_uid"] = fromUId })
}

// WithToUId to_uid获取 接收方UID
func (obj *_MessageMgr) WithToUId(toUId int64) Option {
	return optionFunc(func(o *options) { o.query["to_uid"] = toUId })
}

// WithContent content获取 内容
func (obj *_MessageMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithType type获取 消息类型
func (obj *_MessageMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAppname appname获取 应用名称
func (obj *_MessageMgr) WithAppname(appname uint8) Option {
	return optionFunc(func(o *options) { o.query["appname"] = appname })
}

// WithSend send获取 发送状态
func (obj *_MessageMgr) WithSend(send bool) Option {
	return optionFunc(func(o *options) { o.query["send"] = send })
}

// WithReaded readed获取 读取状态
func (obj *_MessageMgr) WithReaded(readed bool) Option {
	return optionFunc(func(o *options) { o.query["readed"] = readed })
}

// WithExt ext获取 扩展字段
func (obj *_MessageMgr) WithExt(ext string) Option {
	return optionFunc(func(o *options) { o.query["ext"] = ext })
}

// WithState state获取 状态
func (obj *_MessageMgr) WithState(state bool) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithAdminID admin_id获取 审核人
func (obj *_MessageMgr) WithAdminID(adminID uint32) Option {
	return optionFunc(func(o *options) { o.query["admin_id"] = adminID })
}

// WithCtime ctime获取 创建时间
func (obj *_MessageMgr) WithCtime(ctime uint32) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// WithMtime mtime获取 更新时间
func (obj *_MessageMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// GetByOption 功能选项模式获取
func (obj *_MessageMgr) GetByOption(opts ...Option) (result Message, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MessageMgr) GetByOptions(opts ...Option) (results []*Message, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 ID
func (obj *_MessageMgr) GetFromID(id int64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找 ID
func (obj *_MessageMgr) GetBatchFromID(ids []int64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMsgSerID 通过msg_ser_id获取内容 服务端消息id
func (obj *_MessageMgr) GetFromMsgSerID(msgSerID uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id = ?", msgSerID).Find(&results).Error

	return
}

// GetBatchFromMsgSerID 批量唯一主键查找 服务端消息id
func (obj *_MessageMgr) GetBatchFromMsgSerID(msgSerIDs []uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id IN (?)", msgSerIDs).Find(&results).Error

	return
}

// GetFromGroupsID 通过groups_id获取内容 群id
func (obj *_MessageMgr) GetFromGroupsID(groupsID int) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("groups_id = ?", groupsID).Find(&results).Error

	return
}

// GetBatchFromGroupsID 批量唯一主键查找 群id
func (obj *_MessageMgr) GetBatchFromGroupsID(groupsIDs []int) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("groups_id IN (?)", groupsIDs).Find(&results).Error

	return
}

// GetFromFromUId 通过from_uid获取内容 发送方UID
func (obj *_MessageMgr) GetFromFromUId(fromUId uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_uid = ?", fromUId).Find(&results).Error

	return
}

// GetBatchFromFromUId 批量唯一主键查找 发送方UID
func (obj *_MessageMgr) GetBatchFromFromUId(fromUIds []uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_uid IN (?)", fromUIds).Find(&results).Error

	return
}

// GetFromToUId 通过to_uid获取内容 接收方UID
func (obj *_MessageMgr) GetFromToUId(toUId int64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_uid = ?", toUId).Find(&results).Error

	return
}

// GetBatchFromToUId 批量唯一主键查找 接收方UID
func (obj *_MessageMgr) GetBatchFromToUId(toUIds []int64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_uid IN (?)", toUIds).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_MessageMgr) GetFromContent(content string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找 内容
func (obj *_MessageMgr) GetBatchFromContent(contents []string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 消息类型
func (obj *_MessageMgr) GetFromType(_type bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 消息类型
func (obj *_MessageMgr) GetBatchFromType(_types []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromAppname 通过appname获取内容 应用名称
func (obj *_MessageMgr) GetFromAppname(appname uint8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("appname = ?", appname).Find(&results).Error

	return
}

// GetBatchFromAppname 批量唯一主键查找 应用名称
func (obj *_MessageMgr) GetBatchFromAppname(appnames []uint8) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("appname IN (?)", appnames).Find(&results).Error

	return
}

// GetFromSend 通过send获取内容 发送状态
func (obj *_MessageMgr) GetFromSend(send bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send = ?", send).Find(&results).Error

	return
}

// GetBatchFromSend 批量唯一主键查找 发送状态
func (obj *_MessageMgr) GetBatchFromSend(sends []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("send IN (?)", sends).Find(&results).Error

	return
}

// GetFromReaded 通过readed获取内容 读取状态
func (obj *_MessageMgr) GetFromReaded(readed bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("readed = ?", readed).Find(&results).Error

	return
}

// GetBatchFromReaded 批量唯一主键查找 读取状态
func (obj *_MessageMgr) GetBatchFromReaded(readeds []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("readed IN (?)", readeds).Find(&results).Error

	return
}

// GetFromExt 通过ext获取内容 扩展字段
func (obj *_MessageMgr) GetFromExt(ext string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ext = ?", ext).Find(&results).Error

	return
}

// GetBatchFromExt 批量唯一主键查找 扩展字段
func (obj *_MessageMgr) GetBatchFromExt(exts []string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ext IN (?)", exts).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态
func (obj *_MessageMgr) GetFromState(state bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找 状态
func (obj *_MessageMgr) GetBatchFromState(states []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromAdminID 通过admin_id获取内容 审核人
func (obj *_MessageMgr) GetFromAdminID(adminID uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("admin_id = ?", adminID).Find(&results).Error

	return
}

// GetBatchFromAdminID 批量唯一主键查找 审核人
func (obj *_MessageMgr) GetBatchFromAdminID(adminIDs []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("admin_id IN (?)", adminIDs).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_MessageMgr) GetFromCtime(ctime uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量唯一主键查找 创建时间
func (obj *_MessageMgr) GetBatchFromCtime(ctimes []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime IN (?)", ctimes).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容 更新时间
func (obj *_MessageMgr) GetFromMtime(mtime time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量唯一主键查找 更新时间
func (obj *_MessageMgr) GetBatchFromMtime(mtimes []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime IN (?)", mtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByMsgSerID  获取多个内容
func (obj *_MessageMgr) FetchIndexByMsgSerID(msgSerID uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id = ?", msgSerID).Find(&results).Error

	return
}

// FetchIndexByGroupsID  获取多个内容
func (obj *_MessageMgr) FetchIndexByGroupsID(groupsID int) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("groups_id = ?", groupsID).Find(&results).Error

	return
}

// FetchIndexByFromUId  获取多个内容
func (obj *_MessageMgr) FetchIndexByFromUId(fromUId uint64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_uid = ?", fromUId).Find(&results).Error

	return
}

// FetchIndexByToUId  获取多个内容
func (obj *_MessageMgr) FetchIndexByToUId(toUId int64) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_uid = ?", toUId).Find(&results).Error

	return
}
