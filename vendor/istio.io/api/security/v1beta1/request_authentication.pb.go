// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: security/v1beta1/request_authentication.proto

// $schema: istio.security.v1beta1.RequestAuthentication
// $title: RequestAuthentication
// $description: Request authentication configuration for workloads.
// $location: https://istio.io/docs/reference/config/security/request_authentication.html
// $aliases: [/docs/reference/config/security/v1beta1/request_authentication]

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1beta1 "istio.io/api/type/v1beta1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RequestAuthentication defines what request authentication methods are supported by a workload.
// It will reject a request if the request contains invalid authentication information, based on the
// configured authentication rules. A request that does not contain any authentication credentials
// will be accepted but will not have any authenticated identity. To restrict access to authenticated
// requests only, this should be accompanied by an authorization rule.
// Examples:
//
// - Require JWT for all request for workloads that have label `app:httpbin`
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//     jwksUri: https://example.com/.well-known/jwks.json
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["*"]
// ```
//
// - A policy in the root namespace ("istio-system" by default) applies to workloads in all namespaces
// in a mesh. The following policy makes all workloads only accept requests that contain a
// valid JWT token.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: req-authn-for-all
//   namespace: istio-system
// spec:
//   jwtRules:
//   - issuer: "issuer-foo"
//     jwksUri: https://example.com/.well-known/jwks.json
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: require-jwt-for-all
//   namespace: istio-system
// spec:
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["*"]
// ```
//
// - The next example shows how to set a different JWT requirement for a different `host`. The `RequestAuthentication`
// declares it can accept JWTs issued by either `issuer-foo` or `issuer-bar` (the public key set is implicitly
// set from the OpenID Connect spec).
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   jwtRules:
//   - issuer: "issuer-foo"
//   - issuer: "issuer-bar"
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["issuer-foo/*"]
//     to:
//     - operation:
//         hosts: ["example.com"]
//   - from:
//     - source:
//         requestPrincipals: ["issuer-bar/*"]
//     to:
//     - operation:
//         hosts: ["another-host.com"]
// ```
//
// - You can fine tune the authorization policy to set different requirement per path. For example,
// to require JWT on all paths, except /healthz, the same `RequestAuthentication` can be used, but the
// authorization policy could be:
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: httpbin
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: httpbin
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["*"]
//   - to:
//     - operation:
//         paths: ["/healthz"]
// ```
//
// [Experimental] Routing based on derived [metadata](https://istio.io/latest/docs/reference/config/security/conditions/)
// is now supported. A prefix '@' is used to denote a match against internal metadata instead of the headers in the request.
// Currently this feature is only supported for the following metadata:
//
// - `request.auth.claims.{claim-name}[.{sub-claim}]*` which are extracted from validated JWT tokens. The claim name
// currently does not support the `.` character. Examples: `request.auth.claims.sub` and `request.auth.claims.name.givenName`.
//
// The use of matches against JWT claim metadata is only supported in Gateways. The following example shows:
//
// - RequestAuthentication to decode and validate a JWT. This also makes the `@request.auth.claims` available for use in the VirtualService.
// - AuthorizationPolicy to check for valid principals in the request. This makes the JWT required for the request.
// - VirtualService to route the request based on the "sub" claim.
//
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: RequestAuthentication
// metadata:
//   name: jwt-on-ingress
//   namespace: istio-system
// spec:
//  selector:
//    matchLabels:
//      app: istio-ingressgateway
//   jwtRules:
//   - issuer: "example.com"
//     jwksUri: https://example.com/.well-known/jwks.json
// ---
// apiVersion: security.istio.io/v1beta1
// kind: AuthorizationPolicy
// metadata:
//   name: require-jwt
//   namespace: istio-system
// spec:
//  selector:
//    matchLabels:
//      app: istio-ingressgateway
//   rules:
//   - from:
//     - source:
//         requestPrincipals: ["*"]
// ---
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: route-jwt
// spec:
//   hosts:
//   - foo.prod.svc.cluster.local
//   gateways:
//   - istio-ingressgateway
//   http:
//   - name: "v2"
//     match:
//     - headers:
//         "@request.auth.claims.sub":
//           exact: "dev"
//     route:
//     - destination:
//         host: foo.prod.svc.cluster.local
//         subset: v2
//   - name: "default"
//     route:
//     - destination:
//         host: foo.prod.svc.cluster.local
//         subset: v1
// ```
//
// <!-- crd generation tags
// +cue-gen:RequestAuthentication:groupName:security.istio.io
// +cue-gen:RequestAuthentication:version:v1beta1
// +cue-gen:RequestAuthentication:storageVersion
// +cue-gen:RequestAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:RequestAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:RequestAuthentication:subresource:status
// +cue-gen:RequestAuthentication:scope:Namespaced
// +cue-gen:RequestAuthentication:resource:categories=istio-io,security-istio-io,shortNames=ra
// +cue-gen:RequestAuthentication:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type RequestAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The selector decides where to apply the request authentication policy. The selector will match with workloads
	// in the same namespace as the request authentication policy. If the request authentication policy is in the root namespace,
	// the selector will additionally match with workloads in all namespaces.
	//
	// If not set, the selector will match all workloads.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Define the list of JWTs that can be validated at the selected workloads' proxy. A valid token
	// will be used to extract the authenticated identity.
	// Each rule will be activated only when a token is presented at the location recognized by the
	// rule. The token will be validated based on the JWT rule config. If validation fails, the request will
	// be rejected.
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	JwtRules []*JWTRule `protobuf:"bytes,2,rep,name=jwt_rules,json=jwtRules,proto3" json:"jwt_rules,omitempty"`
}

