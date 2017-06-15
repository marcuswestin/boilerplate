#!/bin/bash
set -x

kubectl rollout status deployment hello-node
