/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import "context"

type TraceMap struct {
	QueryCondition string `json:"query_condition"`
	TimeStart      int    `json:"time_start" binding:"required"`
	TimeEnd        int    `json:"time_end" binding:"required"`
	Debug          bool   `json:"debug"`
	Context        context.Context
	OrgID          string
}

type Debug struct {
	IP        string `json:"ip"`
	Sql       string `json:"sql"`
	SqlCH     string `json:"sql_CH"`
	QueryTime string `json:"query_time"`
	QueryUUID string `json:"query_uuid"`
	Error     string `json:"error"`
}

type TraceMapDebug struct {
	QuerierDebug []Debug `json:"querier_debug"`
	FormatTime   string  `json:"format_time"`
}

type TraceMapTree struct {
}