package integration

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestPing(t *testing.T) {
	cmd := newGedisCmd()

	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	time.Sleep(2 * time.Second)

	ping := r.Ping(ctx)

	if ping.Err() != nil {
		t.Errorf("can't ping gedis: %s", ping.Err())
	}

	if ping.Val() != "PONG" {
		t.Errorf("Expected \"PONG\", got %s", ping.Val())
	}

	err := cmd.Process.Signal(syscall.SIGTERM)

	if err != nil {
		t.Error("failed to terminate gedis process")
	}
}

func TestPingMultipleClients(t *testing.T) {
	cmd := newGedisCmd()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rdb2 := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	time.Sleep(2 * time.Second)

	ping := rdb.Ping(ctx)
	ping2 := rdb2.Ping(ctx)

	if ping.Err() != nil {
		t.Errorf("can't ping gedis: %s", ping.Err())
	}

	if ping.Val() != "PONG" {
		t.Errorf("Expected \"PONG\", got %s", ping.Val())
	}

	if ping2.Err() != nil {
		t.Errorf("can't ping gedis: %s", ping.Err())
	}

	if ping2.Val() != "PONG" {
		t.Errorf("Expected \"PONG\", got %s", ping.Val())
	}

	err := cmd.Process.Signal(syscall.SIGTERM)

	if err != nil {
		t.Error("failed to terminate gedis process")
	}
}
