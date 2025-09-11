# Almost drop-in replacements for big.Int, big.Rat, big.Float

* [GMP](https://gmplib.org/)
* [MPFR](https://www.mpfr.org/)

## Zero-value is **NOT** ready to use.

This will be fixed in a future version but for now, the zero-value is not ready to use.

In the Stdlib versions, the zero value is ready is use and set to 0. 

```
var x big.Int
x.Sign()
```

The BigNum versions will panic with a nil pointer exception.

Use one of the New** variants to make a value, or initialize a value use Set**.

## Requires Go 1.24+

These use the Go 1.24 [runtime.AddCleanup](https://go.dev/blog/cleanups-and-weak) calls to reclaim resources.  This design usings more allocations than the old finalizer methods, but maintains performance and reclaims memory faster.

## Performance

* Under 1,000 digits, the native libraries are likely to be faster.
* Very large numbers are likely to 8-20x faster using GMP/MPFC.

## Future work

* The wrapper classes and memory management could be simplified future if [Proposal 70224](https://github.com/golang/go/issues/70224) is implemented.
* [MPC](https://www.multiprecision.org/mpc/) complex number support

