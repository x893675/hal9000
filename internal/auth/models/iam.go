package models

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"hal9000/internal/auth/authentication"
	"hal9000/pkg/pb"
	"time"
)

func Login(ctx context.Context, username, password, ip string) (*pb.Token, error) {
	//Fixme: check auth failed rate limit, move to business layer
	//redisClient, err := client.ClientSets().Redis()
	//if err != nil {
	//	return nil, err
	//}
	//records, err := redisClient.Keys(fmt.Sprintf("hal9000:authfailed:%s:*", username)).Result()
	//if err != nil {
	//	logger.Error(ctx, err.Error())
	//	return nil, err
	//}
	//if len(records) >= maxAuthFailed {
	//	return nil, grpcerr.New(ctx, grpcerr.ResourceExhausted, grpcerr.ErrorAuthRateLimitExceed, maxAuthFailed, len(records))
	//}

	//TODO: ldap auth
	//client, err := client.ClientSets().Ldap()
	//if err != nil {
	//	return nil, err
	//}
	//conn, err := client.NewConn()
	//if err != nil {
	//	return nil, err
	//}
	//defer conn.Close()
	//
	//userSearchRequest := ldap.NewSearchRequest(
	//	client.UserSearchBase(),
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//	fmt.Sprintf("(&(objectClass=inetOrgPerson)(|(uid=%s)(mail=%s)))", username, username),
	//	[]string{"uid", "mail"},
	//	nil,
	//)
	//
	//result, err := conn.Search(userSearchRequest)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(result.Entries) != 1 {
	//	logger.Error(ctx, "incorrect password! ldap result code is [%d]", ldap.LDAPResultInvalidCredentials)
	//	return nil, grpcerr.New(ctx, grpcerr.Unauthenticated, grpcerr.ErrorAuthFailure)
	//}
	//
	//uid := result.Entries[0].GetAttributeValue("uid")
	//email := result.Entries[0].GetAttributeValue("mail")
	//dn := result.Entries[0].DN
	//
	//// bind as the user to verify their password
	//err = conn.Bind(dn, password)
	//
	//if err != nil {
	//	logger.Info(ctx, "[%v] auth failed, err is [%v]", username, err.Error())
	//	if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
	//		loginFailedRecord := fmt.Sprintf("hal9000:authfailed:%s:%d", uid, time.Now().UnixNano())
	//		redisClient.Set(loginFailedRecord, "", authTimeInterval)
	//	}
	//	return nil, err
	//}
	//
	//claims := jwt.MapClaims{}
	//
	//// token without expiration time will auto sliding
	//claims["username"] = uid
	//claims["email"] = email
	//claims["iat"] = time.Now().Unix()
	//
	//token := authentication.MustSigned(claims)
	//
	//if !enableMultiLogin {
	//	// multi login not allowed, remove the previous token
	//	sessions, err := redisClient.Keys(fmt.Sprintf("hal9000:users:%s:token:*", uid)).Result()
	//	if err != nil {
	//		logger.Error(ctx, "redis get multi login session error, err is [%v]",err.Error())
	//		//internal service error
	//		return nil, err
	//	}
	//	if len(sessions) > 0 {
	//		logger.Info(ctx, "revoke token, [%v]", sessions)
	//		err = redisClient.Del(sessions...).Err()
	//		if err != nil{
	//			logger.Error(ctx, "redis del sessions error [%v]", err.Error())
	//			return nil, err
	//		}
	//	}
	//}
	//
	//// cache token with expiration time
	//if err = redisClient.Set(fmt.Sprintf("hal9000:users:%s:token:%s", uid, token), token, tokenIdleTimeout).Err(); err != nil {
	//	logger.Error(ctx, "redis set sessions error [%v]", err.Error())
	//	return nil, err
	//}
	//
	//loginLog(uid, ip)

	//Fixme: only for test
	claims := jwt.MapClaims{}
	claims["username"] = "hal9000"
	claims["email"] = "hal9000@example.com"
	claims["iat"] = time.Now().Unix()
	token := authentication.MustSigned(claims)

	return &pb.Token{
		TokenType:   "bearer",
		ExpiresIn:   7200,
		AccessToken: token,
	}, nil
}

//oauth client credential model
func ClientCredentialGrant() {

}