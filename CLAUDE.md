# CLAUDE.md

## Repository policy

Read and follow `AGENTS.md` before making changes.

`AGENTS.md` contains the project-wide architecture, engineering, TDD, reduced SDD, Git, validation, and scope rules.

Do not reinterpret or duplicate those policies here.

## Working style

* Investigate before editing.
* Read only the parts of the repository required to understand the current task, expanding investigation when evidence requires it.
* Prefer references to existing code and files over assumptions.
* Keep implementations incremental.
* Do not turn focused tasks into repository-wide redesigns.
* When behavior changes, follow the TDD workflow defined in `AGENTS.md`.

## Planning

Do not use Plan Mode mechanically for every task.

For small, local, obvious changes, investigate briefly and implement directly.

Use Plan Mode when the task involves meaningful architectural decisions, unfamiliar or cross-cutting code, multiple interacting modules, persistence semantics, concurrency, significant refactoring, meaningful risk, or unclear requirements.

A plan should resolve uncertainty before implementation rather than restating obvious coding steps.

If investigation reveals a decision with material trade-offs that has not already been approved, stop implementation and present the alternatives rather than silently choosing a new architecture.

## Specs

Follow the reduced SDD policy in `AGENTS.md`.

* Do not create specs for trivial work.
* When an approved spec exists, treat its behavior, invariants, scope, and architectural decisions as authoritative.
* Solve unconstrained implementation details locally and simply.
* If implementation requires violating or materially changing the spec, surface the conflict before proceeding.

## Context management

* Keep the working context focused.
* Prefer targeted file reads and searches over loading large unrelated portions of the repository.
* Do not repeatedly re-read unchanged files without a reason.
* Preserve important discoveries in code, tests, specs, or the final report rather than relying solely on transient conversational memory.
* Do not modify `AGENTS.md` or `CLAUDE.md` merely because you found a different preference while implementing a task.
* Propose durable process changes for review unless the task explicitly asks for those files to be changed.

## Implementation

* Favor straightforward Go over clever abstractions.
* Keep public APIs small.
* Keep vendor-specific behavior inside adapters.
* Do not create speculative abstractions for future agents or features.
* Do not hide failures, weaken tests, or silently skip validation.
* Avoid unrelated formatting, renaming, cleanup, or refactoring.

## Completion

Before declaring work complete:

* inspect the final diff;
* run validation appropriate to the change;
* report the behavior implemented;
* report tests and validation and their results;
* report significant implementation decisions;
* report commits created;
* report anything that still requires attention.

Do not merge changes unless explicitly instructed.
