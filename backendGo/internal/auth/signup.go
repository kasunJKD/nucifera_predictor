package auth

import (
	"context"
	"database/sql"
	"log"
	db "nucifera_backend/internal/db"
	"nucifera_backend/internal/jwt"
	pb "nucifera_backend/protos/membership"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SignUp (ctx context.Context, database *sql.DB, req *pb.Request) (*pb.Response, error) {
	//check user exists -- returns true/false
	cc := db.DBConfig{DB: database}
	userExist, err := cc.CheckAuthUserExists(&pb.Request{Email: req.Email})

	if err != nil {
		log.Fatal(err)
	}

	if userExist {
		return nil, status.Errorf(codes.Internal, "User already exists")
	}

	//create new user
	userCreateResponse, err := cc.CreateNewUser(req)

	if err != nil {
		log.Fatal(err)
	}

	userToken, err := jwt.CreateToken(&pb.Request{LocalId: userCreateResponse.Users.UserId})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	//if err != nil {
	//	log.Fatal(err)
	//}

	res := &pb.Response{
		Users: &pb.Users{
			UserId:		  		userCreateResponse.Users.GetUserId(),
			Email:			  	userCreateResponse.Users.GetEmail(),
			EmailVerified:    	userCreateResponse.Users.GetEmailVerified(),
			DisplayName:      	userCreateResponse.Users.GetDisplayName(),
			PhotoUrl:         	userCreateResponse.Users.GetPhotoUrl(),
			UpdatedAt:        	userCreateResponse.Users.GetUpdatedAt(),
			CreatedAt:        	userCreateResponse.Users.GetCreatedAt(),
			Gender: 			userCreateResponse.Users.GetGender(),
			Address: 			userCreateResponse.Users.GetAddress(),
			Age: 			 	userCreateResponse.Users.GetAge(),
			Experience: 	 	userCreateResponse.Users.GetExperience(),
			PlayingTime: 	 	userCreateResponse.Users.GetPlayingTime(),
			PreferredPlatforms: userCreateResponse.Users.GetPreferredPlatforms(),
		},
		FirstName:        userCreateResponse.GetFirstName(),
		LastName:         userCreateResponse.GetLastName(),
		FullName:         userCreateResponse.GetFullName(),
		OauthAccessToken: userToken.OauthAccessToken,
		ExpiresIn:        userToken.GetExpiresIn(),
		IsNewUser:		  userCreateResponse.GetIsNewUser(),
		RefreshToken: 	  userToken.GetRefreshToken(),
	}

	return res, err

}