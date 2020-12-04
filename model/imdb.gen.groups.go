package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GroupsMgr struct {
	*_BaseMgr
}

// GroupsMgr open func
func GroupsMgr(db *gorm.DB) *_GroupsMgr {
	if db == nil {
		panic(fmt.Errorf("GroupsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("groups"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GroupsMgr) GetTableName() string {
	return "groups"
}

// Get 获取
func (obj *_GroupsMgr) Get() (result Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GroupsMgr) Gets() (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_GroupsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_GroupsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAppid appid获取 应用id
func (obj *_GroupsMgr) WithAppid(appid uint8) Option {
	return optionFunc(func(o *options) { o.query["appid"] = appid })
}

// WithState state获取 状态(1:有效，0无效，9删除)
func (obj *_GroupsMgr) WithState(state int8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCtime ctime获取 创建时间
func (obj *_GroupsMgr) WithCtime(ctime int) Option {
	return optionFunc(func(o *options) { o.query["ctime"] = ctime })
}

// WithMtime mtime获取 修改时间
func (obj *_GroupsMgr) WithMtime(mtime time.Time) Option {
	return optionFunc(func(o *options) { o.query["mtime"] = mtime })
}

// GetByOption 功能选项模式获取
func (obj *_GroupsMgr) GetByOption(opts ...Option) (result Groups, err error) {
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
func (obj *_GroupsMgr) GetByOptions(opts ...Option) (results []*Groups, err error) {
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
func (obj *_GroupsMgr) GetFromID(id int64) (result Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 主键
func (obj *_GroupsMgr) GetBatchFromID(ids []int64) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_GroupsMgr) GetFromName(name string) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 名称
func (obj *_GroupsMgr) GetBatchFromName(names []string) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAppid 通过appid获取内容 应用id
func (obj *_GroupsMgr) GetFromAppid(appid uint8) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("appid = ?", appid).Find(&results).Error

	return
}

// GetBatchFromAppid 批量唯一主键查找 应用id
func (obj *_GroupsMgr) GetBatchFromAppid(appids []uint8) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("appid IN (?)", appids).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态(1:有效，0无效，9删除)
func (obj *_GroupsMgr) GetFromState(state int8) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找 状态(1:有效，0无效，9删除)
func (obj *_GroupsMgr) GetBatchFromState(states []int8) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromCtime 通过ctime获取内容 创建时间
func (obj *_GroupsMgr) GetFromCtime(ctime int) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime = ?", ctime).Find(&results).Error

	return
}

// GetBatchFromCtime 批量唯一主键查找 创建时间
func (obj *_GroupsMgr) GetBatchFromCtime(ctimes []int) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ctime IN (?)", ctimes).Find(&results).Error

	return
}

// GetFromMtime 通过mtime获取内容 修改时间
func (obj *_GroupsMgr) GetFromMtime(mtime time.Time) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime = ?", mtime).Find(&results).Error

	return
}

// GetBatchFromMtime 批量唯一主键查找 修改时间
func (obj *_GroupsMgr) GetBatchFromMtime(mtimes []time.Time) (results []*Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mtime IN (?)", mtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GroupsMgr) FetchByPrimaryKey(id int64) (result Groups, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
