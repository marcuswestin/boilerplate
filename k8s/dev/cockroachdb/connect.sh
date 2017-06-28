#!/bin/bash

kubectl run cockroachdb -it --image=cockroachdb/cockroach --rm --restart=Never -- sql --insecure --host=cockroachdb-public \
	--execute="`cat $PROJECTNAME_ROOT/db/schema.sql`" \
	--execute="`cat $PROJECTNAME_ROOT/db/test-data.sql`" \

