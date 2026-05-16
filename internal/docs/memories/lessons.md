# Lessons Memory

Store reusable lessons and pitfalls here.

## Pitfalls

- The default Go build cache under the user Library path can be blocked by sandbox permissions. If package listing or tests fail with cache permission errors, retry with `GOCACHE=/private/tmp/go-build-cache-tapd` before treating it as a repository failure.
- Creating or switching Git branches may require escalated permissions when Git needs to write `.git/refs/...` lock files in this environment.
- Branch names using a path prefix can fail if a flat ref with the same first segment already exists. If `git switch -c docs/...` reports it cannot create a directory under `.git/refs/heads`, choose a flat branch name such as `feature-update-features-api-reference`.

## Fix Patterns

- For local project memory initialization in this repo, use `.agents/skills/memory-init/scripts/init_project_memory.py <project-root> --memory-dir internal/docs/memories`. The script creates missing memory files only and appends a managed Project Memory block to `AGENTS.md`.
- When the memory-init script fails creating the memory directory with `Operation not permitted`, rerun the same command with the required filesystem escalation instead of hand-creating the files.
