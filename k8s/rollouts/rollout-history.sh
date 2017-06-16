#!/bin/bash
set -x

APP_NAME=$1

kubectl rollout history deployment/$APP_NAME
