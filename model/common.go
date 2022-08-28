package model

import "time"

type MetaInfo struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
