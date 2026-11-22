# Mika — Project and source analysis

Audit date: 2026-11-22

## Source of truth used
- User fork: `project10dollar-png/Mika`
- Working branch: `feature/hey-oka-branding`
- Exact source commit: `98c0756c2e62f6aa3d1c42fa568793cc5211d1b0`
- Original upstream: `wangzan101/Mika`, branch `main`, commit `2ca391c4c868a39d98d06639f327489baaebade8`

## Repository inventory
The recursive Git tree for the selected user-fork commit was retrieved completely (not truncated):
- 4,690 tree entries
- 4,096 files
- 594 directories
- 109,039,253 bytes of file content
- Full path/size/blob-SHA inventory is in [Mika file manifest](MIGRATION_FILE_MANIFEST.tsv).

The upstream original tree has 4,685 entries, 4,091 files, 594 directories, and 109,011,996 bytes. The user fork contains five additional entries and 27,257 additional bytes at the inspected commit.

## Project identity and build system
- Native Android application implemented with Kotlin and Gradle Kotlin DSL.
- Gradle modules: `:app` and `:core`.
- Root project name: `Mika`.
- Android namespace and application ID: `com.wangzan101.mika`.
- compileSdk 36, targetSdk 35, minSdk 29.
- Java compatibility and Kotlin toolchain: Java 17.
- Main app build configuration is `app/build.gradle.kts`.
- App version observed in that file: versionCode 52, versionName 0.1.51.
- Compose, AIDL, BuildConfig, Kotlin serialization, AndroidX, OkHttp, XZ, CommonMark, ML Kit, ZXing and Shizuku dependencies are configured.
- Main CI workflow: `.github/workflows/ci.yml`.

## Main project areas
- `app/`: Android application, UI, services, resources, assets, native libraries and Android tests.
- `core/`: reusable core implementation and JVM-side tests.
- `docs/`: user guides, feature explanations, setup instructions and demo media.
- `eval/`: evaluation tasks and test assets.
- `examples/`: examples and integration samples.
- `plugins/`: plugin-related implementation/resources.
- `tools/`: developer utilities, sandbox/runtime build scripts, CLI and relay-related tooling.
- `openspec/`: change proposals, specifications and supporting evidence.
- `licenses/`: third-party license notices and attribution.
- `gradle/`: Gradle wrapper and version catalog/configuration.

## Product capabilities described by the source README
- AI assistant that operates Android apps and web pages through screen interaction.
- Chat and model-provider configuration, with user-supplied credentials/accounts.
- Background-screen interaction, floating capsule and user interruption/continuation.
- Built-in browser, file handling and per-conversation Linux sandbox.
- Scheduled and conditional tasks.
- Phone-control setup using wireless debugging.
- CLI/Skill/MCP handoff from a computer.
- Ask and Autonomous operating modes.
- User data is described as staying on device; no Mika server is required for the core app.

## CI pipeline observed
The source CI workflow defines jobs for:
1. Core JVM tests.
2. Python CLI contract and translation/log-export tests.
3. Android unit tests, debug APK assembly and Android test compilation.
4. Go relay unit tests and Kotlin-to-relay integration tests.

## User-requested HEY OKA work
The user fork includes the branch `feature/hey-oka-branding`. Prior work in this conversation records the requested branding and voice changes:
- HEY OKA naming and robotic mosaic visual identity.
- Vosk Android offline speech recognition in place of Picovoice/Porcupine.
- Acoustic phrase grammar approximating “Hey Oka” with phrase `hey okay`; this is not a fully trained custom wake-word model.
- Foreground microphone service and persistent notification with stop action.
- Voice activation under Settings → Advanced → HEY OKA.
- No automatic boot start.
- Android background activity launch limits mean notification interaction may be required to open the voice flow.
These items are based on prior project work notes; the complete source code still needs a detailed diff/build review before asserting each behavior is present and correct in this exact commit.

## Important large assets
The upstream tree includes a large runtime archive at `app/src/main/assets/runtime/rootfs.tar.xz` (about 59 MB), along with demo media and launcher artwork. Do not omit binary assets during a real migration.

## Migration status and limitations
This document and the TSV manifest record the repository analysis and every path in the selected source Git tree. They do NOT constitute a full source-code copy.
A complete migration still requires transferring all Git objects/files (including binary assets), preserving the desired branch/commit history, and verifying the target tree against this source commit. Do not report the full migration complete until the destination tree matches the selected source tree (or a documented intentional delta).
Never put access tokens in source files, commits, issues, or chat.
