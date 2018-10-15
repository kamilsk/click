#!/usr/bin/env bash

set -euo pipefail

export CLICK_TOKEN=10000000-2000-4000-8000-160000000003

click ctl create -f env/client/grpc/link.create.yml
click ctl create -f env/client/grpc/namespace.create.global.yml
click ctl create -f env/client/grpc/namespace.create.support.yml
click ctl create -f env/client/grpc/alias.create.issue.yml
click ctl create -f env/client/grpc/alias.create.promo.yml
click ctl create -f env/client/grpc/alias.create.src.yml
click ctl create -f env/client/grpc/target.create.issue.yml
click ctl create -f env/client/grpc/target.create.promo.yml
click ctl create -f env/client/grpc/target.create.src.yml

click ctl read -f env/client/grpc/link.read.yml
click ctl read -f env/client/grpc/namespace.read.global.yml
click ctl read -f env/client/grpc/namespace.read.support.yml
click ctl read -f env/client/grpc/alias.read.issue.yml
click ctl read -f env/client/grpc/alias.read.promo.yml
click ctl read -f env/client/grpc/alias.read.src.yml
click ctl read -f env/client/grpc/target.read.issue.yml
click ctl read -f env/client/grpc/target.read.promo.yml
click ctl read -f env/client/grpc/target.read.src.yml

click ctl update -f env/client/grpc/link.update.yml
click ctl update -f env/client/grpc/namespace.update.global.yml
click ctl update -f env/client/grpc/namespace.update.support.yml
click ctl update -f env/client/grpc/alias.update.issue.yml
click ctl update -f env/client/grpc/alias.update.promo.yml
click ctl update -f env/client/grpc/alias.update.src.yml
click ctl update -f env/client/grpc/target.update.issue.yml
click ctl update -f env/client/grpc/target.update.promo.yml
click ctl update -f env/client/grpc/target.update.src.yml

click ctl delete -f env/client/grpc/link.delete.yml
click ctl delete -f env/client/grpc/namespace.delete.global.yml
click ctl delete -f env/client/grpc/namespace.delete.support.yml
click ctl delete -f env/client/grpc/alias.delete.issue.yml
click ctl delete -f env/client/grpc/alias.delete.promo.yml
click ctl delete -f env/client/grpc/alias.delete.src.yml
click ctl delete -f env/client/grpc/target.delete.issue.yml
click ctl delete -f env/client/grpc/target.delete.promo.yml
click ctl delete -f env/client/grpc/target.delete.src.yml
