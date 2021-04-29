#!/usr/bin/env bash

set -e
set -o nounset
set -o pipefail

script_dir=$(cd "$(dirname "$0")" ; pwd -P)

goal_run() {
  pushd "${script_dir}" > /dev/null
    docker run -d \
    -p 8080:8080 \
    -v ${script_dir}/dags:/usr/local/airflow/dags \
    -v ${script_dir}/src:/usr/local/airflow/src \
    -v ${script_dir}/requirements.txt/:/requirements.txt \
    puckel/docker-airflow
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
