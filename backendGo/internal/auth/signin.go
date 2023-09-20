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
	"nucifera_backend/internal/utils"
)

func PasswordSignIn (ctx context.Context, database *sql.DB, req *pb.Request) (*pb.Response, error) {
	//check user exists -- returns true/false
	cc := db.DBConfig{DB: database}
	userExist, err := cc.CheckAuthUserExists(&pb.Request{Email: req.Email})
	log.Printf("hit1")
	if err != nil {
		log.Fatal(err)
	}

	if userExist {
		//getUser info
		userInfo, err := cc.GetAccountInfo(ctx, &pb.Request{Email: req.Email})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("hit2")
		//check request password with stored password
		err = utils.CheckPassword(req.Password, userInfo.Users.PasswordHash)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Password incorrect")
		}
		log.Printf("hit3")
		// Create user access token from id
		userToken, err := jwt.CreateToken(&pb.Request{LocalId: userInfo.Users.UserId})
		log.Printf("hit4")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate access token")
		}

		//return response
		res := &pb.Response{
			Users: &pb.Users{
				UserId:		  	  userInfo.Users.GetUserId(),
				Email:			  userInfo.Users.GetEmail(),
				EmailVerified:    userInfo.Users.GetEmailVerified(),
				DisplayName:      userInfo.Users.GetDisplayName(),
				PhotoUrl:         userInfo.Users.GetPhotoUrl(),
				UpdatedAt:        userInfo.Users.GetUpdatedAt(),
				CreatedAt:        userInfo.Users.GetCreatedAt(),
				Gender: 		  userInfo.Users.GetGender(),
				Address: 		  userInfo.Users.GetAddress(),
				Age: 			  userInfo.Users.GetAge(),
				Experience: 	  userInfo.Users.GetExperience(),
				PlayingTime: 	  userInfo.Users.GetPlayingTime(),
				PreferredPlatforms: userInfo.Users.GetPreferredPlatforms(),
				//LinkedAccounts: &pb.LinkedAccounts {
				//	ProviderId:       "providerId",
				//	FederatedId:      "federatedId",
				//	Email:            "email",
				//},
			},
			FirstName:        userInfo.GetFirstName(),
			LastName:         userInfo.GetLastName(),
			FullName:         userInfo.GetFullName(),
			OauthAccessToken: userToken.GetOauthAccessToken(),
			ExpiresIn:        userToken.GetExpiresIn(),
			//IsNewUser:		  userInfo.GetIsNewUser(),
			//RefreshToken: 	  userToken.GetRefreshToken(),
		}

		return res, err

	}

	return nil, status.Errorf(codes.Internal, "User not valid")

}