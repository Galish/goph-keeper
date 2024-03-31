package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetCardsList(ctx context.Context, _ *emptypb.Empty) (*pb.GetListResponse, error) {
	var response pb.GetListResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	cards, err := s.keeper.GetCards(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.List = make([]*pb.ListItem, len(cards))

	for i, c := range cards {
		response.List[i] = &pb.ListItem{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
		}
	}

	return &response, nil
}
