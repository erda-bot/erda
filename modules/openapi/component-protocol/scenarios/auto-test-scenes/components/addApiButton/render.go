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

package addApiButton

import (
	"context"

	"github.com/erda-project/erda/apistructs"
	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
)

type ComponentAction struct{}

func (ca *ComponentAction) Render(ctx context.Context, c *apistructs.Component, scenario apistructs.ComponentProtocolScenario, event apistructs.ComponentEvent, gs *apistructs.GlobalStateData) error {

	bdl := ctx.Value(protocol.GlobalInnerKeyCtxBundle.String()).(protocol.ContextBundle)

	switch event.Operation {
	case apistructs.ExecuteAddApiOperationKey:
		var req apistructs.AutotestSceneRequest
		req.Target = -1
		req.GroupID = -1
		req.Type = apistructs.StepTypeAPI

		var autotestSceneRequest apistructs.AutotestSceneRequest
		autotestSceneRequest.UserID = bdl.Identity.UserID
		autotestSceneRequest.ID = uint64(c.State["sceneId"].(float64))
		autotestSceneRequest.SceneID = uint64(c.State["sceneId"].(float64))
		result, err := bdl.Bdl.GetAutoTestScene(autotestSceneRequest)
		if err != nil {
			return err
		}

		req.SceneID = result.ID
		req.SpaceID = result.SpaceID
		req.UserID = bdl.Identity.UserID
		req.CreatorID = bdl.Identity.UserID
		req.UpdaterID = bdl.Identity.UserID
		stepID, err := bdl.Bdl.CreateAutoTestSceneStep(req)
		if err != nil {
			return err
		}
		c.State["createStepID"] = stepID
		c.State["showApiEditorDrawer"] = true
		c.State["isClick"] = true
	case apistructs.InitializeOperation, apistructs.RenderingOperation:
		c.Type = "Button"
		c.Props = map[string]interface{}{
			"text": "+ 接口",
			"menu": []interface{}{
				map[string]interface{}{
					"key": "addApi",
					"operations": map[string]interface{}{
						"click": map[string]interface{}{
							"key":    "addApi",
							"reload": true,
						},
					},
					"text": "表单添加",
				},
				map[string]interface{}{
					"key":  "addByCopyText",
					"text": "粘贴文本添加",
					"operations": map[string]interface{}{
						"click": map[string]interface{}{
							"key":    "addByCopyText",
							"reload": false,
							"command": map[string]interface{}{
								"key":    "set",
								"target": "addCopyApiFormModal",
								"state": map[string]interface{}{
									"visible": true,
								},
							},
						},
					},
				},
			},
		}
	}
	return nil
}

func RenderCreator() protocol.CompRender {
	return &ComponentAction{}
}