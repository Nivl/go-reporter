package sentryreporter_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dchest/uniuri"

	uuid "github.com/google/uuid"

	"github.com/stretchr/testify/require"

	reporter "github.com/Nivl/go-reporter"
	"github.com/Nivl/go-reporter/implementations/sentryreporter"
)

func TestSentryHappyPath(t *testing.T) {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		t.Skip("sentry not set")
	}

	creator, err := sentryreporter.NewCreator(sentryDSN)
	require.NoError(t, err, "NewCreator() should not have failed")

	r, err := creator.New()
	require.NoError(t, err, "creator.New() should not have failed")

	// Set some data
	r.SetUser(&reporter.User{
		ID:       uniuri.New(),
		Username: uniuri.New(),
		Email:    uniuri.New(),
		IP:       uniuri.New(),
	})
	r.AddTag("endpoint", "TestSentryHappyPath")
	r.AddTag("Req ID", uuid.Must(uuid.NewRandom()).String())

	// send the request
	r.ReportErrorAndWait(fmt.Errorf("New test error %s", uniuri.NewLen(4)))
}

func TestSentryNilUser(t *testing.T) {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		t.Skip("sentry not set")
	}

	creator, err := sentryreporter.NewCreator(sentryDSN)
	require.NoError(t, err, "NewCreator() should not have failed")

	r, err := creator.New()
	require.NoError(t, err, "creator.New() should not have failed")

	// Set some data
	r.SetUser(nil)
	r.AddTag("endpoint", "TestSentryHappyPath")
	r.AddTag("Req ID", uuid.Must(uuid.NewRandom()).String())

	// send the request
	r.ReportErrorAndWait(fmt.Errorf("New test error %s", uniuri.NewLen(4)))
}
