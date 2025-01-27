---
sidebar_position: 3
sidebar_label: "Post-Deployment Check"
---

# Post-Deployment Check

The below example is from a 4-node kubernetes cluster:

```console
$ kubectl get no
NAME           STATUS   ROLES   AGE   VERSION
k8s-master-1   Ready    master  82d   v1.24.3-2+63243a96d1c393
k8s-worker-1   Ready    worker  36d   v1.24.3-2+63243a96d1c393
k8s-worker-2   Ready    worker  59d   v1.24.3-2+63243a96d1c393
k8s-worker-3   Ready    worker  36d   v1.24.3-2+63243a96d1c393
```

## Check the pods

The following pods should be up and running:

```console
$ kubectl -n hwameistor get pod
NAME                                                       READY   STATUS                  RESTARTS   AGE
hwameistor-local-disk-csi-controller-665bb7f47d-6227f      2/2     Running                 0          30s
hwameistor-local-disk-manager-5ph2d                        2/2     Running                 0          30s
hwameistor-local-disk-manager-jhj59                        2/2     Running                 0          30s
hwameistor-local-disk-manager-k9cvj                        2/2     Running                 0          30s
hwameistor-local-disk-manager-kxwww                        2/2     Running                 0          30s
hwameistor-local-storage-csi-controller-667d949fbb-k488w   3/3     Running                 0          30s
hwameistor-local-storage-csqqv                             2/2     Running                 0          30s
hwameistor-local-storage-gcrzm                             2/2     Running                 0          30s
hwameistor-local-storage-v8g7t                             2/2     Running                 0          30s
hwameistor-local-storage-zkwmn                             2/2     Running                 0          30s
hwameistor-scheduler-58dfcf79f5-lswkt                      1/1     Running                 0          30s
hwameistor-webhook-986479678-278cr                         1/1     Running                 0          30s
```

:::info
`local-disk-manager` and `local-storage` are `DaemonSets`. They should have one pod on each Kubernetes node.
:::

## Check the APIs

HwameiStor CRDs create the following APIs.

```console
$ kubectl api-resources --api-group hwameistor.io

NAME                       SHORTNAMES   APIVERSION               NAMESPACED   KIND
localdiskclaims            ldc          hwameistor.io/v1alpha1   false        LocalDiskClaim
localdisknodes             ldn          hwameistor.io/v1alpha1   false        LocalDiskNode
localdisks                 ld           hwameistor.io/v1alpha1   false        LocalDisk
localdiskvolumes           ldv          hwameistor.io/v1alpha1   false        LocalDiskVolume
localstoragenodes          lsn          hwameistor.io/v1alpha1   false        LocalStorageNode
localvolumeconverts        lvconvert    hwameistor.io/v1alpha1   true         LocalVolumeConvert
localvolumeexpands         lvexpand     hwameistor.io/v1alpha1   false        LocalVolumeExpand
localvolumegroupconverts   lvgconvert   hwameistor.io/v1alpha1   true         LocalVolumeGroupConvert
localvolumegroupmigrates   lvgmigrate   hwameistor.io/v1alpha1   true         LocalVolumeGroupMigrate
localvolumegroups          lvg          hwameistor.io/v1alpha1   true         LocalVolumeGroup
localvolumemigrates        lvmigrate    hwameistor.io/v1alpha1   true         LocalVolumeMigrate
localvolumereplicas        lvr          hwameistor.io/v1alpha1   false        LocalVolumeReplica
localvolumes               lv           hwameistor.io/v1alpha1   false        LocalVolume
```

For the details about CRDs, please also refer to the chapter [CRDs](../../architecture/apis.md).

## Check the `LocalDiskNode` and `localDisks`

HwameiStor autoscans each node and registers each disk as CRD `LocalDisk(ld)`. The unused disks are displayed with `PHASE: Unclaimed`.

```console
$ kubectl get localdisknodes
NAME           TOTALDISK   FREEDISK
k8s-master-1   5           3
k8s-worker-1   5           2
k8s-worker-2   5           2
k8s-worker-3   5           2

$ kubectl get localdisks
NAME               NODEMATCH      CLAIM   PHASE
k8s-master-1-sda   k8s-master-1           Inuse
k8s-worker-1-sda   k8s-worker-1           Inuse
k8s-worker-1-sdb   k8s-worker-1           Unclaimed
k8s-worker-1-sdc   k8s-worker-1           Unclaimed
k8s-worker-2-sda   k8s-worker-2           Inuse
k8s-worker-2-sdb   k8s-worker-2           Unclaimed
k8s-worker-2-sdc   k8s-worker-2           Unclaimed
k8s-worker-3-sda   k8s-worker-3           Inuse
k8s-worker-3-sdb   k8s-worker-3           Unclaimed
k8s-worker-3-sdc   k8s-worker-3           Unclaimed
```
