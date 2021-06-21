package api

import (
	"context"
	"errors"
	"time"

	"github.com/ozoncp/ocp-resume-api/internal/metrics"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-resume-api/internal/producer"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errSql      = "sql error"
	errNotFound = "resume not found"
)

type api struct {
	desc.UnimplementedOcpResumeApiServer
	repo repo.RepoResume
	prod producer.Producer
}

func (a *api) MultiCreateResumesV1(
	ctx context.Context,
	req *desc.MultiCreateResumesV1Request,
) (*desc.MultiCreateResumesV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Create resumes multiple")
	defer span.Finish()
	metrics.IncrementCreateRequests(len(req.Resumes))

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

	metrics.IncrementCreates(len(req.Resumes))
	a.prod.SendEvent(newCrudEvent(producer.EventTypeCreated, insertedIds))
	return &desc.MultiCreateResumesV1Response{
		ResumeIds: insertedIds,
	}, nil
}

func (a *api) UpdateResumeV1(
	ctx context.Context,
	req *desc.UpdateResumeV1Request,
) (*desc.UpdateResumeV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Update resume")
	defer span.Finish()
	metrics.IncrementUpdateRequests(1)

	err := a.repo.UpdateResumeById(ctx, uint(req.ResumeId), resume.Resume{
		Id:         uint(req.Resume.GetId()),
		DocumentId: uint(req.Resume.GetDocumentId()),
	})
	if err != nil {
		if errors.Is(err, &repo.ResumeNotFoundError{}) {
			return &desc.UpdateResumeV1Response{
				Found: false,
			}, status.Error(codes.DataLoss, errNotFound)
		}
		return &desc.UpdateResumeV1Response{
			Found: false,
		}, status.Error(codes.NotFound, errSql)
	}
	log.Info().Msgf("Resume %v updated returned", req.ResumeId)

	metrics.IncrementUpdates(1)
	a.prod.SendEvent(newCrudEvent(producer.EventTypeUpdated, []uint64{req.ResumeId}))
	return &desc.UpdateResumeV1Response{
		Found: true,
	}, nil
}

func (a *api) DescribeResumeV1(
	ctx context.Context,
	req *desc.DescribeResumeV1Request,
) (*desc.DescribeResumeV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Describe resume")
	defer span.Finish()
	metrics.IncrementDescribeRequests(1)

	resume, err := a.repo.GetResumeById(ctx, uint(req.ResumeId))
	if err != nil || resume == nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume description returned")

	metrics.IncrementDescribe(1)
	a.prod.SendEvent(newCrudEvent(producer.EventTypeDescribed, []uint64{req.ResumeId}))
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "Create resume")
	defer span.Finish()
	metrics.IncrementCreateRequests(1)

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

	if len(insertedIds) != 1 {
		log.Warn().Msgf("Must be created 1 resume, but created %v", len(insertedIds))
		return nil, status.Error(codes.Internal, errSql)
	}

	log.Info().Msg("Resume created")
	metrics.IncrementCreates(1)
	a.prod.SendEvent(newCrudEvent(producer.EventTypeCreated, insertedIds))
	return &desc.CreateResumeV1Response{
		ResumeId: uint64(insertedIds[0]),
	}, nil
}

func (a *api) RemoveResumeV1(
	ctx context.Context,
	req *desc.RemoveResumeV1Request,
) (*desc.RemoveResumeV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Remove resume")
	defer span.Finish()
	metrics.IncrementRemoveRequests(1)

	err := a.repo.RemoveResumeById(ctx, uint(req.ResumeId))
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	log.Info().Msg("Resume removed")

	metrics.IncrementRemoves(1)
	a.prod.SendEvent(newCrudEvent(producer.EventTypeRemoved, []uint64{req.ResumeId}))
	return &desc.RemoveResumeV1Response{
		Found: true,
	}, nil
}

func (a *api) ListResumesV1(
	ctx context.Context,
	req *desc.ListResumesV1Request,
) (*desc.ListResumesV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "List resumes")
	defer span.Finish()
	metrics.IncrementDescribeRequests(int(req.Limit - req.Offset - 1))

	resumesArr, err := a.repo.ListResumes(ctx, req.Offset, req.Limit)
	if err != nil {
		log.Err(err).Msg(errSql)
		return nil, status.Error(codes.DataLoss, errSql)
	}
	resumesProto := make([]*desc.Resume, 0)
	listedIds := make([]uint64, 0)
	for _, resume := range resumesArr {
		resumeProto := &desc.Resume{
			Id:         uint64(resume.Id),
			DocumentId: uint64(resume.DocumentId),
		}
		listedIds = append(listedIds, uint64(resume.Id))
		resumesProto = append(resumesProto, resumeProto)
	}
	log.Info().Msg("Resumes list returned")

	metrics.IncrementDescribe(int(req.Limit - req.Offset - 1))
	a.prod.SendEvent(newCrudEvent(producer.EventTypeDescribed, listedIds))
	return &desc.ListResumesV1Response{
		Resumes: resumesProto,
	}, nil
}

func NewOcpResumeApi(repo repo.Repo, prod producer.Producer) desc.OcpResumeApiServer {
	return &api{
		repo: repo,
		prod: prod,
	}
}

func newCrudEvent(eventType producer.EventType, ids []uint64) producer.Event {
	return producer.Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Body: map[string]interface{}{
			"Ids": ids,
		},
	}
}
