package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupsMessageMgr struct {
	*_BaseMgr
}

// GroupsMessageMgr open func
func GroupsMessageMgr(db *gorm.DB) *_GroupsMessageMgr {
	if db == nil {
		panic(fmt.Errorf("GroupsMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupsMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("groups_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupsMessageMgr) GetTableName() string {
	return "groups_message"
}

// Get 获取
func (obj *_GroupsMessageMgr) Get() (result GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupsMessageMgr) Gets() (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GroupsMessageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMsgSerID msg_ser_id获取 服务端消息id
func (obj *_GroupsMessageMgr) WithMsgSerID(msgSerID int64) Option {
	return optionFunc(func(o *options) { o.query["msg_ser_id"] = msgSerID })
}

// WithUId uid获取 用户id
func (obj *_GroupsMessageMgr) WithUId(uid int64) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithReaded readed获取 读取状态（0:未读,1:已读）
func (obj *_GroupsMessageMgr) WithReaded(readed int8) Option {
	return optionFunc(func(o *options) { o.query["readed"] = readed })
}

// WithMtime mtime获取 修改时间
func (obj *_GroupsMessageMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// WithCtime ctime获取 创建时间
func (obj *_GroupsMessageMgr) WithCtime(ctime int) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// GetByOption 功能选项模式获取
func (obj *_GroupsMessageMgr) GetByOption(opts ...Option) (result GroupsMessage, err error) {
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
func (obj *_GroupsMessageMgr) GetByOptions(opts ...Option) (results []*GroupsMessage, err error) {
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

// GetFromID 通过id获取内容
func (obj *_GroupsMessageMgr) GetFromID(id int) (result GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_GroupsMessageMgr) GetBatchFromID(ids []int) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromMsgSerID 通过msg_ser_id获取内容 服务端消息id
func (obj *_GroupsMessageMgr) GetFromMsgSerID(msgSerID int64) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id = ?", msgSerID).Find(&results).Error

	return
}

// GetBatchFromMsgSerID 批量唯一主键查找 服务端消息id
func (obj *_GroupsMessageMgr) GetBatchFromMsgSerID(msgSerIDs []int64) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id IN (?)", msgSerIDs).Find(&results).Error

	return
}

// GetFromUId 通过uid获取内容 用户id
func (obj *_GroupsMessageMgr) GetFromUId(uid int64) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUId 批量唯一主键查找 用户id
func (obj *_GroupsMessageMgr) GetBatchFromUId(uids []int64) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid IN (?)", uids).Find(&results).Error

	return
}

// GetFromReaded 通过readed获取内容 读取状态（0:未读,1:已读）
func (obj *_GroupsMessageMgr) GetFromReaded(readed int8) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("readed = ?", readed).Find(&results).Error

	return
}

// GetBatchFromReaded 批量唯一主键查找 读取状态（0:未读,1:已读）
func (obj *_GroupsMessageMgr) GetBatchFromReaded(readeds []int8) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("readed IN (?)", readeds).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容 修改时间
func (obj *_GroupsMessageMgr) GetFromMtime(mtime time.Time) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量唯一主键查找 修改时间
func (obj *_GroupsMessageMgr) GetBatchFromMtime(mtimes []time.Time) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime IN (?)", mtimes).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_GroupsMessageMgr) GetFromCtime(ctime int) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量唯一主键查找 创建时间
func (obj *_GroupsMessageMgr) GetBatchFromCtime(ctimes []int) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime IN (?)", ctimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupsMessageMgr) FetchByPrimaryKey(id int) (result GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByMsgSerID  获取多个内容
func (obj *_GroupsMessageMgr) FetchIndexByMsgSerID(msgSerID int64, uid int64) (results []*GroupsMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("msg_ser_id = ? AND uid = ?", msgSerID, uid).Find(&results).Error

	return
}
