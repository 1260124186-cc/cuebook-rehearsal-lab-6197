# 项目与标准命令

Cuebook Rehearsal Lab 是一个用于组装、复核和发布演出排练提示册的 Go 命令行项目。此故障使用以下标准命令复现：

```sh
go test -race -count=1 -run TestConcurrentDepartmentReviewKeepsReviewPathSafe ./internal/app
```

## 环境构建与编译

在当前机器（macOS arm64）上，使用 Go 1.26.2。先在项目根目录执行：

```sh
go build ./...
```

编译应成功。

## 故障触发步骤

1. 在项目根目录准备复现用的并发复核场景。
2. 同时让 32 个部门复核请求进入提示册复核工作流。
3. 执行标准命令：

```sh
go test -race -count=1 -run TestConcurrentDepartmentReviewKeepsReviewPathSafe ./internal/app
```

## 实际错误输出

以下是故障基线一次实际执行的完整输出（退出码 1）：

```text
==================
WARNING: DATA RACE
Read at 0x00c0001043f0 by goroutine 14:
  runtime.mapdelete_fast64()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_fast64.go:505 +0x8c
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x3c0
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c0001043f0 by goroutine 15:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x3fc
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 14 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 15 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Read at 0x00c0001043c0 by goroutine 14:
  runtime.mapdelete_fast64()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_fast64.go:505 +0x8c
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x1f8
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c0001043c0 by goroutine 15:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x234
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 14 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 15 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Read at 0x00c000104420 by goroutine 15:
  runtime.mapdelete_fast64()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_fast64.go:505 +0x8c
  example.com/cuebook-rehearsal-lab/internal/store.NewEvent()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/store/event.go:22 +0x198
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:48 +0xac0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c000104420 by goroutine 14:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/store.NewEvent()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/store/event.go:22 +0x1d4
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:48 +0xac0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 15 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 14 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Read at 0x00c0001069d8 by goroutine 15:
  example.com/cuebook-rehearsal-lab/internal/store.NewEvent()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/store/event.go:22 +0x1a0
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:48 +0xac0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c0001069d8 by goroutine 14:
  example.com/cuebook-rehearsal-lab/internal/store.NewEvent()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/store/event.go:22 +0x1e0
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:48 +0xac0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 15 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 14 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Read at 0x00c000106aa8 by goroutine 31:
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x3c8
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c000106aa8 by goroutine 23:
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x408
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 31 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 23 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Read at 0x00c000106b78 by goroutine 31:
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x200
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c000106b78 by goroutine 23:
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x240
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 31 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 23 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Write at 0x00c0001043f0 by goroutine 28:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x3fc
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c0001043f0 by goroutine 24:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x3fc
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 28 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 24 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Write at 0x00c000106af0 by goroutine 28:
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x408
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c000106af0 by goroutine 12:
  example.com/cuebook-rehearsal-lab/internal/engine.trackReviewVisit()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:21 +0x408
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:44 +0x368
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 28 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 12 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Write at 0x00c0001043c0 by goroutine 12:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x234
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c0001043c0 by goroutine 24:
  runtime.mapaccess2_faststr()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/internal/runtime/maps/runtime_faststr.go:161 +0x2ac
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x234
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 12 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 24 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
==================
WARNING: DATA RACE
Write at 0x00c000106bc0 by goroutine 12:
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x240
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Previous write at 0x00c000106bc0 by goroutine 24:
  example.com/cuebook-rehearsal-lab/internal/model.NewReviewNote()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/model/note.go:29 +0x240
  example.com/cuebook-rehearsal-lab/internal/engine.AcceptNext()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/engine/reviewer.go:49 +0x4f8
  example.com/cuebook-rehearsal-lab/internal/app.(*Service).Review()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/review.go:37 +0x468
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.func1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:23 +0xb0
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe.gowrap1()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:24 +0x38

Goroutine 12 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34

Goroutine 24 (running) created at:
  example.com/cuebook-rehearsal-lab/internal/app_test.TestConcurrentDepartmentReviewKeepsReviewPathSafe()
      /Users/zhangchengcheng/work/ai-project/Trea/go-cc/go-cc-00068/2026-08-24/cuebook-rehearsal-lab__001/env/internal/app/concurrent_review_test.go:20 +0x188
  testing.tRunner()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2036 +0x164
  testing.(*T).Run.gowrap1()
      /opt/homebrew/Cellar/go/1.26.2/libexec/src/testing/testing.go:2101 +0x34
==================
--- FAIL: TestConcurrentDepartmentReviewKeepsReviewPathSafe (0.00s)
    testing.go:1712: race detected during execution of test
FAIL
FAIL	example.com/cuebook-rehearsal-lab/internal/app	0.463s
FAIL
```

## 期望行为

多位部门负责人同时确认提示时，复核工作流应保持可并行执行；并发检测不应报告 data race，已提交的复核和后续运行查看都应保持可靠。
