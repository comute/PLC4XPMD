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

public class PositiveIntegerValue implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final OptionalProperty description;
  protected final ReadableProperty presentValue;
  protected final ReadableProperty statusFlags;
  protected final OptionalProperty eventState;
  protected final OptionalProperty reliability;
  protected final OptionalProperty outOfService;
  protected final ReadableProperty units;
  protected final OptionalProperty priorityArray;
  protected final OptionalProperty relinquishDefault;
  protected final OptionalProperty cOVIncrement;
  protected final OptionalProperty timeDelay;
  protected final OptionalProperty notificationClass;
  protected final OptionalProperty highLimit;
  protected final OptionalProperty lowLimit;
  protected final OptionalProperty deadband;
  protected final OptionalProperty limitEnable;
  protected final OptionalProperty eventEnable;
  protected final OptionalProperty ackedTransitions;
  protected final OptionalProperty notifyType;
  protected final OptionalProperty eventTimeStamps;
  protected final OptionalProperty eventMessageTexts;
  protected final OptionalProperty eventMessageTextsConfig;
  protected final OptionalProperty eventDetectionEnable;
  protected final OptionalProperty eventAlgorithmInhibitRef;
  protected final OptionalProperty eventAlgorithmInhibit;
  protected final OptionalProperty timeDelayNormal;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final OptionalProperty minPresValue;
  protected final OptionalProperty maxPresValue;
  protected final OptionalProperty resolution;
  protected final ReadableProperty propertyList;
  protected final OptionalProperty faultHighLimit;
  protected final OptionalProperty faultLowLimit;
  protected final OptionalProperty currentCommandPriority;
  protected final OptionalProperty valueSource;
  protected final OptionalProperty valueSourceArray;
  protected final OptionalProperty lastCommandTime;
  protected final OptionalProperty commandTimeArray;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public PositiveIntegerValue(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      OptionalProperty description,
      ReadableProperty presentValue,
      ReadableProperty statusFlags,
      OptionalProperty eventState,
      OptionalProperty reliability,
      OptionalProperty outOfService,
      ReadableProperty units,
      OptionalProperty priorityArray,
      OptionalProperty relinquishDefault,
      OptionalProperty cOVIncrement,
      OptionalProperty timeDelay,
      OptionalProperty notificationClass,
      OptionalProperty highLimit,
      OptionalProperty lowLimit,
      OptionalProperty deadband,
      OptionalProperty limitEnable,
      OptionalProperty eventEnable,
      OptionalProperty ackedTransitions,
      OptionalProperty notifyType,
      OptionalProperty eventTimeStamps,
      OptionalProperty eventMessageTexts,
      OptionalProperty eventMessageTextsConfig,
      OptionalProperty eventDetectionEnable,
      OptionalProperty eventAlgorithmInhibitRef,
      OptionalProperty eventAlgorithmInhibit,
      OptionalProperty timeDelayNormal,
      OptionalProperty reliabilityEvaluationInhibit,
      OptionalProperty minPresValue,
      OptionalProperty maxPresValue,
      OptionalProperty resolution,
      ReadableProperty propertyList,
      OptionalProperty faultHighLimit,
      OptionalProperty faultLowLimit,
      OptionalProperty currentCommandPriority,
      OptionalProperty valueSource,
      OptionalProperty valueSourceArray,
      OptionalProperty lastCommandTime,
      OptionalProperty commandTimeArray,
      OptionalProperty tags,
      OptionalProperty profileLocation,
      OptionalProperty profileName) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.objectName = objectName;
    this.objectType = objectType;
    this.description = description;
    this.presentValue = presentValue;
    this.statusFlags = statusFlags;
    this.eventState = eventState;
    this.reliability = reliability;
    this.outOfService = outOfService;
    this.units = units;
    this.priorityArray = priorityArray;
    this.relinquishDefault = relinquishDefault;
    this.cOVIncrement = cOVIncrement;
    this.timeDelay = timeDelay;
    this.notificationClass = notificationClass;
    this.highLimit = highLimit;
    this.lowLimit = lowLimit;
    this.deadband = deadband;
    this.limitEnable = limitEnable;
    this.eventEnable = eventEnable;
    this.ackedTransitions = ackedTransitions;
    this.notifyType = notifyType;
    this.eventTimeStamps = eventTimeStamps;
    this.eventMessageTexts = eventMessageTexts;
    this.eventMessageTextsConfig = eventMessageTextsConfig;
    this.eventDetectionEnable = eventDetectionEnable;
    this.eventAlgorithmInhibitRef = eventAlgorithmInhibitRef;
    this.eventAlgorithmInhibit = eventAlgorithmInhibit;
    this.timeDelayNormal = timeDelayNormal;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.minPresValue = minPresValue;
    this.maxPresValue = maxPresValue;
    this.resolution = resolution;
    this.propertyList = propertyList;
    this.faultHighLimit = faultHighLimit;
    this.faultLowLimit = faultLowLimit;
    this.currentCommandPriority = currentCommandPriority;
    this.valueSource = valueSource;
    this.valueSourceArray = valueSourceArray;
    this.lastCommandTime = lastCommandTime;
    this.commandTimeArray = commandTimeArray;
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

  public ReadableProperty getPresentValue() {
    return presentValue;
  }

  public ReadableProperty getStatusFlags() {
    return statusFlags;
  }

  public OptionalProperty getEventState() {
    return eventState;
  }

  public OptionalProperty getReliability() {
    return reliability;
  }

  public OptionalProperty getOutOfService() {
    return outOfService;
  }

  public ReadableProperty getUnits() {
    return units;
  }

  public OptionalProperty getPriorityArray() {
    return priorityArray;
  }

  public OptionalProperty getRelinquishDefault() {
    return relinquishDefault;
  }

  public OptionalProperty getCOVIncrement() {
    return cOVIncrement;
  }

  public OptionalProperty getTimeDelay() {
    return timeDelay;
  }

  public OptionalProperty getNotificationClass() {
    return notificationClass;
  }

  public OptionalProperty getHighLimit() {
    return highLimit;
  }

  public OptionalProperty getLowLimit() {
    return lowLimit;
  }

  public OptionalProperty getDeadband() {
    return deadband;
  }

  public OptionalProperty getLimitEnable() {
    return limitEnable;
  }

  public OptionalProperty getEventEnable() {
    return eventEnable;
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

  public OptionalProperty getEventDetectionEnable() {
    return eventDetectionEnable;
  }

  public OptionalProperty getEventAlgorithmInhibitRef() {
    return eventAlgorithmInhibitRef;
  }

  public OptionalProperty getEventAlgorithmInhibit() {
    return eventAlgorithmInhibit;
  }

  public OptionalProperty getTimeDelayNormal() {
    return timeDelayNormal;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public OptionalProperty getMinPresValue() {
    return minPresValue;
  }

  public OptionalProperty getMaxPresValue() {
    return maxPresValue;
  }

  public OptionalProperty getResolution() {
    return resolution;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public OptionalProperty getFaultHighLimit() {
    return faultHighLimit;
  }

  public OptionalProperty getFaultLowLimit() {
    return faultLowLimit;
  }

  public OptionalProperty getCurrentCommandPriority() {
    return currentCommandPriority;
  }

  public OptionalProperty getValueSource() {
    return valueSource;
  }

  public OptionalProperty getValueSourceArray() {
    return valueSourceArray;
  }

  public OptionalProperty getLastCommandTime() {
    return lastCommandTime;
  }

  public OptionalProperty getCommandTimeArray() {
    return commandTimeArray;
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
    writeBuffer.pushContext("PositiveIntegerValue");

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

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (eventState)
    writeSimpleField("eventState", eventState, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (outOfService)
    writeSimpleField("outOfService", outOfService, writeComplex(writeBuffer));

    // Simple Field (units)
    writeSimpleField("units", units, writeComplex(writeBuffer));

    // Simple Field (priorityArray)
    writeSimpleField("priorityArray", priorityArray, writeComplex(writeBuffer));

    // Simple Field (relinquishDefault)
    writeSimpleField("relinquishDefault", relinquishDefault, writeComplex(writeBuffer));

    // Simple Field (cOVIncrement)
    writeSimpleField("cOVIncrement", cOVIncrement, writeComplex(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, writeComplex(writeBuffer));

    // Simple Field (notificationClass)
    writeSimpleField("notificationClass", notificationClass, writeComplex(writeBuffer));

    // Simple Field (highLimit)
    writeSimpleField("highLimit", highLimit, writeComplex(writeBuffer));

    // Simple Field (lowLimit)
    writeSimpleField("lowLimit", lowLimit, writeComplex(writeBuffer));

    // Simple Field (deadband)
    writeSimpleField("deadband", deadband, writeComplex(writeBuffer));

    // Simple Field (limitEnable)
    writeSimpleField("limitEnable", limitEnable, writeComplex(writeBuffer));

    // Simple Field (eventEnable)
    writeSimpleField("eventEnable", eventEnable, writeComplex(writeBuffer));

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

    // Simple Field (eventDetectionEnable)
    writeSimpleField("eventDetectionEnable", eventDetectionEnable, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibitRef)
    writeSimpleField(
        "eventAlgorithmInhibitRef", eventAlgorithmInhibitRef, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibit)
    writeSimpleField("eventAlgorithmInhibit", eventAlgorithmInhibit, writeComplex(writeBuffer));

    // Simple Field (timeDelayNormal)
    writeSimpleField("timeDelayNormal", timeDelayNormal, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (minPresValue)
    writeSimpleField("minPresValue", minPresValue, writeComplex(writeBuffer));

    // Simple Field (maxPresValue)
    writeSimpleField("maxPresValue", maxPresValue, writeComplex(writeBuffer));

    // Simple Field (resolution)
    writeSimpleField("resolution", resolution, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

    // Simple Field (faultHighLimit)
    writeSimpleField("faultHighLimit", faultHighLimit, writeComplex(writeBuffer));

    // Simple Field (faultLowLimit)
    writeSimpleField("faultLowLimit", faultLowLimit, writeComplex(writeBuffer));

    // Simple Field (currentCommandPriority)
    writeSimpleField("currentCommandPriority", currentCommandPriority, writeComplex(writeBuffer));

    // Simple Field (valueSource)
    writeSimpleField("valueSource", valueSource, writeComplex(writeBuffer));

    // Simple Field (valueSourceArray)
    writeSimpleField("valueSourceArray", valueSourceArray, writeComplex(writeBuffer));

    // Simple Field (lastCommandTime)
    writeSimpleField("lastCommandTime", lastCommandTime, writeComplex(writeBuffer));

    // Simple Field (commandTimeArray)
    writeSimpleField("commandTimeArray", commandTimeArray, writeComplex(writeBuffer));

    // Simple Field (tags)
    writeSimpleField("tags", tags, writeComplex(writeBuffer));

    // Simple Field (profileLocation)
    writeSimpleField("profileLocation", profileLocation, writeComplex(writeBuffer));

    // Simple Field (profileName)
    writeSimpleField("profileName", profileName, writeComplex(writeBuffer));

    writeBuffer.popContext("PositiveIntegerValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PositiveIntegerValue _value = this;
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

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (eventState)
    lengthInBits += eventState.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (outOfService)
    lengthInBits += outOfService.getLengthInBits();

    // Simple field (units)
    lengthInBits += units.getLengthInBits();

    // Simple field (priorityArray)
    lengthInBits += priorityArray.getLengthInBits();

    // Simple field (relinquishDefault)
    lengthInBits += relinquishDefault.getLengthInBits();

    // Simple field (cOVIncrement)
    lengthInBits += cOVIncrement.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (notificationClass)
    lengthInBits += notificationClass.getLengthInBits();

    // Simple field (highLimit)
    lengthInBits += highLimit.getLengthInBits();

    // Simple field (lowLimit)
    lengthInBits += lowLimit.getLengthInBits();

    // Simple field (deadband)
    lengthInBits += deadband.getLengthInBits();

    // Simple field (limitEnable)
    lengthInBits += limitEnable.getLengthInBits();

    // Simple field (eventEnable)
    lengthInBits += eventEnable.getLengthInBits();

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

    // Simple field (eventDetectionEnable)
    lengthInBits += eventDetectionEnable.getLengthInBits();

    // Simple field (eventAlgorithmInhibitRef)
    lengthInBits += eventAlgorithmInhibitRef.getLengthInBits();

    // Simple field (eventAlgorithmInhibit)
    lengthInBits += eventAlgorithmInhibit.getLengthInBits();

    // Simple field (timeDelayNormal)
    lengthInBits += timeDelayNormal.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (minPresValue)
    lengthInBits += minPresValue.getLengthInBits();

    // Simple field (maxPresValue)
    lengthInBits += maxPresValue.getLengthInBits();

    // Simple field (resolution)
    lengthInBits += resolution.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

    // Simple field (faultHighLimit)
    lengthInBits += faultHighLimit.getLengthInBits();

    // Simple field (faultLowLimit)
    lengthInBits += faultLowLimit.getLengthInBits();

    // Simple field (currentCommandPriority)
    lengthInBits += currentCommandPriority.getLengthInBits();

    // Simple field (valueSource)
    lengthInBits += valueSource.getLengthInBits();

    // Simple field (valueSourceArray)
    lengthInBits += valueSourceArray.getLengthInBits();

    // Simple field (lastCommandTime)
    lengthInBits += lastCommandTime.getLengthInBits();

    // Simple field (commandTimeArray)
    lengthInBits += commandTimeArray.getLengthInBits();

    // Simple field (tags)
    lengthInBits += tags.getLengthInBits();

    // Simple field (profileLocation)
    lengthInBits += profileLocation.getLengthInBits();

    // Simple field (profileName)
    lengthInBits += profileName.getLengthInBits();

    return lengthInBits;
  }

  public static PositiveIntegerValue staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PositiveIntegerValue");
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

    ReadableProperty presentValue =
        readSimpleField(
            "presentValue",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    ReadableProperty statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetStatusFlags")),
                readBuffer));

    OptionalProperty eventState =
        readSimpleField(
            "eventState",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetEventState")),
                readBuffer));

    OptionalProperty reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetReliability")),
                readBuffer));

    OptionalProperty outOfService =
        readSimpleField(
            "outOfService",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty units =
        readSimpleField(
            "units",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetEngineeringUnits")),
                readBuffer));

    OptionalProperty priorityArray =
        readSimpleField(
            "priorityArray",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetPriorityArray")),
                readBuffer));

    OptionalProperty relinquishDefault =
        readSimpleField(
            "relinquishDefault",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty cOVIncrement =
        readSimpleField(
            "cOVIncrement",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty timeDelay =
        readSimpleField(
            "timeDelay",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty notificationClass =
        readSimpleField(
            "notificationClass",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty highLimit =
        readSimpleField(
            "highLimit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty lowLimit =
        readSimpleField(
            "lowLimit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty deadband =
        readSimpleField(
            "deadband",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty limitEnable =
        readSimpleField(
            "limitEnable",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetLimitEnable")),
                readBuffer));

    OptionalProperty eventEnable =
        readSimpleField(
            "eventEnable",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
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

    OptionalProperty eventDetectionEnable =
        readSimpleField(
            "eventDetectionEnable",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty eventAlgorithmInhibitRef =
        readSimpleField(
            "eventAlgorithmInhibitRef",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetObjectPropertyReference")),
                readBuffer));

    OptionalProperty eventAlgorithmInhibit =
        readSimpleField(
            "eventAlgorithmInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty timeDelayNormal =
        readSimpleField(
            "timeDelayNormal",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty reliabilityEvaluationInhibit =
        readSimpleField(
            "reliabilityEvaluationInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty minPresValue =
        readSimpleField(
            "minPresValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty maxPresValue =
        readSimpleField(
            "maxPresValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty resolution =
        readSimpleField(
            "resolution",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    ReadableProperty propertyList =
        readSimpleField(
            "propertyList",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetPropertyIdentifier")),
                readBuffer));

    OptionalProperty faultHighLimit =
        readSimpleField(
            "faultHighLimit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty faultLowLimit =
        readSimpleField(
            "faultLowLimit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty currentCommandPriority =
        readSimpleField(
            "currentCommandPriority",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetOptionalUnsigned")),
                readBuffer));

    OptionalProperty valueSource =
        readSimpleField(
            "valueSource",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetValueSource")),
                readBuffer));

    OptionalProperty valueSourceArray =
        readSimpleField(
            "valueSourceArray",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[16] of BACnetValueSource")),
                readBuffer));

    OptionalProperty lastCommandTime =
        readSimpleField(
            "lastCommandTime",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetTimeStamp")),
                readBuffer));

    OptionalProperty commandTimeArray =
        readSimpleField(
            "commandTimeArray",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[16] of BACnetTimeStamp")),
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

    readBuffer.closeContext("PositiveIntegerValue");
    // Create the instance
    PositiveIntegerValue _positiveIntegerValue;
    _positiveIntegerValue =
        new PositiveIntegerValue(
            objectIdentifier,
            objectName,
            objectType,
            description,
            presentValue,
            statusFlags,
            eventState,
            reliability,
            outOfService,
            units,
            priorityArray,
            relinquishDefault,
            cOVIncrement,
            timeDelay,
            notificationClass,
            highLimit,
            lowLimit,
            deadband,
            limitEnable,
            eventEnable,
            ackedTransitions,
            notifyType,
            eventTimeStamps,
            eventMessageTexts,
            eventMessageTextsConfig,
            eventDetectionEnable,
            eventAlgorithmInhibitRef,
            eventAlgorithmInhibit,
            timeDelayNormal,
            reliabilityEvaluationInhibit,
            minPresValue,
            maxPresValue,
            resolution,
            propertyList,
            faultHighLimit,
            faultLowLimit,
            currentCommandPriority,
            valueSource,
            valueSourceArray,
            lastCommandTime,
            commandTimeArray,
            tags,
            profileLocation,
            profileName);
    return _positiveIntegerValue;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PositiveIntegerValue)) {
      return false;
    }
    PositiveIntegerValue that = (PositiveIntegerValue) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getDescription() == that.getDescription())
        && (getPresentValue() == that.getPresentValue())
        && (getStatusFlags() == that.getStatusFlags())
        && (getEventState() == that.getEventState())
        && (getReliability() == that.getReliability())
        && (getOutOfService() == that.getOutOfService())
        && (getUnits() == that.getUnits())
        && (getPriorityArray() == that.getPriorityArray())
        && (getRelinquishDefault() == that.getRelinquishDefault())
        && (getCOVIncrement() == that.getCOVIncrement())
        && (getTimeDelay() == that.getTimeDelay())
        && (getNotificationClass() == that.getNotificationClass())
        && (getHighLimit() == that.getHighLimit())
        && (getLowLimit() == that.getLowLimit())
        && (getDeadband() == that.getDeadband())
        && (getLimitEnable() == that.getLimitEnable())
        && (getEventEnable() == that.getEventEnable())
        && (getAckedTransitions() == that.getAckedTransitions())
        && (getNotifyType() == that.getNotifyType())
        && (getEventTimeStamps() == that.getEventTimeStamps())
        && (getEventMessageTexts() == that.getEventMessageTexts())
        && (getEventMessageTextsConfig() == that.getEventMessageTextsConfig())
        && (getEventDetectionEnable() == that.getEventDetectionEnable())
        && (getEventAlgorithmInhibitRef() == that.getEventAlgorithmInhibitRef())
        && (getEventAlgorithmInhibit() == that.getEventAlgorithmInhibit())
        && (getTimeDelayNormal() == that.getTimeDelayNormal())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getMinPresValue() == that.getMinPresValue())
        && (getMaxPresValue() == that.getMaxPresValue())
        && (getResolution() == that.getResolution())
        && (getPropertyList() == that.getPropertyList())
        && (getFaultHighLimit() == that.getFaultHighLimit())
        && (getFaultLowLimit() == that.getFaultLowLimit())
        && (getCurrentCommandPriority() == that.getCurrentCommandPriority())
        && (getValueSource() == that.getValueSource())
        && (getValueSourceArray() == that.getValueSourceArray())
        && (getLastCommandTime() == that.getLastCommandTime())
        && (getCommandTimeArray() == that.getCommandTimeArray())
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
        getStatusFlags(),
        getEventState(),
        getReliability(),
        getOutOfService(),
        getUnits(),
        getPriorityArray(),
        getRelinquishDefault(),
        getCOVIncrement(),
        getTimeDelay(),
        getNotificationClass(),
        getHighLimit(),
        getLowLimit(),
        getDeadband(),
        getLimitEnable(),
        getEventEnable(),
        getAckedTransitions(),
        getNotifyType(),
        getEventTimeStamps(),
        getEventMessageTexts(),
        getEventMessageTextsConfig(),
        getEventDetectionEnable(),
        getEventAlgorithmInhibitRef(),
        getEventAlgorithmInhibit(),
        getTimeDelayNormal(),
        getReliabilityEvaluationInhibit(),
        getMinPresValue(),
        getMaxPresValue(),
        getResolution(),
        getPropertyList(),
        getFaultHighLimit(),
        getFaultLowLimit(),
        getCurrentCommandPriority(),
        getValueSource(),
        getValueSourceArray(),
        getLastCommandTime(),
        getCommandTimeArray(),
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