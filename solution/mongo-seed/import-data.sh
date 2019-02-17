#!/usr/bin/env bash

set -euo pipefail

export MONGOIMPORT_CSV_FILE_FIELDS="survived.boolean(),passengerClass.int32(),name.string(),sex.string(),age.double(),siblingsOrSpousesAboard.int32(),parentsOrChildrenAboard.int32(),fare.double()"
export MONGODB_URL=${MONGODB_URL:-"mongodb://localhost:27017/titanic"}

while ! mongo ${MONGODB_URL} --eval "db.version()" > /dev/null 2>&1; do sleep 1; done

mongoimport --uri ${MONGODB_URL} --collection passengers --columnsHaveTypes \
    --fields ${MONGOIMPORT_CSV_FILE_FIELDS} \
    --type csv --file /mongo-seed/titanic.csv \
    --parseGrace skipRow
