// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_GetWorkflowExecutionHistory_Args represents the arguments for the WorkflowService.GetWorkflowExecutionHistory function.
//
// The arguments for GetWorkflowExecutionHistory are sent and received over the wire as this struct.
type WorkflowService_GetWorkflowExecutionHistory_Args struct {
	GetRequest *shared.GetWorkflowExecutionHistoryRequest `json:"getRequest,omitempty"`
}

// ToWire translates a WorkflowService_GetWorkflowExecutionHistory_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.GetRequest != nil {
		w, err = v.GetRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _GetWorkflowExecutionHistoryRequest_Read(w wire.Value) (*shared.GetWorkflowExecutionHistoryRequest, error) {
	var v shared.GetWorkflowExecutionHistoryRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_GetWorkflowExecutionHistory_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_GetWorkflowExecutionHistory_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_GetWorkflowExecutionHistory_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.GetRequest, err = _GetWorkflowExecutionHistoryRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_GetWorkflowExecutionHistory_Args
// struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.GetRequest != nil {
		fields[i] = fmt.Sprintf("GetRequest: %v", v.GetRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_GetWorkflowExecutionHistory_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_GetWorkflowExecutionHistory_Args match the
// provided WorkflowService_GetWorkflowExecutionHistory_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) Equals(rhs *WorkflowService_GetWorkflowExecutionHistory_Args) bool {
	if !((v.GetRequest == nil && rhs.GetRequest == nil) || (v.GetRequest != nil && rhs.GetRequest != nil && v.GetRequest.Equals(rhs.GetRequest))) {
		return false
	}

	return true
}

// GetGetRequest returns the value of GetRequest if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) GetGetRequest() (o *shared.GetWorkflowExecutionHistoryRequest) {
	if v.GetRequest != nil {
		return v.GetRequest
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "GetWorkflowExecutionHistory" for this struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) MethodName() string {
	return "GetWorkflowExecutionHistory"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_GetWorkflowExecutionHistory_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.GetWorkflowExecutionHistory
// function.
var WorkflowService_GetWorkflowExecutionHistory_Helper = struct {
	// Args accepts the parameters of GetWorkflowExecutionHistory in-order and returns
	// the arguments struct for the function.
	Args func(
		getRequest *shared.GetWorkflowExecutionHistoryRequest,
	) *WorkflowService_GetWorkflowExecutionHistory_Args

	// IsException returns true if the given error can be thrown
	// by GetWorkflowExecutionHistory.
	//
	// An error can be thrown by GetWorkflowExecutionHistory only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for GetWorkflowExecutionHistory
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// GetWorkflowExecutionHistory into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by GetWorkflowExecutionHistory
	//
	//   value, err := GetWorkflowExecutionHistory(args)
	//   result, err := WorkflowService_GetWorkflowExecutionHistory_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from GetWorkflowExecutionHistory: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*shared.GetWorkflowExecutionHistoryResponse, error) (*WorkflowService_GetWorkflowExecutionHistory_Result, error)

	// UnwrapResponse takes the result struct for GetWorkflowExecutionHistory
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if GetWorkflowExecutionHistory threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := WorkflowService_GetWorkflowExecutionHistory_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_GetWorkflowExecutionHistory_Result) (*shared.GetWorkflowExecutionHistoryResponse, error)
}{}

func init() {
	WorkflowService_GetWorkflowExecutionHistory_Helper.Args = func(
		getRequest *shared.GetWorkflowExecutionHistoryRequest,
	) *WorkflowService_GetWorkflowExecutionHistory_Args {
		return &WorkflowService_GetWorkflowExecutionHistory_Args{
			GetRequest: getRequest,
		}
	}

	WorkflowService_GetWorkflowExecutionHistory_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.ServiceBusyError:
			return true
		default:
			return false
		}
	}

	WorkflowService_GetWorkflowExecutionHistory_Helper.WrapResponse = func(success *shared.GetWorkflowExecutionHistoryResponse, err error) (*WorkflowService_GetWorkflowExecutionHistory_Result, error) {
		if err == nil {
			return &WorkflowService_GetWorkflowExecutionHistory_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetWorkflowExecutionHistory_Result.BadRequestError")
			}
			return &WorkflowService_GetWorkflowExecutionHistory_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetWorkflowExecutionHistory_Result.InternalServiceError")
			}
			return &WorkflowService_GetWorkflowExecutionHistory_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetWorkflowExecutionHistory_Result.EntityNotExistError")
			}
			return &WorkflowService_GetWorkflowExecutionHistory_Result{EntityNotExistError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_GetWorkflowExecutionHistory_Result.ServiceBusyError")
			}
			return &WorkflowService_GetWorkflowExecutionHistory_Result{ServiceBusyError: e}, nil
		}

		return nil, err
	}
	WorkflowService_GetWorkflowExecutionHistory_Helper.UnwrapResponse = func(result *WorkflowService_GetWorkflowExecutionHistory_Result) (success *shared.GetWorkflowExecutionHistoryResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// WorkflowService_GetWorkflowExecutionHistory_Result represents the result of a WorkflowService.GetWorkflowExecutionHistory function call.
//
// The result of a GetWorkflowExecutionHistory execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type WorkflowService_GetWorkflowExecutionHistory_Result struct {
	// Value returned by GetWorkflowExecutionHistory after a successful execution.
	Success              *shared.GetWorkflowExecutionHistoryResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError                     `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError                `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError                `json:"entityNotExistError,omitempty"`
	ServiceBusyError     *shared.ServiceBusyError                    `json:"serviceBusyError,omitempty"`
}

