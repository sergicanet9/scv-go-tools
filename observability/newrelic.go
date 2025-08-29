package observability

import (
	"log"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/event"
)

// SetupNewRelic configures the application and configures the singleton logger to forward logs to New Relic
func SetupNewRelic(appName, newrelicKey string) (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(newrelicKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return app, err
	}
	if err := app.WaitForConnection(5 * time.Second); err != nil {
		return app, err
	}

	writer := logWriter.New(os.Stdout, app)
	logger = log.New(&writer, "", log.Default().Flags())
	return app, nil
}

// NewRelicMongoMonitor returns a new new relic monitor for MongoDB
func NewRelicMongoMonitor() *event.CommandMonitor {
	return nrmongo.NewCommandMonitor(nil)
}
