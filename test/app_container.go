package test

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"path/filepath"
)

type AppContainer struct {
	tc.Container
	URL string
}

type AppConfig struct {
	Port nat.Port
}

func (ap *AppContainer) Start(ctx context.Context, appConfig AppConfig, psgConfig PostgresConfig, networkName string) (*AppContainer, error) {
	dir, _ := os.Getwd()
	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{Context: filepath.Join(dir, "./.."), PrintBuildLog: true},
			Networks:       []string{networkName},
			NetworkAliases: map[string][]string{networkName: {"app"}},
			Env: map[string]string{
				"DB_HOST":        "postgres",
				"DB_DRIVER":      "postgres",
				"DB_USER":        psgConfig.User,
				"DB_PASSWORD":    psgConfig.Password,
				"DB_NAME":        psgConfig.DB,
				"DB_PORT":        psgConfig.Port.Port(),
				"DB_SSLMODE":     "disable",
				"Time_Zone":      "Asia/Tehran",
				"DB_LOG":         "true",
				"JTW_SECRET_KEY": "44f205c828243e8da7ccf7ac883c6c26c81dbaf825cb8c918d91e69afade874e",
			},
			ExposedPorts: []string{appConfig.Port.Port()},
			WaitingFor: wait.ForAll(
				wait.ForLog("http server started"),
			),
		},
		Started: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	ip, _ := container.Host(ctx)

	if err != nil {
		log.Fatal(err)
	}

	mappedPort, _ := container.MappedPort(ctx, appConfig.Port)
	fmt.Println("mappedPort", mappedPort)
	return &AppContainer{
		URL: fmt.Sprintf("http://%s:%s", ip, mappedPort.Port()),
	}, nil
}

func NewAppContainer() *AppContainer {
	return &AppContainer{}
}
