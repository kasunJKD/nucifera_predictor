package jwt

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	pb "nucifera_backend/protos/membership"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")
var RedisConn *redis.Client

type Claims struct {
	LocalId string `json:"localId"`
	//jwt.StandardClaims // 'StandardClaims' is deprecated
	Issuer string `json:"iss,omitempty"`
	Subject string `json:"sub,omitempty"`
	Audience string `json:"aud,omitempty"`
	ExpiresAt int64 `json:"exp,omitempty"`
	NotBefore int64 `json:"nbf,omitempty"`
	IssuedAt int64 `json:"iat,omitempty"`
	Id string `json:"jti,omitempty"`
}

func (a AuthCodeToken) Valid() error {
	//TODO implement me
	panic("implement me")
}

func (c Claims) Valid() error {
	//TODO implement me
	panic("implement me")
}

//type MD map[string][]string

type Metadata struct {
	MdMainAccessToken string
}

func CreateToken(req *pb.Request) (*pb.Response, error) {
	refreshToken := req.RefreshToken
	expirationTime := time.Now().Add(5 * time.Minute)
	issuedTime := time.Now()

	//Create new accessToken
	claims := &Claims{
		LocalId: req.LocalId,
		Issuer: "kasun",
		Subject: "Test",
		Audience: "nucifero",
		ExpiresAt: expirationTime.Unix(),
		NotBefore: 0,
		IssuedAt: issuedTime.Unix(),
		Id: "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal("InternalServerError", err)
	}

	if refreshToken != "" {
		//Revoke the current refreshToken as we will create and return a new one
		_, redis_err := RevokeToken(context.Background(), &pb.Request{RefreshToken: refreshToken})
		if redis_err != nil {
			fmt.Println("Error: old refreshToken is not revoked; ", redis_err)
		}

		//Get the list of the all refresh tokens for the user
		val, redis_err := RedisConn.Get(context.Background(), req.LocalId).Result()

		if redis_err == nil && val != "" {
			//Remove the old revoked refreshToken from the user's refresh tokens list
			val = strings.Replace(val, ", " + refreshToken, "", 1)
			redisCmd := RedisConn.Set(context.Background(), req.LocalId, val, 0)
			if redisCmd.Err() != nil {
				log.Fatal("Error removing the revoked token from the user's refresh tokens list: ", redisCmd.Err())
			}
		}
	}

	//Create new refreshToken
	expirationTime = time.Now().Add(7 * 24 * time.Hour)

	claims = &Claims{
		LocalId: req.LocalId,
		Issuer: "kasun",
		Subject: "Test",
		Audience: "nucifero",
		ExpiresAt: expirationTime.Unix(),
		NotBefore: 0,
		IssuedAt: issuedTime.Unix(),
		Id: "",
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err = token.SignedString(jwtKey)

	if err != nil {
		log.Fatal("InternalServerError", err)
	}

	redisExp := 7 * 24 * time.Hour

	//Store the refreshToken in redis
	redisCmd := RedisConn.Set(context.Background(), refreshToken, req.LocalId, redisExp)

	if redisCmd.Err() != nil {
		log.Fatal("Error adding the token to redis: ", redisCmd.Err())
	}

	//Get the list of the all refresh tokens for the user
	val, redis_err := RedisConn.Get(context.Background(), req.LocalId).Result()

	if redis_err != nil || val == "" {
		//Create a list for the user's refresh tokens
		redisCmd = RedisConn.Set(context.Background(), req.LocalId, refreshToken, 0)
	} else {
		//Add the new refreshToken to the user's refresh tokens list
		val += ", " + refreshToken
		redisCmd = RedisConn.Set(context.Background(), req.LocalId, val, 0)
	}
	if redisCmd.Err() != nil {
		log.Fatal("Error adding the token to redis: ", redisCmd.Err())
	}

	expirationTime = time.Now().Add(5 * time.Minute)

	res := &pb.Response{
		//Users: &pb.Users{
		//	LocalId:          req.GetLocalId(),
		//	Email:            req.GetEmail(),
		//	EmailVerified:    req.GetEmailVerified(),
		//	DisplayName:      req.GetDisplayName(),
		//	PhotoUrl:         req.GetPhotoUrl(),
		//	ProviderUserInfo: &pb.ProviderUserInfo {
		//		ProviderId:       req.GetProviderId(),
		//		FederatedId:      req.GetFederatedId(),
		//		Email:            req.GetEmail(),
		//	},
		//},
		OauthAccessToken: tokenString,
		//FirstName:        req.GetFirstName(),
		//LastName:         req.GetLastName(),
		//FullName:         req.GetFullName(),
		//IdToken:          req.GetIdToken(),
		ExpiresIn:        expirationTime.String(),
		RefreshToken: 	  refreshToken,
	}

	return res, err
}

func VerifyToken(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	//isValid := true
	AccessToken := req.AccessToken
	claims := jwt.MapClaims{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Error: metadata is not provided")
	}

	found := false
	values := md["authorization"]
	// check the rest request header and grpc metadata
	if len(values) > 0 {
		AccessToken = values[0]
		found = true
	} else {
		//check the grpc metadata
		for key := range claims {
			if key == "authorization"  {
				AccessToken = md.Get("authorization")[0]
				found = true
				break
			}
		}
	}
	if !found {
		log.Println("authorization token is not provided")
	}

	token, err:= jwt.ParseWithClaims(AccessToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	fmt.Println(token)
	log.Println(AccessToken)

	if err != nil {
		// Refresh Token
		//if err.Error() == "Token is expired" {
		//	res, err := CreateToken(req)
		//	return res, err
		//}
		fmt.Println("Token is invalid")
		//isValid = false
	}

	// Identity Authentication
	//if req.LocalId != "" {
	//	if req.LocalId != claims["localId"] {
	//		fmt.Printf("Token is invalid; req.LocalId: %v, id not equal to jwtLocalId: %v\n", req.LocalId, claims["localId"])
	//		err = fmt.Errorf("identity authentication failed")
	//		//isValid = false
	//	}
	//}

	exp := fmt.Sprintf("%v", claims["exp"])

	res := &pb.Response{
		//Users: &pb.Users{
		//	LocalId:          req.GetLocalId(),
		//	Email:            req.GetEmail(),
		//	EmailVerified:    req.GetEmailVerified(),
		//	DisplayName:      req.GetDisplayName(),
		//	PhotoUrl:         req.GetPhotoUrl(),
		//	ProviderUserInfo: &pb.ProviderUserInfo {
		//		ProviderId:       req.GetProviderId(),
		//		FederatedId:      req.GetFederatedId(),
		//		Email:            req.GetEmail(),
		//	},
		//},
		OauthAccessToken: AccessToken,
		//FirstName:        req.GetFirstName(),
		//LastName:         req.GetLastName(),
		//FullName:         req.GetFullName(),
		//IdToken:          req.GetIdToken(),
		ExpiresIn:        exp,
		//IsValid: 		  isValid,
	}

	return res, err
}

func RefreshToken(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	//_, verify_err := VerifyToken(ctx, req)
	//
	//if verify_err == nil {

	refreshToken := req.RefreshToken

	if refreshToken == "" {
		err := fmt.Errorf("refresh token is not provided")
		return nil, err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	fmt.Println(token)
	log.Println(refreshToken)

	if err != nil {
		fmt.Println("refreshToken is invalid")
		return nil, err
	}

	//Check in redis if the refreshToken is revoked
	if err == nil {
		val, redis_err := RedisConn.Get(ctx, refreshToken).Result()

		if redis_err != nil || val == "" {
			fmt.Println("refreshToken is revoked; not found in redis")
			err = fmt.Errorf("refreshToken is revoked")
			return nil, err
		}
	}

	localId := fmt.Sprintf("%v", claims["localId"])

	// Create user new access token from id
	userToken, err := CreateToken(&pb.Request{RefreshToken: refreshToken, LocalId: localId})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate new access token")
	}

	res := &pb.Response{
		Users: &pb.Users{
			UserId: localId,
			//	Email:            req.GetEmail(),
			//	EmailVerified:    req.GetEmailVerified(),
			//	DisplayName:      req.GetDisplayName(),
			//	PhotoUrl:         req.GetPhotoUrl(),
			//	ProviderUserInfo: &pb.ProviderUserInfo {
			//		ProviderId:       req.GetProviderId(),
			//		FederatedId:      req.GetFederatedId(),
			//		Email:            req.GetEmail(),
			//	},
		},
		OauthAccessToken: userToken.GetOauthAccessToken(),
		//FirstName:        req.GetFirstName(),
		//LastName:         req.GetLastName(),
		//FullName:         req.GetFullName(),
		//IdToken:          req.GetIdToken(),
		ExpiresIn: userToken.GetExpiresIn(),
		//IsValid: 		  isValid,
		RefreshToken: userToken.GetRefreshToken(),
	}
	return res, err

	//}
	//res := &pb.Response{
	//	//IsValid: 		  verify_res.IsValid,
	//}
	//
	//return res, verify_err
}

func RevokeToken(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	redisCmd := RedisConn.Del(ctx, req.RefreshToken)

	if redisCmd.Err() != nil {
		log.Fatal("Error deleting the token from redis: ", redisCmd.Err())
		return nil, redisCmd.Err()
	}

	log.Printf("RefreshToken: \"%s\" is revoked", req.RefreshToken)
	res := &pb.Response{
		RefreshToken: "revoked",

	}

	return res, nil
}