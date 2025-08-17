package data

import (
	"context"
	"gorm.io/gorm"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todoRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todoRepo) Save(ctx context.Context, t *biz.Todo) (*biz.Todo, error) {
	t.MessageInfo = t.MessageInfo
	err := r.data.db.Create(t).Error
	return t, err
}

func (r *todoRepo) Update(ctx context.Context, t *biz.Todo) (*biz.Todo, error) {
	// 根据 ID 更新整个记录
	if err := r.data.db.Model(&biz.Todo{}).Where("id = ?", t.ID).Updates(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *todoRepo) FindByID(ctx context.Context, id int64) (*biz.Todo, error) {
	var todo biz.Todo
	// 根据主键 ID 查询
	if err := r.data.db.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 或返回自定义错误，表示未找到
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepo) DeleteByID(ctx context.Context, id int64) error {
	// 根据 ID 删除记录
	if err := r.data.db.Delete(&biz.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *todoRepo) ListByHello(context.Context, string) ([]*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) ListAll(context.Context) ([]*biz.Todo, error) {
	var todos []*biz.Todo
	if err := r.data.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
