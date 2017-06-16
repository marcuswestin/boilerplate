#!/bin/bash
set -x

APP_NAME=$1

kubectl rollout status deployment $APP_NAME
