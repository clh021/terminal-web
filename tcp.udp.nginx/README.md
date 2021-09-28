# 测试 tcp/udp 在 nginx 代理的可行性

## 可以通过一个端口监听转发到不同容器服务
```conf
stream {
    server {
        listen 7564;
        proxy_pass server1:7990;
        # server_name  tcp.localhost; # 不允许域名设置
    }
    server {
        listen 7564 udp;
        proxy_pass server2:7990;
        # server_name  udp.localhost; # 不允许域名设置
    }
}
```
设置 nginx 代理后的情形
```log
-> % docker-compose up
WARNING: Found orphan containers (tcpudpnginx_server_1, tcpudpnginx_client1_1) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
Starting tcpudpnginx_server2_1 ... done
Starting tcpudpnginx_server1_1 ... done
Starting tcpudpnginx_nginx_1   ... done
Starting tcpudpnginx_client2_1 ... done
Attaching to tcpudpnginx_server1_1, tcpudpnginx_server2_1, tcpudpnginx_nginx_1, tcpudpnginx_client2_1
client2_1  | Start Type: client
client2_1  | send Host:  nginx
client2_1  | send Port:  7564
client2_1  | [UDP]ParseIP:  nginx  =>  172.23.0.4
nginx_1    | /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
nginx_1    | /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
client2_1  | 2021/09/29 00:02:18 udp sendto ip: 172.23.0.4
server2_1  | Start Type: server
server2_1  | Server Port:  7990
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 2, 18, 3179806, time.Local)
server1_1  | Start Type: server
server1_1  | Server Port:  7990
nginx_1    | 10-listen-on-ipv6-by-default.sh: info: IPv6 listen already enabled
server2_1  | [UDP]Listen: <[::]:7990> 
server1_1  | [UDP]Listen: <[::]:7990> 
server1_1  | [TCP]Listen:  [::]:7990
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
server2_1  | [TCP]Listen:  [::]:7990
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
server2_1  | [UDP]Recv: bytes= 66  from= 172.23.0.4:44519  data: time.Date(2021, time.September, 29, 0, 2, 18, 3179806, time.Local)
nginx_1    | /docker-entrypoint.sh: Configuration complete; ready for start up
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: using the "epoll" event method
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: nginx/1.21.3
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: built by gcc 8.3.0 (Debian 8.3.0-6) 
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: OS: Linux 5.10.69-1-lts
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: start worker processes
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: start worker process 24
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: start worker process 25
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: start worker process 26
nginx_1    | 2021/09/29 00:02:17 [notice] 1#1: start worker process 27
client2_1  | [UDP]read 5 from <172.23.0.4:7564> recv:[104 101 108 108 111] 
server2_1  | [UDP]Recv: bytes= 66  from= 172.23.0.4:44519  data: time.Date(2021, time.September, 29, 0, 2, 20, 3741332, time.Local)
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 2, 20, 3741332, time.Local)
client2_1  | [UDP]read 5 from <172.23.0.4:7564> recv:[104 101 108 108 111] 
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 2, 22, 4348847, time.Local)
server2_1  | [UDP]Recv: bytes= 66  from= 172.23.0.4:44519  data: time.Date(2021, time.September, 29, 0, 2, 22, 4348847, time.Local)
client2_1  | [UDP]read 5 from <172.23.0.4:7564> recv:[104 101 108 108 111] 
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 2, 24, 4814156, time.Local)
server2_1  | [UDP]Recv: bytes= 66  from= 172.23.0.4:44519  data: time.Date(2021, time.September, 29, 0, 2, 24, 4814156, time.Local)
client2_1  | [UDP]read 5 from <172.23.0.4:7564> recv:[104 101 108 108 111] 
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 2, 26, 5454086, time.Local)
server2_1  | [UDP]Recv: bytes= 66  from= 172.23.0.4:44519  data: time.Date(2021, time.September, 29, 0, 2, 26, 5454086, time.Local)
client2_1  | [UDP]read 5 from <172.23.0.4:7564> recv:[104 101 108 108 111] 
client2_1  | [TCP]Send: time.Date(2021, time.September, 29, 0, 2, 28, 7736177, time.Local)
server1_1  | [TCP]Recv: bytes= 68  data: time.Date(2021, time.September, 29, 0, 2, 28, 7736177, time.Local)
server1_1  | 
client2_1  | [TCP]Message received. data: time.Date(2021, time.September, 29, 0, 2, 28, 7736177, time.Local)
client2_1  | 
server1_1  | [TCP]Recv: bytes= 68  data: time.Date(2021, time.September, 29, 0, 2, 30, 8873861, time.Local)
server1_1  | 
client2_1  | [TCP]Send: time.Date(2021, time.September, 29, 0, 2, 30, 8873861, time.Local)
client2_1  | [TCP]Message received. data: time.Date(2021, time.September, 29, 0, 2, 30, 8873861, time.Local)
client2_1  | 
client2_1  | [TCP]Send: time.Date(2021, time.September, 29, 0, 2, 32, 9393610, time.Local)
server1_1  | [TCP]Recv: bytes= 68  data: time.Date(2021, time.September, 29, 0, 2, 32, 9393610, time.Local)
server1_1  | 
client2_1  | [TCP]Message received. data: time.Date(2021, time.September, 29, 0, 2, 32, 9393610, time.Local)
client2_1  | 
client2_1  | [TCP]Send: time.Date(2021, time.September, 29, 0, 2, 34, 10933750, time.Local)
server1_1  | [TCP]Recv: bytes= 69  data: time.Date(2021, time.September, 29, 0, 2, 34, 10933750, time.Local)
server1_1  | 
client2_1  | [TCP]Message received. data: time.Date(2021, time.September, 29, 0, 2, 34, 10933750, time.Local)
client2_1  | 
client2_1  | [TCP]Send: time.Date(2021, time.September, 29, 0, 2, 36, 12340049, time.Local)
server1_1  | [TCP]Recv: bytes= 69  data: time.Date(2021, time.September, 29, 0, 2, 36, 12340049, time.Local)
server1_1  | 
client2_1  | [TCP]Message received. data: time.Date(2021, time.September, 29, 0, 2, 36, 12340049, time.Local)
client2_1  | 
server1_1  | [TCP]Error reading: EOF
tcpudpnginx_client2_1 exited with code 0
^CGracefully stopping... (press Ctrl+C again to force)
Stopping tcpudpnginx_nginx_1   ... done
Stopping tcpudpnginx_server1_1 ... done
Stopping tcpudpnginx_server2_1 ... done

```


