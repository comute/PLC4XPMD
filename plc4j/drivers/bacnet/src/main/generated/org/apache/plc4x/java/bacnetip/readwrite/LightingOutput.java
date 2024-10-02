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

public class LightingOutput implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final WritableProperty presentValue;
  protected final ReadableProperty trackingValue;
  protected final WritableProperty lightingCommand;
  protected final ReadableProperty inProgress;
  protected final OptionalProperty description;
  protected final ReadableProperty statusFlags;
  protected final OptionalProperty reliability;
  protected final ReadableProperty outOfService;
  protected final ReadableProperty blinkWarnEnable;
  protected final ReadableProperty egressTime;
  protected final ReadableProperty egressActive;
  protected final ReadableProperty defaultFadeTime;
  protected final ReadableProperty defaultRampRate;
  protected final ReadableProperty defaultStepIncrement;
  protected final OptionalProperty transition;
  protected final OptionalProperty feedbackValue;
  protected final ReadableProperty priorityArray;
  protected final ReadableProperty relinquishDefault;
  protected final OptionalProperty power;
  protected final OptionalProperty instantaneousPower;
  protected final OptionalProperty minActualValue;
  protected final OptionalProperty maxActualValue;
  protected final ReadableProperty lightingCommandDefaultPriority;
  protected final OptionalProperty cOVIncrement;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final ReadableProperty propertyList;
  protected final ReadableProperty currentCommandPriority;
  protected final OptionalProperty valueSource;
  protected final OptionalProperty valueSourceArray;
  protected final OptionalProperty lastCommandTime;
  protected final OptionalProperty commandTimeArray;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public LightingOutput(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      WritableProperty presentValue,
      ReadableProperty trackingValue,
      WritableProperty lightingCommand,
      ReadableProperty inProgress,
      OptionalProperty description,
      ReadableProperty statusFlags,
      OptionalProperty reliability,
      ReadableProperty outOfService,
      ReadableProperty blinkWarnEnable,
      ReadableProperty egressTime,
      ReadableProperty egressActive,
      ReadableProperty defaultFadeTime,
      ReadableProperty defaultRampRate,
      ReadableProperty defaultStepIncrement,
      OptionalProperty transition,
      OptionalProperty feedbackValue,
      ReadableProperty priorityArray,
      ReadableProperty relinquishDefault,
      OptionalProperty power,
      OptionalProperty instantaneousPower,
      OptionalProperty minActualValue,
      OptionalProperty maxActualValue,
      ReadableProperty lightingCommandDefaultPriority,
      OptionalProperty cOVIncrement,
      OptionalProperty reliabilityEvaluationInhibit,
      ReadableProperty propertyList,
      ReadableProperty currentCommandPriority,
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
    this.presentValue = presentValue;
    this.trackingValue = trackingValue;
    this.lightingCommand = lightingCommand;
    this.inProgress = inProgress;
    this.description = description;
    this.statusFlags = statusFlags;
    this.reliability = reliability;
    this.outOfService = outOfService;
    this.blinkWarnEnable = blinkWarnEnable;
    this.egressTime = egressTime;
    this.egressActive = egressActive;
    this.defaultFadeTime = defaultFadeTime;
    this.defaultRampRate = defaultRampRate;
    this.defaultStepIncrement = defaultStepIncrement;
    this.transition = transition;
    this.feedbackValue = feedbackValue;
    this.priorityArray = priorityArray;
    this.relinquishDefault = relinquishDefault;
    this.power = power;
    this.instantaneousPower = instantaneousPower;
    this.minActualValue = minActualValue;
    this.maxActualValue = maxActualValue;
    this.lightingCommandDefaultPriority = lightingCommandDefaultPriority;
    this.cOVIncrement = cOVIncrement;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.propertyList = propertyList;
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

  public WritableProperty getPresentValue() {
    return presentValue;
  }

  public ReadableProperty getTrackingValue() {
    return trackingValue;
  }

  public WritableProperty getLightingCommand() {
    return lightingCommand;
  }

  public ReadableProperty getInProgress() {
    return inProgress;
  }

  public OptionalProperty getDescription() {
    return description;
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

  public ReadableProperty getBlinkWarnEnable() {
    return blinkWarnEnable;
  }

  public ReadableProperty getEgressTime() {
    return egressTime;
  }

  public ReadableProperty getEgressActive() {
    return egressActive;
  }

  public ReadableProperty getDefaultFadeTime() {
    return defaultFadeTime;
  }

  public ReadableProperty getDefaultRampRate() {
    return defaultRampRate;
  }

  public ReadableProperty getDefaultStepIncrement() {
    return defaultStepIncrement;
  }

  public OptionalProperty getTransition() {
    return transition;
  }

  public OptionalProperty getFeedbackValue() {
    return feedbackValue;
  }

  public ReadableProperty getPriorityArray() {
    return priorityArray;
  }

  public ReadableProperty getRelinquishDefault() {
    return relinquishDefault;
  }

  public OptionalProperty getPower() {
    return power;
  }

  public OptionalProperty getInstantaneousPower() {
    return instantaneousPower;
  }

  public OptionalProperty getMinActualValue() {
    return minActualValue;
  }

  public OptionalProperty getMaxActualValue() {
    return maxActualValue;
  }

  public ReadableProperty getLightingCommandDefaultPriority() {
    return lightingCommandDefaultPriority;
  }

  public OptionalProperty getCOVIncrement() {
    return cOVIncrement;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public ReadableProperty getCurrentCommandPriority() {
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
    writeBuffer.pushContext("LightingOutput");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (objectName)
    writeSimpleField("objectName", objectName, writeComplex(writeBuffer));

    // Simple Field (objectType)
    writeSimpleField("objectType", objectType, writeComplex(writeBuffer));

    // Simple Field (presentValue)
    writeSimpleField("presentValue", presentValue, writeComplex(writeBuffer));

    // Simple Field (trackingValue)
    writeSimpleField("trackingValue", trackingValue, writeComplex(writeBuffer));

    // Simple Field (lightingCommand)
    writeSimpleField("lightingCommand", lightingCommand, writeComplex(writeBuffer));

    // Simple Field (inProgress)
    writeSimpleField("inProgress", inProgress, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (outOfService)
    writeSimpleField("outOfService", outOfService, writeComplex(writeBuffer));

    // Simple Field (blinkWarnEnable)
    writeSimpleField("blinkWarnEnable", blinkWarnEnable, writeComplex(writeBuffer));

    // Simple Field (egressTime)
    writeSimpleField("egressTime", egressTime, writeComplex(writeBuffer));

    // Simple Field (egressActive)
    writeSimpleField("egressActive", egressActive, writeComplex(writeBuffer));

    // Simple Field (defaultFadeTime)
    writeSimpleField("defaultFadeTime", defaultFadeTime, writeComplex(writeBuffer));

    // Simple Field (defaultRampRate)
    writeSimpleField("defaultRampRate", defaultRampRate, writeComplex(writeBuffer));

    // Simple Field (defaultStepIncrement)
    writeSimpleField("defaultStepIncrement", defaultStepIncrement, writeComplex(writeBuffer));

    // Simple Field (transition)
    writeSimpleField("transition", transition, writeComplex(writeBuffer));

    // Simple Field (feedbackValue)
    writeSimpleField("feedbackValue", feedbackValue, writeComplex(writeBuffer));

    // Simple Field (priorityArray)
    writeSimpleField("priorityArray", priorityArray, writeComplex(writeBuffer));

    // Simple Field (relinquishDefault)
    writeSimpleField("relinquishDefault", relinquishDefault, writeComplex(writeBuffer));

    // Simple Field (power)
    writeSimpleField("power", power, writeComplex(writeBuffer));

    // Simple Field (instantaneousPower)
    writeSimpleField("instantaneousPower", instantaneousPower, writeComplex(writeBuffer));

    // Simple Field (minActualValue)
    writeSimpleField("minActualValue", minActualValue, writeComplex(writeBuffer));

    // Simple Field (maxActualValue)
    writeSimpleField("maxActualValue", maxActualValue, writeComplex(writeBuffer));

    // Simple Field (lightingCommandDefaultPriority)
    writeSimpleField(
        "lightingCommandDefaultPriority",
        lightingCommandDefaultPriority,
        writeComplex(writeBuffer));

    // Simple Field (cOVIncrement)
    writeSimpleField("cOVIncrement", cOVIncrement, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

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

    writeBuffer.popContext("LightingOutput");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LightingOutput _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (objectName)
    lengthInBits += objectName.getLengthInBits();

    // Simple field (objectType)
    lengthInBits += objectType.getLengthInBits();

    // Simple field (presentValue)
    lengthInBits += presentValue.getLengthInBits();

    // Simple field (trackingValue)
    lengthInBits += trackingValue.getLengthInBits();

    // Simple field (lightingCommand)
    lengthInBits += lightingCommand.getLengthInBits();

    // Simple field (inProgress)
    lengthInBits += inProgress.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (outOfService)
    lengthInBits += outOfService.getLengthInBits();

    // Simple field (blinkWarnEnable)
    lengthInBits += blinkWarnEnable.getLengthInBits();

    // Simple field (egressTime)
    lengthInBits += egressTime.getLengthInBits();

    // Simple field (egressActive)
    lengthInBits += egressActive.getLengthInBits();

    // Simple field (defaultFadeTime)
    lengthInBits += defaultFadeTime.getLengthInBits();

    // Simple field (defaultRampRate)
    lengthInBits += defaultRampRate.getLengthInBits();

    // Simple field (defaultStepIncrement)
    lengthInBits += defaultStepIncrement.getLengthInBits();

    // Simple field (transition)
    lengthInBits += transition.getLengthInBits();

    // Simple field (feedbackValue)
    lengthInBits += feedbackValue.getLengthInBits();

    // Simple field (priorityArray)
    lengthInBits += priorityArray.getLengthInBits();

    // Simple field (relinquishDefault)
    lengthInBits += relinquishDefault.getLengthInBits();

    // Simple field (power)
    lengthInBits += power.getLengthInBits();

    // Simple field (instantaneousPower)
    lengthInBits += instantaneousPower.getLengthInBits();

    // Simple field (minActualValue)
    lengthInBits += minActualValue.getLengthInBits();

    // Simple field (maxActualValue)
    lengthInBits += maxActualValue.getLengthInBits();

    // Simple field (lightingCommandDefaultPriority)
    lengthInBits += lightingCommandDefaultPriority.getLengthInBits();

    // Simple field (cOVIncrement)
    lengthInBits += cOVIncrement.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

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

  public static LightingOutput staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LightingOutput");
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

    WritableProperty presentValue =
        readSimpleField(
            "presentValue",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    ReadableProperty trackingValue =
        readSimpleField(
            "trackingValue",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    WritableProperty lightingCommand =
        readSimpleField(
            "lightingCommand",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("BACnetLightingCommand")),
                readBuffer));

    ReadableProperty inProgress =
        readSimpleField(
            "inProgress",
            readComplex(
                () ->
                    ReadableProperty.staticParse(readBuffer, (String) ("BACnetLightingInProgress")),
                readBuffer));

    OptionalProperty description =
        readSimpleField(
            "description",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
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

    ReadableProperty blinkWarnEnable =
        readSimpleField(
            "blinkWarnEnable",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty egressTime =
        readSimpleField(
            "egressTime",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    ReadableProperty egressActive =
        readSimpleField(
            "egressActive",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty defaultFadeTime =
        readSimpleField(
            "defaultFadeTime",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    ReadableProperty defaultRampRate =
        readSimpleField(
            "defaultRampRate",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    ReadableProperty defaultStepIncrement =
        readSimpleField(
            "defaultStepIncrement",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    OptionalProperty transition =
        readSimpleField(
            "transition",
            readComplex(
                () ->
                    OptionalProperty.staticParse(readBuffer, (String) ("BACnetLightingTransition")),
                readBuffer));

    OptionalProperty feedbackValue =
        readSimpleField(
            "feedbackValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    ReadableProperty priorityArray =
        readSimpleField(
            "priorityArray",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetPriorityArray")),
                readBuffer));

    ReadableProperty relinquishDefault =
        readSimpleField(
            "relinquishDefault",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    OptionalProperty power =
        readSimpleField(
            "power",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    OptionalProperty instantaneousPower =
        readSimpleField(
            "instantaneousPower",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    OptionalProperty minActualValue =
        readSimpleField(
            "minActualValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    OptionalProperty maxActualValue =
        readSimpleField(
            "maxActualValue",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

    ReadableProperty lightingCommandDefaultPriority =
        readSimpleField(
            "lightingCommandDefaultPriority",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty cOVIncrement =
        readSimpleField(
            "cOVIncrement",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("REAL")), readBuffer));

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

    ReadableProperty currentCommandPriority =
        readSimpleField(
            "currentCommandPriority",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetOptionalUnsigned")),
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

    readBuffer.closeContext("LightingOutput");
    // Create the instance
    LightingOutput _lightingOutput;
    _lightingOutput =
        new LightingOutput(
            objectIdentifier,
            objectName,
            objectType,
            presentValue,
            trackingValue,
            lightingCommand,
            inProgress,
            description,
            statusFlags,
            reliability,
            outOfService,
            blinkWarnEnable,
            egressTime,
            egressActive,
            defaultFadeTime,
            defaultRampRate,
            defaultStepIncrement,
            transition,
            feedbackValue,
            priorityArray,
            relinquishDefault,
            power,
            instantaneousPower,
            minActualValue,
            maxActualValue,
            lightingCommandDefaultPriority,
            cOVIncrement,
            reliabilityEvaluationInhibit,
            propertyList,
            currentCommandPriority,
            valueSource,
            valueSourceArray,
            lastCommandTime,
            commandTimeArray,
            tags,
            profileLocation,
            profileName);
    return _lightingOutput;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LightingOutput)) {
      return false;
    }
    LightingOutput that = (LightingOutput) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getPresentValue() == that.getPresentValue())
        && (getTrackingValue() == that.getTrackingValue())
        && (getLightingCommand() == that.getLightingCommand())
        && (getInProgress() == that.getInProgress())
        && (getDescription() == that.getDescription())
        && (getStatusFlags() == that.getStatusFlags())
        && (getReliability() == that.getReliability())
        && (getOutOfService() == that.getOutOfService())
        && (getBlinkWarnEnable() == that.getBlinkWarnEnable())
        && (getEgressTime() == that.getEgressTime())
        && (getEgressActive() == that.getEgressActive())
        && (getDefaultFadeTime() == that.getDefaultFadeTime())
        && (getDefaultRampRate() == that.getDefaultRampRate())
        && (getDefaultStepIncrement() == that.getDefaultStepIncrement())
        && (getTransition() == that.getTransition())
        && (getFeedbackValue() == that.getFeedbackValue())
        && (getPriorityArray() == that.getPriorityArray())
        && (getRelinquishDefault() == that.getRelinquishDefault())
        && (getPower() == that.getPower())
        && (getInstantaneousPower() == that.getInstantaneousPower())
        && (getMinActualValue() == that.getMinActualValue())
        && (getMaxActualValue() == that.getMaxActualValue())
        && (getLightingCommandDefaultPriority() == that.getLightingCommandDefaultPriority())
        && (getCOVIncrement() == that.getCOVIncrement())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getPropertyList() == that.getPropertyList())
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
        getPresentValue(),
        getTrackingValue(),
        getLightingCommand(),
        getInProgress(),
        getDescription(),
        getStatusFlags(),
        getReliability(),
        getOutOfService(),
        getBlinkWarnEnable(),
        getEgressTime(),
        getEgressActive(),
        getDefaultFadeTime(),
        getDefaultRampRate(),
        getDefaultStepIncrement(),
        getTransition(),
        getFeedbackValue(),
        getPriorityArray(),
        getRelinquishDefault(),
        getPower(),
        getInstantaneousPower(),
        getMinActualValue(),
        getMaxActualValue(),
        getLightingCommandDefaultPriority(),
        getCOVIncrement(),
        getReliabilityEvaluationInhibit(),
        getPropertyList(),
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