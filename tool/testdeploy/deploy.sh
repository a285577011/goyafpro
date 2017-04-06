#!/bin/bash

programDir="/Users/zenghui/go/mushu-youxing"

$programDir/tool/compile.sh

#$programDir/tool/testdeploy/rm.sh

cd $programDir/bin

tar -jcf mushu-youxing-linux64.tar.bz mushu-youxing-linux64

scp $programDir/bin/mushu-youxing-linux64.tar.bz root@192.168.3.233:/data/website/mushu-youxing

$programDir/tool/testdeploy/reload.sh