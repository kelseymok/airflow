#!/usr/bin/env bash

set -e
set -o nounset
set -o pipefail

script_dir=$(cd "$(dirname "$0")" ; pwd -P)

goal_run() {
  pushd "${script_dir}" > /dev/null
    curl -LfO 'https://airflow.apache.org/docs/apache-airflow/2.0.2/docker-compose.yaml'

    docker-compose up airflow-init
    docker-compose up
  popd > /dev/null
}

goal_destroy() {
  pushd "${script_dir}" > /dev/null
    docker-compose down
  popd > /dev/null
}



TARGET=${1:-}
if type -t "goal_${TARGET}" &>/dev/null; then
  "goal_${TARGET}" ${@:2}
else
  echo "Usage: $0 <goal>

goal:
    run                     - Runs container
"
  exit 1
fi
