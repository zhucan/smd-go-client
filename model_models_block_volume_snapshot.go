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

type ModelsBlockVolumeSnapshot struct {
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int64 `json:"ID"`
	Name string `json:"Name"`
	Protected bool `json:"Protected"`
	Size int64 `json:"Size"`
	SnapID int64 `json:"SnapID"`
	SnapName string `json:"SnapName"`
	Status string `json:"Status"`
	Timestamp string `json:"Timestamp"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Volume *ModelsVolume `json:"Volume"`
	VolumeID int64 `json:"VolumeID"`
}
