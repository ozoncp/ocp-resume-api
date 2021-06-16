# ocp-resume-api
Resume Api service


protoc -I vendor.protogen --go_out=pkg/ocp-resume-api --go_opt=paths=import --go-grpc_out=pkg/ocp-resume-api --go-grpc_opt=paths=import --grpc-gateway_out=pkg/ocp-resume-api --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import --validate_out lang=go:pkg/ocp-resume-api --swagger_out=allow_merge=true,merge_file_name=api:swagger api/ocp-resume-api/ocp-resume-api.proto