# Anime Hostess

基于Web的自托管番剧库，支持调用B站弹幕。（弹弹play web版）

![](https://i.imgur.com/U3XfrCj.png)

## 后端部署

后端采用 Go + gRPC 开发。

### 编译

需求：Go 1.16+

```bash
cd anime-hostess-backend
go build
```

### 部署gRPC服务端

gRPC服务端提供弹幕搜索和获取。由于B站番剧的IP限制，境外IP无法浏览番剧及获取弹幕，故将该服务与主服务端拆分。建议部署在境内服务器，或运行在本地。

将`config.grpc.ini.example`复制一份为`config.grpc.ini`，并根据配置文件中的注释进行修改后启动。

```bash
./anime-hostess -c config.grpc.ini
```

### 部署主服务端

主服务端提供自托管番剧的存储和索引服务。建议部署在大容量硬盘的服务器上，或运行在本地。

将`config.ini.example`复制一份为`config.ini`，并根据配置文件中的注释进行修改后启动。

配置文件的gRPC服务地址需与实际部署的相同。在运行主服务端前，需要先运行gRPC服务。

```bash
./anime-hostess -c config.ini
```

## 前端部署

前端采用 Vue 开发。

### 编译

需求：Node.js 14+

```bash
cd anime-hostess-frontend
npm install
npm run build
```

然后将`dist/`中生成的文件部署到`nginx`等 Web 服务器的路径下。

### 配置

在设定界面，填入主服务端的地址。点击更新后即可使用。

## 截图

![](https://i.imgur.com/Rfo2zi5.png)

![](https://i.imgur.com/75DFluU.png)