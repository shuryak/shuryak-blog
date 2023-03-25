package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/shuryak-blog/internal/user/entity"
	"github.com/shuryak/shuryak-blog/internal/user/usecase"
	"github.com/shuryak/shuryak-blog/pkg/errors"
	"github.com/shuryak/shuryak-blog/pkg/logger"
	pb "github.com/shuryak/shuryak-blog/proto/user"
)

type UserHandler struct {
	users       usecase.Users
	sessions    usecase.UserSessions
	jwt         JWTManager
	serviceName string // TODO: ?
	l           logger.Interface
}

func NewUsersHandler(u usecase.Users, us usecase.UserSessions, jwt JWTManager, name string, l logger.Interface) *UserHandler {
	return &UserHandler{u, us, jwt, name, l}
}

// Check for implementation
var _ pb.UserHandler = (*UserHandler)(nil)

var globalErrors errors.ServerError

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest, resp *pb.RegisterResponse) error {
	// TODO: hash password in usecase
	user, err := entity.NewUser(req.GetUsername(), req.GetPlainPassword(), req.GetRole())
	if err != nil {
		h.l.Error(err)
		return err // TODO: global errors
	}

	storedUser, err := h.users.Create(ctx, *user) // TODO: handle already exists
	if err != nil {
		h.l.Error(err)
		return err
	}

	resp.Id = storedUser.Id
	resp.Username = storedUser.Username
	resp.Role = storedUser.Role
	return nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest, resp *pb.TokenPairResponse) error {
	user, err := h.users.GetByUsername(ctx, req.GetUsername())
	// TODO: global errors
	if err != nil {
		return fmt.Errorf("cannot find user: %w", err)
	}

	if !user.IsPasswordCorrect(req.GetPlainPassword()) {
		return fmt.Errorf("incorrect username or password")
	}

	session, err := h.sessions.Add(ctx, user.Id)
	if err != nil {
		return err
	}

	accessToken, err := h.jwt.Generate(user)
	if err != nil {
		return fmt.Errorf("cannot generate access token")
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = session.RefreshToken
	return nil
}

func (h *UserHandler) RefreshSession(ctx context.Context, req *pb.RefreshSessionRequest, resp *pb.TokenPairResponse) error {
	user, err := h.users.GetByUsername(ctx, req.GetUsername())

	// TODO: global errors
	if err != nil {
		return fmt.Errorf("cannot find user: %w", err)
	}

	us, err := h.sessions.Refresh(ctx, user.Id) // TODO: check refresh token in usecase
	if err != nil {
		return err
	}

	accessToken, err := h.jwt.Generate(user)
	if err != nil {
		return fmt.Errorf("cannot generate access token")
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = us.RefreshToken
	return nil
}

func (h *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest, resp *pb.ValidateResponse) error {
	claims, err := h.jwt.Decode(req.AccessToken)
	if err != nil {
		return err // TODO: check Decode internal errors
	}
	if claims == nil {
		return fmt.Errorf(globalErrors.AuthInvalidClaim(h.serviceName))
	}
	if claims.Username == "" { // TODO: check issuer
		return fmt.Errorf(globalErrors.AuthInvalidClaim(h.serviceName))
	}

	user, err := h.users.GetByUsername(ctx, claims.Username)
	if err != nil {
		return fmt.Errorf(globalErrors.AuthNoUserInToken(err))
	}

	resp.UserId = user.Id
	resp.Username = user.Username
	resp.IsValid = true

	h.l.Info(fmt.Sprintf("Response: %v", *resp))

	return nil
}
