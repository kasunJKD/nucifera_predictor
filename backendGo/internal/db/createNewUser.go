package db

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	//"database/sql"
	"log"
	pb "nucifera_backend/protos/membership"

	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nucifera_backend/internal/utils"
)

func (c DBConfig) CreateNewUser(req *pb.Request) (*pb.Response, error) {
	//Inserting new User
	log.Println("Creating new user account------------>")
	sqlStatement := `
	WITH inUser AS (
		INSERT INTO users(userId, email, emailVerified, createdAt, updatedAt, passwordHash)
		VALUES (gen_random_uuid(), $1, $2, (select current_timestamp at time zone ('utc')), (select current_timestamp at time zone ('utc')), $3)
		RETURNING userId, createdAt, updatedAt
	)INSERT INTO userinfo (userId, displayName, firstName, lastName, gender, address)
		SELECT IU.userId, $4, $5, $6, $7, $8 FROM inUser IU
		RETURNING (SELECT inUser.userId FROM inUser), (SELECT inUser.createdAt FROM inUser), (SELECT inUser.updatedAt FROM inUser)
	`

	var (
		userId string
		createdAt time.Time
		updatedAt time.Time
	)

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "password cannot be hashed")
	}

	inserterror := c.DB.QueryRow(sqlStatement,
		req.GetEmail(),
		req.GetEmailVerified(),
		hashedPassword,
		req.GetDisplayName(),
		req.GetFirstName(),
		req.GetLastName(),
		req.GetGender(),
		req.GetAddress()).Scan(&userId, &createdAt, &updatedAt)

	if inserterror != nil {
		panic(inserterror)
	}


	log.Println("New User created: userId = " + userId)

	//Create user access token from id
	//userToken, err := jwt.CreateToken(&pb.Request{LocalId: localId})

	res := &pb.Response{
		Users: &pb.Users{
			UserId:				userId,
			Email:			 	req.GetEmail(),
			EmailVerified:   	req.GetEmailVerified(),
			DisplayName:     	req.GetDisplayName(),
			CreatedAt:       	timestamppb.New(createdAt),
			UpdatedAt:       	timestamppb.New(updatedAt),
			Gender: 			req.GetGender(),
			Address: 			req.GetAddress(),
			Age: 			 	req.GetAge(),
			Experience: 	 	req.GetExperience(),
			PlayingTime: 	 	req.GetPlayingTime(),
			PreferredPlatforms: req.GetPreferredPlatforms(),
			//LinkedAccounts: &pb.LinkedAccounts {
			//	ProviderId:       req.GetProviderId(),
			//	FederatedId:      req.GetFederatedId(),
			//	Email:req.GetEmail(),
			//},
		},
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		FullName:  req.GetFullName(),
		//OauthAccessToken: userToken.OauthAccessToken,
		//ExpiresIn: userToken.ExpiresIn,
		IsNewUser: true,
	}

	return res, nil

}
