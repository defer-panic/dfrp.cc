package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/defer-panic/url-shortener-api/internal/config"
	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/google/go-github/v48/github"
)

type Storage interface {
	CreateOrUpdate(context.Context, model.User) (*model.User, error)
	Update(context.Context, model.User) error
	GetByGitHubLogin(context.Context, string) (*model.User, error)
	Deactivate(context.Context, string) error
}

//go:generate moq --out=mock_github_client.gen.go --pkg=auth . GitHubClient
type GitHubClient interface {
	ExchangeCodeToAccessKey(ctx context.Context, clientID, clientSecret, code string) (string, error)
	IsMember(ctx context.Context, accessKey, org, user string) (bool, error)
	GetUser(ctx context.Context, accessKey, user string) (*github.User, error)
}

type Service struct {
	github  GitHubClient
	storage Storage

	ghClientID     string
	ghClientSecret string
}

func NewService(githubClient GitHubClient, storage Storage, ghClientID, ghClientSecret string) *Service {
	return &Service{
		storage:        storage,
		github:         githubClient,
		ghClientID:     ghClientID,
		ghClientSecret: ghClientSecret,
	}
}

func (s *Service) GitHubAuthLink() string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?scopes=user,read:org&client_id=%s", s.ghClientID)
}

func (s *Service) GitHubAuthCallback(ctx context.Context, sessionCode string) (*model.User, string, error) {
	accessKey, err := s.github.ExchangeCodeToAccessKey(ctx, s.ghClientID, s.ghClientSecret, sessionCode)
	if err != nil {
		return nil, "", err
	}

	ghUser, err := s.github.GetUser(ctx, accessKey, "")
	if err != nil {
		return nil, "", err
	}

	user, err := s.RegisterUser(ctx, ghUser, accessKey)
	if err != nil {
		return nil, "", err
	}

	jwt, err := MakeJWT(*user)
	if err != nil {
		log.Printf("failed to make jwt: %v", err)
		return nil, "", err
	}

	return user, jwt, nil
}

func (s *Service) RegisterUser(ctx context.Context, ghUser *github.User, accessKey string) (*model.User, error) {
	isMember, err := s.github.IsMember(ctx, accessKey, config.Get().Auth.AllowedGitHubOrg, ghUser.GetLogin())
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, fmt.Errorf("%w %q", model.ErrUserIsNotMember, config.Get().Auth.AllowedGitHubOrg)
	}

	user := model.User{
		GitHubLogin:     ghUser.GetLogin(),
		IsActive:        true,
		GitHubAccessKey: accessKey,
	}

	return s.storage.CreateOrUpdate(ctx, user)
}
