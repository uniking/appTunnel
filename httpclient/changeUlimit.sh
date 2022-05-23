#!/bin/bash
sysctl -w fs.file-max=2000500
sysctl -w fs.nr_open=2000500
sysctl -w net.nf_conntrack_max=2000500
ulimit -n 2000500
