package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errSql = "sql error"
)

type api struct {
	desc.UnimplementedOcpResumeApiServer
	repo repo.RepoResume
}

func (a *api) DescribeResumeV1(
	ctx context.Context,
	req *desc.DescribeResumeV1Request,
) (*desc.DescribeResumeV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resume, err := a.repo.GetResumeById(ctx, uint(req.ResumeId))
	if err != nil || resume == nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume description returned")
	return &desc.DescribeResumeV1Response{
		Resume: &desc.Resume{
			Id:         uint64(resume.Id),
			DocumentId: uint64(resume.DocumentId),
		},
	}, nil
}

func (a *api) CreateResumeV1(
	ctx context.Context,
	req *desc.CreateResumeV1Request,
) (*desc.CreateResumeV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resumeArr := make([]resume.Resume, 0)
	resumeArr = append(resumeArr, resume.Resume{
		Id:         0,
		DocumentId: uint(req.DocumentId),
	})
	err := a.repo.AddResumes(ctx, resumeArr)
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume created")
	return &desc.CreateResumeV1Response{
		ResumeId: 1,
	}, nil
}

func (a *api) RemoveResumeV1(
	ctx context.Context,
	req *desc.RemoveResumeV1Request,
) (*desc.RemoveResumeV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := a.repo.RemoveResumeById(ctx, uint(req.ResumeId))
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume removed")
	return &desc.RemoveResumeV1Response{
		Found: true,
	}, nil
}

func (a *api) ListResumesV1(
	ctx context.Context,
	req *desc.ListResumesV1Request,
) (*desc.ListResumesV1Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resumesArr, err := a.repo.ListResumes(ctx, req.Offset, req.Limit)
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	var resumesProto []*desc.Resume
	for _, resume := range resumesArr {
		resumeProto := &desc.Resume{
			Id:         uint64(resume.Id),
			DocumentId: uint64(resume.DocumentId),
		}
		resumesProto = append(resumesProto, resumeProto)
	}
	log.Info().Msg("Resumes list returned")
	return &desc.ListResumesV1Response{
		Resumes: resumesProto,
	}, nil
}

func NewOcpResumeApi(repo repo.Repo) desc.OcpResumeApiServer {
	return &api{
		repo: repo,
	}
}
