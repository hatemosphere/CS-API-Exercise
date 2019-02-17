#!/usr/bin/env bash

set -euo pipefail

export MONGOIMPORT_CSV_FILE_FIELDS="survived.boolean(),passengerClass.int32(),name.string(),sex.string(),age.double(),siblingsOrSpousesAboard.int32(),parentsOrChildrenAboard.int32(),fare.double()"

while ! mongo mongodb:27017/titanic --eval "db.version()" > /dev/null 2>&1; do sleep 1; done

mongoimport --host mongodb --db titanic --collection passengers --columnsHaveTypes \
    --fields ${MONGOIMPORT_CSV_FILE_FIELDS} \
    --type csv --file /mongo-seed/titanic.csv \
    --parseGrace skipRow
