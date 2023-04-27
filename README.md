# Golang项目数据库迁移Sample

## 环境启动

1. 初始化数据库环境

    ```bash
docker-compose up -d

    ```

2. 启动项目

    ```bash
    go run main.go
    ```

启动时执行内容有:

* 创建 MySQL 数据库, 并暴露 3400 本地端口
* 执行 Liquibase 数据库迁移脚本, 创建数据库表以及插入默认数据

> [Liquibase](https://www.liquibase.com) 没有放到应用程序中
>
> 1. 可以用于数据库SQL审计，由DBA执行。能够建表的数据库用户权限会更高，实际运行的时候权限只会有CRUD权限
> 2. [liquibase 作为 Kubernetes init container](https://www.liquibase.com/blog/using-liquibase-in-kubernetes)
     ，在系统启动前执行。避免由于耗时操作导致 Kubernetes 健康检查失败，从而杀死 Pod

## 为什么是Liquibase

