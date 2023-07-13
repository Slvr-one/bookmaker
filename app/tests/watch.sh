#!/usr/bin/env bash
set -e

make skaffold-run
reflex -r "\.go$" -R "vendor.*" make skaffold-run