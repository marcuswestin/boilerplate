# Commands
##########

setup: install-dependencies
run: run-cockroach run-influxdb
vet:
	go tool vet projectname
test: vet
	go test -p 1 -v projectname/...


# Dependencies
##############
install-dependencies: install-brew install-cockroach install-influxdb

install-brew:
	@$(call install-check, brew, ruby -e "$$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"; brew update && brew doctor && brew upgrade)

# Relational db
install-cockroach:
	@$(call install-check, cockroach, wget https://binaries.cockroachdb.com/cockroach-latest.darwin-10.9-amd64.tgz; tar xfz cockroach-latest.darwin-10.9-amd64.tgz; mv ./cockroach-latest.darwin-10.9-amd64/cockroach /usr/local/bin)
run-cockroach: install-cockroach
	cockroach start --store=/tmp/projectname/cockroachdb/node1 --port=26257 --http-port=8080 --background
	cockroach start --store=/tmp/projectname/cockroachdb/node2 --port=26258 --http-port=8081 --join=localhost:26257 --background
	cockroach start --store=/tmp/projectname/cockroachdb/node3 --port=26259 --http-port=8082 --join=localhost:26257 --background
stop-cockroach:
	cockroach quit --port=26257 || echo "already dead"
	cockroach quit --port=26258 || echo "already dead"
	cockroach quit --port=26259 || echo "already dead"
rm-cockroach: stop-cockroach
	rm -rf /tmp/projectname/cockroachdb

# Timeseries db
install-influxdb:
	@$(call install-check, influx, brew install influxdb)
run-influxdb:
	nohup influxd &
stop-influxdb:
	killall influxd


# Util
######

# Install a program if it isn't already
define install-check
	if ! type$(1) > /dev/null 2>&1; \
		then echo "Installing$(1)"; $(2); \
	fi
endef
# Run a program if it isn't already running
define run-check
	if [ `ps ax | grep "$(1)" | grep -v grep | grep -v "make run" | wc -l` -eq 0 ]; \
		then echo "run $(1)"; $(2) &; sleep 0.1; \
	fi
endef
