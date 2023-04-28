package middleware

import (
	"context"
	"github.com/shuryak/shuryak-blog/pkg/constants"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/user"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"strconv"
	"strings"
)

type AuthWrapper struct {
	c client.Client
	l logger.Interface
}

func NewAuthWrapper(c client.Client, l logger.Interface) *AuthWrapper {
	return &AuthWrapper{c, l}
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

		md, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.Unauthorized("context", "no token")
		}

		auth, ok := md.Get(constants.AuthMetadataName)
		if !ok {
			return errors.Unauthorized("metadata", "no token")
		}
		authSplit := strings.Split(auth, "Bearer ")

		if len(authSplit) != 2 {
			return errors.Unauthorized("length", "invalid bearer payload")
		}

		token := strings.TrimSpace(authSplit[1])

		uc := pb.NewUserService("user", w.c)
		outToken, err := uc.Validate(ctx, &pb.ValidateRequest{AccessToken: token})
		if err != nil {
			return errors.Unauthorized("token", "invalid token structure") // TODO: inner errors from user service
		}
		if outToken.IsValid != true {
			return errors.Unauthorized("expired", "token is expired")
		}

		newCtx := metadata.Set(ctx, constants.UsernameMetadataName, outToken.Username)
		newCtx2 := metadata.Set(newCtx, constants.UserIdMetadataName, strconv.Itoa(int(outToken.UserId))) // TODO: one newCtx

		return fn(newCtx2, req, resp)
	}
}
