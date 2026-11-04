# Dark factory policy

- Take the human's goal once and deliver a reviewed PR; the human merges.
- Use OpenAI only, newest GPT-6 generation and latest variant in each tier. Verified 2026-11-04 against the [official model list](https://learn.chatgpt.com/docs/models) and [API pricing](https://developers.openai.com/api/docs/pricing). Prices below are standard USD per million uncached input / output tokens; ascending price: Luna < Terra < Sol < Astra. No price tie; do not reuse models across roles unless no alternative exists. Ask the human only to resolve a tie in model selection.
- Orchestrator/Planner: `gpt-6.1-terra`, high reasoning, $1 / $5. Plan and coordinate; never code; below the priciest tier.
- Executor: `gpt-6.1-sol`, $2 / $10. Best coding model below the priciest tier; implement the planned behavior.
- Auditor: `gpt-6-astra`, high effort, $10 / $50. Most capable model; fresh context. Return ACCEPT or REJECT; never fix. REJECT if tests echo implementation, anything is weakened or skipped, a test is deleted, a protected file changes without explicit human sign-off, a dependency is unflagged, or any policy or goal rule is broken. Also try to break the build with inputs missing from tests.
- Tooler: `gpt-6-luna`, $0.10 / $0.50. Cheapest model; handles setup and CI on its own, without Executor or Auditor.
- Give each subagent only its task, with the goal, applicable rules and necessary evidence; never inherited conversation. Subagents must not spawn subagents. Run independent tasks in parallel, each in its own `herdr` pane. Use herdr's blocking wait (`herdr agent prompt <agent> '<task>' --wait` or `herdr agent wait <agent>`); never poll in a loop.
- Work only in worktrees created and removed via `wt`, never in the primary checkout. Start every PR branch from freshly fetched `main`: `git fetch origin main`, then `wt switch --create <branch> --base origin/main --no-hooks --path .worktrees/<branch>`. Remove with `wt remove <branch> --no-hooks --foreground`. These commands were checked against installed `wt` help; do not configure or run hooks.
- Use `gh` for GitHub operations. Never push to `main`; leave merging to the human.
- Before any PR, `task gate` must pass locally. Include its result and a per-file coverage table in the PR, using existing standard tools; no custom tooling.
- Go: one table-driven test using `t.Parallel()` per exported function; one blackbox (`package <name>_test`) test file per source file. Assert the goal's behavior, not implementation output. Commit tests before the code they test. Keep CLI `main` thin. Require 90% coverage. Never weaken tests.
- Protected; explicit human sign-off required: Taskfile, lint config, CI, Dependabot, `AGENTS.md`, `spec/target.html`, repository settings and rulesets.
- Guardrails and all these rules hold unless the human explicitly allows an exception. If the same failure persists after two rounds, stop and escalate.
- Codex reads `AGENTS.md` directly. If another tool reads a different instructions filename, symlink that filename to `AGENTS.md`; use a copy only if symlinks fail.
