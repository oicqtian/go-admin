package model

import (
	"time"
)

type BaseModel struct {
	ID int64 `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`

}


type PageModel struct {
	CurrPage int  `json:"currPage"`
	List interface{} `json:"list"`
	PageSize int `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	TotalPage int `json:"totalPage"`
}