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

## Repository documentation

- `README.md` — short project entry point and links.
- `docs/PROJECT_PLAN.md` — vision, scope, architecture, milestones, and delivery criteria.
- `docs/GITHUB_WORKFLOW.md` — this planning and documentation workflow.
- `docs/decisions/` — accepted architecture decision records.
- `docs/api/` — OpenAPI specification and API examples.
- `docs/operations/` — deployment, backup, restore, and troubleshooting guides.
