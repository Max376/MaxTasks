## Change

Describe the change and the user or engineering outcome it delivers.

## Related issue

Closes #<number>

## Agent workflow

- Issue: #<number>
- Branch: `agent/issue-<number>-<short-name>` or approved tracked workflow branch `chore/<short-name>`
- Attempt: <1|2|3>/3
- Agent phase: Coding / Testing / Fix Required / Accepted / Failed

### Acceptance criteria

Copy the complete acceptance criteria from the issue so the testing agent can review the same contract.

### Validation commands

List the exact commands the testing agent must run, including working-directory requirements.

### Coding result

State the commit SHA and summarize the implementation. Include known limitations.

### Testing result

Testing agent: return exactly `PASS` or `FAIL`.

### Evidence

Record command output, test results, and relevant file or line references.

### Feedback

For `FAIL`, copy the actionable repairs for the coding agent. For `PASS`, write `None`.

## Verification

- [ ] Tests or checks were added or updated where appropriate.
- [ ] Documentation was updated when behavior or setup changed.
- [ ] I checked that no credentials or sensitive data are included.

## Review notes

Mention tradeoffs, follow-up work, migration steps, or screenshots.
