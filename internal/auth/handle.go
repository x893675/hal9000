package auth

import (
	"context"
	"hal9000/internal/auth/models"
	"hal9000/pkg/pb"
	"hal9000/pkg/utils/ctxutils"
)

func (a *AuthhService) GetOauthToken(ctx context.Context, req *pb.Oauth2Request) (*pb.Token, error) {
	return &pb.Token{}, nil
}
func (a *AuthhService) GetAuthToken(ctx context.Context, req *pb.AuthRequest) (*pb.Token, error) {
	token, err := models.Login(ctx, req.Username, req.Password, ctxutils.GetClientIP(ctx))
	if err != nil {
		return nil, err
	}
	return token, nil
}
