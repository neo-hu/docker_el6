package daemon

import (

	// therefore they register themselves to the logdriver factory.
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/fluentd"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/gcplogs"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/gelf"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/journald"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/jsonfilelog"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/logentries"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/splunk"
	_ "github.com/neo-hu/docker_el6/docker/daemon/logger/syslog"
)
