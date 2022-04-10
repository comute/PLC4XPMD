/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#ifndef PLC4C_MQTT_READ_WRITE_MQT_T__CONTROL_PACKET_H_
#define PLC4C_MQTT_READ_WRITE_MQT_T__CONTROL_PACKET_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/driver_mqtt_static.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>
#include "mqt_t__reason_code.h"
#include "filter.h"
#include "mqt_t__string.h"
#include "mqt_t__qos.h"
#include "mqt_t__property.h"
#include "mqt_t__control_packet_type.h"

// Code generated by code-generation. DO NOT EDIT.

#ifdef __cplusplus
extern "C" {
#endif


// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_mqtt_read_write_mqt_t__control_packet_discriminator {
  enum plc4c_mqtt_read_write_mqt_t__control_packet_type packetType;
};
typedef struct plc4c_mqtt_read_write_mqt_t__control_packet_discriminator plc4c_mqtt_read_write_mqt_t__control_packet_discriminator;

// Enum assigning each sub-type an individual id.
enum plc4c_mqtt_read_write_mqt_t__control_packet_type {
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__connect = 0,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__connack = 1,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__publish = 2,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__puback = 3,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__pubrec = 4,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__pubrel = 5,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__pubcomp = 6,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__subscribe = 7,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__suback = 8,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__unsubscribe = 9,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__unsuback = 10,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__pingreq = 11,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__pingresp = 12,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__disconnect = 13,
  plc4c_mqtt_read_write_mqt_t__control_packet_type_plc4c_mqtt_read_write_mqt_t__control_packet__auth = 14};
typedef enum plc4c_mqtt_read_write_mqt_t__control_packet_type plc4c_mqtt_read_write_mqt_t__control_packet_type;

// Function to get the discriminator values for a given type.
plc4c_mqtt_read_write_mqt_t__control_packet_discriminator plc4c_mqtt_read_write_mqt_t__control_packet_get_discriminator(plc4c_mqtt_read_write_mqt_t__control_packet_type type);

