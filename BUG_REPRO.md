# 修复前故障复现（Docker）

## 项目与标准命令

Cuebook Rehearsal Lab 用于组装、复核、读取和发布演出排练提示册。标准复现命令：

```sh
go test -count=1 -run TestIndependentRunsDoNotReuseEarlierShowState ./internal/app
```

## 环境构建与编译

在当前机器的 arm64 平台上，使用官方 `golang:1.26.2` 镜像构建项目镜像；镜像构建完成后，在容器内执行：

```sh
go build ./...
```

构建成功。

## 故障触发步骤

1. 构造对应演出的排练运行和读取或完成流程。
2. 执行标准命令。
3. 观察运行结果。

## 实际错误输出

以下是修复前实际执行的完整输出（退出码 1）：

```text
--- FAIL: TestIndependentRunsDoNotReuseEarlierShowState (0.00s)
    batch2_regression_test.go:23: app reused show "aurora"
FAIL
FAIL	example.com/cuebook-rehearsal-lab/internal/app	0.4s
FAIL
```

## 期望行为

连续准备不同演出时，每场排练都应保留自己的节目名称和运行身份，不能继承前一场的状态。
