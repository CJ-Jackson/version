#!/bin/dash
v=${GO_VER-go} . path-go               
$(basename $0) "$@"
