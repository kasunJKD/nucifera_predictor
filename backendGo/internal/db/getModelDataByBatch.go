package db

import (
	"context"
	"log"
	pb "nucifera_backend/protos/membership"
	"strconv"
)

func (c DBConfigFlask) GetModelDataByBatch(ctx context.Context, req *pb.BatchRequest) (*pb.BatchResponseList, error) {
    batchNumber := req.GetBatchNumber()
    bntostring := strconv.Itoa(int(batchNumber))
    sqlStatement := "select u.model_Id, u.model_Name, u.plot_Fit, u.plot_Validation, u.actual_Precited_Graph, mse, mape from batch"+bntostring+".models u"

    log.Println(sqlStatement)

    rows, err := c.DB.Query(sqlStatement)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }
    defer rows.Close()

    // Create a slice to hold the batch responses
    var batchResponses []*pb.BatchResponse

    for rows.Next() {
        var (
            model_Id             int32
            model_Name          string
            plot_Fit            []byte
            plot_Validation     []byte
            actual_Precited_Graph []byte
            mse                 float32
            mape                float32
        )

        err := rows.Scan(&model_Id, &model_Name, &plot_Fit, &plot_Validation, &actual_Precited_Graph, &mse, &mape)
        if err != nil {
            log.Println(err)
            continue // Skip this row and continue with the next one
        }

        // Create a BatchResponse for each row and add it to the slice
        batchResponses = append(batchResponses, &pb.BatchResponse{
            ModelId:         model_Id,
            ModelName:       model_Name,
            PlotFit:         plot_Fit,
            PlotValidation:  plot_Validation,
            TestPredictions: actual_Precited_Graph,
            Mse:             mse,
            Mape:            mape,
        })
    }

    if err := rows.Err(); err != nil {
        log.Fatalln(err)
        return nil, err
    }

    // Create and return the BatchResponseList
    res := &pb.BatchResponseList{
        BatchResponse: batchResponses,
    }

    return res, nil
}
