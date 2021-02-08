/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <stdio.h>
#include <string.h>
#include <time.h>
#include <plc4c/data.h>
#include <plc4c/utils/list.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_s7.h>
#include "data_item.h"

// Parse function.
plc4c_return_code plc4c_s7_read_write_data_item_parse(plc4c_spi_read_buffer* io, char* dataProtocolId, int32_t stringLength, plc4c_data** data_item) {
    uint16_t startPos = plc4c_spi_read_get_pos(io);
    uint16_t curPos;
    plc4c_return_code _res = OK;

    if(strcmp(dataProtocolId, "IEC61131_BOOL") == 0) { /* BOOL */

                // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
                {
                    unsigned int _reserved = 0;
                    _res = plc4c_spi_read_unsigned_byte(io, 7, (uint8_t*) &_reserved);
                    if(_res != OK) {
                        return _res;
                    }
                    if(_reserved != 0x00) {
                      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x00, _reserved);
                    }
                }

                // Simple Field (value)
                bool value = false;
                _res = plc4c_spi_read_bit(io, (bool*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_bool_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_BYTE") == 0) { /* List */

        // Array field (value)
        // Count array
        plc4c_list* value;
        plc4c_utils_list_create(&value);
        int itemCount = (int) 8;
        for(int curItem = 0; curItem < itemCount; curItem++) {
            bool* _val = malloc(sizeof(bool) * 1);
            _res = plc4c_spi_read_bit(io, (bool*) _val);
            if(_res != OK) {
                return _res;
            }
            plc4c_data* _item = plc4c_data_create_bool_data(*_val);
            plc4c_utils_list_insert_head_value(value, _item);
        }
        *data_item = plc4c_data_create_list_data(*value);

    } else     if(strcmp(dataProtocolId, "IEC61131_WORD") == 0) { /* List */

        // Array field (value)
        // Count array
        plc4c_list* value;
        plc4c_utils_list_create(&value);
        int itemCount = (int) 16;
        for(int curItem = 0; curItem < itemCount; curItem++) {
            bool* _val = malloc(sizeof(bool) * 1);
            _res = plc4c_spi_read_bit(io, (bool*) _val);
            if(_res != OK) {
                return _res;
            }
            plc4c_data* _item = plc4c_data_create_bool_data(*_val);
            plc4c_utils_list_insert_head_value(value, _item);
        }
        *data_item = plc4c_data_create_list_data(*value);

    } else     if(strcmp(dataProtocolId, "IEC61131_DWORD") == 0) { /* List */

        // Array field (value)
        // Count array
        plc4c_list* value;
        plc4c_utils_list_create(&value);
        int itemCount = (int) 32;
        for(int curItem = 0; curItem < itemCount; curItem++) {
            bool* _val = malloc(sizeof(bool) * 1);
            _res = plc4c_spi_read_bit(io, (bool*) _val);
            if(_res != OK) {
                return _res;
            }
            plc4c_data* _item = plc4c_data_create_bool_data(*_val);
            plc4c_utils_list_insert_head_value(value, _item);
        }
        *data_item = plc4c_data_create_list_data(*value);

    } else     if(strcmp(dataProtocolId, "IEC61131_LWORD") == 0) { /* List */

        // Array field (value)
        // Count array
        plc4c_list* value;
        plc4c_utils_list_create(&value);
        int itemCount = (int) 64;
        for(int curItem = 0; curItem < itemCount; curItem++) {
            bool* _val = malloc(sizeof(bool) * 1);
            _res = plc4c_spi_read_bit(io, (bool*) _val);
            if(_res != OK) {
                return _res;
            }
            plc4c_data* _item = plc4c_data_create_bool_data(*_val);
            plc4c_utils_list_insert_head_value(value, _item);
        }
        *data_item = plc4c_data_create_list_data(*value);

    } else     if(strcmp(dataProtocolId, "IEC61131_SINT") == 0) { /* SINT */

                // Simple Field (value)
                int8_t value = 0;
                _res = plc4c_spi_read_signed_byte(io, 8, (int8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int8_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_USINT") == 0) { /* USINT */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_INT") == 0) { /* INT */

                // Simple Field (value)
                int16_t value = 0;
                _res = plc4c_spi_read_signed_short(io, 16, (int16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int16_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_UINT") == 0) { /* UINT */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint16_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_DINT") == 0) { /* DINT */

                // Simple Field (value)
                int32_t value = 0;
                _res = plc4c_spi_read_signed_int(io, 32, (int32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int32_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_UDINT") == 0) { /* UDINT */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(io, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint32_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_LINT") == 0) { /* LINT */

                // Simple Field (value)
                int64_t value = 0;
                _res = plc4c_spi_read_signed_long(io, 64, (int64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int64_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_ULINT") == 0) { /* ULINT */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(io, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint64_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_REAL") == 0) { /* REAL */

                // Simple Field (value)
                float value = 0.0f;
                _res = plc4c_spi_read_float(io, 32, (float*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_float_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_LREAL") == 0) { /* LREAL */

                // Simple Field (value)
                double value = 0.0f;
                _res = plc4c_spi_read_double(io, 64, (double*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_double_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_CHAR") == 0) { /* CHAR */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_char(io, "UTF-8"));

                *data_item = plc4c_data_create_char_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_WCHAR") == 0) { /* CHAR */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_char(io, "UTF-16"));

                *data_item = plc4c_data_create_char_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_STRING") == 0) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_string(io, stringLength, "UTF-8"));

                *data_item = plc4c_data_create_string_data(stringLength, value);

    } else     if(strcmp(dataProtocolId, "IEC61131_WSTRING") == 0) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_string(io, stringLength, "UTF-16"));

                *data_item = plc4c_data_create_string_data(stringLength, value);

    } else     if(strcmp(dataProtocolId, "IEC61131_TIME") == 0) { /* TIME */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(io, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint32_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_LTIME") == 0) { /* LTIME */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(io, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint64_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_DATE") == 0) { /* DATE */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint16_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_TIME_OF_DAY") == 0) { /* TIME_OF_DAY */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(io, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint32_t_data(value);

    } else     if(strcmp(dataProtocolId, "IEC61131_DATE_AND_TIME") == 0) { /* DATE_AND_TIME */

                // Simple Field (year)
                uint16_t year = 0;
                _res = plc4c_spi_read_unsigned_short(io, 16, (uint16_t*) &year);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint16_t_data(year);


                // Simple Field (month)
                uint8_t month = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &month);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(month);


                // Simple Field (day)
                uint8_t day = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &day);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(day);


                // Simple Field (dayOfWeek)
                uint8_t dayOfWeek = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &dayOfWeek);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(dayOfWeek);


                // Simple Field (hour)
                uint8_t hour = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &hour);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(hour);


                // Simple Field (minutes)
                uint8_t minutes = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &minutes);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(minutes);


                // Simple Field (seconds)
                uint8_t seconds = 0;
                _res = plc4c_spi_read_unsigned_byte(io, 8, (uint8_t*) &seconds);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint8_t_data(seconds);


                // Simple Field (nanos)
                uint32_t nanos = 0;
                _res = plc4c_spi_read_unsigned_int(io, 32, (uint32_t*) &nanos);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint32_t_data(nanos);

    }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_data_item_serialize(plc4c_spi_write_buffer* io, plc4c_data** data_item) {
  plc4c_return_code _res = OK;

  return OK;
}

uint16_t plc4c_s7_read_write_data_item_length_in_bytes(plc4c_data* data_item) {
  return plc4c_s7_read_write_data_item_length_in_bits(data_item) / 8;
}

uint16_t plc4c_s7_read_write_data_item_length_in_bits(plc4c_data* data_item) {
  uint16_t lengthInBits = 0;

  return lengthInBits;
}

