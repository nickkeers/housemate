package handlers

import (
    "github.com/gin-gonic/gin"
)

func (h *HandlerContext) HealthcheckHandler(ctx *gin.Context) {
    dbAlive := h.db.Ping()

    ctx.JSON(200, gin.H{
        "database": dbAlive,
    })
}