## 不被允许设置域名
```conf
    server {
        listen 7564 udp;
        proxy_pass server2:7990;
        server_name  udp.localhost; # 不允许域名设置
    }
```
设置域名后的情形
```log
WARNING: Found orphan containers (tcpudpnginx_client1_1, tcpudpnginx_server_1) for this project. If you removed or renamed this service in your compose file, you can run this command with the --remove-orphans flag to clean it up.
Starting tcpudpnginx_server1_1 ... done
Starting tcpudpnginx_server2_1 ... done
Starting tcpudpnginx_nginx_1   ... done
Recreating tcpudpnginx_client2_1 ... done
Attaching to tcpudpnginx_server1_1, tcpudpnginx_server2_1, tcpudpnginx_nginx_1, tcpudpnginx_client2_1
client2_1  | Start Type: client
client2_1  | send Host:  nginx
client2_1  | send Port:  7564
nginx_1    | /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
nginx_1    | /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
client2_1  | [UDP]ParseIPError:  lookup nginx on 127.0.0.11:53: no such host
client2_1  | [UDP]ParseIP:  nginx  =>  <nil>
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
server1_1  | Start Type: server
server1_1  | Server Port:  7990
client2_1  | 2021/09/29 00:07:35 udp sendto ip: <nil>
nginx_1    | 10-listen-on-ipv6-by-default.sh: info: IPv6 listen already enabled
server2_1  | Start Type: server
server2_1  | Server Port:  7990
server1_1  | [UDP]Listen: <[::]:7990> 
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
server2_1  | [UDP]Listen: <[::]:7990> 
server2_1  | [TCP]Listen:  [::]:7990
nginx_1    | /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
nginx_1    | /docker-entrypoint.sh: Configuration complete; ready for start up
server1_1  | [TCP]Listen:  [::]:7990
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 7, 35, 63051475, time.Local)
nginx_1    | 2021/09/29 00:07:34 [emerg] 1#1: "server_name" directive is not allowed here in /etc/nginx/nginx.conf:37
nginx_1    | nginx: [emerg] "server_name" directive is not allowed here in /etc/nginx/nginx.conf:37
tcpudpnginx_nginx_1 exited with code 1
client2_1  | [UDP]ReadError:  read udp 127.0.0.1:50032->127.0.0.1:7564: recvfrom: connection refused
client2_1  | [UDP]read 0 from <<nil>> recv:[] 
client2_1  | [UDP]Send: time.Date(2021, time.September, 29, 0, 7, 37, 63987404, time.Local)
^CGracefully stopping... (press Ctrl+C again to force)
Stopping tcpudpnginx_client2_1   ... done
Stopping tcpudpnginx_server1_1   ... done
Stopping tcpudpnginx_server2_1   ... done
```