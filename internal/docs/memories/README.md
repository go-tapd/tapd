# Project Memories

This directory stores durable, reusable project knowledge for future agents.

Memory belongs here when it is project-specific and likely to help across multiple tasks. Do not store one-off task notes, temporary debugging details, raw chat logs, secrets, credentials, tokens, or private keys.

## How To Use

1. Read `index.md` before non-trivial work.
2. Open only the memory files relevant to the current task.
3. Update memories when the user requests memory maintenance or when the current task explicitly includes maintaining reusable project knowledge.
4. Recommend a memory update instead of editing memory files when memory maintenance is outside the current task scope.
5. Keep entries concise, actionable, and easy to review in git.
6. Update `index.md` when adding, removing, renaming, or materially changing memory files.

## Default Files

- `project.md`: stable project context, conventions, and structure.
- `workflows.md`: reusable commands, release steps, CI routines, and maintenance procedures.
- `lessons.md`: pitfalls, verified fixes, and reusable debugging lessons.

Memory directory: `internal/docs/memories`
