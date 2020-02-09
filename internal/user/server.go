package user

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbUser "hal9000/pb/user"
	"sync"
	"time"
)

type UserServer struct {
	mu    *sync.RWMutex
	users []*pbUser.User
}

var _ pbUser.UserServiceServer = (*UserServer)(nil)

func New() *UserServer {
	return &UserServer{
		mu: &sync.RWMutex{},
	}
}

func (u *UserServer) AddUser(ctx context.Context, user *pbUser.User) (*types.Empty, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if len(u.users) == 0 && user.GetRole() != pbUser.Role_ADMIN {
		st := status.New(codes.InvalidArgument, "First user created must be an admin")
		detSt, err := st.WithDetails(&rpc.BadRequest{
			FieldViolations: []*rpc.BadRequest_FieldViolation{
				{
					Field:       "role",
					Description: "The first user created must have the role of an ADMIN",
				},
			},
		})
		if err == nil {
			return nil, detSt.Err()
		}
		return nil, st.Err()
	}

	// Check user ID doesn't already exist
	for _, u := range u.users {
		if u.GetID() == user.GetID() {
			return nil, status.Error(codes.FailedPrecondition, "user already exists")
		}
	}

	if user.GetCreateDate() == nil {
		now := time.Now()
		user.CreateDate = &now
	}

	u.users = append(u.users, user)

	return new(types.Empty), nil
}

func (u *UserServer) ListUsers(req *pbUser.ListUsersRequest, srv pbUser.UserService_ListUsersServer) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if len(u.users) == 0 {
		st := status.New(codes.FailedPrecondition, "No users have been created")
		detSt, err := st.WithDetails(&rpc.PreconditionFailure{
			Violations: []*rpc.PreconditionFailure_Violation{
				{
					Type:        "USER",
					Subject:     "no users created",
					Description: "No users have been created",
				},
			},
		})
		if err == nil {
			return detSt.Err()
		}
		return st.Err()
	}

	for _, user := range u.users {
		switch {
		case req.GetCreatedSince() != nil && user.GetCreateDate().Before(*req.GetCreatedSince()):
			continue
		case req.GetOlderThan() != nil && time.Since(*user.GetCreateDate()) <= *req.GetOlderThan():
			continue
		}
		err := srv.Send(user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UserServer) ListUsersByRole(req *pbUser.UserRole, srv pbUser.UserService_ListUsersByRoleServer) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, user := range u.users {
		if user.GetRole() == req.GetRole() {
			err := srv.Send(user)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	var user *pbUser.User
	for _, item := range u.users {
		if item.GetID() == req.GetUser().GetID() {
			user = item
		}
	}

	if user == nil {
		return nil, status.Error(codes.NotFound, "user was not found")
	}

	st := structs.New(user)
	for _, path := range req.GetUpdateMask().GetPaths() {
		if path == "id" {
			return nil, status.Error(codes.InvalidArgument, "cannot update id field")
		}
		// This doesn't translate properly if a CustomName setting is used,
		// but none of the fields except ID has that set, so NO WORRIES.
		fname := generator.CamelCase(path)
		field, ok := st.FieldOk(fname)
		if !ok {
			st := status.New(codes.InvalidArgument, "invalid field specified")
			st, err := st.WithDetails(&rpc.BadRequest{
				FieldViolations: []*rpc.BadRequest_FieldViolation{{
					Field:       "update_mask",
					Description: fmt.Sprintf("The user message type does not have a field called %q", path),
				}},
			})
			if err != nil {
				panic(err)
			}
			return nil, st.Err()
		}

		in := structs.New(req.GetUser())
		err := field.Set(in.Field(fname).Value())
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
