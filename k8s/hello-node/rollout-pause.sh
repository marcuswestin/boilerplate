#!/bin/bash
set -x

kubectl rollout pause deployment hello-node
