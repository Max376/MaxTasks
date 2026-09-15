# GitHub Planning and Documentation Workflow

GitHub is the project-management source of truth for MaxTasks. The repository documentation explains the product and technical decisions, while GitHub Projects, Issues, and Milestones show what is planned and what is in progress.

## Project board

The [MaxTasks Roadmap project](https://github.com/users/Max376/projects/4) tracks all planned work. Every issue should be added to the project and assigned to a milestone.

The project uses these fields:

- **Status** — Todo, In Progress, In Review, Done, or the built-in project status options.
- **Priority** — Now, Next, or Later.
- **Area** — Product, Backend, Frontend, DevOps, Docs, or Security.
- **Size** — XS, S, M, or L.
- **Release** — Foundation, MVP, PWA, Quality, or Release.

Recommended project views:

- **Roadmap** — grouped by milestone or Release.
- **Current work** — filtered to Todo, In Progress, and In Review.
- **Backlog** — open issues grouped by Area and sorted by Priority.
- **Release readiness** — Quality and Release milestones only.

## Milestones

Milestones represent delivery phases rather than individual calendar sprints:

1. **Foundation** — repository, local runtime, API, database, and migrations.
2. **MVP** — authentication and core task management.
3. **PWA** — responsive web experience and Android installation.
4. **Quality** — API documentation, tests, CI, logging, and security review.
5. **Release** — demo, recovery procedure, screenshots, and tagged MVP release.

## Issue workflow

1. Create an issue from the appropriate template.
2. Write a concrete goal and testable acceptance criteria.
3. Add labels for type, area, and priority.
4. Assign the issue to a milestone and the Roadmap project.
5. Move the project status as work progresses.
6. Close the issue only when the acceptance criteria and definition of done are satisfied.

## Pull requests

Implementation work uses a short-lived branch and a pull request linked to its issue. The pull-request template records the change, verification, documentation impact, and review notes.

## Architecture decisions

Architecture decisions use the ADR workflow in
[`docs/decisions/README.md`](decisions/README.md):

1. Copy [`docs/decisions/0000-adr-template.md`](decisions/0000-adr-template.md)
   to the next numbered ADR file.
2. Set the new ADR to `Proposed`, describe its context, decision, alternatives,
   and consequences, and link the related GitHub issue.
3. Open a pull request for the ADR and explain the design evidence and affected
   documentation. Review happens in the pull request before the decision is
   accepted.
4. Resolve review feedback, change the status to `Accepted`, and update the
   accepted-decisions index in the same pull request.
5. If a later decision replaces it, mark the old ADR `Superseded` and link the
   replacement.

The authentication and session model is the first accepted ADR. It establishes
the security contract that the MVP authentication implementation in issue #5
must follow; this issue does not implement authentication.

## Repository documentation

- `README.md` — short project entry point and links.
- `docs/PROJECT_PLAN.md` — vision, scope, architecture, milestones, and delivery criteria.
- `docs/GITHUB_WORKFLOW.md` — this planning and documentation workflow.
- `docs/decisions/README.md` — accepted architecture decisions and the ADR review process.
- `docs/decisions/` — accepted architecture decision records.
- `docs/api/` — OpenAPI specification and API examples.
- `docs/operations/` — deployment, backup, restore, and troubleshooting guides.
