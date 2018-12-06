// common.go
package common

import (
	//"time"
)


type Invoice struct {
	_id int64		`json:"_id,omitempty"`
	Id int64			`json:"id,omitempty"`
	Uid uint32		`json:"uid,omitempty"`
	Utime int64		`json:"utime,omitempty"`
	Client float64	`json:"client,omitempty"`
	Items float64	`json:"items,omitempty"`
	Amount float64	`json:"amount,omitempty"`
}

const A_PORT = 12001
const I_PORT = 12002
const R_PORT = 2003
