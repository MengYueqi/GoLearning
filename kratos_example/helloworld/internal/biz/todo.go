package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Todo struct {
	ID          int64
	MessageInfo string
	Status      bool
}

// GreeterRepo is a Greater repo.
// biz 对数据操作层提出 TodoRepo 中规定的要求
type TodoRepo interface {
	Save(context.Context, *Todo) (*Todo, error)
	Update(context.Context, *Todo) (*Todo, error)
	FindByID(context.Context, int64) (*Todo, error)
	ListByHello(context.Context, string) ([]*Todo, error)
	DeleteByID(context.Context, int64) error
	ListAll(context.Context) ([]*Todo, error)
}

// TodoUsecase is a Todoer usecase.
type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *TodoUsecase) CreateTodo(ctx context.Context, t *Todo) (*Todo, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", t)
	return uc.repo.Save(ctx, t)
}
