## v0.11.0
FEATURES:
- Support coverage report
- Support cleanup report

ENHANCEMENTS:

BUG FIXES:

## v0.10.0

FEATURES:
- Support generate data source from Swagger GET example.
- Support generate dependency automatically for identity.

ENHANCEMENTS:

BUG FIXES:
- Fix error message in reports.

## v0.9.0

ENHANCEMENTS:
- Update mapping of azurerm v3.50.0

## v0.8.1

BUG FIXES:
- Fix invalid characters in folder name on Windows

## v0.8.0

FEATURES:
- Support error report

ENHANCEMENTS:
- Update mapping of azurerm v3.41.0

## v0.7.0

FEATURES:
- Generate a passed report or partially passed report

ENHANCEMENTS:
- Update mapping of azurerm v3.30.0

## v0.6.0

BUG FIXES:
- Dependency detection failed when working-dir is specified

## v0.5.0

ENHANCEMENTS:
- Update mapping of azurerm v3.24.0

BUG FIXES:
- Wrap keys which start with numbers.

## v0.4.0

ENHANCEMENTS:
- Update mapping of azurerm v3.22.0

## v0.3.0

FEATURES:
- Supports validate command to preview the resource changes.
- Supports -working-dir option to specify the working directory.

## v0.2.1

FEATURES:
- Generated document improvement: now differences are highlighted.

## v0.2.0

FEATURES:
- Support install Terraform automatically.
- Support `-raw` option, which allows user to use raw json payload. The default payload will use jsonencode function.
- Support `-overwrite` option, which allows user to overwrite existing configurations. The default behavior is appending test cases on the existing configurations.
- Support markdown report generation: The `test` command can generate markdown report when it found bugs for each test cases.

## v0.1.0
Initial release.