module github.com/ozoncp/ocp-resume-api

go 1.15

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/rs/zerolog v1.22.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	google.golang.org/genproto v0.0.0-20210608205507-b6d2f5bf0d7d
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/ozoncp/ocp-resume-api/pkg/ocp-resume-api => ./pkg/ocp-resume-api
