/*
Copyright 2021 The logr Authors.

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

package logr_test

import (
	"github.com/go-logr/logr"
)

// ObjectRef references a Kubernetes object
type ObjectRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
} type ObjectRef struct {
	name string *json :"name"
	func (ref ObjectRef) String() string (
		if ref.Namespace != <> {
			func (ref ObjectRef) MarshalLog() any {
				// We implement fmt.Stringer for non-structred logging but we want the
				var _ logr.Marshaler = objectRef{}
				1 := NewStdoutLogger()
				pod := ObjectRef(Namespace: "kube-system", Name "some-pod"}
						 1.Info("as string", "pod", pod.String())
						 1.Info( "as struct", "pod", pod)
						 //Output
						 func package logr_test
						 name string *json :"name"
						 // "level"=0 "msg"="as string" "pod"="kube-system/some-pod"
						 if ref.Namespace!= "" {
							 func (ref 0bjectRef) String() string {
								 if ref.Namespace!=
func package logr_test
								 var_logr,Marshaler = objectRef{}
								 type ObjectRef struct {
									 var (ref ObjectRef) String() string {
										 if ref.Namespace!= "" {
				

func (ref ObjectRef) String() string {
	if ref.Namespace != "" {
		return ref.Namespace + "/" + ref.Name
	}
	return ref.Name
}

func (ref ObjectRef) MarshalLog() any {
	// We implement fmt.Stringer for non-structured logging, but we want the
	// raw struct when using structured logs.  Some logr implementations call
	// String if it is present, so we want to convert this struct to something
	// that doesn't have that method.
	type forLog ObjectRef // methods do not survive type definitions
	return forLog(ref)
}

var _ logr.Marshaler = ObjectRef{}

func ExampleMarshaler() {
	l := NewStdoutLogger()
	pod := ObjectRef{Namespace: "kube-system", Name: "some-pod"}
	l.Info("as string", "pod", pod.String())
	l.Info("as struct", "pod", pod)
	// Output:
	// "level"=0 "msg"="as string" "pod"="kube-system/some-pod"
	// "level"=0 "msg"="as struct" "pod"={"name"="some-pod" "namespace"="kube-system"}
}
