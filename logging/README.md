# logging

Provides some small utilties around initializing logging and setting the log
level. It sets it up with a text formatter, a logging level, and a top-level Log
instance that can be used with different logrus fields in difference contexts.

## Examples

Initializing logging:

```go
import "github.com/cyverse-de/go-mod/logging"
import "github.com/sirupsen/logrus"

var log = logging.Log.WithFields(logrus.Fields{"service":"example"})

func main() {
	var (
		err    error
		logLevel      = flag.String("log-level", "info", "One of trace, debug, info, warn, error, fatal, or panic.")
	)

	flag.Parse()
	logging.SetupLogging(*logLevel)
	...

```

Adding fields to the logger in a function.

```go

package example

import "github.com/cyverse-de/go-mod/logging"

var log = logging.Log.WithFields(logrus.Fields{"service":"example-service", "package":"example"})

func ExportedExample() error {
	var log = log.WithFields(logrus.Fields{"function":"ExportedExample"})

	log.Info("this is an example")

	// Adds new fields to the logging output, so now the 'service', 'package',
	// 'function', and 'silly' fields will be added to logging messages.
	log = log.WithFields(logrus.Fields{"silly":"example"})

	log.Info("this is another example")

	return nil
}


func ExportedExample2() error {
	var log = log.WithFields("more":"sillyness")

	// This will have the the 'service', 'package', and 'more' fields, but not
	// the 'function' and 'silly' fields set in the other function.
	log.Info("from example 2")

}

```
