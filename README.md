# _n_-vector

_Functions for performing geographical position calculations using n-vectors_

[![Current version][badge-version-image]][badge-version-link]
[![Build status][badge-build-image]][badge-build-link]
[![Test coverage][badge-coverage-image]][badge-coverage-link]

[badge-build-image]:
  https://img.shields.io/github/actions/workflow/status/ezzatron/nvector-go/ci.yml?branch=main&style=for-the-badge
[badge-build-link]:
  https://github.com/ezzatron/nvector-go/actions/workflows/ci.yml
[badge-coverage-image]:
  https://img.shields.io/codecov/c/gh/ezzatron/nvector-go?style=for-the-badge
[badge-coverage-link]: https://codecov.io/gh/ezzatron/nvector-go
[badge-version-image]:
  https://img.shields.io/github/v/tag/ezzatron/nvector-go?include_prereleases&sort=semver&logo=go&label=github.com%2Fezzatron%2Fnvector-go&style=for-the-badge

[badge-version-link]: https://pkg.go.dev/github.com/ezzatron/nvector-go

This library is a port of the [Matlab n-vector library] by [Kenneth Gade]. All
original functions are included, although the names of the functions and
arguments have been changed in an attempt to clarify their purpose. In addition,
this library plus some extras for vector and matrix operations needed to solve
the [10 examples from the n-vector page].

[matlab n-vector library]: https://github.com/FFI-no/n-vector
[kenneth gade]: https://github.com/KennethGade
[10 examples from the n-vector page]: https://www.ffi.no/en/research/n-vector

## Installation

```sh
go get github.com/ezzatron/nvector-go
```

## Methodology

If you look at the test suite for this library, you'll see that there are very
few concrete test cases. Instead, this library uses model-based testing, powered
by [rapid], and using the [Python nvector library] as the "model", or reference
implementation.

[rapid]: https://github.com/flyingmutant/rapid
[python nvector library]: https://nvector.readthedocs.io/

In other words, this library is tested by generating large amounts of "random"
inputs, and then comparing the output with the Python library. This allowed me
to quickly port the library with a high degree of confidence in its correctness,
without a deep understanding of the underlying mathematics.

If you find any issues with the implementations, there's a good chance that the
issue will also be present in the Python library, and an equally good chance
that I won't personally understand how to fix it 😅 Still, don't let that stop
you from opening an issue or a pull request!

## References

- Gade, K. (2010). [A Non-singular Horizontal Position Representation], The
  Journal of Navigation, Volume 63, Issue 03, pp 395-417, July 2010.
- [The n-vector page]
- Ellipsoid data taken from [chrisveness/geodesy]

[a non-singular horizontal position representation]:
  https://www.navlab.net/Publications/A_Nonsingular_Horizontal_Position_Representation.pdf
[the n-vector page]: https://www.ffi.no/en/research/n-vector
[chrisveness/geodesy]: https://github.com/chrisveness/geodesy
