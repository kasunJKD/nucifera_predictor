package db

import (
	"context"
	"log"
	pb "nucifera_backend/protos/membership"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c DBConfigFlask) GetOriginalData(ctx context.Context, req *pb.OriginalDataRequest) (*pb.OriginalDataList, error) {
    batchNumber := req.GetBatchId()
    bntostring := strconv.Itoa(int(batchNumber))
    sqlStatement := "select u.date, u.average_price, u.rainfall_kurunegala, u.rainfall_puttalam, u.rainfall_colombo, u.min_temp_kurunegala, u.min_temp_puttalam, u.min_temp_colombo, u.max_temp_kurunegala, u.max_temp_puttalam, u.max_temp_colombo from batch"+bntostring+".original u"

    rows, err := c.DB.Query(sqlStatement)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }
    defer rows.Close()

    // Create a slice to hold the batch responses
    var batchResponses []*pb.OriginalDataResponse

    for rows.Next() {
        var (
			date int64
  			prices float32
  			rainfall_colombo float32
			rainfall_puttalam float32
			rainfall_kurunegala float32
			min_temp_colombo float32
			min_temp_puttalam float32
			min_temp_kurunegala float32
			max_temp_colombo float32
			max_temp_puttalam float32
			max_temp_kurunegala float32
        )

        err := rows.Scan(&date, &prices, &rainfall_kurunegala, &rainfall_puttalam, &rainfall_colombo, &min_temp_kurunegala, &min_temp_puttalam, &min_temp_colombo, &max_temp_kurunegala, &max_temp_puttalam, &max_temp_colombo)
        if err != nil {
            log.Println(err)
            continue // Skip this row and continue with the next one
        }

        // Create a BatchResponse for each row and add it to the slice
        batchResponses = append(batchResponses, &pb.OriginalDataResponse{
            Date:         timestamppb.New(time.Unix(date, 0)),
            Prices:  prices,
			RainfallColombo: rainfall_colombo,
			RainfallPuttalam: rainfall_puttalam,
			RainfallKurunegala: rainfall_kurunegala,
			MinTempColombo: min_temp_colombo,
			MinTempPuttalam: min_temp_puttalam,
			MinTempKurunegala: min_temp_kurunegala,
			MaxTempColombo: max_temp_colombo,
			MaxTempPuttalam: max_temp_puttalam,
			MaxTempKurunegala: max_temp_kurunegala,
        })
    }

    if err := rows.Err(); err != nil {
        log.Fatalln(err)
        return nil, err
    }

    // Create and return the BatchResponseList
    res := &pb.OriginalDataList{
        OriginalDataList: batchResponses,
    }

    return res, nil
}
