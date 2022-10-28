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

type ModelsCephNfsExport struct {
	CephNFS *ModelsCephNfs `json:"CephNFS"`
	CephNFSID int64 `json:"CephNFSID"`
	Cephfs *ModelsCephfs `json:"Cephfs"`
	CephfsID int64 `json:"CephfsID"`
	CephfsSubvolumeID int64 `json:"CephfsSubvolumeID"`
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	ID int64 `json:"ID"`
	Name string `json:"Name"`
	PseudoPath string `json:"PseudoPath"`
	Status string `json:"Status"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}