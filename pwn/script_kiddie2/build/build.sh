#!/bin/sh

cd `dirname $0`
cd ../src
make
mv main ../dist/script_kiddie2
