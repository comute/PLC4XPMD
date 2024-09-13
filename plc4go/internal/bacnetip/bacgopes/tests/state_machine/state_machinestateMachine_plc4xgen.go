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

// Code generated by "plc4xGenerator -type=stateMachine -prefix=state_machine"; DO NOT EDIT.

package state_machine

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *stateMachine) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *stateMachine) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("stateMachine"); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.interceptor)

		if err := writeBuffer.WriteString("interceptor", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteBit("stateDecorator", d.stateDecorator != nil); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("states", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.states {
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString("states", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("states", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("name", uint32(len(d.name)*8), d.name); err != nil {
		return err
	}
	if d.machineGroup != nil {
		{
			_value := fmt.Sprintf("%v", d.machineGroup)

			if err := writeBuffer.WriteString("machineGroup", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	{
		_value := fmt.Sprintf("%v", d.stateSubStruct)

		if err := writeBuffer.WriteString("stateSubStruct", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	{
		_value := fmt.Sprintf("%v", d.startState)

		if err := writeBuffer.WriteString("startState", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	{
		_value := fmt.Sprintf("%v", d.unexpectedReceiveState)

		if err := writeBuffer.WriteString("unexpectedReceiveState", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}

	_transitionQueue_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.transitionQueue))
	if err := writeBuffer.WriteString("transitionQueue", uint32(len(_transitionQueue_plx4gen_description)*8), _transitionQueue_plx4gen_description); err != nil {
		return err
	}
	if d.stateTimeoutTask != nil {
		{
			_value := fmt.Sprintf("%v", d.stateTimeoutTask)

			if err := writeBuffer.WriteString("stateTimeoutTask", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("timeout", uint32(len(fmt.Sprintf("%s", d.timeout))*8), fmt.Sprintf("%s", d.timeout)); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.timeoutState)

		if err := writeBuffer.WriteString("timeoutState", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteString("stateMachineTimeout", uint32(len(fmt.Sprintf("%s", d.stateMachineTimeout))*8), fmt.Sprintf("%s", d.stateMachineTimeout)); err != nil {
		return err
	}
	if d.timeoutTask != nil {
		{
			_value := fmt.Sprintf("%v", d.timeoutTask)

			if err := writeBuffer.WriteString("timeoutTask", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteBit("running", d.running); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("startupFlag", d.startupFlag); err != nil {
		return err
	}
	if d.isSuccessState != nil {
		if err := writeBuffer.WriteBit("isSuccessState", *d.isSuccessState); err != nil {
			return err
		}
	}
	if d.isFailState != nil {
		if err := writeBuffer.WriteBit("isFailState", *d.isFailState); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteInt64("stateTransitioning", 64, int64(d.stateTransitioning)); err != nil {
		return err
	}
	{
		_value := fmt.Sprintf("%v", d.currentState)

		if err := writeBuffer.WriteString("currentState", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PushContext("transactionLog", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.transactionLog {
		if err := writeBuffer.WriteString("", uint32(len(elem)*8), elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("transactionLog", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("stateMachine"); err != nil {
		return err
	}
	return nil
}

func (d *stateMachine) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
