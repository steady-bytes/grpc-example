// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commandhandler.proto

package commandhandler

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/steady-bytes/grpc-example/proto/commands"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Command) Validate() error {
	if this.PlaceOrder != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlaceOrder); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlaceOrder", err)
		}
	}
	return nil
}
func (this *Response) Validate() error {
	return nil
}
