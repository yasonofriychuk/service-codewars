package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yasonofriychuk/service-codewars/internal/codewars_api"
	"github.com/yasonofriychuk/service-codewars/internal/retry_client"
)

func main() {
	ctx := context.Background()

	api := codewars_api.New(retry_client.New(http.DefaultClient))

	fmt.Println(api.ListCompletedChallenges(ctx, "some_user", 0))
}