func (x *RequestAuthentication) Reset() {
	*x = RequestAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_request_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthentication) ProtoMessage() {}

func (x *RequestAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_request_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthentication.ProtoReflect.Descriptor instead.
func (*RequestAuthentication) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_request_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAuthentication) GetSelector() *v1beta1.WorkloadSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *RequestAuthentication) GetJwtRules() []*JWTRule {
	if x != nil {
		return x.JwtRules
	}
	return nil
}

var File_security_v1beta1_request_authentication_proto protoreflect.FileDescriptor

var file_security_v1beta1_request_authentication_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x09,
	0x6a, 0x77, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x08, 0x6a, 0x77, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_security_v1beta1_request_authentication_proto_rawDescOnce sync.Once
	file_security_v1beta1_request_authentication_proto_rawDescData = file_security_v1beta1_request_authentication_proto_rawDesc
)

func file_security_v1beta1_request_authentication_proto_rawDescGZIP() []byte {
	file_security_v1beta1_request_authentication_proto_rawDescOnce.Do(func() {
		file_security_v1beta1_request_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_v1beta1_request_authentication_proto_rawDescData)
	})
	return file_security_v1beta1_request_authentication_proto_rawDescData
}

var file_security_v1beta1_request_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_security_v1beta1_request_authentication_proto_goTypes = []interface{}{
	(*RequestAuthentication)(nil),    // 0: istio.security.v1beta1.RequestAuthentication
	(*v1beta1.WorkloadSelector)(nil), // 1: istio.type.v1beta1.WorkloadSelector
	(*JWTRule)(nil),                  // 2: istio.security.v1beta1.JWTRule
}
var file_security_v1beta1_request_authentication_proto_depIdxs = []int32{
	1, // 0: istio.security.v1beta1.RequestAuthentication.selector:type_name -> istio.type.v1beta1.WorkloadSelector
	2, // 1: istio.security.v1beta1.RequestAuthentication.jwt_rules:type_name -> istio.security.v1beta1.JWTRule
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_security_v1beta1_request_authentication_proto_init() }
func file_security_v1beta1_request_authentication_proto_init() {
	if File_security_v1beta1_request_authentication_proto != nil {
		return
	}
	file_security_v1beta1_jwt_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_security_v1beta1_request_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthentication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_v1beta1_request_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_v1beta1_request_authentication_proto_goTypes,
		DependencyIndexes: file_security_v1beta1_request_authentication_proto_depIdxs,
		MessageInfos:      file_security_v1beta1_request_authentication_proto_msgTypes,
	}.Build()
	File_security_v1beta1_request_authentication_proto = out.File
	file_security_v1beta1_request_authentication_proto_rawDesc = nil
	file_security_v1beta1_request_authentication_proto_goTypes = nil
	file_security_v1beta1_request_authentication_proto_depIdxs = nil
}
