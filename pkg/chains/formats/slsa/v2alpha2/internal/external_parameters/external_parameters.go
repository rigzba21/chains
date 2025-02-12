/*
Copyright 2023 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externalparameters

import (
	"fmt"

	"github.com/tektoncd/chains/pkg/chains/objects"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

func buildConfigSource(provenance *v1beta1.Provenance) map[string]string {
	ref := ""
	for alg, hex := range provenance.RefSource.Digest {
		ref = fmt.Sprintf("%s:%s", alg, hex)
		break
	}
	buildConfigSource := map[string]string{
		"ref":        ref,
		"repository": provenance.RefSource.URI,
		"path":       provenance.RefSource.EntryPoint,
	}
	return buildConfigSource
}

// PipelineRun adds the pipeline run spec and provenance if available
func PipelineRun(pro *objects.PipelineRunObject) map[string]any {
	externalParams := make(map[string]any)

	if provenance := pro.GetRemoteProvenance(); provenance != nil {
		externalParams["buildConfigSource"] = buildConfigSource(provenance)
	}
	externalParams["runSpec"] = pro.Spec
	return externalParams
}

// TaskRun adds the task run spec and provenance if available
func TaskRun(tro *objects.TaskRunObject) map[string]any {
	externalParams := make(map[string]any)

	if provenance := tro.GetRemoteProvenance(); provenance != nil {
		externalParams["buildConfigSource"] = buildConfigSource(provenance)
	}
	externalParams["runSpec"] = tro.Spec
	return externalParams
}
