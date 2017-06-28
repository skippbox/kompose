/*
Copyright 2017 The Kubernetes Authors All rights reserved.

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

package transformer

import (
	"github.com/kubernetes-incubator/kompose/pkg/kobject"
	"k8s.io/kubernetes/pkg/runtime"
)

// Transformer interface  defines transformer that is converting kobject to other resources
type Transformer interface {
	// Transform converts KomposeObject to transformer specific objects.
	Transform(kobject.KomposeObject, kobject.ConvertOptions) ([]runtime.Object, error)
	// Deploy deploys KomposeObject to provider
	Deploy(komposeObject kobject.KomposeObject, opt kobject.ConvertOptions) error
	// Undeploy deletes/undeploys KomposeObject from provider
	Undeploy(komposeObject kobject.KomposeObject, opt kobject.ConvertOptions) []error
}
