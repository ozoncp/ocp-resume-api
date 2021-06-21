package api_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-resume-api/internal/api"
	"github.com/ozoncp/ocp-resume-api/internal/mocks"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	desc "github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api"
)

var _ = Describe("Api", func() {
	var (
		ctx       context.Context
		testApi   desc.OcpResumeApiServer
		mock      sqlmock.Sqlmock
		ctrl      *gomock.Controller
		prod      *mocks.MockProducer
		db        *sql.DB
		sqlxDB    *sqlx.DB
		repo_test repo.Repo
		err       error
	)
	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		prod = mocks.NewMockProducer(ctrl)
		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		repo_test = repo.NewRepo(sqlxDB, 32)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
		ctrl.Finish()
	})

	Context("test api functions", func() {

		It("Test create resume", func() {
			request := &desc.CreateResumeV1Request{
				DocumentId: 111,
			}
			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1)
			mock.ExpectQuery("insert into resumes").
				WithArgs(request.DocumentId).WillReturnRows(rows)

			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())
			response, err := testApi.CreateResumeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.ResumeId).Should(BeEquivalentTo(1))
		})

		It("Test incorrect create resume", func() {
			request := &desc.CreateResumeV1Request{
				DocumentId: 111,
			}
			mock.ExpectQuery("insert into resumes").
				WithArgs(request.DocumentId).WillReturnError(errors.New("can't insert resume"))
			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.CreateResumeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test remove resume", func() {
			request := &desc.RemoveResumeV1Request{
				ResumeId: 0,
			}
			rows := sqlmock.NewRows([]string{"found"}).
				AddRow(true)
			mock.ExpectQuery("delete from resumes").
				WithArgs(request.ResumeId).WillReturnRows(rows)

			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.RemoveResumeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Found).Should(BeEquivalentTo(true))
		})

		It("Test incorrect remove resume", func() {
			request := &desc.RemoveResumeV1Request{
				ResumeId: 1,
			}
			mock.ExpectQuery("delete from resumes").
				WithArgs(request.ResumeId).WillReturnError(errors.New("can't remove resume"))
			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.RemoveResumeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test describe resume", func() {
			request := &desc.DescribeResumeV1Request{
				ResumeId: 1,
			}
			rows := sqlmock.NewRows([]string{"id", "document_id"}).
				AddRow(request.ResumeId, 222)
			mock.ExpectQuery("select id, document_id from resumes where").
				WithArgs(request.ResumeId).
				WillReturnRows(rows)

			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.DescribeResumeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Resume.Id).Should(BeEquivalentTo(1))
			Expect(response.Resume.DocumentId).Should(BeEquivalentTo(222))
		})

		It("Test incorrect describe resume", func() {
			request := &desc.DescribeResumeV1Request{
				ResumeId: 1,
			}
			mock.ExpectQuery("select id, document_id from resumes where").
				WithArgs(request.ResumeId).
				WillReturnError(errors.New("can't select resumes"))
			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.DescribeResumeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

		It("Test list resumes", func() {
			request := &desc.ListResumesV1Request{
				Limit:  3,
				Offset: 0,
			}
			rows := sqlmock.NewRows([]string{"id", "document_id"}).
				AddRow(1, 222).AddRow(2, 333)
			mock.ExpectQuery("select id, document_id from resumes limit 3 offset 0").
				WillReturnRows(rows)
			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			response, err := testApi.ListResumeV1(ctx, request)
			Expect(err).Should(BeNil())
			Expect(response.Resumes[0].Id).Should(BeEquivalentTo(1))
			Expect(response.Resumes[1].Id).Should(BeEquivalentTo(2))
			Expect(response.Resumes[0].DocumentId).Should(BeEquivalentTo(222))
			Expect(response.Resumes[1].DocumentId).Should(BeEquivalentTo(333))
		})

		It("Test incorrect list resumes", func() {
			request := &desc.ListResumesV1Request{
				Limit:  3,
				Offset: 0,
			}
			mock.ExpectQuery("select id, document_id from resumes limit 3 offset 0").
				WillReturnError(errors.New("can't list resumes"))
			testApi = api.NewOcpResumeApi(repo_test, prod)
			Expect(testApi).ShouldNot(BeNil())

			_, err := testApi.ListResumeV1(ctx, request)
			Expect(err).ShouldNot(BeNil())
		})

	})
})
