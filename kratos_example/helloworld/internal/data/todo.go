package data

import (
	"context"

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
	t.MessageInfo = "Hello" + t.MessageInfo
	return t, nil
}

func (r *todoRepo) Update(ctx context.Context, t *biz.Todo) (*biz.Todo, error) {
	return t, nil
}

func (r *todoRepo) FindByID(context.Context, int64) (*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) DeleteByID(context.Context, int64) error {
	return nil
}

func (r *todoRepo) ListByHello(context.Context, string) ([]*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) ListAll(context.Context) ([]*biz.Todo, error) {
	return nil, nil
}
