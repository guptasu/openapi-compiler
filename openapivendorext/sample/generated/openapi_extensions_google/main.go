// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/googleapis/openapi-compiler/compiler"
	"github.com/googleapis/openapi-compiler/openapivendorext/plugin"
	"github.com/googleapis/openapi-compiler/openapivendorext/sample/generated/openapi_extensions_google/proto"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Version() string {
	return "main"
}

type documentHandler func(name string, version string, extensionName string, document string)

func forInputYamlFromOpenapic(handler documentHandler) {
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("File error:", err.Error())
		os.Exit(1)
	}
	request := &openapiextension_plugin_v1.VendorExtensionHandlerRequest{}
	err = proto.Unmarshal(data, request)
	handler(request.Wrapper.Name, request.Wrapper.Version, request.Wrapper.ExtensionName, request.Wrapper.Yaml)
}

func main() {
	response := &openapiextension_plugin_v1.VendorExtensionHandlerResponse{}
	forInputYamlFromOpenapic(
		func(name string, version string, extensionName string, yamlInput string) {
			var info yaml.MapSlice
			var newObject proto.Message
			var err error
			err = yaml.Unmarshal([]byte(yamlInput), &info)
			if err != nil {
				response.Error = append(response.Error, err.Error())
				responseBytes, _ := proto.Marshal(response)
				os.Stdout.Write(responseBytes)
				os.Exit(0)
			}

			switch extensionName {
			// All supported extensions

			case "x-book":
				newObject, err = googleextensions.NewBook(info, compiler.NewContextWithCustomAnyProtoGenerators("$root", nil, nil))

			case "x-shelve":
				newObject, err = googleextensions.NewShelve(info, compiler.NewContextWithCustomAnyProtoGenerators("$root", nil, nil))

			default:
				responseBytes, _ := proto.Marshal(response)
				os.Stdout.Write(responseBytes)
				os.Exit(0)
			}
			// If we reach hear, then the extension is handled
			response.Handled = true
			if err != nil {
				response.Error = append(response.Error, err.Error())
				responseBytes, _ := proto.Marshal(response)
				os.Stdout.Write(responseBytes)
				os.Exit(0)
			}
			response.Value, err = ptypes.MarshalAny(newObject)
			if err != nil {
				response.Error = append(response.Error, err.Error())
				responseBytes, _ := proto.Marshal(response)
				os.Stdout.Write(responseBytes)
				os.Exit(0)
			}
		})

	responseBytes, _ := proto.Marshal(response)
	os.Stdout.Write(responseBytes)
}