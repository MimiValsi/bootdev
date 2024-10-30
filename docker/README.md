# Docker

## cmd

- docker exec <\container> netstat -ltnp
netstat:
    l: listening -> Show only listening sockets
    t: tcp
    n: numeric -> show only numeric addresses
    p: program -> Show the PID and name of the program to which each socket belongs

- docker exec -it
i: makes exec cmd interactive
t: give tty interface

- p flag XX:YY
XX: host port
YY: port withing the container
