package codewars_api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"

	"github.com/yasonofriychuk/service-codewars/internal/codewars_api/models"
)

const (
	basePath = "https://www.codewars.com/api/v1"
)

type Api struct {
	client HttpClient
}

func New(client HttpClient) *Api {
	return &Api{
		client: client,
	}
}

func (a *Api) GetUser(ctx context.Context, user string) (*models.GetUserResponse, error) {
	path := fmt.Sprintf("%s/users/%s", basePath, url.PathEscape(user))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	return doTypedRequest[models.GetUserResponse](ctx, a.client, req)
}

func (a *Api) GetCodeChallenge(ctx context.Context, challenge string) (*models.GetCodeChallengeResponse, error) {
	path := fmt.Sprintf("%s/code-challenges/%s", basePath, url.PathEscape(challenge))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	return doTypedRequest[models.GetCodeChallengeResponse](ctx, a.client, req)
}

func (a *Api) ListCompletedChallenges(ctx context.Context, user string, page int) (*models.ListCompletedChallengesResponse, error) {
	path := fmt.Sprintf("%s/users/%s/code-challenges/completed?page=%d", basePath, url.PathEscape(user), page)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	return doTypedRequest[models.ListCompletedChallengesResponse](ctx, a.client, req)
}

func (a *Api) ListAuthoredChallenges(ctx context.Context, user string) (*models.ListAuthoredChallengesResponse, error) {
	path := fmt.Sprintf("%s/users/%s/code-challenges/authored", basePath, url.PathEscape(user))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	return doTypedRequest[models.ListAuthoredChallengesResponse](ctx, a.client, req)
}

func doTypedRequest[T any](ctx context.Context, client HttpClient, req *http.Request) (*T, error) {
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("DoRequst: %w", err)
	}

	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("empty response")
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read resp body: %w", err)
	}

	var response T
	if err := json.Unmarshal(b, &response); err != nil {
		return nil, fmt.Errorf("unmarshal resp body: %w", err)
	}

	return &response, nil
}
