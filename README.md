# CUP

[![Build Status](https://travis-ci.org/jiajunhuang/cup.svg?branch=master)](https://travis-ci.org/jiajunhuang/cup)
[![codecov](https://codecov.io/gh/jiajunhuang/cup/branch/master/graph/badge.svg)](https://codecov.io/gh/jiajunhuang/cup)


cup is a container :)

# Set it up

- clone this repo:

```bash
$ git clone https://github.com/jiajunhuang/cup
$ cd cup
```

- prepare rootfs & command

```bash
$ # download busybox binary in: https://busybox.net/downloads/binaries/1.21.1/
$ wget https://busybox.net/downloads/binaries/1.21.1/busybox-x86_64 -o busybox
$ mkdir -p rootfs/{bin,proc}
$ mv busybox rootfs/bin/
```

- change rootfs's owner to root:root(because we're gonna run cup in root)

```bash
$ sudo chown -R root:root rootfs/
```

- compile cup and spawn it!

```bash
$ make && sudo ./cup
go fmt ./...
go vet -v .
github.com/jiajunhuang/cup
# github.com/jiajunhuang/cup
Checking file ./main.go
go test -cover  -race   ./...
?       github.com/jiajunhuang/cup      [no test files]
go build -o cup
2018/03/09 10:15:47 main start...
2018/03/09 10:15:48 childProcess start...uid: 0, gid: 0
2018/03/09 10:15:48 child: hostname: idea
2018/03/09 10:15:48 child: hostname: cup-host
/ # busybox ps
PID   USER     TIME   COMMAND
    1 0          0:00 {exe} childProcess
    6 0          0:00 /bin/busybox sh
    7 0          0:00 busybox ps
/ # busybox cat /proc/self/mounts
/proc /proc proc rw,relatime 0 0
/ # busybox id
uid=0 gid=0 groups=0,65534,65534,65534,65534,65534,65534,65534
/ # exit
2018/03/09 10:16:17 main: hostname: idea
```
