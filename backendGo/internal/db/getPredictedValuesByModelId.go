package db

import (
	"context"
	"log"
	pb "nucifera_backend/protos/membership"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c DBConfig) GetPredictedValuesByModelId(ctx context.Context, req *pb.PredictedRequest) (*pb.PredictedResponseList, error) {
    modelId := req.GetModelId()

    sqlStatement := `select u.date, u.price
                    from batch1.predictions u
                    where u.model_Id = $1`

    rows, err := c.DB.Query(sqlStatement, modelId)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }
    defer rows.Close()

    // Create a slice to hold the predicted responses
    var predictedResponses []*pb.PredictedResponse

    for rows.Next() {
        var (
            date  time.Time
            price float32
        )

        err := rows.Scan(&date, &price)
        if err != nil {
            log.Println(err)
            continue // Skip this row and continue with the next one
        }

        // Create a PredictedResponse for each row and add it to the slice
        predictedResponses = append(predictedResponses, &pb.PredictedResponse{
            Date:   timestamppb.New(date),
            Values: price,
        })
    }

    if err := rows.Err(); err != nil {
        log.Fatalln(err)
        return nil, err
    }

    // Create and return the PredictedResponseList
    res := &pb.PredictedResponseList{
        PredictedResponseList: predictedResponses,
    }

    return res, nil
}