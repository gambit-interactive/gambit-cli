#  Gambit Framework CLI

[![Go version](https://img.shields.io/badge/go-1.18%2B-blue.svg)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/gambit-interactive/gambit-cli/actions)

**Gambit CLI** is the official command-line interface for the Gambit Framework. It provides a set of powerful tools to help you create, manage, and scale your Gambit applications with ease.

Built with Go and powered by Cobra, this CLI is designed to be fast, cross-platform, and intuitive.

## ✨ Features

- **Project Scaffolding:** Quickly create new Gambit projects with a standard directory structure.
- **Code Generation:** (Future) Generate boilerplate code for controllers, models, and more.
- **Cross-Platform:** A single, statically-linked binary that runs on Windows, macOS, and Linux.
- **Intuitive Commands:** A clean, `dotnet`-inspired command structure that is easy to remember.

## 🚀 Installation

### From Source

To install the CLI from the source code, you'll need Go (version 1.18 or higher) installed on your system.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/gambit-interactive/gambit-cli.git
    cd gambit-cli
    ```

2.  **Build the binary:**
    ```bash
    go build -o gambit
    ```
    On Windows, use:
    ```bash
    go build -o gambit.exe
    ```

3.  **(Optional) Move the binary to your system's PATH:**
    To make the `gambit` command available globally, move the compiled binary to a directory included in your system's `PATH`.

    - **Linux/macOS:**
      ```bash
      sudo mv gambit /usr/local/bin/
      ```
    - **Windows:**
      Move `gambit.exe` to a folder like `C:\Windows\System32` or another directory that is in your `PATH` environment variable.

### From Releases (Recommended)

Once official releases are available, you can download the pre-compiled binary for your operating system directly from the [Releases](https://github.com/gambit-interactive/gambit-cli/releases) page.

## 💻 Usage

The Gambit CLI provides a rich set of commands to streamline your development workflow.

### Project Management

| Command                                     | Description                                                    |
| ------------------------------------------- | -------------------------------------------------------------- |
| `gambit create project <name>`              | Creates a new project by cloning the Gambit Framework.         |
| `gambit create project <name> -o <path>`    | Creates the project in a specific output directory.            |
| `gambit create project <name> --branch dev` | Clones a specific branch (e.g., `dev`, `master`). Default: `stable`. |

### Development Workflow

| Command        | Description                                                              |
| -------------- | ------------------------------------------------------------------------ |
| `gambit dev`   | Runs the app with hot-reloading (requires `air` to be installed).        |
| `gambit run`   | Compiles and runs the application once without hot-reloading.            |
| `gambit build` | Compiles and builds a production-ready binary named `gambit-app`.        |
| `gambit test`  | Runs all tests for the project (`go test ./...`).                        |

### Docker Integration

| Command              | Description                                                      |
| -------------------- | ---------------------------------------------------------------- |
| `gambit docker:build`| Builds the Docker image for the application (needs a `Dockerfile`).|
| `gambit docker:up`   | Starts all services defined in `docker-compose.yml`.             |

### General Commands

- **Get help:**
  ```bash
  gambit --help