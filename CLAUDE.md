# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go test ./...                       # run tests
go test -run TestBuildXls            # run single test
go test -bench=. -tags=benchmark     # size-controlled benchmark (50k row CSV) — needs `benchmark` build tag
go test -bench=BenchmarkBuildXls_DiskV -run '^$'   # run one in-repo benchmark (no build tag needed)
go build                             # build csv2xlsx binary
```

CI (`.github/workflows/test.yaml`) just runs `go test` on push, Go version pinned via `go.mod`.

## Architecture

Single Go module, four files at package root, no internal packages:

- `main.go` — entrypoint; opens CSV via `getCsvData` (uses `encoding/csv`, `ReuseRecord = true` for low allocations).
- `command_line.go` — CLI flag parsing via `urfave/cli/v2`. `checkAndReturnParams` validates raw flags into a `params` struct (paths resolved to absolute, delimiter as single rune, template existence checked). This is the only place flags are read.
- `write.go` — core conversion logic, all built on `codeberg.org/tealeg/xlsx/v4`:
  - `buildXls` — top-level orchestration: opens template or creates new `xlsx.File`, calls `writeAllSheets`, saves output.
  - `writeAllSheets` / `getSheet` — maps each input CSV file to a sheet by position, reusing an existing template sheet by name if `sheetNames[i]` matches, otherwise creating `Sheet %d`.
  - `exampleRow` mechanic: if `--exampleRow` is set, that row is read from the template sheet, removed, and its per-cell style/number-format is copied onto every generated row (`writeRowToXls`). This is how template-based cell formatting/styling survives into the generated data rows.
  - `setCellValue` / `isNumeric` — decides whether a CSV field becomes a numeric or string cell. A leading `'` forces string type (spreadsheet-style escape). Custom `isNumeric` (not `strconv.ParseFloat`) deliberately rejects things like `"Inf"`, `"1.."`, `"-"` that Go's parser would otherwise treat specially — see `TestStrangeBehaviorOnInfValue`/`TestInNumeric` before changing this logic.
- `--use-cache` flag switches `xlsx.NewFile`/`xlsx.OpenFile` to `xlsx.UseDiskVCellStore` for disk-backed cell storage on large files (trades speed for memory).

## Notes

- `example/` holds fixture CSVs/templates used directly by tests (`main_test.go` references `./example/data.csv`, `./example/template.xlsx`).
- Sheets are matched to input files positionally: Nth CSV file → Nth `--sheets` name (or `Sheet %d` if not enough names given).
- Release builds are handled by GoReleaser (`.goreleaser.yml`) across linux/windows/darwin, 386/amd64/arm/arm64.
