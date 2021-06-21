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

func (a *api) MultiCreateResumesV1(
	ctx context.Context,
	req *desc.MultiCreateResumesV1Request,
) (*desc.MultiCreateResumesV1Response, error) {
	resumeArr := make([]resume.Resume, 0)
	for _, resumeInReq := range req.Resumes {
		resumeArr = append(resumeArr, resume.Resume{
			Id:         0,
			DocumentId: uint(resumeInReq.DocumentId),
		})
	}
	insertedIds, err := a.repo.AddResumes(ctx, resumeArr)
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msgf("Created %v resumes", len(insertedIds))
	return &desc.MultiCreateResumesV1Response{
		ResumeIds: insertedIds,
	}, nil
}

func (a *api) UpdateResumeV1(
	ctx context.Context,
	req *desc.UpdateResumeV1Request,
) (*desc.UpdateResumeV1Response, error) {
	err := a.repo.UpdateResumeById(ctx, uint(req.ResumeId), resume.Resume{
		Id:         uint(req.Resume.Id),
		DocumentId: uint(req.Resume.DocumentId),
	})
	if err != nil {
		return &desc.UpdateResumeV1Response{
			Found: false,
		}, status.Error(codes.NotFound, errSql)
	} else {
		return &desc.UpdateResumeV1Response{
			Found: true,
		}, nil
	}
}

func (a *api) DescribeResumeV1(
	ctx context.Context,
	req *desc.DescribeResumeV1Request,
) (*desc.DescribeResumeV1Response, error) {
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
	resumeArr := make([]resume.Resume, 0)
	resumeArr = append(resumeArr, resume.Resume{
		Id:         0,
		DocumentId: uint(req.DocumentId),
	})
	insertedIds, err := a.repo.AddResumes(ctx, resumeArr)
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume created")
	return &desc.CreateResumeV1Response{
		ResumeId: uint64(insertedIds[0]),
	}, nil
}

func (a *api) RemoveResumeV1(
	ctx context.Context,
	req *desc.RemoveResumeV1Request,
) (*desc.RemoveResumeV1Response, error) {
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