// ToWire translates a WorkflowService_GetWorkflowExecutionHistory_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) ToWire() (wire.Value, error) {
	var (
		fields [5]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_GetWorkflowExecutionHistory_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _GetWorkflowExecutionHistoryResponse_Read(w wire.Value) (*shared.GetWorkflowExecutionHistoryResponse, error) {
	var v shared.GetWorkflowExecutionHistoryResponse
	err := v.FromWire(w)
	return &v, err
}

func _ServiceBusyError_Read(w wire.Value) (*shared.ServiceBusyError, error) {
	var v shared.ServiceBusyError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_GetWorkflowExecutionHistory_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_GetWorkflowExecutionHistory_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_GetWorkflowExecutionHistory_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _GetWorkflowExecutionHistoryResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_GetWorkflowExecutionHistory_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_GetWorkflowExecutionHistory_Result
// struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [5]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}

	return fmt.Sprintf("WorkflowService_GetWorkflowExecutionHistory_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_GetWorkflowExecutionHistory_Result match the
// provided WorkflowService_GetWorkflowExecutionHistory_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) Equals(rhs *WorkflowService_GetWorkflowExecutionHistory_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) GetSuccess() (o *shared.GetWorkflowExecutionHistoryResponse) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetBadRequestError returns the value of BadRequestError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) GetBadRequestError() (o *shared.BadRequestError) {
	if v.BadRequestError != nil {
		return v.BadRequestError
	}

	return
}

// GetInternalServiceError returns the value of InternalServiceError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) GetInternalServiceError() (o *shared.InternalServiceError) {
	if v.InternalServiceError != nil {
		return v.InternalServiceError
	}

	return
}

// GetEntityNotExistError returns the value of EntityNotExistError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) GetEntityNotExistError() (o *shared.EntityNotExistsError) {
	if v.EntityNotExistError != nil {
		return v.EntityNotExistError
	}

	return
}

// GetServiceBusyError returns the value of ServiceBusyError if it is set or its
// zero value if it is unset.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) GetServiceBusyError() (o *shared.ServiceBusyError) {
	if v.ServiceBusyError != nil {
		return v.ServiceBusyError
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "GetWorkflowExecutionHistory" for this struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) MethodName() string {
	return "GetWorkflowExecutionHistory"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_GetWorkflowExecutionHistory_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
