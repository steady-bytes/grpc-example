// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commandhandler.proto

package commandhandler_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "commands"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CommandRequest) Validate() error {
	if oneOfNester, ok := this.GetCommand().(*CommandRequest_PlaceOrder); ok {
		if oneOfNester.PlaceOrder != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PlaceOrder); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PlaceOrder", err)
			}
		}
	}
	return nil
}
func (this *CommandResponse) Validate() error {
	return nil
}