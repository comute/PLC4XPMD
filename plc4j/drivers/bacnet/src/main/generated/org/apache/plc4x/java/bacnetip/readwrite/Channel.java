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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class Channel implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final OptionalProperty description;
  protected final WritableProperty presentValue;
  protected final ReadableProperty lastPriority;
  protected final ReadableProperty writeStatus;
  protected final ReadableProperty statusFlags;
  protected final OptionalProperty reliability;
  protected final ReadableProperty outOfService;
  protected final WritableProperty listOfObjectPropertyReferences;
  protected final OptionalProperty executionDelay;
  protected final OptionalProperty allowGroupDelayInhibit;
  protected final WritableProperty channelNumber;
  protected final WritableProperty controlGroups;
  protected final OptionalProperty eventDetectionEnable;
  protected final OptionalProperty notificationClass;
  protected final OptionalProperty eventEnable;
  protected final OptionalProperty eventState;
  protected final OptionalProperty ackedTransitions;
  protected final OptionalProperty notifyType;
  protected final OptionalProperty eventTimeStamps;
  protected final OptionalProperty eventMessageTexts;
  protected final OptionalProperty eventMessageTextsConfig;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final ReadableProperty propertyList;
  protected final OptionalProperty valueSource;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public Channel(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      OptionalProperty description,
      WritableProperty presentValue,
      ReadableProperty lastPriority,
      ReadableProperty writeStatus,
      ReadableProperty statusFlags,
      OptionalProperty reliability,
      ReadableProperty outOfService,
      WritableProperty listOfObjectPropertyReferences,
      OptionalProperty executionDelay,
      OptionalProperty allowGroupDelayInhibit,
      WritableProperty channelNumber,
      WritableProperty controlGroups,
      OptionalProperty eventDetectionEnable,
      OptionalProperty notificationClass,
      OptionalProperty eventEnable,
      OptionalProperty eventState,
      OptionalProperty ackedTransitions,
      OptionalProperty notifyType,
      OptionalProperty eventTimeStamps,
      OptionalProperty eventMessageTexts,
      OptionalProperty eventMessageTextsConfig,
      OptionalProperty reliabilityEvaluationInhibit,
      ReadableProperty propertyList,
      OptionalProperty valueSource,
      OptionalProperty tags,
      OptionalProperty profileLocation,
      OptionalProperty profileName) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.objectName = objectName;
    this.objectType = objectType;
    this.description = description;
    this.presentValue = presentValue;
    this.lastPriority = lastPriority;
    this.writeStatus = writeStatus;
    this.statusFlags = statusFlags;
    this.reliability = reliability;
    this.outOfService = outOfService;
    this.listOfObjectPropertyReferences = listOfObjectPropertyReferences;
    this.executionDelay = executionDelay;
    this.allowGroupDelayInhibit = allowGroupDelayInhibit;
    this.channelNumber = channelNumber;
    this.controlGroups = controlGroups;
    this.eventDetectionEnable = eventDetectionEnable;
    this.notificationClass = notificationClass;
    this.eventEnable = eventEnable;
    this.eventState = eventState;
    this.ackedTransitions = ackedTransitions;
    this.notifyType = notifyType;
    this.eventTimeStamps = eventTimeStamps;
    this.eventMessageTexts = eventMessageTexts;
    this.eventMessageTextsConfig = eventMessageTextsConfig;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.propertyList = propertyList;
    this.valueSource = valueSource;
    this.tags = tags;
    this.profileLocation = profileLocation;
    this.profileName = profileName;
  }

  public ReadableProperty getObjectIdentifier() {
    return objectIdentifier;
  }

  public ReadableProperty getObjectName() {
    return objectName;
  }

  public ReadableProperty getObjectType() {
    return objectType;
  }

  public OptionalProperty getDescription() {
    return description;
  }

  public WritableProperty getPresentValue() {
    return presentValue;
  }

  public ReadableProperty getLastPriority() {
    return lastPriority;
  }

  public ReadableProperty getWriteStatus() {
    return writeStatus;
  }

  public ReadableProperty getStatusFlags() {
    return statusFlags;
  }

  public OptionalProperty getReliability() {
    return reliability;
  }

  public ReadableProperty getOutOfService() {
    return outOfService;
  }

  public WritableProperty getListOfObjectPropertyReferences() {
    return listOfObjectPropertyReferences;
  }

  public OptionalProperty getExecutionDelay() {
    return executionDelay;
  }

  public OptionalProperty getAllowGroupDelayInhibit() {
    return allowGroupDelayInhibit;
  }

  public WritableProperty getChannelNumber() {
    return channelNumber;
  }

  public WritableProperty getControlGroups() {
    return controlGroups;
  }

  public OptionalProperty getEventDetectionEnable() {
    return eventDetectionEnable;
  }

  public OptionalProperty getNotificationClass() {
    return notificationClass;
  }

  public OptionalProperty getEventEnable() {
    return eventEnable;
  }

  public OptionalProperty getEventState() {
    return eventState;
  }

  public OptionalProperty getAckedTransitions() {
    return ackedTransitions;
  }

  public OptionalProperty getNotifyType() {
    return notifyType;
  }

  public OptionalProperty getEventTimeStamps() {
    return eventTimeStamps;
  }

  public OptionalProperty getEventMessageTexts() {
    return eventMessageTexts;
  }

  public OptionalProperty getEventMessageTextsConfig() {
    return eventMessageTextsConfig;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public OptionalProperty getValueSource() {
    return valueSource;
  }

  public OptionalProperty getTags() {
    return tags;
  }

  public OptionalProperty getProfileLocation() {
    return profileLocation;
  }

  public OptionalProperty getProfileName() {
    return profileName;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Channel");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (objectName)
    writeSimpleField("objectName", objectName, writeComplex(writeBuffer));

    // Simple Field (objectType)
    writeSimpleField("objectType", objectType, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (presentValue)
    writeSimpleField("presentValue", presentValue, writeComplex(writeBuffer));

    // Simple Field (lastPriority)
    writeSimpleField("lastPriority", lastPriority, writeComplex(writeBuffer));

    // Simple Field (writeStatus)
    writeSimpleField("writeStatus", writeStatus, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (outOfService)
    writeSimpleField("outOfService", outOfService, writeComplex(writeBuffer));

    // Simple Field (listOfObjectPropertyReferences)
    writeSimpleField(
        "listOfObjectPropertyReferences",
        listOfObjectPropertyReferences,
        writeComplex(writeBuffer));

    // Simple Field (executionDelay)
    writeSimpleField("executionDelay", executionDelay, writeComplex(writeBuffer));

    // Simple Field (allowGroupDelayInhibit)
    writeSimpleField("allowGroupDelayInhibit", allowGroupDelayInhibit, writeComplex(writeBuffer));

    // Simple Field (channelNumber)
    writeSimpleField("channelNumber", channelNumber, writeComplex(writeBuffer));

    // Simple Field (controlGroups)
    writeSimpleField("controlGroups", controlGroups, writeComplex(writeBuffer));

    // Simple Field (eventDetectionEnable)
    writeSimpleField("eventDetectionEnable", eventDetectionEnable, writeComplex(writeBuffer));

    // Simple Field (notificationClass)
    writeSimpleField("notificationClass", notificationClass, writeComplex(writeBuffer));

    // Simple Field (eventEnable)
    writeSimpleField("eventEnable", eventEnable, writeComplex(writeBuffer));

    // Simple Field (eventState)
    writeSimpleField("eventState", eventState, writeComplex(writeBuffer));

    // Simple Field (ackedTransitions)
    writeSimpleField("ackedTransitions", ackedTransitions, writeComplex(writeBuffer));

    // Simple Field (notifyType)
    writeSimpleField("notifyType", notifyType, writeComplex(writeBuffer));

    // Simple Field (eventTimeStamps)
    writeSimpleField("eventTimeStamps", eventTimeStamps, writeComplex(writeBuffer));

    // Simple Field (eventMessageTexts)
    writeSimpleField("eventMessageTexts", eventMessageTexts, writeComplex(writeBuffer));

    // Simple Field (eventMessageTextsConfig)
    writeSimpleField("eventMessageTextsConfig", eventMessageTextsConfig, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

    // Simple Field (valueSource)
    writeSimpleField("valueSource", valueSource, writeComplex(writeBuffer));

    // Simple Field (tags)
    writeSimpleField("tags", tags, writeComplex(writeBuffer));

    // Simple Field (profileLocation)
    writeSimpleField("profileLocation", profileLocation, writeComplex(writeBuffer));

    // Simple Field (profileName)
    writeSimpleField("profileName", profileName, writeComplex(writeBuffer));

    writeBuffer.popContext("Channel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Channel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (objectName)
    lengthInBits += objectName.getLengthInBits();

    // Simple field (objectType)
    lengthInBits += objectType.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (presentValue)
    lengthInBits += presentValue.getLengthInBits();

    // Simple field (lastPriority)
    lengthInBits += lastPriority.getLengthInBits();

    // Simple field (writeStatus)
    lengthInBits += writeStatus.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (outOfService)
    lengthInBits += outOfService.getLengthInBits();

    // Simple field (listOfObjectPropertyReferences)
    lengthInBits += listOfObjectPropertyReferences.getLengthInBits();

    // Simple field (executionDelay)
    lengthInBits += executionDelay.getLengthInBits();

    // Simple field (allowGroupDelayInhibit)
    lengthInBits += allowGroupDelayInhibit.getLengthInBits();

    // Simple field (channelNumber)
    lengthInBits += channelNumber.getLengthInBits();

    // Simple field (controlGroups)
    lengthInBits += controlGroups.getLengthInBits();

    // Simple field (eventDetectionEnable)
    lengthInBits += eventDetectionEnable.getLengthInBits();

    // Simple field (notificationClass)
    lengthInBits += notificationClass.getLengthInBits();

    // Simple field (eventEnable)
    lengthInBits += eventEnable.getLengthInBits();

    // Simple field (eventState)
    lengthInBits += eventState.getLengthInBits();

    // Simple field (ackedTransitions)
    lengthInBits += ackedTransitions.getLengthInBits();

    // Simple field (notifyType)
    lengthInBits += notifyType.getLengthInBits();

    // Simple field (eventTimeStamps)
    lengthInBits += eventTimeStamps.getLengthInBits();

    // Simple field (eventMessageTexts)
    lengthInBits += eventMessageTexts.getLengthInBits();

    // Simple field (eventMessageTextsConfig)
    lengthInBits += eventMessageTextsConfig.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

    // Simple field (valueSource)
    lengthInBits += valueSource.getLengthInBits();

    // Simple field (tags)
    lengthInBits += tags.getLengthInBits();

    // Simple field (profileLocation)
    lengthInBits += profileLocation.getLengthInBits();

    // Simple field (profileName)
    lengthInBits += profileName.getLengthInBits();

    return lengthInBits;
  }

  public static Channel staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Channel");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReadableProperty objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectIdentifier")),
                readBuffer));

    ReadableProperty objectName =
        readSimpleField(
            "objectName",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    ReadableProperty objectType =
        readSimpleField(
            "objectType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectType")),
                readBuffer));

    OptionalProperty description =
        readSimpleField(
            "description",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    WritableProperty presentValue =
        readSimpleField(
            "presentValue",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("BACnetChannelValue")),
                readBuffer));

    ReadableProperty lastPriority =
        readSimpleField(
            "lastPriority",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    ReadableProperty writeStatus =
        readSimpleField(
            "writeStatus",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetWriteStatus")),
                readBuffer));

    ReadableProperty statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetStatusFlags")),
                readBuffer));

    OptionalProperty reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetReliability")),
                readBuffer));

    ReadableProperty outOfService =
        readSimpleField(
            "outOfService",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    WritableProperty listOfObjectPropertyReferences =
        readSimpleField(
            "listOfObjectPropertyReferences",
            readComplex(
                () ->
                    WritableProperty.staticParse(
                        readBuffer,
                        (String) ("BACnetARRAY[N] of BACnetDeviceObjectPropertyReference")),
                readBuffer));

    OptionalProperty executionDelay =
        readSimpleField(
            "executionDelay",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of Unsigned")),
                readBuffer));

    OptionalProperty allowGroupDelayInhibit =
        readSimpleField(
            "allowGroupDelayInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    WritableProperty channelNumber =
        readSimpleField(
            "channelNumber",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("Unsigned16")),
                readBuffer));

    WritableProperty controlGroups =
        readSimpleField(
            "controlGroups",
            readComplex(
                () ->
                    WritableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of Unsigned32")),
                readBuffer));

    OptionalProperty eventDetectionEnable =
        readSimpleField(
            "eventDetectionEnable",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty notificationClass =
        readSimpleField(
            "notificationClass",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty eventEnable =
        readSimpleField(
            "eventEnable",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty eventState =
        readSimpleField(
            "eventState",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetEventState")),
                readBuffer));

    OptionalProperty ackedTransitions =
        readSimpleField(
            "ackedTransitions",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty notifyType =
        readSimpleField(
            "notifyType",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetNotifyType")),
                readBuffer));

    OptionalProperty eventTimeStamps =
        readSimpleField(
            "eventTimeStamps",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of BACnetTimeStamp")),
                readBuffer));

    OptionalProperty eventMessageTexts =
        readSimpleField(
            "eventMessageTexts",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty eventMessageTextsConfig =
        readSimpleField(
            "eventMessageTextsConfig",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty reliabilityEvaluationInhibit =
        readSimpleField(
            "reliabilityEvaluationInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty propertyList =
        readSimpleField(
            "propertyList",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetPropertyIdentifier")),
                readBuffer));

    OptionalProperty valueSource =
        readSimpleField(
            "valueSource",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetValueSource")),
                readBuffer));

    OptionalProperty tags =
        readSimpleField(
            "tags",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetNameValue")),
                readBuffer));

    OptionalProperty profileLocation =
        readSimpleField(
            "profileLocation",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty profileName =
        readSimpleField(
            "profileName",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    readBuffer.closeContext("Channel");
    // Create the instance
    Channel _channel;
    _channel =
        new Channel(
            objectIdentifier,
            objectName,
            objectType,
            description,
            presentValue,
            lastPriority,
            writeStatus,
            statusFlags,
            reliability,
            outOfService,
            listOfObjectPropertyReferences,
            executionDelay,
            allowGroupDelayInhibit,
            channelNumber,
            controlGroups,
            eventDetectionEnable,
            notificationClass,
            eventEnable,
            eventState,
            ackedTransitions,
            notifyType,
            eventTimeStamps,
            eventMessageTexts,
            eventMessageTextsConfig,
            reliabilityEvaluationInhibit,
            propertyList,
            valueSource,
            tags,
            profileLocation,
            profileName);
    return _channel;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Channel)) {
      return false;
    }
    Channel that = (Channel) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getDescription() == that.getDescription())
        && (getPresentValue() == that.getPresentValue())
        && (getLastPriority() == that.getLastPriority())
        && (getWriteStatus() == that.getWriteStatus())
        && (getStatusFlags() == that.getStatusFlags())
        && (getReliability() == that.getReliability())
        && (getOutOfService() == that.getOutOfService())
        && (getListOfObjectPropertyReferences() == that.getListOfObjectPropertyReferences())
        && (getExecutionDelay() == that.getExecutionDelay())
        && (getAllowGroupDelayInhibit() == that.getAllowGroupDelayInhibit())
        && (getChannelNumber() == that.getChannelNumber())
        && (getControlGroups() == that.getControlGroups())
        && (getEventDetectionEnable() == that.getEventDetectionEnable())
        && (getNotificationClass() == that.getNotificationClass())
        && (getEventEnable() == that.getEventEnable())
        && (getEventState() == that.getEventState())
        && (getAckedTransitions() == that.getAckedTransitions())
        && (getNotifyType() == that.getNotifyType())
        && (getEventTimeStamps() == that.getEventTimeStamps())
        && (getEventMessageTexts() == that.getEventMessageTexts())
        && (getEventMessageTextsConfig() == that.getEventMessageTextsConfig())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getPropertyList() == that.getPropertyList())
        && (getValueSource() == that.getValueSource())
        && (getTags() == that.getTags())
        && (getProfileLocation() == that.getProfileLocation())
        && (getProfileName() == that.getProfileName())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getObjectIdentifier(),
        getObjectName(),
        getObjectType(),
        getDescription(),
        getPresentValue(),
        getLastPriority(),
        getWriteStatus(),
        getStatusFlags(),
        getReliability(),
        getOutOfService(),
        getListOfObjectPropertyReferences(),
        getExecutionDelay(),
        getAllowGroupDelayInhibit(),
        getChannelNumber(),
        getControlGroups(),
        getEventDetectionEnable(),
        getNotificationClass(),
        getEventEnable(),
        getEventState(),
        getAckedTransitions(),
        getNotifyType(),
        getEventTimeStamps(),
        getEventMessageTexts(),
        getEventMessageTextsConfig(),
        getReliabilityEvaluationInhibit(),
        getPropertyList(),
        getValueSource(),
        getTags(),
        getProfileLocation(),
        getProfileName());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}