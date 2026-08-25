# cuebook-rehearsal-lab-6197 Docker 交付说明

## 项目概览

Cuebook Rehearsal Lab 为现场演出排练团队准备可追溯的 cue book。命令行入口仍保留 `compose`、`review`、`publish` 和 `inspect` 工作流；镜像默认启动一个常驻健康检查接口，以便容器运行状态可被自动验收。

- Go module: `example.com/cuebook-rehearsal-lab`
- HTTP health endpoint: `GET /healthz`
- 默认端口: `8080`（可通过 `PORT` 环境变量覆盖）

## 标准命令

```bash
go build ./...
go test ./...
go run ./cmd/cuebook compose --show lantern --director Mira
```

## Docker 构建与运行

```bash
./build_benzhi_docker.sh cuebook-rehearsal-lab-6197-benzhi linux/amd64
docker run --rm -d --name cuebook-rehearsal-lab -p 8080:8080 cuebook-rehearsal-lab-6197-benzhi
curl -i http://127.0.0.1:8080/healthz
```

成功响应为 HTTP `200 OK`，响应体如下：

```json
{"status":"ok","service":"cuebook-rehearsal-lab"}
```

镜像在没有附加终端的 `docker run -d` 模式下会持续运行；Dockerfile 声明了 `EXPOSE 8080`，也可使用 `docker run -P` 和 `docker port` 获取映射端口。

## 环境

- 基础镜像: `golang:1.26.2`
- 二进制文件: `/usr/local/bin/cuebook`
- 代码目录: `/app`
