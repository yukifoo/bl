# bl - Backlog CLI

A GitHub CLI-like tool for Backlog project management system.

## Overview

`bl` provides a command-line interface for Backlog that replicates the user experience of GitHub CLI (`gh`). It brings issues, pull requests, wikis, and other Backlog concepts to the terminal.

## Features

- 🚀 GitHub CLI-like command structure
- 📋 Comprehensive issue management
- 🔄 Pull request operations
- 📚 Wiki management
- 🗂️ Project navigation
- 🎨 Multiple output formats (table, JSON, CSV)
- 🤖 Interactive mode for complex operations
- 🔐 Secure authentication with API keys

## Installation

### Using Homebrew (macOS/Linux)

```bash
brew install bl
```

### Using Go

```bash
go install github.com/yourusername/bl@latest
```

### Download Binary

Download the latest release from the [releases page](https://github.com/yourusername/bl/releases).

## Quick Start

### Authentication

```bash
# Login to your Backlog space
bl auth login

# Check authentication status
bl auth status
```

### Basic Usage

```bash
# List issues assigned to you
bl issue list --assignee @me

# View issue details
bl issue view PROJ-123

# Create a new issue interactively
bl issue create

# List pull requests
bl pr list --state open

# View project information
bl project view PROJ
```

## Command Structure

```
bl <resource> <action> [flags]
```

### Available Resources

- `auth` - Authentication management
- `issue` - Issue operations
- `pr` - Pull request management
- `project` - Project information
- `wiki` - Wiki pages
- `repo` - Git repository management
- `config` - Configuration management

## Configuration

Configuration file is located at `~/.config/bl/config.yml`:

```yaml
spaces:
  default: myspace
  myspace:
    api_key: your-api-key
    space_key: myspace
    default_project: PROJ

format: table
editor: vim
```

## Examples

### Issue Management

```bash
# List all open issues in a project
bl issue list --project PROJ --state open

# Create issue with specific parameters
bl issue create --title "Bug fix" --project PROJ --type bug --priority high

# Add comment to an issue
bl issue comment PROJ-123 --body "Working on this"

# Close an issue
bl issue close PROJ-123 --comment "Fixed in PR #456"
```

### Pull Request Workflow

```bash
# List your pull requests
bl pr list --assignee @me

# Create a pull request
bl pr create --title "Feature: Add new API" --base main --head feature/new-api

# Review a pull request
bl pr review 123 --approve
bl pr review 123 --request-changes --body "Please update tests"

# Merge a pull request
bl pr merge 123 --squash
```

## Development

### Prerequisites

- Go 1.21 or later
- Make

### Building from Source

```bash
git clone https://github.com/yourusername/bl.git
cd bl
make build
```

### Running Tests

```bash
make test
```

### Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [GitHub CLI](https://github.com/cli/cli)
- Built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)
- Uses [backlog](https://github.com/kenzo0107/backlog) Go client library