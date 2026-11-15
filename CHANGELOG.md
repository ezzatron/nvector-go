# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to [Semantic
Versioning].

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## Unreleased

### Changed

- Documented that n-vectors passed to the ECEF and rotation matrix conversion
  functions must have unit length.

### Fixed

- Fixed the sign of the Y Euler angle returned by `RotationMatrixToEulerZYX` at
  the Euler singularity.
- Fixed `Vector.Norm` and `Vector.Normalize` for zero and very large vectors.

## [v0.2.0] - 2024-05-28

[v0.2.0]: https://github.com/ezzatron/nvector-go/releases/tag/v0.2.0

### Added

- Added a `Sphere` function for creating spherical ellipsoids.

### Removed

- **\[BREAKING\]** The `WGS84Sphere` variable has been removed. Use the `Sphere`
  function with your chosen radius instead.

## [v0.1.3] - 2024-05-28

[v0.1.3]: https://github.com/ezzatron/nvector-go/releases/tag/v0.1.3

### Fixed

- Fixed minor documentation issues.

## [v0.1.2] - 2024-05-27

[v0.1.2]: https://github.com/ezzatron/nvector-go/releases/tag/v0.1.2

### Fixed

- Added a workflow to publish the reference docs to pkg.go.dev.

## [v0.1.1] - 2024-05-27

[v0.1.1]: https://github.com/ezzatron/nvector-go/releases/tag/v0.1.1

### Fixed

- Added examples with illustrations to README.
- Fixed a typo in the README.

## [v0.1.0] - 2024-05-27

[v0.1.0]: https://github.com/ezzatron/nvector-go/releases/tag/v0.1.0

### Added

- Initial release.
