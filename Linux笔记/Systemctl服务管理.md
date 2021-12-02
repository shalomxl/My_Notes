先给出几篇优秀文章：

[Systemd 入门教程：命令篇](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

[Systemd 入门教程：实战篇](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html)

[systemd服务配置文件编写(1)](https://www.junmajinlong.com/linux/systemd/service_1/)

[systemd服务配置文件编写(2)](https://www.junmajinlong.com/linux/systemd/service_2/)

## systemd文件夹

开机启动时，systemd只执行`/etc/systemd/system` 文件夹下的内容

```bash
systemctl enable httpd
```

以上命令相当于在`/etc/systemd/system`目录添加一个符号链接，指向`/usr/lib/systemd/system`里面的`httpd.service`文件。`/usr/lib/systemd/system`或者`/lib/systemd/system`是存放service文件的地方。

一共就这3个目录。

使用systemctl部署一个web应用，让它能够开机自启动

1. 要保证web应用的依赖也能开机自启动
2. 要在 .service 的配置文件中指明服务的依赖关系
3. 当服务停止后，自启动的设置

目前web应用依赖的服务有：

1. Nginx
2. mysql
3. redis

以上三个均为强需求依赖，必须保证他们正常运行，否则web服务将不能启动

**Nginx**

查看当前是否运行了 nginx.service 服务，可以使用查看 nginx.service 状态的命令：

```sh
systemctl status nginx.service
```

若不存在，则会提示 `Unit nginx.service could not be found.`

安装Nginx的时候，软件是否会自动创建 nginx.service ？不会。

三个部分：

- Unit
- Service
- Install

**Unit**

最简单的是 Description 属性，写上让自己和他人能够辨认这个服务就可以

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
```

定义服务依赖关系。查阅若干个 nginx.service 的案例后，并没有看到 nginx 的依赖服务，只定义了 nginx.service服务要在 network.target 之后启动。表示“我应该在谁后面启动”应该使用 After 属性。

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
```

定义依赖关系的属性还有：

- Before：表示“我应该在谁前面启动”，服务之间相互没有影响
- Wants：弱依赖关系；当前服务启动时，会先启动 Wants 中指定的服务，这里代表的是依赖关系；若 Wants 中服务启动失败，不影响当前服务的启动，这里强调的是 **弱**依赖
- Requires：强依赖关系；当前服务启动时，会先启动 Requires 中指定的服务，若 Requires 中的服务启动失败，那么当前服务也不会启动，若在运行时 Requires 中的服务中断，那么当前服务也会中断。

**Install**

开机启动将在这里配置。

WantedBy 属性表示当前服务所在的 Target。Systemd中的Target代表一个服务组，也就是包含多个服务。而Systemd 有默认开机启动的服务组，那就是 multi-user.target ，所以想让 nginx.service 服务开机启动，那么就必须设置 WantedBy 属性为 multi-user.target。

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
...
[Install]
WantedBy=multi-user.target
```

查看systemd默认开机启动的target：

```bash
# 查看启动时的默认 Target
$ systemctl get-default
```

**Service**

这是 service 文件中最核心的环节。

首先选择 **Type** 属性。Type属性有两个常用的值：

1. forking
2. simple

当然不止这两个 value，常用的可以在文章首部给出的链接中查阅。

forking只适用于 以守护进程的方式启动的服务，nginx启动时，可以选择以守护进程的方式启动，也可以不选择，在服务器中一般都是以守护进程的方式启动，所以选择 forking。

在选择forking的时候，需要配置PID文件路径

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
[Install]
WantedBy=multi-user.target
```

还有如下命令需要设置：

- `ExecStart`：启动当前服务的命令
- `ExecStartPre`：启动当前服务之前执行的命令
- `ExecStartPost`：启动当前服务之后执行的命令
- `ExecReload`：重启当前服务时执行的命令
- `ExecStop`：停止当前服务时执行的命令
- `ExecStopPost`：停止当其服务之后执行的命令

在上面的配置中，可根据具体需求来选择性的配置。

首先，启动命令一定要配置，也就是 ExecStart

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
[Install]
WantedBy=multi-user.target
```

然后我们可能还会需要在启动服务前，判断配置文件是否正确，那我们就需要将判断命令配置在ExecStartPre命令中：

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
[Install]
WantedBy=multi-user.target
```

最后配置重启和关闭命令：

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload -c /usr/local/nginx/conf/nginx.conf
ExecStop=/usr/local/nginx/sbin/nginx -s quit -c /usr/local/nginx/conf/nginx.conf
[Install]
WantedBy=multi-user.target
```

因为nginx服务是web应用的强需求，所以我们需要配置它的自启动，保证nginx停了后能重新启动，继续工作，配置Restart属性：

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload -c /usr/local/nginx/conf/nginx.conf
ExecStop=/usr/local/nginx/sbin/nginx -s quit -c /usr/local/nginx/conf/nginx.conf
Restart=always
[Install]
WantedBy=multi-user.target
```

由于服务进程需要时间将PID写入PIDFile中，所以最好在服务启动后，停止一小段时间，让服务进程将PID写入：

```sh
[Unit]
Description=A nginx.service created by Shalom on 2020.08.16
After=network.target
[Service]
Type=forking
PIDFile=/var/run/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload -c /usr/local/nginx/conf/nginx.conf
ExecStop=/usr/local/nginx/sbin/nginx -s quit -c /usr/local/nginx/conf/nginx.conf
ExecStartPost=/usr/bin/sleep 0.1
Restart=always
[Install]
WantedBy=multi-user.target
```

重新加载systemd的配置文件

```sh
systemctl daemon-reload
```

重载相关服务

```
systemctl restart nginx.service
```

最后将自己的web应用当做一个练习：

```sh
[Unit]
Description=My Web App DistributeBookKeeping service Create on 2020.08.16
After=nginx.service
Requires=nginx.service
[Service]
ExecStart=/root/project/DistributedBookkeeping_front/distribute
Restart=always
[Install]
WantedBy=multi-user.target
```

## 查看后台运行服务日志

```sh
journalctl -f -u geth.service
```
