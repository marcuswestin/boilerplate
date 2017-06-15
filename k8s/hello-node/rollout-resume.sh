#!/bin/bash
set -x

kubectl rollout resume deployment hello-node
