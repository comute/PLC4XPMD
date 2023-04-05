/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// RequestReset is the corresponding interface of RequestReset
type RequestReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Request
	// GetTildePeek returns TildePeek (property field)
	GetTildePeek() RequestType
	// GetSecondTilde returns SecondTilde (property field)
	GetSecondTilde() *RequestType
	// GetTildePeek2 returns TildePeek2 (property field)
	GetTildePeek2() RequestType
	// GetThirdTilde returns ThirdTilde (property field)
	GetThirdTilde() *RequestType
}

// RequestResetExactly can be used when we want exactly this type and not a type which fulfills RequestReset.
// This is useful for switch cases.
type RequestResetExactly interface {
	RequestReset
	isRequestReset() bool
}

// _RequestReset is the data-structure of this message
type _RequestReset struct {
	*_Request
	TildePeek   RequestType
	SecondTilde *RequestType
	TildePeek2  RequestType
	ThirdTilde  *RequestType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestReset) InitializeParent(parent Request, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestReset) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestReset) GetTildePeek() RequestType {
	return m.TildePeek
}

func (m *_RequestReset) GetSecondTilde() *RequestType {
	return m.SecondTilde
}

func (m *_RequestReset) GetTildePeek2() RequestType {
	return m.TildePeek2
}

func (m *_RequestReset) GetThirdTilde() *RequestType {
	return m.ThirdTilde
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestReset factory function for _RequestReset
func NewRequestReset(tildePeek RequestType, secondTilde *RequestType, tildePeek2 RequestType, thirdTilde *RequestType, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestReset {
	_result := &_RequestReset{
		TildePeek:   tildePeek,
		SecondTilde: secondTilde,
		TildePeek2:  tildePeek2,
		ThirdTilde:  thirdTilde,
		_Request:    NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestReset(structType interface{}) RequestReset {
	if casted, ok := structType.(RequestReset); ok {
		return casted
	}
	if casted, ok := structType.(*RequestReset); ok {
		return *casted
	}
	return nil
}

func (m *_RequestReset) GetTypeName() string {
	return "RequestReset"
}

func (m *_RequestReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (secondTilde)
	if m.SecondTilde != nil {
		lengthInBits += 8
	}

	// Optional Field (thirdTilde)
	if m.ThirdTilde != nil {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *_RequestReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestResetParse(theBytes []byte, cBusOptions CBusOptions) (RequestReset, error) {
	return RequestResetParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func RequestResetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (RequestReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (tildePeek)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("tildePeek"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tildePeek")
	}
	tildePeek, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tildePeek' field of RequestReset")
	}
	if closeErr := readBuffer.CloseContext("tildePeek"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tildePeek")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (secondTilde) (Can be skipped, if a given expression evaluates to false)
	var secondTilde *RequestType = nil
	if bool((tildePeek) == (RequestType_RESET)) {
		if pullErr := readBuffer.PullContext("secondTilde"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for secondTilde")
		}
		_val, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'secondTilde' field of RequestReset")
		}
		secondTilde = &_val
		if closeErr := readBuffer.CloseContext("secondTilde"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for secondTilde")
		}
	}

	// Peek Field (tildePeek2)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("tildePeek2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for tildePeek2")
	}
	tildePeek2, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'tildePeek2' field of RequestReset")
	}
	if closeErr := readBuffer.CloseContext("tildePeek2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for tildePeek2")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (thirdTilde) (Can be skipped, if a given expression evaluates to false)
	var thirdTilde *RequestType = nil
	if bool((tildePeek2) == (RequestType_RESET)) {
		if pullErr := readBuffer.PullContext("thirdTilde"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for thirdTilde")
		}
		_val, _err := RequestTypeParseWithBuffer(ctx, readBuffer)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'thirdTilde' field of RequestReset")
		}
		thirdTilde = &_val
		if closeErr := readBuffer.CloseContext("thirdTilde"); closeErr != nil {
			return nil, errors.Wrap(closeErr, "Error closing for thirdTilde")
		}
	}

	if closeErr := readBuffer.CloseContext("RequestReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestReset")
	}

	// Create a partially initialized instance
	_child := &_RequestReset{
		_Request: &_Request{
			CBusOptions: cBusOptions,
		},
		TildePeek:   tildePeek,
		SecondTilde: secondTilde,
		TildePeek2:  tildePeek2,
		ThirdTilde:  thirdTilde,
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestReset")
		}

		// Optional Field (secondTilde) (Can be skipped, if the value is null)
		var secondTilde *RequestType = nil
		if m.GetSecondTilde() != nil {
			if pushErr := writeBuffer.PushContext("secondTilde"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for secondTilde")
			}
			secondTilde = m.GetSecondTilde()
			_secondTildeErr := writeBuffer.WriteSerializable(ctx, secondTilde)
			if popErr := writeBuffer.PopContext("secondTilde"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for secondTilde")
			}
			if _secondTildeErr != nil {
				return errors.Wrap(_secondTildeErr, "Error serializing 'secondTilde' field")
			}
		}

		// Optional Field (thirdTilde) (Can be skipped, if the value is null)
		var thirdTilde *RequestType = nil
		if m.GetThirdTilde() != nil {
			if pushErr := writeBuffer.PushContext("thirdTilde"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for thirdTilde")
			}
			thirdTilde = m.GetThirdTilde()
			_thirdTildeErr := writeBuffer.WriteSerializable(ctx, thirdTilde)
			if popErr := writeBuffer.PopContext("thirdTilde"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for thirdTilde")
			}
			if _thirdTildeErr != nil {
				return errors.Wrap(_thirdTildeErr, "Error serializing 'thirdTilde' field")
			}
		}

		if popErr := writeBuffer.PopContext("RequestReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestReset")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestReset) isRequestReset() bool {
	return true
}

func (m *_RequestReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
