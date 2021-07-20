#!/usr/bin/env bash
docker build -t rest_app .
docker run -itd -p 8000:8000 rest_app