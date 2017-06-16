#!/bin/bash
set -x

APP_NAME=$1

kubectl rollout resume deployment $APP_NAME
