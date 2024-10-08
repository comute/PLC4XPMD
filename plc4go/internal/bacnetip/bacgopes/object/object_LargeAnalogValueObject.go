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

package object

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type LargeAnalogValueObject struct {
	Object
}

func NewLargeAnalogValueObject(options ...Option) (*LargeAnalogValueObject, error) {
	l := new(LargeAnalogValueObject)
	objectType := "largeAnalogValue"
	_object_supports_cov := true
	properties := []Property{
		NewReadableProperty("presentValue", V2P(NewDouble)),
		NewReadableProperty("statusFlags", V2P(NewStatusFlags)),
		NewOptionalProperty("eventState", V2P(NewEventState)),
		NewOptionalProperty("reliability", V2P(NewReliability)),
		NewOptionalProperty("outOfService", V2P(NewBoolean)),
		NewReadableProperty("units", V2P(NewEngineeringUnits)),
		NewOptionalProperty("priorityArray", V2P(NewPriorityArray)),
		NewOptionalProperty("relinquishDefault", V2P(NewInteger)),
		NewOptionalProperty("covIncrement", V2P(NewUnsigned)),
		NewOptionalProperty("timeDelay", V2P(NewUnsigned)),
		NewOptionalProperty("notificationClass", V2P(NewUnsigned)),
		NewOptionalProperty("highLimit", V2P(NewDouble)),
		NewOptionalProperty("lowLimit", V2P(NewDouble)),
		NewOptionalProperty("deadband", V2P(NewDouble)),
		NewOptionalProperty("limitEnable", V2P(NewLimitEnable)),
		NewOptionalProperty("eventEnable", V2P(NewEventTransitionBits)),
		NewOptionalProperty("ackedTransitions", V2P(NewEventTransitionBits)),
		NewOptionalProperty("notifyType", V2P(NewNotifyType)),
		NewOptionalProperty("eventTimeStamps", ArrayOfP(NewTimeStamp, 3, 0)),
		NewOptionalProperty("eventMessageTexts", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventMessageTextsConfig", ArrayOfP(NewCharacterString, 3, 0)),
		NewOptionalProperty("eventDetectionEnable", V2P(NewBoolean)),
		NewOptionalProperty("eventAlgorithmInhibitRef", V2P(NewObjectPropertyReference)),
		NewOptionalProperty("eventAlgorithmInhibit", V2P(NewBoolean)),
		NewOptionalProperty("timeDelayNormal", V2P(NewUnsigned)),
		NewOptionalProperty("reliabilityEvaluationInhibit", V2P(NewBoolean)),
		NewOptionalProperty("minPresValue", V2P(NewDouble)),
		NewOptionalProperty("maxPresValue", V2P(NewDouble)),
		NewOptionalProperty("resolution", V2P(NewDouble)),
		NewOptionalProperty("faultHighLimit", V2P(NewDouble)),
		NewOptionalProperty("faultLowLimit", V2P(NewDouble)),
		NewOptionalProperty("currentCommandPriority", V2P(NewOptionalUnsigned)),
		NewOptionalProperty("valueSource", V2P(NewValueSource)),
		NewOptionalProperty("valueSourceArray", ArrayOfP(NewValueSource, 16, 0)),
		NewOptionalProperty("lastCommandTime", V2P(NewTimeStamp)),
		NewOptionalProperty("commandTimeArray", ArrayOfP(NewTimeStamp, 16, 0)),
		NewOptionalProperty("auditablePriorityFilter", V2P(NewOptionalPriorityFilter)),
	}
	var err error
	l.Object, err = NewObject(Combine(options, WithObjectType(objectType), WithObjectExtraProperties(properties), WithObjectSupportsCov(_object_supports_cov))...)
	if err != nil {
		return nil, errors.Wrap(err, "error creating object")
	}
	if _, err := RegisterObjectType(NKW(KWCls, l)); err != nil {
		return nil, errors.Wrap(err, "error registering object type")
	}
	return l, nil
}
