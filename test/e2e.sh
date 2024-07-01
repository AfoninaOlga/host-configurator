#!/bin/bash

set -e

BASEDIR=$(realpath "$(dirname "$0")")
ROOTDIR=$(realpath "$BASEDIR/..")
CLIENTDIR="$ROOTDIR/client"
SERVICEDIR="$ROOTDIR/service"

BLUE='\033[0;34m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

RUNNING="${BLUE}[RUNNING]${NC}"
PASSED="${GREEN}[PASSED]${NC} "

function start_service() {
	pushd "$SERVICEDIR"
	docker compose up -d --build
	popd
}

function finish_service() {
	pushd "$SERVICEDIR"
	docker compose down
	popd
}

trap finish_service EXIT

start_service

pushd "$CLIENTDIR"
echo Build client
make

# Wait service up
while [[ -z $(./configurator get-hostname | grep "Hostname is") ]]; do sleep 2; done

echo -e "${RUNNING} Get hostname test"
./configurator get-hostname | grep "Hostname is"
echo -e "${PASSED}  Get hostname test"

echo -e "${RUNNING} Set hostname test"
NEWHOSTNAME="my-pc"
./configurator set-hostname "$NEWHOSTNAME" | grep "Hostname is set to '${NEWHOSTNAME}'"
./configurator get-hostname | grep "Hostname is '${NEWHOSTNAME}'"
echo -e "${PASSED} Set hostname test"

echo -e "${RUNNING} Set invalid hostname test"
INVALIDHOSTNAME="my=pc"
./configurator set-hostname "$INVALIDHOSTNAME" | grep "Couldn't update hostname. Reason: error changing hostname: invalid hostname"
./configurator get-hostname | grep "Hostname is '${NEWHOSTNAME}'"
echo -e "${PASSED} Set invalid hostname test"

echo -e "${RUNNING} Get DNS servers list test"
./configurator dns-servers-list | grep "DNS servers:"
echo -e "${PASSED} Get DNS servers list test"

echo -e "${RUNNING} Add DNS server test"
NEWDNSSERVER="1.1.1.1"
./configurator dns-servers-add "$NEWDNSSERVER" | grep "Server ${NEWDNSSERVER} added"
./configurator dns-servers-list | grep "$NEWDNSSERVER"
echo -e "${PASSED} Add DNS server test"

echo -e "${RUNNING} Add invalid DNS server test"
INVALIDDNSSERVER="1.1.1.1.1"
./configurator dns-servers-add "$INVALIDDNSSERVER" | grep "Couldn't add server. Reason: error getting servers: invalid DNS server"
echo -e "${PASSED} Add invalid DNS server test"

echo -e "${RUNNING} Remove DNS server test"
./configurator dns-servers-delete "$NEWDNSSERVER" | grep "Server ${NEWDNSSERVER} is deleted"
! ./configurator dns-servers-list | grep "$NEWDNSSERVER"
echo -e "${PASSED} Remove DNS server test"

echo -e "${GREEN}All tests passed${NC}"

popd

