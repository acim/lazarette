#!/usr/bin/env bash

cd /app/frontend
npm run dev &

cd /app
reflex -d none -c /home/acim/etc/reflex.conf
