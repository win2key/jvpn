#!/usr/bin/expect

set timeout  -1
set user     $env(VPN_LOGIN)
set password $env(VPN_PASSWORD)
set server   $env(VPN_SERVER)

# Handle termination signals
trap {
    send_user "\nCaught signal, terminating openconnect and exiting...\n"
    send "\003" ;# Ctrl+C to terminate the openconnect session
    exit
} SIGINT SIGTERM

spawn openconnect --protocol=AnyConnect --user=$user $server

expect "Password:"
send "$password\r"

expect eof