struct plc4c_mqtt_read_write_mqt_t__control_packet {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_mqtt_read_write_mqt_t__control_packet_type _type;
  /* Properties */
  union {
    struct { /* MQTT_ControlPacket_CONNECT */
      uint8_t mqt_t__control_packet__connect_remaining_length;
      plc4c_mqtt_read_write_mqt_t__string* mqt_t__control_packet__connect_protocol_name;
      uint8_t mqt_t__control_packet__connect_protocol_version;
      bool mqt_t__control_packet__connect_user_name_flag_set : 1;
      bool mqt_t__control_packet__connect_password_flag_set : 1;
      bool mqt_t__control_packet__connect_will_retain_flag_set : 1;
      uint8_t mqt_t__control_packet__connect_will_qos_level : 2;
      bool mqt_t__control_packet__connect_will_flag_set : 1;
      bool mqt_t__control_packet__connect_clean_start_flag_set : 1;
      uint16_t mqt_t__control_packet__connect_keep_alive;
      uint32_t mqt_t__control_packet__connect_property_length;
      plc4c_list* mqt_t__control_packet__connect_properties;
      plc4c_mqtt_read_write_mqt_t__string* mqt_t__control_packet__connect_client_id;
      plc4c_mqtt_read_write_mqt_t__string* mqt_t__control_packet__connect_username;
      plc4c_mqtt_read_write_mqt_t__string* mqt_t__control_packet__connect_password;
    };
    struct { /* MQTT_ControlPacket_CONNACK */
      uint8_t mqt_t__control_packet__connack_remaining_length;
      bool mqt_t__control_packet__connack_session_present_flag_set : 1;
      plc4c_mqtt_read_write_mqt_t__reason_code mqt_t__control_packet__connack_reason_code;
      uint32_t* mqt_t__control_packet__connack_property_length;
      plc4c_list* mqt_t__control_packet__connack_properties;
    };
    struct { /* MQTT_ControlPacket_PUBLISH */
      bool mqt_t__control_packet__publish_dup : 1;
      plc4c_mqtt_read_write_mqt_t__qos mqt_t__control_packet__publish_qos;
      bool mqt_t__control_packet__publish_retain : 1;
      uint8_t mqt_t__control_packet__publish_remaining_length;
      plc4c_mqtt_read_write_mqt_t__string* mqt_t__control_packet__publish_topic_name;
      uint16_t* mqt_t__control_packet__publish_packet_identifier;
      uint32_t* mqt_t__control_packet__publish_property_length;
      plc4c_list* mqt_t__control_packet__publish_properties;
      plc4c_list* mqt_t__control_packet__publish_payload;
    };
    struct { /* MQTT_ControlPacket_PUBACK */
      uint8_t mqt_t__control_packet__puback_remaining_length;
      uint16_t mqt_t__control_packet__puback_packet_identifier;
      plc4c_mqtt_read_write_mqt_t__reason_code* mqt_t__control_packet__puback_reason_code;
      uint32_t* mqt_t__control_packet__puback_property_length;
      plc4c_list* mqt_t__control_packet__puback_properties;
    };
    struct { /* MQTT_ControlPacket_PUBREC */
      uint8_t mqt_t__control_packet__pubrec_remaining_length;
      uint16_t mqt_t__control_packet__pubrec_packet_identifier;
      plc4c_mqtt_read_write_mqt_t__reason_code* mqt_t__control_packet__pubrec_reason_code;
      uint32_t* mqt_t__control_packet__pubrec_property_length;
      plc4c_list* mqt_t__control_packet__pubrec_properties;
    };
    struct { /* MQTT_ControlPacket_PUBREL */
      uint8_t mqt_t__control_packet__pubrel_remaining_length;
      uint16_t mqt_t__control_packet__pubrel_packet_identifier;
      plc4c_mqtt_read_write_mqt_t__reason_code* mqt_t__control_packet__pubrel_reason_code;
      uint32_t* mqt_t__control_packet__pubrel_property_length;
      plc4c_list* mqt_t__control_packet__pubrel_properties;
    };
    struct { /* MQTT_ControlPacket_PUBCOMP */
      uint8_t mqt_t__control_packet__pubcomp_remaining_length;
      uint16_t mqt_t__control_packet__pubcomp_packet_identifier;
      plc4c_mqtt_read_write_mqt_t__reason_code* mqt_t__control_packet__pubcomp_reason_code;
      uint32_t* mqt_t__control_packet__pubcomp_property_length;
      plc4c_list* mqt_t__control_packet__pubcomp_properties;
    };
    struct { /* MQTT_ControlPacket_SUBSCRIBE */
      uint8_t mqt_t__control_packet__subscribe_remaining_length;
      uint16_t mqt_t__control_packet__subscribe_packet_identifier;
      uint32_t* mqt_t__control_packet__subscribe_property_length;
      plc4c_list* mqt_t__control_packet__subscribe_properties;
      plc4c_list* mqt_t__control_packet__subscribe_filters;
    };
    struct { /* MQTT_ControlPacket_SUBACK */
      uint8_t mqt_t__control_packet__suback_remaining_length;
      uint16_t mqt_t__control_packet__suback_packet_identifier;
      uint32_t* mqt_t__control_packet__suback_property_length;
      plc4c_list* mqt_t__control_packet__suback_properties;
      plc4c_list* mqt_t__control_packet__suback_results;
    };
    struct { /* MQTT_ControlPacket_UNSUBSCRIBE */
      uint8_t mqt_t__control_packet__unsubscribe_remaining_length;
      uint16_t mqt_t__control_packet__unsubscribe_packet_identifier;
      uint32_t* mqt_t__control_packet__unsubscribe_property_length;
      plc4c_list* mqt_t__control_packet__unsubscribe_properties;
      plc4c_list* mqt_t__control_packet__unsubscribe_filters;
    };
    struct { /* MQTT_ControlPacket_UNSUBACK */
      uint8_t mqt_t__control_packet__unsuback_remaining_length;
      uint16_t mqt_t__control_packet__unsuback_packet_identifier;
      uint32_t* mqt_t__control_packet__unsuback_property_length;
      plc4c_list* mqt_t__control_packet__unsuback_properties;
      plc4c_list* mqt_t__control_packet__unsuback_results;
    };
    struct { /* MQTT_ControlPacket_PINGREQ */
      uint8_t mqt_t__control_packet__pingreq_remaining_length;
    };
    struct { /* MQTT_ControlPacket_PINGRESP */
      uint8_t mqt_t__control_packet__pingresp_remaining_length;
    };
    struct { /* MQTT_ControlPacket_DISCONNECT */
      uint8_t mqt_t__control_packet__disconnect_remaining_length;
      plc4c_mqtt_read_write_mqt_t__reason_code mqt_t__control_packet__disconnect_reason;
    };
    struct { /* MQTT_ControlPacket_AUTH */
      uint8_t mqt_t__control_packet__auth_remaining_length;
      plc4c_mqtt_read_write_mqt_t__reason_code mqt_t__control_packet__auth_reason;
      uint32_t* mqt_t__control_packet__auth_property_length;
      plc4c_list* mqt_t__control_packet__auth_properties;
    };
  };
};
typedef struct plc4c_mqtt_read_write_mqt_t__control_packet plc4c_mqtt_read_write_mqt_t__control_packet;

// Create an empty NULL-struct
plc4c_mqtt_read_write_mqt_t__control_packet plc4c_mqtt_read_write_mqt_t__control_packet_null();

plc4c_return_code plc4c_mqtt_read_write_mqt_t__control_packet_parse(plc4c_spi_read_buffer* readBuffer, plc4c_mqtt_read_write_mqt_t__control_packet** message);

plc4c_return_code plc4c_mqtt_read_write_mqt_t__control_packet_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_mqtt_read_write_mqt_t__control_packet* message);

uint16_t plc4c_mqtt_read_write_mqt_t__control_packet_length_in_bytes(plc4c_mqtt_read_write_mqt_t__control_packet* message);

uint16_t plc4c_mqtt_read_write_mqt_t__control_packet_length_in_bits(plc4c_mqtt_read_write_mqt_t__control_packet* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_MQTT_READ_WRITE_MQT_T__CONTROL_PACKET_H_
