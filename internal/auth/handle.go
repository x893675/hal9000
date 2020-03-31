package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"hal9000/internal/auth/authentication"
	"hal9000/internal/auth/models"
	"hal9000/internal/auth/models/entity"
	"hal9000/pkg/constants"
	"hal9000/pkg/grpcerr"
	"hal9000/pkg/pb"
	"hal9000/pkg/utils/ctxutils"
	"time"
)

func (a *AuthhService) GetOauthToken(ctx context.Context, req *pb.Oauth2Request) (*pb.Token, error) {
	switch req.GrantType {
	case constants.GrantTypeClientCredentials:
	case constants.GrantTypeRefreshToken:
		if req.RefreshToken == "" {
			return nil, grpcerr.New(ctx, grpcerr.InvalidArgument, grpcerr.ErrorParameterShouldNotBeEmpty, "refresh_token")
		}
	default:
		return nil, grpcerr.New(ctx, grpcerr.InvalidArgument, grpcerr.ErrorUnsupportedParameterValue, "grant_type", req.GrantType)
	}

	userClient, err := validateClientCredentials(ctx, req)
	if err != nil {
		return nil, err
	}

	userId := userClient.UserId
	var token *entity.ClientToken
	if req.GrantType == constants.GrantTypeRefreshToken {
		token, err = models.GetClientTokenByRefreshToken(ctx, req.RefreshToken)
		if err != nil {
			return nil, grpcerr.New(ctx, grpcerr.PermissionDenied, grpcerr.ErrorAuthFailure)
		}
		//Fixme: expire time move to config
		if token.CreateTime.Add(time.Duration(7200)).Unix() <= time.Now().Unix() {
			return nil, grpcerr.New(ctx, grpcerr.PermissionDenied, grpcerr.ErrorRefreshTokenExpired)
		}
	} else {
		// reuse exist token
		token, err = models.GetClientTokenClientId(ctx, userClient.ClientId, userId, req.Scope)
		if err != nil {
			return nil, grpcerr.NewWithDetail(ctx, grpcerr.Internal, err, grpcerr.ErrorInternalError)
		}
		// token not exists or expired
		if token == nil || token.CreateTime.Add(time.Duration(7200)).Unix() <= time.Now().Unix() {
			// generate access token
			token, err = models.CreateClientToken(ctx, userClient.ClientId, req.Scope, userId)
			if err != nil {
				return nil, grpcerr.NewWithDetail(ctx, grpcerr.Internal, err, grpcerr.ErrorInternalError)
			}
		}
	}

	//gen access_token
	claims := jwt.MapClaims{}
	claims["username"] = "hal9000"
	claims["email"] = "hal9000@example.com"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(7200)*time.Second).Unix()
	claims["sub"] = token.ClientId
	accessToken := authentication.MustSigned(claims)

	return &pb.Token{
		TokenType:            "bearer",
		ExpiresIn:            7200,
		AccessToken:          accessToken,
		RefreshToken:         token.RefreshToken,
		IdToken:              "",
	}, nil
}
func (a *AuthhService) GetAuthToken(ctx context.Context, req *pb.AuthRequest) (*pb.Token, error) {
	token, err := models.Login(ctx, req.Username, req.Password, ctxutils.GetClientIP(ctx))
	if err != nil {
		return nil, err
	}
	return token, nil
}



type clientIface interface {
	GetClientId() string
	GetClientSecret() string
}

func validateClientCredentials(ctx context.Context, client clientIface) (*entity.UserClient, error) {
	//TODO: verify client id and client secret
	userClient, err := models.GetUserClient(ctx, client.GetClientId(), client.GetClientSecret())
	if err != nil {
		return nil, err
	}
	return userClient, nil
}