package util

import (
	"os"
	"strconv"

	"github.com/kayac/go-katsubushi/v2"
)

var (
	idGenerator *katsubushi.App
)

func init() {
	workerID := os.Getenv("WORKER_ID")
	if workerID == "" {
		panic("env variable WORKER_ID is not set")
	}
	i, err := strconv.Atoi(workerID)
	if err != nil {
		panic(err)
	}

	idGenerator, err = katsubushi.New(uint(i))
	if err != nil {
		panic(err)
	}
}

func GenerateUniqueID() (int64, error) {
	id, err := idGenerator.NextID()
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}
