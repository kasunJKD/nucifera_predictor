package main

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"log"
	"net"
	"os"
	"nucifera_backend/internal/auth"
	"nucifera_backend/internal/gateway"
	"nucifera_backend/internal/jwt"

	//"net/http"

	//"net/http"
	grpc "google.golang.org/grpc"

	db "nucifera_backend/internal/db"
	rd "nucifera_backend/internal/redis"
	pb "nucifera_backend/protos/membership"
)

var dbConn *sql.DB
var redisConn *redis.Client

var (
	grpcPort = getEnv("GRPC_PORT", ":8081")
	httpPort = getEnv("HTTP_PORT", ":11001")
	host = getEnv("HOST", "0.0.0.0")
	mem_host = getEnv("MEM_HOST", "0.0.0.0")
	mem_port = getEnv("MEM_PORT", "15452")
	redisHost = getEnv("REDIS_HOST", "0.0.0.0")
	redisPort = getEnv("REDIS_PORT", "16389")
	user = getEnv("DB_USER", "postgres")
	password = getEnv("PASSWORD", "9221")
	mem_dbname = getEnv("MEM_DBNAME", "membership")
)

func getEnv(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	log.Println("Key not found: ", key)
	os.Setenv(key, fallback)
	return fallback
}

type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
}

func main() {
	log.Println("Welcome to the server")

	//connecting to Membership Database 
	dbConn = db.Connect(mem_host, mem_port, user, password, mem_dbname)
	defer dbConn.Close()

	//connecting to redis
	redisConn = rd.Connect(redisHost, redisPort, password)
	jwt.RedisConn = redisConn
	defer redisConn.Close()

	//start listening for grpc
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcserver := grpc.NewServer()

	//Register DataService
	pb.RegisterDataServiceServer(grpcserver, new(DataServiceServer))
	log.Println("Starting grpc connection on port " + grpcPort)

	//startServing requests
	go grpcserver.Serve(listen)

	//start http server for rest
	log.Println("start http server on port" + httpPort)
	err = gateway.Run(host + grpcPort)
	log.Fatalln(err)

}

func (s *DataServiceServer) SignUp(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	data, err := auth.SignUp(ctx, dbConn, req)
	return data, err
}

func (s *DataServiceServer) PasswordSignIn(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	data, err := auth.PasswordSignIn(ctx, dbConn, req)
	return data, err
}