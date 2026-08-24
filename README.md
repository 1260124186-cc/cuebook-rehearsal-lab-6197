# Cuebook Rehearsal Lab

Cuebook Rehearsal Lab helps a stage-management team prepare a consistent cue book before a live performance. A stage manager assembles a rehearsal run, department leads review unresolved cues, and the show caller publishes a ready brief once every required cue has a decision.

## Roles and workflows

- **Stage manager** assembles a named rehearsal run from lighting, sound, deck, and projection cues.
- **Department lead** records a review decision and a concise note for a selected cue.
- **Show caller** publishes a cue book only after required departments have accepted the run.

## Layout

- `cmd/cuebook`: the public program entry point.
- `internal/app`: workflow-oriented application services.
- `internal/model`: rehearsal, cue, review, revision, and policy concepts.
- `internal/engine`: planning, review, readiness, and publishing decisions.
- `internal/store`: in-memory snapshots and event journal.
- `internal/format`: readable program output.

## Run locally

```bash
go run ./cmd/cuebook compose --show lantern --director Mira
go run ./cmd/cuebook review --show lantern --department sound
go run ./cmd/cuebook publish --show lantern
```

Build and test with `go build ./...` and `go test ./...`. The program uses deterministic built-in rehearsal material and does not require environment variables or remote services.
