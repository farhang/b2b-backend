package test

import (
	"context"
	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresConfig struct {
	User     string
	Password string
	DB       string
	Port     nat.Port
}

type PostgresContainer struct {
	tc.Container
	IpHost string
	Port   string
}

func (pc *PostgresContainer) Start(ctx context.Context, psConfig PostgresConfig, networkName string) (*PostgresContainer, error) {
	postgresPort := nat.Port("5432/tcp")
	postgres, err := tc.GenericContainer(context.Background(),
		tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				Image:          "postgres:14.3-alpine",
				ExposedPorts:   []string{psConfig.Port.Port()},
				Networks:       []string{networkName},
				NetworkAliases: map[string][]string{networkName: {"postgres"}},
				Env: map[string]string{
					"POSTGRES_USER":     psConfig.User,
					"POSTGRES_PASSWORD": psConfig.Password,
					"POSTGRES_DB":       psConfig.DB,
				},
				WaitingFor: wait.ForAll(
					wait.ForLog("database system is ready to accept connections"),
					wait.ForListeningPort(psConfig.Port),
				),
			},
			Started: true,
		})
	ip, err := postgres.Host(ctx)
	if err != nil {
		return nil, err
	}
	mappedPort, err := postgres.MappedPort(ctx, postgresPort)

	return &PostgresContainer{
		IpHost: ip,
		Port:   mappedPort.Port(),
	}, err
}
func NewPostgresContainer() *PostgresContainer {
	return &PostgresContainer{}
}
