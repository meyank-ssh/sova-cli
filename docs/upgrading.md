# Upgrading Sova CLI

This guide provides instructions for upgrading your Sova CLI installation to the latest version.

## Upgrading from 0.1.0 to 0.1.1

### Quick Upgrade

If you installed Sova CLI using `go install`:

```bash
go install github.com/go-sova/sova-cli@latest
```

### Manual Upgrade

1. Download the latest release from the [Releases page](https://github.com/go-sova/sova-cli/releases)
2. Replace your existing binary with the new version
3. Verify the installation:
   ```bash
   sova-cli version
   ```

### Post-Upgrade Steps

1. Update existing projects:
   ```bash
   cd your-project
   sova-cli update
   ```
   This will:
   - Update `.gitignore` files
   - Fix import paths
   - Update Docker Compose configurations

2. Review and apply changes:
   - Check the new `.gitignore` patterns
   - Verify Docker Compose configurations
   - Update import paths if necessary

### Breaking Changes

Version 0.1.1 includes no breaking changes. However, some improvements require attention:

1. Docker Compose:
   - The `version` attribute has been removed
   - Review your existing `docker-compose.yml` files

2. Import Paths:
   - Routes have been moved to `internal/routes`
   - Update your imports if you've modified the generated code

### New Features

1. Enhanced `.gitignore` templates:
   - API projects now include Docker-specific ignores
   - CLI projects include build and distribution ignores

2. Improved Documentation:
   - New getting started guide
   - Detailed template documentation
   - Project structure guidelines

## Troubleshooting

If you encounter issues after upgrading:

1. Clean Go module cache:
   ```bash
   go clean -modcache
   ```

2. Regenerate Go module files:
   ```bash
   go mod tidy
   ```

3. Verify template updates:
   ```bash
   sova-cli doctor
   ```

## Support

If you need help with the upgrade:

1. Check the [documentation](README.md)
2. Open an issue on GitHub
3. Join our community channels

## Rolling Back

If you need to roll back to a previous version:

```bash
# Using go install
go install github.com/go-sova/sova-cli@v0.1.0

# Or download the specific release from GitHub
# https://github.com/go-sova/sova-cli/releases/tag/v0.1.0
``` 