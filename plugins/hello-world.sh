#!/usr/bin/env bash

set -e

object=${1}
if [ -z "${object}" ]; then
  echo "OBJECT not set. Exiting."
  exit 1
fi

echo "Hello ${object}"