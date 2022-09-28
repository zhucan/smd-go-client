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

type Disks struct {
	ClusterId int64 `json:"cluster_id"`
	DiskIds []int64 `json:"disk_ids"`
	MetadataDiskIds []int64 `json:"metadata_disk_ids"`
}
