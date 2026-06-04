# Agent guide — build-tags module

Module path: `car-system` (see [`go.mod`](go.mod)). Work from the **`build-tags/`** directory unless noted.

## What this module demonstrates

- **Compile-time brand selection** via Go build tags (`audi`, `benz`, `bmw`).
- **Two version entrypoints** only: [`cmd/v1`](cmd/v1/main.go), [`cmd/v2`](cmd/v2/main.go).
- **Function-based APIs** — no interfaces, no `New()` constructors in brand packages.

Sibling reference implementation: [`../sub-commands/`](../sub-commands/) (six binaries, interfaces).

## Audi is the default brand

When the user builds or runs **without** `-tags`, the toolchain compiles **Audi** implementations. This is intentional — not an error state.

Audi files use:

```go
//go:build audi || !(benz || bmw)
```

Benz and BMW files use `//go:build benz` and `//go:build bmw` only.

**Do not** reintroduce stub files (e.g. `engine.go` with `panic("implement me")`) to satisfy untagged builds; the constraint above replaces stubs.

## Build commands (copy for users)

From `build-tags/`:

```bash
# Audi (default) — no -tags needed
go build -o bin/v1_audi ./cmd/v1
go build -o bin/v2_audi ./cmd/v2
go run ./cmd/v1
go run ./cmd/v2

# Other brands
go build -tags benz -o bin/v1_benz ./cmd/v1
go build -tags bmw  -o bin/v1_bmw  ./cmd/v1
```

`go build ./...` without tags must succeed and select Audi in all three brand-sensitive packages.

## Where to change behavior

| Change | Location |
|--------|----------|
| Engine message per brand | `internal/engine/engine_{audi,benz,bmw}.go` |
| v1 navigation per brand | `internal/v1/infotainment/infotainment_*.go` |
| v2 navigation + validation | `internal/v2/infotainment/infotainment_*.go` |
| v1/v2 orchestration | `internal/v1/app/app.go`, `internal/v2/app/app.go` |
| Entrypoints | `cmd/v1/main.go`, `cmd/v2/main.go` — **no brand imports** |

`cmd/*` must stay thin: import only `internal/v1/app` or `internal/v2/app` and call `Run()`.

## Adding or editing a brand file

1. Add or edit `*_brand.go` with `//go:build brand` (single tag).
2. Keep the same function signatures as the other brand files in that package.
3. If adding a **new** brand tag, update **every** `*_audi.go` constraint to exclude it, e.g. `audi || !(benz || bmw || porsche)`.
4. Never leave two brand files active for the same package in one build.

## IDE / gopls

[`.vscode/settings.json`](.vscode/settings.json) sets `gopls.build.buildFlags` to `-tags=audi` by default.

When editing Benz or BMW tagged files, switch the setting to `-tags=benz` or `-tags=bmw` so analysis and diagnostics match that brand’s compiled file set.

## Code conventions

- **Go file header** on every `.go` file:

```go
// (C) Copyright 2026 GTN Group. All Rights Reserved.
// Created by Dilanka Rathnasiri on YYYY-MM-DD
```

- Use **today’s date** in the `Created by` line when creating new files.
- Dummy / example data: prefer car brands (Audi, BMW, Mercedes-Benz/Benz); Porsche 911 Turbo S is the author’s reference favorite.
- Preserve parity with `sub-commands` output strings unless the task explicitly changes behavior.

## Isolation

- `internal/v1` must not import `internal/v2`.
- Do not duplicate `internal/engine` under version trees.
- Do not add a shared `internal/infotainment` — infotainment stays versioned under `v1` and `v2`.

## Verification checklist

After tag or brand changes:

1. Untagged `go build ./cmd/v1` and `./cmd/v2` succeed with Audi output.
2. `-tags audi` matches untagged output.
3. `-tags benz` and `-tags bmw` select the correct files only (no duplicate symbol errors).
4. v2 empty destination still returns `destination is required`.
5. Exactly three tagged files per brand-sensitive package; no stubs.

Full user-facing docs: [`README.md`](../README.md#build-tags).
