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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoredSALShortFormBasicMode is the corresponding interface of MonitoredSALShortFormBasicMode
type MonitoredSALShortFormBasicMode interface {
	utils.LengthAware
	utils.Serializable
	MonitoredSAL
	// GetCounts returns Counts (property field)
	GetCounts() byte
	// GetBridgeCount returns BridgeCount (property field)
	GetBridgeCount() BridgeCount
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() NetworkNumber
	// GetNoCounts returns NoCounts (property field)
	GetNoCounts() *byte
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
}

// MonitoredSALShortFormBasicModeExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALShortFormBasicMode.
// This is useful for switch cases.
type MonitoredSALShortFormBasicModeExactly interface {
	MonitoredSALShortFormBasicMode
	isMonitoredSALShortFormBasicMode() bool
}

// _MonitoredSALShortFormBasicMode is the data-structure of this message
type _MonitoredSALShortFormBasicMode struct {
	*_MonitoredSAL
	Counts        byte
	BridgeCount   BridgeCount
	NetworkNumber NetworkNumber
	NoCounts      *byte
	Application   ApplicationIdContainer
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALShortFormBasicMode) InitializeParent(parent MonitoredSAL, salType byte, salData SALData) {
	m.SalType = salType
	m.SalData = salData
}

func (m *_MonitoredSALShortFormBasicMode) GetParent() MonitoredSAL {
	return m._MonitoredSAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALShortFormBasicMode) GetCounts() byte {
	return m.Counts
}

func (m *_MonitoredSALShortFormBasicMode) GetBridgeCount() BridgeCount {
	return m.BridgeCount
}

func (m *_MonitoredSALShortFormBasicMode) GetNetworkNumber() NetworkNumber {
	return m.NetworkNumber
}

func (m *_MonitoredSALShortFormBasicMode) GetNoCounts() *byte {
	return m.NoCounts
}

func (m *_MonitoredSALShortFormBasicMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSALShortFormBasicMode factory function for _MonitoredSALShortFormBasicMode
func NewMonitoredSALShortFormBasicMode(counts byte, bridgeCount BridgeCount, networkNumber NetworkNumber, noCounts *byte, application ApplicationIdContainer, salType byte, salData SALData) *_MonitoredSALShortFormBasicMode {
	_result := &_MonitoredSALShortFormBasicMode{
		Counts:        counts,
		BridgeCount:   bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts:      noCounts,
		Application:   application,
		_MonitoredSAL: NewMonitoredSAL(salType, salData),
	}
	_result._MonitoredSAL._MonitoredSALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALShortFormBasicMode(structType interface{}) MonitoredSALShortFormBasicMode {
	if casted, ok := structType.(MonitoredSALShortFormBasicMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALShortFormBasicMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALShortFormBasicMode) GetTypeName() string {
	return "MonitoredSALShortFormBasicMode"
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Optional Field (bridgeCount)
	if m.BridgeCount != nil {
		lengthInBits += m.BridgeCount.GetLengthInBits()
	}

	// Optional Field (networkNumber)
	if m.NetworkNumber != nil {
		lengthInBits += m.NetworkNumber.GetLengthInBits()
	}

	// Optional Field (noCounts)
	if m.NoCounts != nil {
		lengthInBits += 8
	}

	// Simple field (application)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MonitoredSALShortFormBasicModeParse(readBuffer utils.ReadBuffer) (MonitoredSALShortFormBasicMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALShortFormBasicMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALShortFormBasicMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (counts)
	currentPos = positionAware.GetPos()
	counts, _err := readBuffer.ReadByte("counts")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'counts' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (bridgeCount) (Can be skipped, if a given expression evaluates to false)
	var bridgeCount BridgeCount = nil
	if bool((counts) != (0x00)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("bridgeCount"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for bridgeCount")
		}
		_val, _err := BridgeCountParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'bridgeCount' field")
		default:
			bridgeCount = _val.(BridgeCount)
			if closeErr := readBuffer.CloseContext("bridgeCount"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for bridgeCount")
			}
		}
	}

	// Optional Field (networkNumber) (Can be skipped, if a given expression evaluates to false)
	var networkNumber NetworkNumber = nil
	if bool((counts) != (0x00)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("networkNumber"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for networkNumber")
		}
		_val, _err := NetworkNumberParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'networkNumber' field")
		default:
			networkNumber = _val.(NetworkNumber)
			if closeErr := readBuffer.CloseContext("networkNumber"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for networkNumber")
			}
		}
	}

	// Optional Field (noCounts) (Can be skipped, if a given expression evaluates to false)
	var noCounts *byte = nil
	if bool((counts) == (0x00)) {
		_val, _err := readBuffer.ReadByte("noCounts")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'noCounts' field")
		}
		noCounts = &_val
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALShortFormBasicMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALShortFormBasicMode")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALShortFormBasicMode{
		Counts:        counts,
		BridgeCount:   bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts:      noCounts,
		Application:   application,
		_MonitoredSAL: &_MonitoredSAL{},
	}
	_child._MonitoredSAL._MonitoredSALChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALShortFormBasicMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALShortFormBasicMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALShortFormBasicMode")
		}

		// Optional Field (bridgeCount) (Can be skipped, if the value is null)
		var bridgeCount BridgeCount = nil
		if m.GetBridgeCount() != nil {
			if pushErr := writeBuffer.PushContext("bridgeCount"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for bridgeCount")
			}
			bridgeCount = m.GetBridgeCount()
			_bridgeCountErr := writeBuffer.WriteSerializable(bridgeCount)
			if popErr := writeBuffer.PopContext("bridgeCount"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for bridgeCount")
			}
			if _bridgeCountErr != nil {
				return errors.Wrap(_bridgeCountErr, "Error serializing 'bridgeCount' field")
			}
		}

		// Optional Field (networkNumber) (Can be skipped, if the value is null)
		var networkNumber NetworkNumber = nil
		if m.GetNetworkNumber() != nil {
			if pushErr := writeBuffer.PushContext("networkNumber"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for networkNumber")
			}
			networkNumber = m.GetNetworkNumber()
			_networkNumberErr := writeBuffer.WriteSerializable(networkNumber)
			if popErr := writeBuffer.PopContext("networkNumber"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for networkNumber")
			}
			if _networkNumberErr != nil {
				return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
			}
		}

		// Optional Field (noCounts) (Can be skipped, if the value is null)
		var noCounts *byte = nil
		if m.GetNoCounts() != nil {
			noCounts = m.GetNoCounts()
			_noCountsErr := writeBuffer.WriteByte("noCounts", *(noCounts))
			if _noCountsErr != nil {
				return errors.Wrap(_noCountsErr, "Error serializing 'noCounts' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredSALShortFormBasicMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALShortFormBasicMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MonitoredSALShortFormBasicMode) isMonitoredSALShortFormBasicMode() bool {
	return true
}

func (m *_MonitoredSALShortFormBasicMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
