package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MemberMgr struct {
	*_BaseMgr
}

// MemberMgr open func
func MemberMgr(db *gorm.DB) *_MemberMgr {
	if db == nil {
		panic(fmt.Errorf("MemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MemberMgr) GetTableName() string {
	return "member"
}

// Get 获取
func (obj *_MemberMgr) Get() (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MemberMgr) Gets() (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MemberMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 姓名
func (obj *_MemberMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithState state获取 状态(1:有效，0无效，9删除)
func (obj *_MemberMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCtime ctime获取 创建时间
func (obj *_MemberMgr) WithCtime(ctime int) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// WithMtime mtime获取
func (obj *_MemberMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// GetByOption 功能选项模式获取
func (obj *_MemberMgr) GetByOption(opts ...Option) (result Member, err error) {
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
func (obj *_MemberMgr) GetByOptions(opts ...Option) (results []*Member, err error) {
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
func (obj *_MemberMgr) GetFromID(id int64) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_MemberMgr) GetBatchFromID(ids []int64) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_MemberMgr) GetFromName(name string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 姓名
func (obj *_MemberMgr) GetBatchFromName(names []string) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态(1:有效，0无效，9删除)
func (obj *_MemberMgr) GetFromState(state int8) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找 状态(1:有效，0无效，9删除)
func (obj *_MemberMgr) GetBatchFromState(states []int8) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_MemberMgr) GetFromCtime(ctime int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量唯一主键查找 创建时间
func (obj *_MemberMgr) GetBatchFromCtime(ctimes []int) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime IN (?)", ctimes).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容
func (obj *_MemberMgr) GetFromMtime(mtime time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量唯一主键查找
func (obj *_MemberMgr) GetBatchFromMtime(mtimes []time.Time) (results []*Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime IN (?)", mtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MemberMgr) FetchByPrimaryKey(id int64) (result Member, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
