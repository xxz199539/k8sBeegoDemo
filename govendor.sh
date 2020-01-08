#! /bin/bash
rm -rf vendor
govendor init
govendor add +e