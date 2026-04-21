package adviser

import (
	"context"

	pb "github.com/effective-security/promptviser/api/pb"
)

// Submit data for analysis.
func (s *Service) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	// TODO: implement
	res := &pb.SubmitResponse{
		ID: "1234567890",
	}
	return res, nil
}
