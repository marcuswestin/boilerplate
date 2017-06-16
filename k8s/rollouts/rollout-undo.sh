#!/bin/bash
set -x

APP_NAME=$1

kubectl rollout undo deployment $APP_NAME
