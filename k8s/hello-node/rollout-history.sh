#!/bin/bash
set -x

kubectl rollout history deployment/hello-node
