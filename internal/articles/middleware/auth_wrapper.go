package middleware

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/articles/handlers"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/proto/user"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"strconv"
	"strings"
)

type Wrappers struct {
	userServerName string
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		// Health checks
		if req.Endpoint() == "Health.Check" || req.Endpoint() == "Health.Watch" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoMetadata(req.Endpoint()))
		}

		auth, ok := meta[constants.AuthMetadataName]
		if !ok {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoToken())
		}
		authSplit := strings.Split(auth, "Bearer ")

		if len(authSplit) != 2 {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoToken())
		}

		token := strings.TrimSpace(authSplit[1])

		userClient := user.NewUserService("user", client.DefaultClient)

		outToken, err := userClient.Validate(ctx, &user.ValidateRequest{AccessToken: token})
		if err != nil {
			return err
		}
		if outToken.IsValid != true {
			return fmt.Errorf(handlers.GlobalErrors.AuthInvalidToken())
		}

		newCtx := metadata.Set(ctx, constants.UsernameMetadataName, outToken.Username)
		newCtx2 := metadata.Set(newCtx, constants.UserIdMetadataName, strconv.Itoa(int(outToken.UserId))) // TODO: one newCtx

		return fn(newCtx2, req, resp)
	}
}
