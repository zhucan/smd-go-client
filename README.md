# Go API client for smdclient

Resource for managing storage

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://www.deeproute.ai/](https://www.deeproute.ai/)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./smdclient"
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionLogsApi* | [**GetActionLog**](docs/ActionLogsApi.md#getactionlog) | **Get** /action-logs/{log-id} | get a action log
*ActionLogsApi* | [**ListActionLogs**](docs/ActionLogsApi.md#listactionlogs) | **Get** /action-logs | get all action logs
*BlockVolumeSnapshotsApi* | [**GetBlockVolumeSnapshot**](docs/BlockVolumeSnapshotsApi.md#getblockvolumesnapshot) | **Get** /block-volume-snapshots/{snapshot-id} | get a snapshot of block volume
*BlockVolumeSnapshotsApi* | [**ListBlockVolumeSnapshots**](docs/BlockVolumeSnapshotsApi.md#listblockvolumesnapshots) | **Get** /block-volume-snapshots | list all snapshots of block volume
*BlockVolumesApi* | [**GetVolume**](docs/BlockVolumesApi.md#getvolume) | **Get** /block-volumes/{volume-id} | get an volume
*BlockVolumesApi* | [**ListVolumes**](docs/BlockVolumesApi.md#listvolumes) | **Get** /block-volumes | get all volumes
*CephClustersApi* | [**CreateCephCluster**](docs/CephClustersApi.md#createcephcluster) | **Post** /ceph-clusters | create ceph cluster
*CephClustersApi* | [**DeleteCephCluster**](docs/CephClustersApi.md#deletecephcluster) | **Delete** /ceph-clusters/{ceph-cluster-id} | delete ceph cluster
*CephClustersApi* | [**GetCephCluster**](docs/CephClustersApi.md#getcephcluster) | **Get** /ceph-clusters/{ceph-cluster-id} | get ceph cluster
*CephClustersApi* | [**ListCephClusters**](docs/CephClustersApi.md#listcephclusters) | **Get** /ceph-clusters | get all ceph cluster
*CephFilesystemsApi* | [**CreateCephFilesystem**](docs/CephFilesystemsApi.md#createcephfilesystem) | **Post** /ceph-filesystems | create a ceph filesystem
*CephFilesystemsApi* | [**DeleteCephFilesystem**](docs/CephFilesystemsApi.md#deletecephfilesystem) | **Delete** /ceph-filesystems/{cephfs-id} | delete a ceph filesystem
*CephFilesystemsApi* | [**GetCephFilesystem**](docs/CephFilesystemsApi.md#getcephfilesystem) | **Get** /ceph-filesystems/{cephfs-id} | get a ceph filesystem
*CephFilesystemsApi* | [**ListCephFilesystems**](docs/CephFilesystemsApi.md#listcephfilesystems) | **Get** /ceph-filesystems | get all ceph filesystems
*CephNfsApi* | [**CreateCephNFS**](docs/CephNfsApi.md#createcephnfs) | **Post** /ceph-nfs | create an ceph nfs
*CephNfsApi* | [**DeleteCephNFS**](docs/CephNfsApi.md#deletecephnfs) | **Delete** /ceph-nfs/{ceph-nfs-id} | delete an ceph nfs
*CephNfsApi* | [**GetCephNFS**](docs/CephNfsApi.md#getcephnfs) | **Get** /ceph-nfs/{ceph-nfs-id} | get an ceph nfs
*CephNfsApi* | [**ListCephNFSes**](docs/CephNfsApi.md#listcephnfses) | **Get** /ceph-nfs | get all ceph nfs
*CephNfsExportsApi* | [**CreateCephNFSExport**](docs/CephNfsExportsApi.md#createcephnfsexport) | **Post** /ceph-nfs-exports | create ceph nfs export
*CephNfsExportsApi* | [**DeleteCephNFSExport**](docs/CephNfsExportsApi.md#deletecephnfsexport) | **Delete** /ceph-nfs-exports/{ceph-nfs-export-id} | delete a ceph nfs export
*CephNfsExportsApi* | [**GetCephNFSExport**](docs/CephNfsExportsApi.md#getcephnfsexport) | **Get** /ceph-nfs-exports/{ceph-nfs-export-id} | get a ceph nfs export
*CephNfsExportsApi* | [**ListCephNFSExports**](docs/CephNfsExportsApi.md#listcephnfsexports) | **Get** /ceph-nfs-exports | get all ceph nfs exports
*CephfsSubvolumegroupsApi* | [**CreateCephfsSubvolumeGroup**](docs/CephfsSubvolumegroupsApi.md#createcephfssubvolumegroup) | **Post** /cephfs-subvolumegroups | create cephfs subvolumegroup
*CephfsSubvolumegroupsApi* | [**DeleteCephfsSubvolumeGroup**](docs/CephfsSubvolumegroupsApi.md#deletecephfssubvolumegroup) | **Delete** /cephfs-subvolumegroups/{cephfs-subvolumegroup-id} | delete a cephfs subvolumegroup
*CephfsSubvolumegroupsApi* | [**GetCephfsSubvolumeGroup**](docs/CephfsSubvolumegroupsApi.md#getcephfssubvolumegroup) | **Get** /cephfs-subvolumegroups/{cephfs-subvolumegroup-id} | get a cephfs subvolumegroup
*CephfsSubvolumegroupsApi* | [**ListCephfsSubvolumeGroups**](docs/CephfsSubvolumegroupsApi.md#listcephfssubvolumegroups) | **Get** /cephfs-subvolumegroups | get all cephfs subvolumegroups
*CephfsSubvolumesApi* | [**CreateCephfsSubvolume**](docs/CephfsSubvolumesApi.md#createcephfssubvolume) | **Post** /cephfs-subvolumes | create cephfs subvolume
*CephfsSubvolumesApi* | [**DeleteCephfsSubvolume**](docs/CephfsSubvolumesApi.md#deletecephfssubvolume) | **Delete** /cephfs-subvolumes/{cephfs-subvolume-id} | delete a cephfs subvolume
*CephfsSubvolumesApi* | [**GetCephfsSubvolume**](docs/CephfsSubvolumesApi.md#getcephfssubvolume) | **Get** /cephfs-subvolumes/{cephfs-subvolume-id} | get a cephfs subvolume
*CephfsSubvolumesApi* | [**ListCephfsSubvolumes**](docs/CephfsSubvolumesApi.md#listcephfssubvolumes) | **Get** /cephfs-subvolumes | get all cephfs subvolumes
*CephfsSubvolumesApi* | [**ResizeCephfsSubvolume**](docs/CephfsSubvolumesApi.md#resizecephfssubvolume) | **Post** /cephfs-subvolumes/{cephfs-subvolume-id}:resize | resize a cephfs subvolume
*CrushRootsApi* | [**AddOsdsToCrushRoot**](docs/CrushRootsApi.md#addosdstocrushroot) | **Post** /crush-roots/{crush-root-id}:addOsds | add osds to crush root
*CrushRootsApi* | [**CreateCrushRoot**](docs/CrushRootsApi.md#createcrushroot) | **Post** /crush-roots | create crush root
*CrushRootsApi* | [**DeleteCrushRoot**](docs/CrushRootsApi.md#deletecrushroot) | **Delete** /crush-roots/{crush-root-id} | delete a crush root
*CrushRootsApi* | [**GetCrushRoot**](docs/CrushRootsApi.md#getcrushroot) | **Get** /crush-roots/{crush-root-id} | get a crush root
*CrushRootsApi* | [**ListCrushRoots**](docs/CrushRootsApi.md#listcrushroots) | **Get** /crush-roots | get all crush roots
*CrushRootsApi* | [**RemoveOsdsFromCrushRoot**](docs/CrushRootsApi.md#removeosdsfromcrushroot) | **Post** /crush-roots/:removeOsds | remove osds from crush root
*DisksApi* | [**GetDisk**](docs/DisksApi.md#getdisk) | **Get** /disks/{disk-id} | get a disk
*DisksApi* | [**ListDisks**](docs/DisksApi.md#listdisks) | **Get** /disks | get all disks
*HostsApi* | [**CreateHost**](docs/HostsApi.md#createhost) | **Post** /hosts | create a host
*HostsApi* | [**DeleteHost**](docs/HostsApi.md#deletehost) | **Delete** /hosts/{host-id} | delete a host
*HostsApi* | [**GetHost**](docs/HostsApi.md#gethost) | **Get** /hosts/{host-id} | get a host
*HostsApi* | [**ListHosts**](docs/HostsApi.md#listhosts) | **Get** /hosts | get all hosts
*NfsGatewaysApi* | [**GetNFSGateway**](docs/NfsGatewaysApi.md#getnfsgateway) | **Get** /nfs-gateways/{nfs-gateway-id} | get a nfs gateway
*NfsGatewaysApi* | [**ListNFSGateways**](docs/NfsGatewaysApi.md#listnfsgateways) | **Get** /nfs-gateways | list all nfs gateways
*ObjectBucketsApi* | [**CreateObjectBucket**](docs/ObjectBucketsApi.md#createobjectbucket) | **Post** /object-buckets | create an object bucket
*ObjectBucketsApi* | [**DeleteObjectBucket**](docs/ObjectBucketsApi.md#deleteobjectbucket) | **Delete** /object-buckets/{object-bucket-id} | delete an object bucket
*ObjectBucketsApi* | [**GetObjectBucket**](docs/ObjectBucketsApi.md#getobjectbucket) | **Get** /object-buckets/{object-bucket-id} | get an object bucket
*ObjectBucketsApi* | [**ListObjectBuckets**](docs/ObjectBucketsApi.md#listobjectbuckets) | **Get** /object-buckets | get all object buckets
*ObjectGatewaysApi* | [**GetObjectGateway**](docs/ObjectGatewaysApi.md#getobjectgateway) | **Get** /object-gateways/{object-gateway-id} | get an object gateway
*ObjectGatewaysApi* | [**ListObjectGateways**](docs/ObjectGatewaysApi.md#listobjectgateways) | **Get** /object-gateways | list all object gateways
*ObjectStorageClassesApi* | [**CreateObjectStorageClass**](docs/ObjectStorageClassesApi.md#createobjectstorageclass) | **Post** /object-storage-classes | create an object storage class
*ObjectStorageClassesApi* | [**DeleteObjectStorageClass**](docs/ObjectStorageClassesApi.md#deleteobjectstorageclass) | **Delete** /object-storage-classes/{object-storage-class-id} | delete an object storage class
*ObjectStorageClassesApi* | [**GetObjectStorageClass**](docs/ObjectStorageClassesApi.md#getobjectstorageclass) | **Get** /object-storage-classes/{object-storage-class-id} | get an object storage classs
*ObjectStorageClassesApi* | [**ListObjectStorageClasses**](docs/ObjectStorageClassesApi.md#listobjectstorageclasses) | **Get** /object-storage-classes | get all object storage classes
*ObjectStoresApi* | [**CreateObjectStore**](docs/ObjectStoresApi.md#createobjectstore) | **Post** /object-stores | create an object store
*ObjectStoresApi* | [**DeleteObjectStore**](docs/ObjectStoresApi.md#deleteobjectstore) | **Delete** /object-stores/{object-store-id} | delete an object store
*ObjectStoresApi* | [**GetObjectStore**](docs/ObjectStoresApi.md#getobjectstore) | **Get** /object-stores/{object-store-id} | get an object store
*ObjectStoresApi* | [**ListObjectStores**](docs/ObjectStoresApi.md#listobjectstores) | **Get** /object-stores | get all object stores
*OsUserApi* | [**CreateOsUser**](docs/OsUserApi.md#createosuser) | **Post** /os-users | create os user
*OsUserApi* | [**DeleteOsUser**](docs/OsUserApi.md#deleteosuser) | **Delete** /os-users/{os-user-id} | delete os user
*OsUserApi* | [**GetOsUser**](docs/OsUserApi.md#getosuser) | **Get** /os-users/{os-user-id} | get an os user
*OsUserApi* | [**ListOsUsers**](docs/OsUserApi.md#listosusers) | **Get** /os-users | list all kind of os users
*OsdsApi* | [**CreateOsds**](docs/OsdsApi.md#createosds) | **Post** /osds | create osds
*OsdsApi* | [**DeleteOsds**](docs/OsdsApi.md#deleteosds) | **Delete** /osds | delete osds from specific ceph cluster
*OsdsApi* | [**GetOsd**](docs/OsdsApi.md#getosd) | **Get** /osds/{osd-id} | get a osd
*OsdsApi* | [**ListOsds**](docs/OsdsApi.md#listosds) | **Get** /osds | get all osds
*PoolsApi* | [**CreatePool**](docs/PoolsApi.md#createpool) | **Post** /pools | create a pool
*PoolsApi* | [**DeletePool**](docs/PoolsApi.md#deletepool) | **Delete** /pools/{pool-id} | delete a pool
*PoolsApi* | [**GetPool**](docs/PoolsApi.md#getpool) | **Get** /pools/{pool-id} | get a pool
*PoolsApi* | [**ListPools**](docs/PoolsApi.md#listpools) | **Get** /pools | get all pools


## Documentation For Models

 - [CephCluster](docs/CephCluster.md)
 - [CephNfs](docs/CephNfs.md)
 - [CephNfsExport](docs/CephNfsExport.md)
 - [Cephfs](docs/Cephfs.md)
 - [CephfsSubvolume](docs/CephfsSubvolume.md)
 - [CephfsSubvolumegroup](docs/CephfsSubvolumegroup.md)
 - [CrushRoot](docs/CrushRoot.md)
 - [Disks](docs/Disks.md)
 - [GormModel](docs/GormModel.md)
 - [Host](docs/Host.md)
 - [ModelsActionLog](docs/ModelsActionLog.md)
 - [ModelsBlockVolumeSnapshot](docs/ModelsBlockVolumeSnapshot.md)
 - [ModelsCephCluster](docs/ModelsCephCluster.md)
 - [ModelsCephNfs](docs/ModelsCephNfs.md)
 - [ModelsCephNfsExport](docs/ModelsCephNfsExport.md)
 - [ModelsCephfs](docs/ModelsCephfs.md)
 - [ModelsCephfsSubvolume](docs/ModelsCephfsSubvolume.md)
 - [ModelsCephfsSubvolumeGroup](docs/ModelsCephfsSubvolumeGroup.md)
 - [ModelsCrushRoot](docs/ModelsCrushRoot.md)
 - [ModelsDisk](docs/ModelsDisk.md)
 - [ModelsHost](docs/ModelsHost.md)
 - [ModelsNfsGateway](docs/ModelsNfsGateway.md)
 - [ModelsObjectBucket](docs/ModelsObjectBucket.md)
 - [ModelsObjectGateway](docs/ModelsObjectGateway.md)
 - [ModelsObjectStorageClass](docs/ModelsObjectStorageClass.md)
 - [ModelsObjectStore](docs/ModelsObjectStore.md)
 - [ModelsOsUser](docs/ModelsOsUser.md)
 - [ModelsOsd](docs/ModelsOsd.md)
 - [ModelsPool](docs/ModelsPool.md)
 - [ModelsPoolProperty](docs/ModelsPoolProperty.md)
 - [ModelsVolume](docs/ModelsVolume.md)
 - [ObjectBucket](docs/ObjectBucket.md)
 - [ObjectStorageClass](docs/ObjectStorageClass.md)
 - [ObjectStore](docs/ObjectStore.md)
 - [OsUser](docs/OsUser.md)
 - [Osds](docs/Osds.md)
 - [Pool](docs/Pool.md)
 - [RoutesCephNfsExportRequestParams](docs/RoutesCephNfsExportRequestParams.md)
 - [RoutesCephNfsRequestParams](docs/RoutesCephNfsRequestParams.md)
 - [RoutesCephfsRequestParams](docs/RoutesCephfsRequestParams.md)
 - [RoutesCephfsSubvolumeGroupRequestParams](docs/RoutesCephfsSubvolumeGroupRequestParams.md)
 - [RoutesCephfsSubvolumeRequestParams](docs/RoutesCephfsSubvolumeRequestParams.md)
 - [RoutesCreateCephClusterRequestParams](docs/RoutesCreateCephClusterRequestParams.md)
 - [RoutesCreateHostRequestParams](docs/RoutesCreateHostRequestParams.md)
 - [RoutesCreateOsUserRequestParams](docs/RoutesCreateOsUserRequestParams.md)
 - [RoutesCrushRootRequestParams](docs/RoutesCrushRootRequestParams.md)
 - [RoutesObjectBucketRequestParams](docs/RoutesObjectBucketRequestParams.md)
 - [RoutesObjectStorageClassRequestParams](docs/RoutesObjectStorageClassRequestParams.md)
 - [RoutesObjectStoreRequestParams](docs/RoutesObjectStoreRequestParams.md)
 - [RoutesOsdsDeleteParams](docs/RoutesOsdsDeleteParams.md)
 - [RoutesOsdsRequestParams](docs/RoutesOsdsRequestParams.md)
 - [RoutesPoolCreateRequestParams](docs/RoutesPoolCreateRequestParams.md)
 - [RoutesPoolRequestParams](docs/RoutesPoolRequestParams.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

canzhu@deeproute.ai

