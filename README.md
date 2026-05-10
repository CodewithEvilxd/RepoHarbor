# RepoHarbor

RepoHarbor is a terminal app for opening GitHub repositories, browsing their files, and downloading only what you need.

It is built for people who want a fast way to inspect a repo without cloning the whole project first.

## What it does

- Opens public GitHub repositories
- Supports branches, tags, commits, compare links, and pull requests
- Browses folders and files in a TUI
- Searches the repository tree with fuzzy matching and simple filters
- Downloads selected files with parallel transfer
- Saves output as regular files, zip, or tar.gz
- Supports GitHub release assets
- Keeps recent repositories and favorites
- Works with GitHub tokens and named profiles

## Install

Use npm:

```bash
npx repoharbor
```

```bash
npm install -g repoharbor
```

Use Go:

```bash
go install github.com/codewithevilxd/repoharbor/cmd/repoharbor@latest
```

Build from source:

```bash
git clone https://github.com/codewithevilxd/repoharbor.git
cd repoharbor
go build -o bin/repoharbor ./cmd/repoharbor
```

## Quick Start

```bash
repoharbor https://github.com/charmbracelet/bubbletea
```

You can also open a branch, tag, commit, folder, compare view, or pull request URL directly.

Examples:

```bash
repoharbor https://github.com/torvalds/linux --ref master
repoharbor https://github.com/rust-lang/rust/tree/master/src/tools
repoharbor https://github.com/owner/repo/compare/main...feature
repoharbor https://github.com/owner/repo/pull/42
```

## Commands

Root command:

```bash
repoharbor [URL] [--cwd] [--no-folder] [--token TOKEN] [--profile NAME] [--ref REF]
```

Config command:

```bash
repoharbor config list
repoharbor config set token TOKEN
repoharbor config set path /your/preferred/path
repoharbor config unset token
repoharbor config unset path
repoharbor config profile list
repoharbor config profile use work
repoharbor config profile set-token work TOKEN
repoharbor config profile import-gh work
repoharbor config favorite add https://github.com/owner/repo
repoharbor config recent clear
```

## Keyboard Use

The TUI is controlled from the keyboard.

Common keys:

- `Enter`, `l`, `Right`: open a folder
- `Backspace`, `h`, `Left`: go back
- `j`, `k`, `Up`, `Down`: move through items
- `Space`: select or unselect an item
- `a`: select all
- `u`: clear selection
- `d`: start download
- `/`: search
- `b`, `t`: switch branch or tag
- `R`: browse release assets
- `m`: open repo summary
- `y`, `Y`, `P`: copy GitHub URL, raw URL, or output path
- `i`: toggle icon mode
- `q`, `Ctrl+C`: quit

## Download Behavior

- Files download in parallel
- Existing files can be skipped, overwritten, renamed, or resumed
- Output can be saved as normal files, a zip archive, or a tar.gz archive
- A manifest file is written beside completed downloads

## Configuration

RepoHarbor stores its config in the user config directory.

Available settings include:

- GitHub token
- Download path
- Active profile
- Named token profiles
- Recent repositories
- Favorite repositories
- Cache settings

## Caching

- Repository trees are cached by default
- Cache entries are keyed by `owner/repo/ref`
- The default cache TTL is `30m`

## Output Files

RepoHarbor may create these outputs during normal use:

- downloaded files and folders
- zip or tar.gz archives
- download manifest JSON
- cache files
- configuration files

## Build

```bash
go test ./...
npm run build
```

## Repository Layout

- `cmd/repoharbor`: CLI entrypoint
- `internal/config`: config and profile storage
- `internal/download`: file and archive download logic
- `internal/github`: GitHub API and URL handling
- `internal/ui`: terminal interface
- `scripts`: npm postinstall helper
- `bin`: Node launcher scripts

## Notes

- RepoHarbor is designed for GitHub repositories only.
- Private repositories need a valid token.
- `--profile` selects a saved token profile for a run.
- `--cwd` downloads into the current working directory.
- `--no-folder` avoids creating a repo-named subfolder.

## License

MIT

## Contact

Portfolio: https://Nishantdev.space

Email: codewithevilxd@gmail.com
