// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service gRPC server types
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package server

import (
	"goa.design/goa/examples/security/gen/grpc/secured_service/pb"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
)

// NewSigninPayload builds the payload of the "signin" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSigninPayload(username string, password string) *securedservice.SigninPayload {
	payload := &securedservice.SigninPayload{}
	payload.Username = username
	payload.Password = password
	return payload
}

// NewSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "secured_service" service.
func NewSigninResponse(result *securedservice.Creds) *pb.SigninResponse {
	message := &pb.SigninResponse{
		Jwt:        result.JWT,
		ApiKey:     result.APIKey,
		OauthToken: result.OauthToken,
	}
	return message
}

// NewSecurePayload builds the payload of the "secure" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSecurePayload(message *pb.SecureRequest, token string) *securedservice.SecurePayload {
	payload := &securedservice.SecurePayload{
		Fail: &message.Fail,
	}
	payload.Token = token
	return payload
}

// NewSecureResponse builds the gRPC response type from the result of the
// "secure" endpoint of the "secured_service" service.
func NewSecureResponse(result string) *pb.SecureResponse {
	message := &pb.SecureResponse{}
	message.Field = result
	return message
}

// NewDoublySecurePayload builds the payload of the "doubly_secure" endpoint of
// the "secured_service" service from the gRPC request type.
func NewDoublySecurePayload(message *pb.DoublySecureRequest, token string) *securedservice.DoublySecurePayload {
	payload := &securedservice.DoublySecurePayload{
		Key: message.Key,
	}
	payload.Token = token
	return payload
}

// NewDoublySecureResponse builds the gRPC response type from the result of the
// "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureResponse(result string) *pb.DoublySecureResponse {
	message := &pb.DoublySecureResponse{}
	message.Field = result
	return message
}

// NewAlsoDoublySecurePayload builds the payload of the "also_doubly_secure"
// endpoint of the "secured_service" service from the gRPC request type.
func NewAlsoDoublySecurePayload(message *pb.AlsoDoublySecureRequest, oauthToken string, token string) *securedservice.AlsoDoublySecurePayload {
	payload := &securedservice.AlsoDoublySecurePayload{
		Username: &message.Username,
		Password: &message.Password,
		Key:      &message.Key,
	}
	payload.OauthToken = &oauthToken
	payload.Token = &token
	return payload
}

// NewAlsoDoublySecureResponse builds the gRPC response type from the result of
// the "also_doubly_secure" endpoint of the "secured_service" service.
func NewAlsoDoublySecureResponse(result string) *pb.AlsoDoublySecureResponse {
	message := &pb.AlsoDoublySecureResponse{}
	message.Field = result
	return message
}
