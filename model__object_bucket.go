/*
 * SmdService
 *
 * Resource for managing storage
 *
 * API version: 1.0.0
 * Contact: canzhu@deeproute.ai
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package smdclient

type ObjectBucket struct {
	BucketName string `json:"bucket_name"`
	Name string `json:"name"`
	ObjectStorageClassId int64 `json:"object_storage_class_id"`
	OsUserId int64 `json:"os_user_id"`
}
