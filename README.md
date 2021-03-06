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
*ActionLogsApi* | [**GetActionLog**](docs/ActionLogsApi.md#getactionlog) | **Get** /action_logs/{log-id} | get a action log
*ActionLogsApi* | [**ListActionLogs**](docs/ActionLogsApi.md#listactionlogs) | **Get** /action_logs | get all action logs
*CephClustersApi* | [**CreateCephCluster**](docs/CephClustersApi.md#createcephcluster) | **Post** /ceph-clusters | create ceph cluster
*CephClustersApi* | [**DeleteCephCluster**](docs/CephClustersApi.md#deletecephcluster) | **Delete** /ceph-clusters | delete ceph cluster
*CephClustersApi* | [**GetCephCluster**](docs/CephClustersApi.md#getcephcluster) | **Get** /ceph-clusters | get ceph cluster
*CrushRootsApi* | [**CreateCrushRoot**](docs/CrushRootsApi.md#createcrushroot) | **Post** /crush-roots | create crush root
*CrushRootsApi* | [**DeleteCrushRoot**](docs/CrushRootsApi.md#deletecrushroot) | **Delete** /crush-roots/{crush-root-id} | delete a crush root
*CrushRootsApi* | [**GetCrushRoot**](docs/CrushRootsApi.md#getcrushroot) | **Get** /crush-roots/{crush-root-id} | get a crush root
*CrushRootsApi* | [**ListCrushRoots**](docs/CrushRootsApi.md#listcrushroots) | **Get** /crush-roots | get all crush roots
*DisksApi* | [**GetDisk**](docs/DisksApi.md#getdisk) | **Get** /disks/{disk-id} | get a disk
*DisksApi* | [**ListDisks**](docs/DisksApi.md#listdisks) | **Get** /disks | get all disks
*HostsApi* | [**CreateHost**](docs/HostsApi.md#createhost) | **Post** /hosts | create a host
*HostsApi* | [**DeleteHost**](docs/HostsApi.md#deletehost) | **Delete** /hosts/{host-id} | delete a host
*HostsApi* | [**GetHost**](docs/HostsApi.md#gethost) | **Get** /hosts/{host-id} | get a host
*HostsApi* | [**ListHosts**](docs/HostsApi.md#listhosts) | **Get** /hosts | get all hosts
*ObjectBucketsApi* | [**CreateObjectBucket**](docs/ObjectBucketsApi.md#createobjectbucket) | **Post** /object-buckets | create an object bucket
*ObjectBucketsApi* | [**DeleteObjectBucket**](docs/ObjectBucketsApi.md#deleteobjectbucket) | **Delete** /object-buckets/{object-bucket-id} | delete an object bucket
*ObjectBucketsApi* | [**GetObjectBucket**](docs/ObjectBucketsApi.md#getobjectbucket) | **Get** /object-buckets/{object-bucket-id} | get an object bucket
*ObjectBucketsApi* | [**ListObjectBuckets**](docs/ObjectBucketsApi.md#listobjectbuckets) | **Get** /object-buckets | get all object buckets
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
*OsdsApi* | [**DeleteOsds**](docs/OsdsApi.md#deleteosds) | **Delete** /osds | delete osds
*OsdsApi* | [**GetOsd**](docs/OsdsApi.md#getosd) | **Get** /osds/{osd-id} | get a osd
*OsdsApi* | [**ListOsds**](docs/OsdsApi.md#listosds) | **Get** /osds | get all osds


## Documentation For Models

 - [CephCluster](docs/CephCluster.md)
 - [CrushRoot](docs/CrushRoot.md)
 - [GormModel](docs/GormModel.md)
 - [Host](docs/Host.md)
 - [ModelsActionLog](docs/ModelsActionLog.md)
 - [ModelsCephCluster](docs/ModelsCephCluster.md)
 - [ModelsCrushRoot](docs/ModelsCrushRoot.md)
 - [ModelsDisk](docs/ModelsDisk.md)
 - [ModelsHost](docs/ModelsHost.md)
 - [ModelsObjectBucket](docs/ModelsObjectBucket.md)
 - [ModelsObjectStorageClass](docs/ModelsObjectStorageClass.md)
 - [ModelsObjectStore](docs/ModelsObjectStore.md)
 - [ModelsOsUser](docs/ModelsOsUser.md)
 - [ModelsOsd](docs/ModelsOsd.md)
 - [ModelsPoolProperty](docs/ModelsPoolProperty.md)
 - [ObjectBucket](docs/ObjectBucket.md)
 - [ObjectStorageClass](docs/ObjectStorageClass.md)
 - [ObjectStore](docs/ObjectStore.md)
 - [OsUser](docs/OsUser.md)
 - [Osds](docs/Osds.md)
 - [RoutesCreateCephClusterRequestParams](docs/RoutesCreateCephClusterRequestParams.md)
 - [RoutesCreateHostRequestParams](docs/RoutesCreateHostRequestParams.md)
 - [RoutesCreateOsUserRequestParams](docs/RoutesCreateOsUserRequestParams.md)
 - [RoutesCrushRootRequestParams](docs/RoutesCrushRootRequestParams.md)
 - [RoutesObjectBucketRequestParams](docs/RoutesObjectBucketRequestParams.md)
 - [RoutesObjectStorageClassRequestParams](docs/RoutesObjectStorageClassRequestParams.md)
 - [RoutesObjectStoreRequestParams](docs/RoutesObjectStoreRequestParams.md)
 - [RoutesOsdsRequestParams](docs/RoutesOsdsRequestParams.md)
 - [RoutesPoolRequestParams](docs/RoutesPoolRequestParams.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

canzhu@deeproute.ai

