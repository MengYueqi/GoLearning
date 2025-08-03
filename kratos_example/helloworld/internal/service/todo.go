package service

import (
	"context"
	"errors"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type TodoService struct {
	pb.UnimplementedTodoServer
	uc *biz.TodoUsecase
}

func NewTodoService(uc *biz.TodoUsecase) *TodoService {
	return &TodoService{uc: uc}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	if len(req.MassageInfo) == 0 {
		return &pb.CreateTodoReply{}, errors.New("massageInfo is empty")
	}

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
	return &pb.UpdateTodoReply{}, nil
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	return &pb.DeleteTodoReply{}, nil
}
func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	return &pb.ListTodoReply{}, nil
}
