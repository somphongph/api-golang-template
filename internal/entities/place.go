package entities

import "github.com/somphongph/lib-golang-packages/xentities"

type Place struct {
	xentities.MongoBaseId `json:",inline" bson:",inline"`

	Name string `json:"name" bson:"name"`

	xentities.MongoBaseRecord `json:",inline" bson:",inline"`
}
