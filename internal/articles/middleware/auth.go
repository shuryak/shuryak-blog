package middleware

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/articles/handlers"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"go-micro.dev/v4/client"
	microLogger "go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"strconv"
	"strings"
)

type AuthWrapper struct {
	с client.Client
	l logger.Interface
}

func NewAuthWrapper(с client.Client, l logger.Interface) *AuthWrapper {
	return &AuthWrapper{с, l}
}

var authEndpoints = map[string]struct{}{
	"Articles.Create": {},
	"Articles.Update": {},
	"Articles.Delete": {},
}

func (w *AuthWrapper) Use(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if _, isAuthNeeded := authEndpoints[req.Endpoint()]; !isAuthNeeded {
			return fn(ctx, req, resp)
		}

		microLogger.Info("Im here")

		md, ok := metadata.FromContext(ctx)
		if !ok {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoMetadata(req.Endpoint()))
		}

		auth, ok := md.Get(constants.AuthMetadataName)
		if !ok {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoToken())
		}
		authSplit := strings.Split(auth, "Bearer ")

		if len(authSplit) != 2 {
			return fmt.Errorf(handlers.GlobalErrors.AuthNoToken())
		}

		token := strings.TrimSpace(authSplit[1])

		uc := pb.NewUserService("user", w.с)
		outToken, err := uc.Validate(ctx, &pb.ValidateRequest{AccessToken: token})
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
