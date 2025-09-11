# Almost drop-in replacements for big.Int, big.Rat, big.Float

* [GMP](https://gmplib.org/)
* [MPFR](https://www.mpfr.org/)
* [MPC](https://www.multiprecision.org/mpc/)

## TODO

* Various functions are indicated by "TODO" in the source
* Complex number support is not started

## Requires Go 1.24+

These use the Go 1.24 [runtime.AddCleanup](https://go.dev/blog/cleanups-and-weak) calls to reclaim resources.  This design usings more allocations than the old finalizer methods, but maintains performance and reclaims memory faster.

## Performance

* Under 1,000 digits, the native libraries are likely to be faster.
* For very large numbers, GMP/MPFR are 8-20x faster 

## Future work

* The wrapper classes and memory management could be simplified future if [Proposal 70224](https://github.com/golang/go/issues/70224) is implemented.

## See Also

* [ncw/gmp](https://github.com/ncw/gmp)

