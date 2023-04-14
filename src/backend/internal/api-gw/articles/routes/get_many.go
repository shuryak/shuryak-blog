package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuryak/shuryak-blog/internal/api-gw/articles/dto"
	"github.com/shuryak/shuryak-blog/internal/api-gw/errors"
	pb "github.com/shuryak/shuryak-blog/proto/articles"
	microErrors "go-micro.dev/v4/errors"
	"net/http"
)

// GetMany godoc
// @Summary     Get an articles by range
// @Description Get an articles by range
// @Tags        Articles
// @Produce     json
// @Param   	offset query    int true "offset to get"
// @Param   	count  query    int true "count to get"
// @Success     200    {object} dto.ManyArticlesResponse
// @Failure     400    {object} errors.Response
// @Failure     500    {object} errors.Response
// @Failure     502    {object} errors.Response
// @Router      /articles/get_many [get]
func (r *Routes) GetMany(ctx *gin.Context) {
	var req dto.GetManyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		r.l.Error(err, "articles - routes - GetMany")
		errors.ValidationErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	articles, err := r.c.GetMany(ctx.Request.Context(), &pb.GetManyRequest{
		Offset:   *req.Offset,
		Count:    *req.Count,
		IsDrafts: *req.IsDrafts,
	})
	if err != nil {
		clientError := microErrors.FromError(err)

		errors.ErrorResponse(ctx, int(clientError.Code), clientError.Detail)
		return
	}

	var resp dto.ManyArticlesResponse
	for _, a := range articles.Articles {
		resp.Articles = append(resp.Articles, &dto.ShortArticleResponse{
			Id:           int(a.Id),
			CustomId:     a.CustomId,
			AuthorId:     int(a.AuthorId),
			Title:        a.Title,
			Thumbnail:    a.Thumbnail,
			ShortContent: a.ShortContent,
			IsDraft:      a.IsDraft,
			CreatedAt:    a.CreatedAt.AsTime(),
			UpdatedAt:    a.UpdatedAt.AsTime(),
		})
	}

	ctx.JSON(http.StatusOK, &resp)
}
