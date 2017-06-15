#!/bin/bash
set -x

kubectl rollout undo deployment hello-node
