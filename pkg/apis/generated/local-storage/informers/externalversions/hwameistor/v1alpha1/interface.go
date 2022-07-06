/*
Copyright 2022 Contributors to The HwameiStor project.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/hwameistor/hwameistor/pkg/apis/generated/local-storage/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// LocalStorageNodes returns a LocalStorageNodeInformer.
	LocalStorageNodes() LocalStorageNodeInformer
	// LocalVolumes returns a LocalVolumeInformer.
	LocalVolumes() LocalVolumeInformer
	// LocalVolumeConverts returns a LocalVolumeConvertInformer.
	LocalVolumeConverts() LocalVolumeConvertInformer
	// LocalVolumeExpands returns a LocalVolumeExpandInformer.
	LocalVolumeExpands() LocalVolumeExpandInformer
	// LocalVolumeGroups returns a LocalVolumeGroupInformer.
	LocalVolumeGroups() LocalVolumeGroupInformer
	// LocalVolumeGroupConverts returns a LocalVolumeGroupConvertInformer.
	LocalVolumeGroupConverts() LocalVolumeGroupConvertInformer
	// LocalVolumeGroupMigrates returns a LocalVolumeGroupMigrateInformer.
	LocalVolumeGroupMigrates() LocalVolumeGroupMigrateInformer
	// LocalVolumeMigrates returns a LocalVolumeMigrateInformer.
	LocalVolumeMigrates() LocalVolumeMigrateInformer
	// LocalVolumeReplicas returns a LocalVolumeReplicaInformer.
	LocalVolumeReplicas() LocalVolumeReplicaInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// LocalStorageNodes returns a LocalStorageNodeInformer.
func (v *version) LocalStorageNodes() LocalStorageNodeInformer {
	return &localStorageNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumes returns a LocalVolumeInformer.
func (v *version) LocalVolumes() LocalVolumeInformer {
	return &localVolumeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeConverts returns a LocalVolumeConvertInformer.
func (v *version) LocalVolumeConverts() LocalVolumeConvertInformer {
	return &localVolumeConvertInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeExpands returns a LocalVolumeExpandInformer.
func (v *version) LocalVolumeExpands() LocalVolumeExpandInformer {
	return &localVolumeExpandInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeGroups returns a LocalVolumeGroupInformer.
func (v *version) LocalVolumeGroups() LocalVolumeGroupInformer {
	return &localVolumeGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeGroupConverts returns a LocalVolumeGroupConvertInformer.
func (v *version) LocalVolumeGroupConverts() LocalVolumeGroupConvertInformer {
	return &localVolumeGroupConvertInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeGroupMigrates returns a LocalVolumeGroupMigrateInformer.
func (v *version) LocalVolumeGroupMigrates() LocalVolumeGroupMigrateInformer {
	return &localVolumeGroupMigrateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeMigrates returns a LocalVolumeMigrateInformer.
func (v *version) LocalVolumeMigrates() LocalVolumeMigrateInformer {
	return &localVolumeMigrateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LocalVolumeReplicas returns a LocalVolumeReplicaInformer.
func (v *version) LocalVolumeReplicas() LocalVolumeReplicaInformer {
	return &localVolumeReplicaInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
