// Code generated by thriftrw v1.15.0. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
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

package echo

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Echo_Echo_Args represents the arguments for the Echo.echo function.
//
// The arguments for echo are sent and received over the wire as this struct.
type Echo_Echo_Args struct {
	Ping *Ping `json:"ping,omitempty"`
}

// ToWire translates a Echo_Echo_Args struct into a Thrift-level intermediate
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
func (v *Echo_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Ping != nil {
		w, err = v.Ping.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Ping_Read(w wire.Value) (*Ping, error) {
	var v Ping
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Echo_Echo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_Echo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_Echo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_Echo_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Ping, err = _Ping_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Echo_Echo_Args
// struct.
func (v *Echo_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Ping != nil {
		fields[i] = fmt.Sprintf("Ping: %v", v.Ping)
		i++
	}

	return fmt.Sprintf("Echo_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_Echo_Args match the
// provided Echo_Echo_Args.
//
// This function performs a deep comparison.
func (v *Echo_Echo_Args) Equals(rhs *Echo_Echo_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Ping == nil && rhs.Ping == nil) || (v.Ping != nil && rhs.Ping != nil && v.Ping.Equals(rhs.Ping))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_Echo_Args.
func (v *Echo_Echo_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Ping != nil {
		err = multierr.Append(err, enc.AddObject("ping", v.Ping))
	}
	return err
}

// GetPing returns the value of Ping if it is set or its
// zero value if it is unset.
func (v *Echo_Echo_Args) GetPing() (o *Ping) {
	if v != nil && v.Ping != nil {
		return v.Ping
	}

	return
}

// IsSetPing returns true if Ping is not nil.
func (v *Echo_Echo_Args) IsSetPing() bool {
	return v != nil && v.Ping != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echo" for this struct.
func (v *Echo_Echo_Args) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Echo_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Echo_Echo_Helper provides functions that aid in handling the
// parameters and return values of the Echo.echo
// function.
var Echo_Echo_Helper = struct {
	// Args accepts the parameters of echo in-order and returns
	// the arguments struct for the function.
	Args func(
		ping *Ping,
	) *Echo_Echo_Args

	// IsException returns true if the given error can be thrown
	// by echo.
	//
	// An error can be thrown by echo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echo
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echo into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echo
	//
	//   value, err := echo(args)
	//   result, err := Echo_Echo_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*Pong, error) (*Echo_Echo_Result, error)

	// UnwrapResponse takes the result struct for echo
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Echo_Echo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Echo_Echo_Result) (*Pong, error)
}{}

func init() {
	Echo_Echo_Helper.Args = func(
		ping *Ping,
	) *Echo_Echo_Args {
		return &Echo_Echo_Args{
			Ping: ping,
		}
	}

	Echo_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Echo_Echo_Helper.WrapResponse = func(success *Pong, err error) (*Echo_Echo_Result, error) {
		if err == nil {
			return &Echo_Echo_Result{Success: success}, nil
		}

		return nil, err
	}
	Echo_Echo_Helper.UnwrapResponse = func(result *Echo_Echo_Result) (success *Pong, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Echo_Echo_Result represents the result of a Echo.echo function call.
//
// The result of a echo execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Echo_Echo_Result struct {
	// Value returned by echo after a successful execution.
	Success *Pong `json:"success,omitempty"`
}

// ToWire translates a Echo_Echo_Result struct into a Thrift-level intermediate
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
func (v *Echo_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
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

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Pong_Read(w wire.Value) (*Pong, error) {
	var v Pong
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Echo_Echo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_Echo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Echo_Echo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Echo_Echo_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Pong_Read(field.Value)
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
	if count != 1 {
		return fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Echo_Echo_Result
// struct.
func (v *Echo_Echo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Echo_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_Echo_Result match the
// provided Echo_Echo_Result.
//
// This function performs a deep comparison.
func (v *Echo_Echo_Result) Equals(rhs *Echo_Echo_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_Echo_Result.
func (v *Echo_Echo_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Echo_Echo_Result) GetSuccess() (o *Pong) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Echo_Echo_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echo" for this struct.
func (v *Echo_Echo_Result) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Echo_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
