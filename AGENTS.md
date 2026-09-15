# MaxTasks agent workflow

This repository uses a coordinator, coding agent, and testing agent to deliver GitHub issues. The workflow is deliberately sequential: only one agent edits the ticket branch at a time, and the testing agent receives the branch after the coding agent reports a completed attempt.

## Roles

### Coordinator

The coordinator selects the issue, starts the coding attempt, hands the branch to testing, records the result, and enforces the three-attempt limit. The coordinator must never treat an implementation report as approval. A testing result with evidence is required.

### Coding agent

The coding agent implements the assigned product issue on a branch named `agent/issue-<number>-<short-name>`. Repository workflow and setup changes owned by the coordinator may use `chore/<short-name>` when they are tracked by a GitHub issue and the pull request links that issue with `Closes #<number>`. The coding agent must read the complete issue, preserve the acceptance criteria, run the documented checks, and report the commit, branch, checks, and any known limitations. It must not merge the branch or close the issue.

### Testing agent

The testing agent reviews the issue, the acceptance criteria, and the coding branch. It runs the relevant checks and inspects the diff for correctness, security, regressions, and documentation impact. It returns exactly one result: `PASS` or `FAIL`.

On `PASS`, the testing agent is authorized by this workflow to merge the pull request, mark the project item as Done, and close the issue. On `FAIL`, it must leave actionable feedback and return the branch to the coding agent. It must not merge a failing attempt.

## Attempt limit

An attempt is exactly one coding-agent implementation followed by one testing-agent review. The maximum is three attempts total:

```text
Attempt 1: initial implementation
Attempt 2: repair after testing feedback
Attempt 3: final repair after testing feedback
```

If attempt 3 fails, stop the loop. Mark the issue `agent:blocked`, set the project phase to `Failed`, and report the failure to the user with the evidence and remaining blocker. Do not merge or silently start a fourth attempt.

## Handoff contract

Every handoff must include:

```text
Issue: #<number>
Branch: agent/issue-<number>-<short-name>
Attempt: <1|2|3>/3
Acceptance criteria: copied from the issue
Validation commands: exact commands to run
Coding result: commit SHA and summary
Testing result: PASS or FAIL
Evidence: command output, test results, and file references
Feedback: required repairs when the result is FAIL
```

The coding agent receives the testing feedback verbatim on a repair attempt. The testing agent reviews the new commit against the original acceptance criteria and the previous feedback.

## GitHub state

Each issue belongs to the `MaxTasks Roadmap` project and a milestone. Keep these project fields current:

- **Agent phase** — Coding, Testing, Fix Required, Accepted, or Failed.
- **Attempt** — 1, 2, or 3.
- **Priority**, **Area**, **Release**, and **Status** — the existing planning fields.

Use the labels `agent:coding`, `agent:testing`, `agent:blocked`, `review:pending`, `review:passed`, or `review:changes-requested` to make the current handoff visible in the issue list.

## Branch and pull-request rules

- Product issue work uses `agent/issue-<number>-<short-name>`; coordinator-owned repository workflow/setup work may use `chore/<short-name>` only when tracked by an issue.
- The coding agent creates one branch for the issue and keeps repair commits on that branch.
- The pull request links the issue with `Closes #<number>`.
- The coding agent never pushes directly to `main`.
- The testing agent merges only after a `PASS` result and successful final checks.
- A failed third attempt leaves the pull request open for a human decision.
