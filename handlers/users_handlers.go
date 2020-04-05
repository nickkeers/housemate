package handlers

import (
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

func (h *HandlerContext) GetUserByIdHandler(ctx *gin.Context) {
    id := ctx.Param("id")

    idInt, err := strconv.ParseInt(id, 10, 64)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Provided ID was not an integer",
        })
        return
    }

    user, err := h.db.GetHouseholdMemberById(int(idInt))

    if err != nil {
        log.Error(err)
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "could not fetch household member",
        })
    } else {
        ctx.JSON(http.StatusOK, user)
    }
}
