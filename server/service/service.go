package service

import (
	pb "github.com/sam-explorex/demo_math_grpc/proto"
	"gorm.io/gorm"
)

type MathServer struct {
	Db *gorm.DB
	pb.UnimplementedMathServiceServer
}
