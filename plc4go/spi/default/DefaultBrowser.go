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

package _default

import (
	"context"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/spi"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/rs/zerolog"
	"runtime/debug"
)

// DefaultBrowserRequirements adds required methods to Browser that are needed when using DefaultBrowser
type DefaultBrowserRequirements interface {
	BrowseQuery(ctx context.Context, interceptor func(result apiModel.PlcBrowseItem) bool, queryName string, query apiModel.PlcQuery) (apiModel.PlcResponseCode, []apiModel.PlcBrowseItem)
}

type DefaultBrowser interface {
	spi.PlcBrowser
}

func NewDefaultBrowser(defaultBrowserRequirements DefaultBrowserRequirements, _options ...options.WithOption) DefaultBrowser {
	customLogger := options.ExtractCustomLoggerOrDefaultToGlobal(_options...)
	return &defaultBrowser{
		DefaultBrowserRequirements: defaultBrowserRequirements,

		log: customLogger,
	}
}

///////////////////////////////////////
///////////////////////////////////////
//
// Internal section
//

type defaultBrowser struct {
	DefaultBrowserRequirements

	log zerolog.Logger
}

//
// Internal section
//
///////////////////////////////////////
///////////////////////////////////////

func (m *defaultBrowser) Browse(ctx context.Context, browseRequest apiModel.PlcBrowseRequest) <-chan apiModel.PlcBrowseRequestResult {
	return m.BrowseWithInterceptor(ctx, browseRequest, func(result apiModel.PlcBrowseItem) bool {
		return true
	})
}

func (m *defaultBrowser) BrowseWithInterceptor(ctx context.Context, browseRequest apiModel.PlcBrowseRequest, interceptor func(result apiModel.PlcBrowseItem) bool) <-chan apiModel.PlcBrowseRequestResult {
	result := make(chan apiModel.PlcBrowseRequestResult, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				m.log.Error().
					Str("stack", string(debug.Stack())).
					Interface("err", err).
					Msg("panic-ed")
			}
		}()
		responseCodes := map[string]apiModel.PlcResponseCode{}
		results := map[string][]apiModel.PlcBrowseItem{}
		for _, queryName := range browseRequest.GetQueryNames() {
			query := browseRequest.GetQuery(queryName)
			responseCodes[queryName], results[queryName] = m.BrowseQuery(ctx, interceptor, queryName, query)
		}
		browseResponse := spiModel.NewDefaultPlcBrowseResponse(browseRequest, results, responseCodes)
		result <- spiModel.NewDefaultPlcBrowseRequestResult(
			browseRequest,
			browseResponse,
			nil,
		)
	}()
	return result
}
