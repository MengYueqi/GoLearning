package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type TodoService struct {
	pb.UnimplementedTodoServer
	uc  *biz.TodoUsecase
	log *log.Helper
}

func NewTodoService(uc *biz.TodoUsecase, logger log.Logger) *TodoService {
	return &TodoService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	// 参数校验
	if len(req.MassageInfo) == 0 {
		return &pb.CreateTodoReply{}, errors.New("massageInfo is empty")
	}

	// 调用 biz 层
	data, err := s.uc.CreateTodo(ctx, &biz.Todo{MessageInfo: req.MassageInfo})
	if err != nil {
		return &pb.CreateTodoReply{}, err
	}
	return &pb.CreateTodoReply{
		Todo: &pb.TodoMessage{
			Id:          data.ID,
			MassageInfo: data.MessageInfo,
			States:      data.Status,
		},
	}, nil
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoReply, error) {
	// 对 ID 的长度进行审查
	if req.Id < 0 {
		return &pb.UpdateTodoReply{}, errors.New("id is 0")
	}
	_, err := s.uc.UpdateTodo(ctx, &biz.Todo{
		ID:          req.Id,
		MessageInfo: req.MassageInfo,
		Status:      req.States,
	})
	return &pb.UpdateTodoReply{}, err
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	err := s.uc.DeleteTodo(ctx, req.Id)
	return &pb.DeleteTodoReply{}, err
}
func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	data, err := s.uc.GetTodo(ctx, req.Id)
	if err != nil {
		return &pb.GetTodoReply{}, err
	}
	return &pb.GetTodoReply{
		Todo: &pb.TodoMessage{
			Id:          data.ID,
			MassageInfo: data.MessageInfo,
			States:      data.Status,
		},
	}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	data, err := s.uc.ListAll(ctx)
	if err != nil {
		return &pb.ListTodoReply{}, err
	}
	var allMassage []*pb.TodoMessage
	// 进行格式转换
	for _, m := range data {
		allMassage = append(allMassage, &pb.TodoMessage{
			Id:          m.ID,
			States:      m.Status,
			MassageInfo: m.MessageInfo,
		})
	}
	return &pb.ListTodoReply{
		Data: allMassage,
	}, nil
}
