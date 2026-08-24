# 项目与标准命令

Cuebook Rehearsal Lab 用于组装、复核和发布演出排练提示册。标准复现命令：

```sh
go test -count=1 -run TestPublishDoesNotHideIncompleteReviewChain ./internal/app
```

## 环境构建与编译

在 macOS arm64、Go 1.26.2 环境的项目根目录执行：

```sh
go build ./...
```

编译成功后执行复现命令。

## 故障触发步骤

1. 准备 lantern 排练册的完整发布流程。
2. 在项目根目录执行标准命令。
3. 观察发布状态、复核记录和审计状态。

## 实际错误输出

以下是故障基线的完整实际输出（退出码 1）：

```text
--- FAIL: TestPublishDoesNotHideIncompleteReviewChain (0.00s)
    review_error_chain_test.go:10: reviews=0 cues=6
FAIL
FAIL	example.com/cuebook-rehearsal-lab/internal/app	0.432s
FAIL
```

## 期望行为

发布工作流应保留完整且一致的运行状态、部门复核和审计记录；发生复核错误时应以可处理错误终止发布，而不能继续发布不完整的排练册。
