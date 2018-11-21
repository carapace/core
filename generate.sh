#!/usr/bin/env bash
COMMAND='./.proto.sh'
find `pwd` -iname ".proto.sh" -printf "%h\n" | sort -u | while read i; do
    cd "$i" && pwd && $COMMAND
done