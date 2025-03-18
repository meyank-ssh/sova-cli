# Changelog

All notable changes to Sova CLI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2025-03-18

### Added
- Comprehensive `.gitignore` templates for both API and CLI projects
- Detailed documentation in `docs/` directory
- Getting started guide with project setup instructions
- Templates documentation explaining available options and features

### Fixed
- Import path issues in generated code
- Removed obsolete `version` attribute from `docker-compose.yml` template
- Project structure organization for better code management
- Synchronization issues with `go.mod` and `go.sum` files

### Changed
- Improved project templates organization
- Enhanced documentation structure
- Updated Docker Compose configuration for better compatibility
- Reorganized internal package structure for cleaner architecture

### Development
- Added `testify` package for enhanced testing capabilities
- Improved test coverage for project components
- Better error handling in project generation
- Enhanced template validation

## [0.1.0] - 2025-03-17

### Added
- Initial release of Sova CLI
- Basic project generation for API and CLI projects
- Docker Compose integration for API projects
- Command management system for CLI projects
- Basic documentation 