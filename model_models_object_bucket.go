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

import (
	"time"
)

type ModelsObjectBucket struct {
	BucketName string `json:"BucketName"`
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	Host string `json:"Host"`
	ID int64 `json:"ID"`
	MaxObjects int64 `json:"MaxObjects"`
	MaxSize string `json:"MaxSize"`
	Name string `json:"Name"`
	NumObjects int64 `json:"NumObjects"`
	ObjectStorageClass *ModelsObjectStorageClass `json:"ObjectStorageClass"`
	ObjectStorageClassID int64 `json:"ObjectStorageClassID"`
	OsUser *ModelsOsUser `json:"OsUser"`
	OsUserID int64 `json:"OsUserID"`
	Passive bool `json:"Passive"`
	Port string `json:"Port"`
	Region string `json:"Region"`
	SizeUtilized int64 `json:"SizeUtilized"`
	Status string `json:"Status"`
	SubRegion string `json:"SubRegion"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
