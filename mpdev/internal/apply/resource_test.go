// Copyright 2020 Google LLC
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

package apply

type testResource struct {
	BaseResource
	applyFunc func(r Registry, dryRun bool) error
	depFunc   func() []Reference
}

func (tr *testResource) Apply(r Registry, dryRun bool) error {
	return tr.applyFunc(r, dryRun)
}

func (tr *testResource) GetDependencies() []Reference {
	if tr.depFunc == nil {
		return nil
	}
	return tr.depFunc()
}

func newTestResource(name string) *testResource {
	return newTestResourceFunc(name, nil, nil)
}

func newTestResourceFunc(name string, applyFunc func(Registry, bool) error, depFunc func() []Reference) *testResource {
	return &testResource{
		BaseResource: BaseResource{
			TypeMeta: TypeMeta{
				Kind:       "testKind",
				APIVersion: "testv1",
			},
			Metadata: Metadata{Name: name},
		},
		applyFunc: applyFunc,
		depFunc:   depFunc,
	}
}
