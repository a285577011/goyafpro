#!/usr/bin/expect -f

set user root
set host 192.168.3.233
set timeout 120

spawn ssh $user@$host

send "tar -xf /data/website/mushu-youxing/mushu-youxing-linux64.tar.bz -C /data/website/mushu-youxing\r"

send "chown -R root:root /data/website/mushu-youxing/mushu-youxing-linux64\r"
send "chmod a+x /data/website/mushu-youxing/mushu-youxing-linux64\r"
send "curl 'http://192.168.3.233:8005/goyaf_upgrade'\r"
send "exit\r"

expect eof