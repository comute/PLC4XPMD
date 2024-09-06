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

// Code generated by "plc4xGenerator -type=requestTransactionManager"; DO NOT EDIT.

package transactions

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *requestTransactionManager) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *requestTransactionManager) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("requestTransactionManager"); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("runningRequests", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.runningRequests {
		var elem any = elem

		if elem != nil {
			if serializableField, ok := any(elem).(utils.Serializable); ok {
				if err := writeBuffer.PushContext("value"); err != nil {
					return err
				}
				if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
					return err
				}
				if err := writeBuffer.PopContext("value"); err != nil {
					return err
				}
			} else {
				stringValue := fmt.Sprintf("%v", elem)
				if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), stringValue); err != nil {
					return err
				}
			}
		}
	}
	if err := writeBuffer.PopContext("runningRequests", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt64("numberOfConcurrentRequests", 64, int64(d.numberOfConcurrentRequests)); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt32("currentTransactionId", 32, int32(d.currentTransactionId)); err != nil {
		return err
	}

	if d.executor != nil {
		if serializableField, ok := any(d.executor).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("executor"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("executor"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.executor)
			if err := writeBuffer.WriteString("executor", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteBit("shutdown", d.shutdown.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("traceTransactionManagerTransactions", d.traceTransactionManagerTransactions); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("requestTransactionManager"); err != nil {
		return err
	}
	return nil
}

func (d *requestTransactionManager) String() string {
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
