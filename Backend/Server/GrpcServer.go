package Server

import (
	"context"
	"errors"
	pb "github.com/rafaelbfs/eSkills/Domain/generated/pb/skills"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PersonGrpcServer struct {
	pb.UnimplementedPersonServiceServer
}

func (s *PersonGrpcServer) CreateNewPerson(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	if p != nil {
		return p, nil
	}
	return nil, errors.New("Person cannot be null")
}
func (s *PersonGrpcServer) SearchPerson(*pb.PersonSearchParams, pb.PersonService_SearchPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchPerson not implemented")
}
func (s *PersonGrpcServer) UpdatePerson(context.Context, *pb.PersonUpdate) (*pb.UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
