# Contributing to Sova CLI

First off, thank you for considering contributing to Sova CLI! It's people like you that make Sova CLI such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

* Use a clear and descriptive title
* Describe the exact steps which reproduce the problem
* Provide specific examples to demonstrate the steps
* Describe the behavior you observed after following the steps
* Explain which behavior you expected to see instead and why
* Include screenshots if possible

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

* Use a clear and descriptive title
* Provide a step-by-step description of the suggested enhancement
* Provide specific examples to demonstrate the steps
* Describe the current behavior and explain which behavior you expected to see instead
* Explain why this enhancement would be useful

### Pull Requests

* Fork the repo and create your branch from `main`
* If you've added code that should be tested, add tests
* If you've changed APIs, update the documentation
* Ensure the test suite passes
* Make sure your code lints
* Issue that pull request!

## Development Setup

1. Fork and clone the repository
   ```bash
   git clone https://github.com/yourusername/go-sova.git
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Run tests
   ```bash
   go test ./...
   ```

4. Build the project
   ```bash
   go build
   ```

## Project Structure

```
.
├── cmd/          # Command implementations
├── internal/     # Private application code
├── pkg/         # Public libraries
├── docs/        # Documentation
└── tests/       # Test files
```

## Coding Style

* Follow standard Go project layout
* Use `gofmt` for formatting
* Follow Go naming conventions
* Write descriptive commit messages
* Add tests for new features

## Testing

* Write unit tests for new features
* Ensure all tests pass before submitting PR
* Include integration tests when needed
* Test edge cases and error conditions

## Documentation

* Update README.md if needed
* Add godoc comments to public functions
* Update wiki pages if needed
* Include examples for new features

## Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

## Pull Request Process

1. Update the README.md with details of changes if needed
2. Update the docs/ with details of changes if needed
3. The PR will be merged once you have the sign-off of at least one maintainer

## Release Process

1. Update version number in relevant files
2. Update CHANGELOG.md
3. Create a new GitHub release
4. Tag the release with version number
5. Update installation instructions if needed

## Questions?

Feel free to open an issue with your question or contact the maintainers directly.

Thank you for contributing to Sova CLI! 🚀 