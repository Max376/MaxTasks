# Architecture decisions

Architecture Decision Records (ADRs) capture technical choices that affect the
future design of MaxTasks. They are short, versioned documents stored beside
the code so contributors can understand why the project works the way it does.

## Accepted decisions

| ADR | Decision | Related issue |
| --- | --- | --- |
| [ADR-0001](0001-authentication-and-session-model.md) | Authentication and session model | [#5](https://github.com/Max376/MaxTasks/issues/5) |

## Proposed decisions

No proposed decisions are currently awaiting review.

## Create and review an ADR

1. Copy [`0000-adr-template.md`](0000-adr-template.md) to the next four-digit
   number and use a short, descriptive filename.
2. Set the status to `Proposed`, complete the context, decision, alternatives,
   consequences, and related issue, then update this index.
3. Open a pull request that links the issue with `Closes #<number>` (or
   `Refs #<number>` when the issue must remain open). Explain the decision and
   include the relevant validation or design evidence.
4. Reviewers check the alternatives, consequences, security impact, and effect
   on existing documentation. Resolve review feedback in the pull request.
5. After approval, change the ADR status to `Accepted` in the same pull request
   and keep the accepted decision listed above. A replacement decision marks
   the old ADR `Superseded` and links the new ADR.

See the [GitHub planning and documentation workflow](../GITHUB_WORKFLOW.md) and
the [pull-request template](../../.github/PULL_REQUEST_TEMPLATE.md) for the
repository-wide review rules.
