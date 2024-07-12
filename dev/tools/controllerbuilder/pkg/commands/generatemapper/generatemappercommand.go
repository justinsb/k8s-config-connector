// Copyright 2024 Google LLC
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

package generatemapper

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/options"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/protoapi"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/spf13/cobra"
)

type GenerateMapperOptions struct {
	*options.GenerateOptions

	APIGoPackagePath      string
	APIDirectory          string
	OutputMapperDirectory string
}

func (o *GenerateMapperOptions) InitDefaults() {

}

func (o *GenerateMapperOptions) BindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.APIGoPackagePath, "api-go-package-path", o.APIGoPackagePath, "package path")
	cmd.Flags().StringVar(&o.APIDirectory, "api-dir", o.APIDirectory, "base directory for reading APIs")
	cmd.Flags().StringVar(&o.OutputMapperDirectory, "output-dir", o.OutputMapperDirectory, "base directory for writing mappers")
}

func BuildCommand(baseOptions *options.GenerateOptions) *cobra.Command {
	opt := &GenerateMapperOptions{
		GenerateOptions: baseOptions,
	}

	opt.InitDefaults()

	cmd := &cobra.Command{
		Use:   "generate-mapper",
		Short: "generate mapper functions for a proto service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := RunGenerateMapper(ctx, opt); err != nil {
				return err
			}
			return nil
		},
	}

	opt.BindFlags(cmd)

	return cmd
}

func RunGenerateMapper(ctx context.Context, o *GenerateMapperOptions) error {
	if o.ServiceName == "" {
		return fmt.Errorf("ServiceName is required")
	}
	if o.GenerateOptions.ProtoSourcePath == "" {
		return fmt.Errorf("ProtoSourcePath is required")
	}
	if o.APIGoPackagePath == "" {
		return fmt.Errorf("GoPackagePath is required")
	}

	api, err := protoapi.LoadProto(o.GenerateOptions.ProtoSourcePath)
	if err != nil {
		return fmt.Errorf("loading proto: %w", err)
	}

	pathForMessage := func(msg protoreflect.MessageDescriptor) (string, bool) {
		fullName := string(msg.FullName())
		if strings.HasSuffix(fullName, "Request") {
			return "", false
		}
		if strings.HasSuffix(fullName, "Response") {
			return "", false
		}
		if strings.HasSuffix(fullName, "Metadata") {
			return "", false
		}
		if !strings.HasPrefix(fullName, o.ServiceName) {
			return "", false
		}

		protoPackagePath := string(msg.ParentFile().Package())
		protoPackagePath = strings.TrimPrefix(protoPackagePath, "google.")
		protoPackagePath = strings.TrimPrefix(protoPackagePath, "cloud.")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".v1")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".v1beta1")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".v2")
		protoPackagePath = strings.TrimSuffix(protoPackagePath, ".admin") // e.g. bigtable.admin.v2
		goPackage := strings.Join(strings.Split(protoPackagePath, "."), "/")

		return goPackage, true
	}
	mapperGenerator := codegen.NewMapperGenerator(pathForMessage)

	if err := mapperGenerator.VisitGoCode(o.APIGoPackagePath, o.APIDirectory); err != nil {
		return err
	}

	if err := mapperGenerator.VisitProto(api); err != nil {
		return err
	}

	if err := mapperGenerator.GenerateMappers(); err != nil {
		return err
	}

	addCopyright := true
	if err := mapperGenerator.WriteFiles(o.OutputMapperDirectory, addCopyright); err != nil {
		return err
	}

	return nil

}
