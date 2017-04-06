#!/usr/bin/expect -f

set user root
set host 192.168.3.233
set timeout 120

spawn ssh $user@$host
send "rm -rf /data/website/pba-fat-scale/linux64\r"
send "exit\r"

expect eof