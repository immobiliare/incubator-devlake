/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package services

import (
	"github.com/apache/incubator-devlake/core/errors"
	"time"
)

type PipelineNotificationParam struct {
	ProjectName string // can be an empty string, if pipeline is created and triggered by API
	PipelineID  uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BeganAt     *time.Time
	FinishedAt  *time.Time
	Status      string
}

type PipelineNotificationService interface {
	PipelineStatusChanged(params PipelineNotificationParam) errors.Error
}

var customPipelineNotificationService PipelineNotificationService

func GetPipelineNotificationService() PipelineNotificationService {
	if customPipelineNotificationService != nil {
		return customPipelineNotificationService
	}
	if defaultNotificationService != nil {
		return defaultNotificationService
	}
	return nil
}
