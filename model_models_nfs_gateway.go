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

type ModelsNfsGateway struct {
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int64 `json:"ID"`
	IP string `json:"IP"`
	Name string `json:"Name"`
	NodeName string `json:"NodeName"`
	Port int32 `json:"Port"`
	Status string `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
