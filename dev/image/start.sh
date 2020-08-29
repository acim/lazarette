#!/usr/bin/env bash

cd /app/frontend
npm run dev &

cd /app
reflex -d none -c /usr/local/etc/reflex.conf
