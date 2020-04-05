package handlers

import (
    "github.com/nickkeers/housemate/datalayer"
)

type HandlerContext struct {
    db *datalayer.PostgresDataAdapter
}

func NewHandlerContext(adapter *datalayer.PostgresDataAdapter) *HandlerContext {
    return &HandlerContext{db: adapter}
}