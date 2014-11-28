lxd-registry
==========================

This project is under initial design and prototype phase for lxd-registry.

Getting started
---------------
'''
go get github.com/kunalkushwaha/lxd-registry

cd $GOPATH/src/github.com/kunalkushwaha/lxd-registry

make

./lxd-registry
'''

API Supported
-------------

```shell
curl -v <IP>:8080/1.0/pingStatus
* About to connect() to 192.168.11.9 port 8080 (#0)
*   Trying 192.168.11.9... connected
> GET /1.0/pingStatus HTTP/1.1
> User-Agent: curl/7.22.0 (i686-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1 zlib/1.2.3.4 libidn/1.23 librtmp/2.3
> Host: 192.168.11.9:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Fri, 28 Nov 2014 19:28:23 GMT
< Content-Length: 61
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 192.168.11.9 left intact
* Closing connection #0
{"type":"Resp","result":"success","metadata":"Server Alive!"}
```

```shell
curl -v <IP>:8080/1.0/listCmd
* About to connect() to 192.168.11.9 port 8080 (#0)
*   Trying 192.168.11.9... connected
> GET /1.0/listCmd HTTP/1.1
> User-Agent: curl/7.22.0 (i686-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1 zlib/1.2.3.4 libidn/1.23 librtmp/2.3
> Host: 192.168.11.9:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Fri, 28 Nov 2014 19:29:51 GMT
< Content-Length: 57
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 192.168.11.9 left intact
* Closing connection #0
{"type":"Resp","result":"success","metadata":"List Cmd!"}
```

```shell
curl <IP>:8080/1.0/getImage/ubuntu32
* About to connect() to 192.168.11.9 port 8080 (#0)
*   Trying 192.168.11.9... connected
> GET /1.0/getImage HTTP/1.1
> User-Agent: curl/7.22.0 (i686-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1 zlib/1.2.3.4 libidn/1.23 librtmp/2.3
> Host: 192.168.11.9:8080
> Accept: */*
> 
< HTTP/1.1 301 Moved Permanently
< Location: http://images.linuxcontainers.org/images/ubuntu/trusty/i386/default/20141128_03:49/rootfs.tar.xz
< Date: Fri, 28 Nov 2014 19:30:22 GMT
< Content-Length: 0
< Content-Type: text/plain; charset=utf-8
< 
* Connection #0 to host 192.168.11.9 left intact
* Closing connection #0
```


