# Loop engineering workflow

MaxTasks uses GitHub as the visible record of work and a coordinator-led loop for implementation and verification. The workflow keeps implementation and approval separate while allowing a ticket to receive up to three coding attempts.

## Starting a ticket

Give the coordinator a GitHub issue number, for example:

```text
Work on issue #3 using the loop engineering workflow.
```

The coordinator reads the issue and its acceptance criteria, assigns project fields, and starts attempt 1. The issue should be in the Roadmap project before coding begins.

## Handoff sequence

```text
Coordinator
    |
    v
Coding agent -- implementation branch and pull request
    |
    v
Testing agent -- checks acceptance criteria and evidence
    |
    +--> PASS: merge PR, mark Done, close issue
    |
    +--> FAIL and attempts remain: actionable feedback to coding agent
    |
    +--> FAIL on attempt 3: mark blocked and report failure
```

The testing agent must review the original issue rather than judging whether the implementation merely looks plausible. A PASS requires the acceptance criteria, relevant tests, and documentation expectations to be satisfied.

## What each agent reports

The coding handoff contains the branch, commit SHA, changed behavior, validation commands, and known limitations. The testing handoff contains the result, commands run, evidence, and file references. A FAIL must identify the exact acceptance criterion that is not satisfied and the smallest useful repair.

## GitHub Project usage

Set **Agent phase** and **Attempt** on the Roadmap item at every handoff:

| Phase | Meaning |
| --- | --- |
| Coding | The coding agent is implementing the current attempt. |
| Testing | The branch is ready for verification. |
| Fix Required | Testing found a problem and feedback was returned. |
| Accepted | Testing passed and the change was merged. |
| Failed | Three attempts failed; human intervention is required. |

The labels make the same state searchable in Issues. The pull request is the review surface; the issue remains the product record.

## Merge and failure rules

Only a testing-agent PASS permits a merge. The coding agent does not merge its own work. After a PASS, the testing agent merges the pull request, updates the project item, and closes the issue. After a third FAIL, the testing agent leaves the pull request open, applies `agent:blocked`, records the evidence, and stops.

The full role contract is in [`AGENTS.md`](../AGENTS.md). Project scope and milestones remain in [`PROJECT_PLAN.md`](PROJECT_PLAN.md).
