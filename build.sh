#!/bin/bash

set -euo pipefail

cd js
npm run build-dev
cd ..

revel run
