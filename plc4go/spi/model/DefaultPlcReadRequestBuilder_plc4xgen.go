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

// Code generated by "plc4xgenerator -type=DefaultPlcReadRequestBuilder"; DO NOT EDIT.

package model

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcReadRequestBuilder) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcReadRequestBuilder) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcReadRequestBuilder"); err != nil {
		return err
	}

	if d.reader != nil {
		if serializableField, ok := d.reader.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("reader"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("reader"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.reader)
			if err := writeBuffer.WriteString("reader", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if d.tagHandler != nil {
		if serializableField, ok := d.tagHandler.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tagHandler"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tagHandler"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.tagHandler)
			if err := writeBuffer.WriteString("tagHandler", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PushContext("tagNames", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.tagNames {
		if err := writeBuffer.WriteString("", uint32(len(elem)*8), "UTF-8", elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("tagNames", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("tagAddresses", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for name, elem := range d.tagAddresses {

		if err := writeBuffer.WriteString(name, uint32(len(elem)*8), "UTF-8", elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("tagAddresses", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("tags", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for name, elem := range d.tags {

		var elem any = elem
		if serializable, ok := elem.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemAsString := fmt.Sprintf("%v", elem)
			if err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), "UTF-8", elemAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("tags", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if d.readRequestInterceptor != nil {
		if serializableField, ok := d.readRequestInterceptor.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("readRequestInterceptor"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("readRequestInterceptor"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.readRequestInterceptor)
			if err := writeBuffer.WriteString("readRequestInterceptor", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("PlcReadRequestBuilder"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcReadRequestBuilder) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
