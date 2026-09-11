# AGENTS.md

## Project

Context is an open-source, local-first tool for preserving work context across AI coding agents.

The immediate goal is simple: when one coding agent reaches a usage limit or must be replaced, another agent should be able to continue the same task without the developer having to reconstruct the work manually.

The project may grow beyond agent handoff, but current needs must drive the architecture. Do not build speculative infrastructure for possible future features.

## Engineering principles

* Keep the system simple, explicit, and easy to reason about.
* Prefer deep modules with small public APIs.
* Complexity should remain behind module boundaries rather than leak into callers.
* Abstract real boundaries, especially external systems.
* Do not introduce abstractions merely because something might change someday.
* Prefer concrete implementations until a second implementation or a genuine testing boundary demonstrates the need for an interface.
* Favor composition over cleverness.
* Optimize for maintainability, correctness, and clarity before minimizing lines of code.
* Do not introduce architectural patterns without a concrete problem they solve.

## Architecture

* The core is written in Go.
* The application starts as a modular monolith.
* Domain and application logic must not depend directly on Claude Code, Codex, Git process details, SQLite implementation details, or a future UI.
* External coding agents must be integrated behind narrow adapters.
* Git is the source of truth for repository code and working-tree state.
* SQLite is intended to persist orchestration state such as tasks, runs, and checkpoints.
* Long-running external processes should be supervised explicitly.
* Use Go concurrency deliberately; goroutines and channels are tools, not architecture by themselves.
* The core must remain independent from any future CLI or desktop UI.
* Do not introduce microservices, a plugin system, event sourcing, a background daemon, distributed messaging, or similar infrastructure until a demonstrated requirement justifies it.

## Development workflow

* Never develop directly on `main`.
* Start each change from an up-to-date `main` and use a short-lived branch.
* Keep pull requests small and cohesive.
* A PR should represent one understandable change or vertical slice.
* Do not combine unrelated cleanup, refactoring, or feature work into the current change.
* Every PR must leave `main` in a valid, buildable state.
* Never merge a PR unless explicitly instructed.
* Do not rewrite, amend, or discard commits created by the user or another agent unless explicitly instructed.
* Commits should be small, coherent, and describe meaningful changes rather than arbitrary file boundaries.

## TDD

Behavioral code is developed with TDD.

Use this cycle:

1. Write a test that describes one desired behavior and observe it fail for the expected reason.
2. Implement the smallest reasonable change that makes the test pass.
3. Refactor while keeping the tests green.

Additional rules:

* Do not write large batches of tests followed by a large implementation.
* Test observable behavior and invariants rather than implementation details.
* Do not weaken, delete, or rewrite a valid test merely to make an implementation pass.
* A failing test that exposes a legitimate defect must be fixed in the implementation.
* Prefer real lightweight infrastructure in integration tests when practical.
* Prefer temporary SQLite databases, temporary Git repositories, and controlled subprocesses over extensive mocking.
* Use fakes at genuine external boundaries, such as AI agent protocols, when invoking the real external service would make tests slow, nondeterministic, expensive, quota-dependent, or network-dependent.
* Do not create meaningless tests solely to increase coverage.
* Documentation-only and configuration-only changes do not require artificial tests.

## Reduced spec-driven development

Specs exist to reduce uncertainty, not to create ceremony.

* For small, local, obvious changes, use TDD directly. Do not create a spec.
* For changes that affect multiple modules, introduce important behavior, change persistence semantics, or require non-trivial design decisions, write a small spec before implementation.
* For architectural changes or decisions with meaningful trade-offs, the design must be discussed and decided before implementation begins.
* A useful spec should capture the problem, expected behavior, invariants, important decisions, out-of-scope items, and small implementation slices.
* Specs describe what the system must guarantee.
* Specs should not prescribe unnecessary internal implementation details.
* If implementation reveals that an approved architectural assumption is wrong, do not silently redesign the system. Surface the issue and its trade-offs before proceeding.

## Go guidelines

* Write idiomatic, boring Go.
* Keep packages cohesive and focused on a domain or clear responsibility.
* Minimize exported identifiers.
* Avoid generic packages such as `utils`, `helpers`, `common`, or `misc`.
* Code should live with the concept that owns it.
* Prefer the standard library when it provides a clear solution.
* Add external dependencies only when they provide meaningful value that would be costly or risky to reproduce.
* Handle expected failures with errors, not panics.
* Add useful context when propagating errors while preserving their underlying meaning when callers may need to inspect them.
* Use `context.Context` for cancellation and lifecycle control where appropriate.
* Do not store contexts in long-lived domain objects.
* Every goroutine must have clear ownership and a termination path.
* Avoid goroutine leaks.
* Use channels when they clarify ownership or communication.
* Do not use channels where a normal function call or synchronous flow is simpler.
* Concurrency must not make domain behavior harder to understand.

## Module boundaries

* Keep external protocol details inside their adapters.
* Claude-specific output, session semantics, errors, and process behavior must not leak into the core domain.
* Codex-specific JSON-RPC, thread semantics, errors, and process behavior must not leak into the core domain.
* Apply the same rule to future agents.
* The domain should speak in Context concepts such as task, run, checkpoint, agent event, completion, failure, and usage limit rather than vendor-specific concepts whenever possible.
* Do not design a universal agent abstraction in advance.
* Grow the shared contract from behaviors actually required by supported agents.

## Scope discipline

* Before editing, inspect the relevant code and repository state.
* Implement the smallest coherent change that satisfies the current task.
* Do not perform opportunistic refactors unless they are necessary for the requested behavior.
* Do not add extension points, configuration, interfaces, or abstractions for hypothetical future requirements.
* If a task requires a materially broader change than requested, surface that before expanding the scope.
* Preserve existing behavior unless the task explicitly changes it.

## Validation

For Go changes, run the checks relevant to the modified code and, before considering a PR ready, normally run:

* `gofmt`
* `go vet ./...`
* `go test ./...`
* `go build ./...`

Use repository-specific commands when they later become more authoritative than these defaults.

Review the final diff for:

* accidental changes;
* scope creep;
* unnecessary abstractions;
* missing tests;
* unintended architectural changes.

## Agent behavior

* Read the repository and existing documentation before making assumptions.
* Repository code and current tests take precedence over guesses about how the project works.
* Do not invent product requirements.
* Do not silently make architectural decisions that contradict this file or an approved spec.
* When uncertainty affects architecture, persistence, public contracts, security, destructive behavior, or scope, surface the uncertainty rather than guessing.
* Do not modify `AGENTS.md` opportunistically.
* Changes to project-wide development policy must be intentional and reviewed.

At completion of implementation tasks, report concisely:

* what changed;
* tests and validation executed;
* relevant design decisions;
* commits created;
* remaining risks, limitations, or follow-up work.
