package user

import (
	"context"
	"log"
	api "msqrd/pkg/api"
	"msqrd/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

type GRPCServer struct {
	Store *Store
}

func (s *GRPCServer) Create(ctx context.Context, r *api.CreateRequest) (*api.CreateResponse, error) {
	hashpb, err := bcrypt.GenerateFromPassword([]byte(r.Password), 7)

	if err != nil {
		return nil, err
	}

	u := &model.User{
		Email:             r.Email,
		Username:          r.Username,
		EncryptedPassword: hashpb,
	}

	u, err = s.Store.CreateUser(ctx, u)
	if err != nil {
		log.Fatal(err)
	}

	return &api.CreateResponse{Id: int32(u.Id)}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, r *api.DeleteRequest) (*api.DeleteResponse, error) {
	res, err := s.Store.DeleteUser(ctx, r.GetId())
	if err != nil {
		return nil, err
	}

	return &api.DeleteResponse{Res: res}, nil
}
