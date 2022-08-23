#!/bin/bash

psql -U postgres -d database -a -f /app/scripts/seed.sql
