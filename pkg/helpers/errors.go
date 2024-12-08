package helpers

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Err_Srv_NilRequest         = status.Error(codes.InvalidArgument, "nil request")
	Err_Srv_Request_KeyMissing = status.Error(codes.InvalidArgument, "request key missing")
)
