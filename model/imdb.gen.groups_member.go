package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupsMemberMgr struct {
	*_BaseMgr
}

// GroupsMemberMgr open func
func GroupsMemberMgr(db *gorm.DB) *_GroupsMemberMgr {
	if db == nil {
		panic(fmt.Errorf("GroupsMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupsMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("groups_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupsMemberMgr) GetTableName() string {
	return "groups_member"
}

// Get 获取
func (obj *_GroupsMemberMgr) Get() (result GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupsMemberMgr) Gets() (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_GroupsMemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGroupsID groups_id获取 群id
func (obj *_GroupsMemberMgr) WithGroupsID(groupsID int64) Option {
	return optionFunc(func(o *options) { o.query["groups_id"] = groupsID })
}

// WithUId uid获取 用户id
func (obj *_GroupsMemberMgr) WithUId(uid int64) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithType type获取 类型（0:普通1群主）
func (obj *_GroupsMemberMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCtime ctime获取 创建时间
func (obj *_GroupsMemberMgr) WithCtime(ctime int) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// WithMtime mtime获取 修改时间
func (obj *_GroupsMemberMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// GetByOption 功能选项模式获取
func (obj *_GroupsMemberMgr) GetByOption(opts ...Option) (result GroupsMember, err error) {
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
func (obj *_GroupsMemberMgr) GetByOptions(opts ...Option) (results []*GroupsMember, err error) {
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

// GetFromID 通过id获取内容 主键
func (obj *_GroupsMemberMgr) GetFromID(id int64) (result GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_GroupsMemberMgr) GetBatchFromID(ids []int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromGroupsID 通过groups_id获取内容 群id
func (obj *_GroupsMemberMgr) GetFromGroupsID(groupsID int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("groups_id = ?", groupsID).Find(&results).Error

	return
}

// GetBatchFromGroupsID 批量唯一主键查找 群id
func (obj *_GroupsMemberMgr) GetBatchFromGroupsID(groupsIDs []int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("groups_id IN (?)", groupsIDs).Find(&results).Error

	return
}

// GetFromUId 通过uid获取内容 用户id
func (obj *_GroupsMemberMgr) GetFromUId(uid int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUId 批量唯一主键查找 用户id
func (obj *_GroupsMemberMgr) GetBatchFromUId(uids []int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid IN (?)", uids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型（0:普通1群主）
func (obj *_GroupsMemberMgr) GetFromType(_type int8) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 类型（0:普通1群主）
func (obj *_GroupsMemberMgr) GetBatchFromType(_types []int8) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_GroupsMemberMgr) GetFromCtime(ctime int) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量唯一主键查找 创建时间
func (obj *_GroupsMemberMgr) GetBatchFromCtime(ctimes []int) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime IN (?)", ctimes).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容 修改时间
func (obj *_GroupsMemberMgr) GetFromMtime(mtime time.Time) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量唯一主键查找 修改时间
func (obj *_GroupsMemberMgr) GetBatchFromMtime(mtimes []time.Time) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime IN (?)", mtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupsMemberMgr) FetchByPrimaryKey(id int64) (result GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByUId  获取多个内容
func (obj *_GroupsMemberMgr) FetchIndexByUId(uid int64) (results []*GroupsMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid = ?", uid).Find(&results).Error

	return
}
