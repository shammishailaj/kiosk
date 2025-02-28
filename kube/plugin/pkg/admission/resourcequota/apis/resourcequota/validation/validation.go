/*
Copyright 2017 The Kubernetes Authors.

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

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	resourcequotaapi "github.com/kiosk-sh/kiosk/kube/plugin/pkg/admission/resourcequota/apis/resourcequota"
)

// ValidateConfiguration validates the configuration.
func ValidateConfiguration(config *resourcequotaapi.Configuration) field.ErrorList {
	allErrs := field.ErrorList{}
	fldPath := field.NewPath("limitedResources")
	for i, limitedResource := range config.LimitedResources {
		idxPath := fldPath.Index(i)
		if len(limitedResource.Resource) == 0 {
			allErrs = append(allErrs, field.Required(idxPath.Child("resource"), ""))
		}
	}
	return allErrs
}
