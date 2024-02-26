/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/apache/camel-k/v2/pkg/apis/camel/v1"
	v1alpha1 "github.com/apache/camel-k/v2/pkg/apis/camel/v1alpha1"
	camelv1 "github.com/apache/camel-k/v2/pkg/client/camel/applyconfiguration/camel/v1"
	camelv1alpha1 "github.com/apache/camel-k/v2/pkg/client/camel/applyconfiguration/camel/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=camel.apache.org, Version=v1
	case v1.SchemeGroupVersion.WithKind("AddonTrait"):
		return &camelv1.AddonTraitApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Artifact"):
		return &camelv1.ArtifactApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BaseTask"):
		return &camelv1.BaseTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Build"):
		return &camelv1.BuildApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildahTask"):
		return &camelv1.BuildahTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildCondition"):
		return &camelv1.BuildConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildConfiguration"):
		return &camelv1.BuildConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuilderTask"):
		return &camelv1.BuilderTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildSpec"):
		return &camelv1.BuildSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("BuildStatus"):
		return &camelv1.BuildStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelArtifact"):
		return &camelv1.CamelArtifactApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelArtifactDependency"):
		return &camelv1.CamelArtifactDependencyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelArtifactExclusion"):
		return &camelv1.CamelArtifactExclusionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelCatalog"):
		return &camelv1.CamelCatalogApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelCatalogCondition"):
		return &camelv1.CamelCatalogConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelCatalogSpec"):
		return &camelv1.CamelCatalogSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelCatalogStatus"):
		return &camelv1.CamelCatalogStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelLoader"):
		return &camelv1.CamelLoaderApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelScheme"):
		return &camelv1.CamelSchemeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CamelSchemeScope"):
		return &camelv1.CamelSchemeScopeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Capability"):
		return &camelv1.CapabilityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigurationSpec"):
		return &camelv1.ConfigurationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DataSpec"):
		return &camelv1.DataSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DataTypeReference"):
		return &camelv1.DataTypeReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DataTypeSpec"):
		return &camelv1.DataTypeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DataTypesSpec"):
		return &camelv1.DataTypesSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Endpoint"):
		return &camelv1.EndpointApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EndpointProperties"):
		return &camelv1.EndpointPropertiesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ErrorHandlerSpec"):
		return &camelv1.ErrorHandlerSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("EventTypeSpec"):
		return &camelv1.EventTypeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalDocumentation"):
		return &camelv1.ExternalDocumentationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Failure"):
		return &camelv1.FailureApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("FailureRecovery"):
		return &camelv1.FailureRecoveryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Flow"):
		return &camelv1.FlowApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HeaderSpec"):
		return &camelv1.HeaderSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HealthCheckResponse"):
		return &camelv1.HealthCheckResponseApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Integration"):
		return &camelv1.IntegrationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationCondition"):
		return &camelv1.IntegrationConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationKit"):
		return &camelv1.IntegrationKitApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationKitCondition"):
		return &camelv1.IntegrationKitConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationKitSpec"):
		return &camelv1.IntegrationKitSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationKitStatus"):
		return &camelv1.IntegrationKitStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationKitTraits"):
		return &camelv1.IntegrationKitTraitsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatform"):
		return &camelv1.IntegrationPlatformApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatformBuildSpec"):
		return &camelv1.IntegrationPlatformBuildSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatformCondition"):
		return &camelv1.IntegrationPlatformConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatformKameletSpec"):
		return &camelv1.IntegrationPlatformKameletSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatformSpec"):
		return &camelv1.IntegrationPlatformSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationPlatformStatus"):
		return &camelv1.IntegrationPlatformStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationSpec"):
		return &camelv1.IntegrationSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("IntegrationStatus"):
		return &camelv1.IntegrationStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JibTask"):
		return &camelv1.JibTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JSON"):
		return &camelv1.JSONApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JSONSchemaProp"):
		return &camelv1.JSONSchemaPropApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JSONSchemaProps"):
		return &camelv1.JSONSchemaPropsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Kamelet"):
		return &camelv1.KameletApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KameletCondition"):
		return &camelv1.KameletConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KameletProperty"):
		return &camelv1.KameletPropertyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KameletRepositorySpec"):
		return &camelv1.KameletRepositorySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KameletSpec"):
		return &camelv1.KameletSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KameletStatus"):
		return &camelv1.KameletStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KanikoTask"):
		return &camelv1.KanikoTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KanikoTaskCache"):
		return &camelv1.KanikoTaskCacheApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MavenArtifact"):
		return &camelv1.MavenArtifactApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MavenBuildSpec"):
		return &camelv1.MavenBuildSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MavenSpec"):
		return &camelv1.MavenSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Pipe"):
		return &camelv1.PipeApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PipeCondition"):
		return &camelv1.PipeConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PipeSpec"):
		return &camelv1.PipeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PipeStatus"):
		return &camelv1.PipeStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodCondition"):
		return &camelv1.PodConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodSpec"):
		return &camelv1.PodSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodSpecTemplate"):
		return &camelv1.PodSpecTemplateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PublishTask"):
		return &camelv1.PublishTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RegistrySpec"):
		return &camelv1.RegistrySpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Repository"):
		return &camelv1.RepositoryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RepositoryPolicy"):
		return &camelv1.RepositoryPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RuntimeSpec"):
		return &camelv1.RuntimeSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("S2iTask"):
		return &camelv1.S2iTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Server"):
		return &camelv1.ServerApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SourceSpec"):
		return &camelv1.SourceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SpectrumTask"):
		return &camelv1.SpectrumTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Task"):
		return &camelv1.TaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Template"):
		return &camelv1.TemplateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TraitConfiguration"):
		return &camelv1.TraitConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Traits"):
		return &camelv1.TraitsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TraitSpec"):
		return &camelv1.TraitSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("UserTask"):
		return &camelv1.UserTaskApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ValueSource"):
		return &camelv1.ValueSourceApplyConfiguration{}

		// Group=camel.apache.org, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("DataTypeReference"):
		return &camelv1alpha1.DataTypeReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DataTypeSpec"):
		return &camelv1alpha1.DataTypeSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DataTypesSpec"):
		return &camelv1alpha1.DataTypesSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Endpoint"):
		return &camelv1alpha1.EndpointApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EndpointProperties"):
		return &camelv1alpha1.EndpointPropertiesApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ErrorHandlerSpec"):
		return &camelv1alpha1.ErrorHandlerSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("EventTypeSpec"):
		return &camelv1alpha1.EventTypeSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ExternalDocumentation"):
		return &camelv1alpha1.ExternalDocumentationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HeaderSpec"):
		return &camelv1alpha1.HeaderSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JSON"):
		return &camelv1alpha1.JSONApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JSONSchemaProp"):
		return &camelv1alpha1.JSONSchemaPropApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("JSONSchemaProps"):
		return &camelv1alpha1.JSONSchemaPropsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Kamelet"):
		return &camelv1alpha1.KameletApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletBinding"):
		return &camelv1alpha1.KameletBindingApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletBindingCondition"):
		return &camelv1alpha1.KameletBindingConditionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletBindingSpec"):
		return &camelv1alpha1.KameletBindingSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletBindingStatus"):
		return &camelv1alpha1.KameletBindingStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletCondition"):
		return &camelv1alpha1.KameletConditionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletProperty"):
		return &camelv1alpha1.KameletPropertyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletSpec"):
		return &camelv1alpha1.KameletSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("KameletStatus"):
		return &camelv1alpha1.KameletStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Template"):
		return &camelv1alpha1.TemplateApplyConfiguration{}

	}
	return nil
}
