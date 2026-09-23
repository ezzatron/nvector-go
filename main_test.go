package nvector_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	if os.Getenv("CI") != "" {
		return m.Run()
	}

	ctx := context.Background()

	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "ghcr.io/ezzatron/nvector-test-api",
				ExposedPorts: []string{"8000/tcp"},
				WaitingFor:   wait.ForListeningPort("8000/tcp"),
			},
			Started: true,
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to start test API container:", err)

		return 1
	}
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			fmt.Fprintln(os.Stderr, "failed to terminate test API container:", err)
		}
	}()

	host, err := container.Host(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get test API container host:", err)

		return 1
	}

	port, err := container.MappedPort(ctx, "8000/tcp")
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get test API container port:", err)

		return 1
	}

	os.Setenv("TEST_API_URL", fmt.Sprintf("ws://%s:%s/", host, port.Port()))

	return m.Run()
}
