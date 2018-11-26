#!/bin/bash

if [[ $EUID -ne 0 ]]; then
	echo "Must be root."
	exit
fi
cp bin/gost.o /usr/local/bin/gost
