# PokéAPI 🏰⚔️

A **modern, protobuf-driven Pokémon API** written in Go.  
It exposes two core domains:

| Domain   | Version | Summary                                                    |
| -------- | ------- | ---------------------------------------------------------- |
| **Dex**  | `v1`    | Catalog & image metadata for every Pokémon                 |
| **Calc** | `v1`    | Battle‐oriented calculators (type match-ups, damage, etc.) |

---

## ✨ Features

- **gRPC + JSON/REST** gateway out of the box (powered by ConnectRPC).
- **Strict schema-first workflow** – Protocol Buffers managed with Buf.
- **Modular micro-services** (`/services`) so you can run only what you need.
- Auto-generated, fully-typed **Go SDK** (see `schema/gen/go`).
- Reproducible builds via Makefile & Docker.

---

## For Developers 🛠️

### Project Layout

```bash
.
├── schema/              # Protobuf sources & code-gen artefacts
│   ├── proto/           # → Authoritative *.proto files
│   ├── buf.*.yaml       # Buf configs (lint, breaking checks, remote registry)
│   ├── gen/go/          # Go code generated by `buf generate`
│   └── Makefile         # Helper targets (lint, gen, clean)
├── services/            # Individual micro-services (Dex, Calc, …)
└── README.md
```

### Prerequisites

| Tool     | Minimum Version | Purpose                        |
| -------- | --------------- | ------------------------------ |
| Go       | **1.22**        | Core language/runtime          |
| Buf      | **1.31**        | Protobuf lint & generation     |
| Docker   | 24.x            | Local orchestration (optional) |
| `protoc` | 3.25            | Fallback if you skip Buf       |

### Common Tasks

| Action                    | Command           |
| ------------------------- | ----------------- |
| **Lint all protos**       | `make proto-lint` |
| **Regenerate Go SDK**     | `make gen`        |
| **Run unit tests**        | `go test ./...`   |
| **Spin up every service** | `make up`         |
| **Tear down containers**  | `make down`       |

## API Reference

| Service | Endpoint                   | Notes                                      |
| ------- | -------------------------- | ------------------------------------------ |
| Dex     | `GET /v1/dex/pokemon/{id}` | Returns full Pokémon record                |
|         | `GET /v1/dex/image/{id}`   | Original & thumbnail sprite URLs           |
| Calc    | `POST /v1/calc/attack`     | JSON body `{ attackerType, defenderType }` |

Full protobuf specification lives in schema/proto/\*\*.

## Contributing 🤝

1. Fork · create feature branch.
2. `make lint test` must pass.
3. Open PR – GitHub Actions will run Buf breaking checks and unit tests.
4. One maintainer approval → merge.

### Coding Guidelines

- Follow Effective Go + goimports.
- Keep protobuf changes backward‐compatible (buf breaking gate).
- One feature / bug-fix per PR.

## Roadmap 🗺️

- Move sets & abilities endpoints
- Competitive match replay parser
- Web UI explorer (Next.js)
- Helm chart & Kustomize overlays for prod deploy
