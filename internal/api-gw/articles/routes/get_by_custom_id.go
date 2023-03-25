package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles/dto"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	microErrors "go-micro.dev/v4/errors"
	"net/http"
)

// GetByCustomId godoc
// @Summary     Get an article by custom_id
// @Description Get an article by custom_id
// @Tags        Articles
// @Produce     json
// @Param   	custom_id query    string true "custom_id to get"
// @Success     200       {object} dto.SingleArticleResponse
// @Failure     400       {object} errors.Response
// @Failure     500       {object} errors.Response
// @Failure     502       {object} errors.Response
// @Router      /articles/get_by_custom_id [get]
func (r *Routes) GetByCustomId(ctx *gin.Context) {
	var req dto.ArticleCustomIdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetByCustomId")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	a, err := r.c.GetByCustomId(ctx.Request.Context(), &pb.ArticleCustomIdRequest{CustomId: req.CustomId})
	if err != nil {
		clientError := microErrors.FromError(err)

		errors.ErrorResponse(ctx, int(clientError.Code), clientError.Detail)
		return
	}

	resp := dto.SingleArticleResponse{
		Id:        int(a.Id),
		CustomId:  a.CustomId,
		AuthorId:  int(a.AuthorId),
		Title:     a.Title,
		Thumbnail: a.Thumbnail,
		Content:   a.Content.AsMap(),
		CreatedAt: a.CreatedAt.AsTime(),
	}

	ctx.JSON(http.StatusOK, &resp)
}
