package model

import "time"

//Common Meta Info used for every response
type MetaInfo struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
