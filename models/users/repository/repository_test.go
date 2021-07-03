package repository_test

import (
	"context"
	"database/sql"
	"net"
	"net/url"
	"runtime"
	"testing"
	"time"

	"github.com/fanfit/user-service/models/users/repository"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

func startDatabase(tb testing.TB) string {
	tb.Helper()

	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("myuser", "mypass"),
		Path:   "mydatabase",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf("Could not connect to docker: %v", err)
	}

	pw, _ := pgURL.User.Password()
	env := []string{
		"POSTGRES_USER=" + pgURL.User.Username(),
		"POSTGRES_PASSWORD=" + pw,
		"POSTGRES_DB=" + pgURL.Path,
	}

	resource, err := pool.Run("postgres", "13-alpine", env)
	if err != nil {
		tb.Fatalf("Could not start postgres container: %v", err)
	}
	tb.Cleanup(func() {
		err = pool.Purge(resource)
		if err != nil {
			tb.Fatalf("Could not purge container: %v", err)
		}
	})

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	// Docker layer network is different on Mac
	if runtime.GOOS == "darwin" {
		pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	logWaiter, err := pool.Client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container: resource.Container.ID,
		Stderr:    true,
		Stdout:    true,
		Stream:    true,
	})
	if err != nil {
		tb.Fatalf("Could not connect to postgres container log output: %v", err)
	}

	tb.Cleanup(func() {
		err = logWaiter.Close()
		if err != nil {
			tb.Fatalf("Could not close container log: %v", err)
		}
		err = logWaiter.Wait()
		if err != nil {
			tb.Fatalf("Could not wait for container log to close: %v", err)
		}
	})

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() (err error) {
		db, err := sql.Open("pgx", pgURL.String())
		if err != nil {
			return err
		}
		defer func() {
			cerr := db.Close()
			if err == nil {
				err = cerr
			}
		}()

		return db.Ping()
	})
	if err != nil {
		tb.Fatalf("Could not connect to postgres container: %v", err)
	}
	return pgURL.String()
}

func TestAddUsers(t *testing.T) {
	t.Parallel()

	userStore, err := repository.NewUserStore(startDatabase(t))
	if err != nil {
		t.Fatalf("Failed to create a new directory: %s", err)
	}
	t.Cleanup(func() {
		err = userStore.Close()
		if err != nil {
			t.Errorf("Failed to close directory: %s", err)
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	t.Run("Add and retrieve a user ", func(t *testing.T) {
		t.Parallel()
		addedUser, err := userStore.Create(ctx, repository.User{
			FirstName:  "test-name",
			Email:      "test-email",
			UserTypeID: 1,
		})
		if err != nil {
			t.Fatalf("Failed to add users: %s", err)
		}

		user, err := userStore.GetByEmail(ctx, "test-email")
		if err != nil {
			t.Fatalf("Failed to list users: %s", err)
		}

		assert.Equal(t, user.Email, addedUser.Email)
	})
}
