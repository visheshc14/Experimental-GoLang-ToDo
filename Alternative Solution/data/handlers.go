package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

// List Function Lists All The Tasks in The List.
func List(ctx context.Context, rdb *redis.Client, filterDone bool) {
	if filterDone {
		list := rdb.LRange(ctx, "done", 0, -1)
		val := list.Val()

		for _, t := range val {
			s := strings.TrimSpace(t)
			fmt.Println("☑️", s)
		}

	} else {
		list := rdb.LRange(ctx, "tasks", 0, -1)
		val := list.Val()

		for _, t := range val {
			s := strings.TrimSpace(t)
			fmt.Println("☑️", s)
		}
	}
}

// Add Function Adds a New Task Into The List.
func Add(ctx context.Context, rdb *redis.Client, t string) {
	rdb.LPush(ctx, "tasks", t)
	fmt.Println("☑️ : The task was added")
}

// Done Function Marks the Given Task For Completion and Moves it Into a New List Called "Done".
func Done(ctx context.Context, rdb *redis.Client, t string) {
	rdb.LRem(ctx, "tasks", 0, t)
	rdb.LPush(ctx, "done", t)
	fmt.Println("☑️ : The task was completed")
}
