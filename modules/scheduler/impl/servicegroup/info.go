// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicegroup

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/scheduler/task"
)

func (s ServiceGroupImpl) Info(ctx context.Context, namespace string, name string) (apistructs.ServiceGroup, error) {
	sg := apistructs.ServiceGroup{}
	if err := s.js.Get(context.Background(), mkServiceGroupKey(namespace, name), &sg); err != nil {
		return sg, err
	}

	result, err := s.handleServiceGroup(ctx, &sg, task.TaskInspect)
	if err != nil {
		return sg, err
	}
	if result.Extra == nil {
		err = errors.Errorf("Cannot get servicegroup(%v/%v) info from TaskInspect", sg.Type, sg.ID)
		logrus.Error(err.Error())
		return sg, err
	}

	newsg := result.Extra.(*apistructs.ServiceGroup)
	return *newsg, nil
}