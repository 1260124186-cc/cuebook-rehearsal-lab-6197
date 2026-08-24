# cuebook-rehearsal-lab__015 Docker 交付说明

基于 Go 实现的现场演出排练提示册 CLI 服务，支持组装排练运行、跨部门复核提示并发布可追溯的演出摘要。

## 项目概览
- Cuebook Rehearsal Lab helps a stage-management team prepare a consistent cue book before a live performance. A stage manager assembles a rehearsal run, department leads review unre
- Go module: `example.com/cuebook-rehearsal-lab`

## 标准命令

```bash
go build ./...
go test ./...
```

## 实际启动入口

```bash
go run ./cmd/cuebook
```

## Docker 构建

```bash
./build_benzhi_docker.sh cuebook-rehearsal-lab__015-benzhi linux/amd64
docker run --rm -it cuebook-rehearsal-lab__015-benzhi bash
```

## 环境

- 基础镜像: `golang:1.26.2`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`
