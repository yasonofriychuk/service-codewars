package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yasonofriychuk/service-codewars/internal/codewars_api"
	"github.com/yasonofriychuk/service-codewars/internal/retry_client"
)

func main() {
	ctx := context.Background()

	api := codewars_api.New(retry_client.New(http.DefaultClient))

	parser := codewars_api.New(
		codewars_api.KataURL,
		codewars_api.AuthoredURL,
		codewars_api.RanksURL,
		codewars_api.LeadersURL,
	)

	names, err := parser.GetAllUniqueNames(ctx)
	if err != nil {
		log.Fatalf("Error parsing names from leaders: %v", err)
	}
	fmt.Println("Уникальные ники из всех таблиц:", names)

	fmt.Println(api.ListCompletedChallenges(ctx, "some_user", 0))
}
