#!/usr/bin/expect

set timeout  -1
set user     $env(VPN_LOGIN)
set password $env(VPN_PASSWORD)
set server   $env(VPN_SERVER)

spawn openconnect --protocol=AnyConnect --user=$user $server

expect "Password:"
send "$password\r"

expect eof
