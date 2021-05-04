// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventstore.proto

package eventstore_v1

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	_ "aggregates"
	_ "events"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Event) Validate() error {
	if this.EventType != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EventType); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EventType", err)
		}
	}
	return nil
}
func (this *StoreEventRequest) Validate() error {
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	return nil
}
func (this *StoreEventResponse) Validate() error {
	return nil
}
func (this *QueryEventStoreRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *QueryEventStoreResponse) Validate() error {
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}