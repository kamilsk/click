#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER "click" WITH PASSWORD 'click';
    CREATE DATABASE "click" WITH OWNER "click";
    \c "click";
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
