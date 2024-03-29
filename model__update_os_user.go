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

type UpdateOsUser struct {
	AccessKey string `json:"access_key"`
	BucketCapability string `json:"bucket_capability"`
	MaxBuckets int32 `json:"max_buckets"`
	MaxObjects int64 `json:"max_objects"`
	MaxSize string `json:"max_size"`
	MetadataCapability string `json:"metadata_capability"`
	Name string `json:"name"`
	SecretKey string `json:"secret_key"`
	UsageCapability string `json:"usage_capability"`
	UserCapability string `json:"user_capability"`
	ZoneCapability string `json:"zone_capability"`
}
