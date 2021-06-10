package api

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errResumeNotFound = "resume not found"
)

type api struct {
	desc.UnimplementedOcpResumeApiServer
}

func (a *api) DescribeResumeV1(
	ctx context.Context,
	req *desc.DescribeResumeV1Request,
) (*desc.DescribeResumeV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Printf("Describing resume with id=%v", req.ResumeId)
	err := status.Error(codes.NotFound, errResumeNotFound)
	return nil, err
}

func (a *api) CreateResumeV1(
	ctx context.Context,
	req *desc.CreateResumeV1Request,
) (*desc.CreateResumeV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Printf("Creating resume with id=%v", req.DocumentId)
	err := status.Error(codes.NotFound, errResumeNotFound)
	return nil, err
}

func (a *api) RemoveResumeV1(
	ctx context.Context,
	req *desc.RemoveResumeV1Request,
) (*desc.RemoveResumeV1Response, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Printf("Removing resume with id=%v", req.ResumeId)
	err := status.Error(codes.NotFound, errResumeNotFound)
	return nil, err
}

func (a *api) ListResumesV1(
	ctx context.Context,
	req *desc.ListResumesV1Request,
) (*desc.ListResumesV1Response, error) {
	fmt.Printf("Listing")
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Printf("Describing resume with from%v to %v", req.Offset, req.Limit)
	err := status.Error(codes.NotFound, errResumeNotFound)
	return nil, err
}

func NewOcpResumeApi() desc.OcpResumeApiServer {
	return &api{}
}
