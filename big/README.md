# Almost drop-in replacements for big.Int, big.Rat, big.Float

## Zero-value is **NOT** ready to use.

In the Stdlib versions, the zero value is ready is use and set to 0. 

```
var x big.Int
x.Sign()
```

The BigNum versions will promply panic with a nil pointer exception.

Use one of the New** variants to make a value, or initialize a value use Set**.

## Requires Go 1.24+

These use the Go 1.24 [https://go.dev/blog/cleanups-and-weak](runtime.AddCleanup) calls to reclaim resources.  This design usings more allocations than the old finalizer methods, but maintains performance and reclaims memory faster.

## Future work

The wrapper classes and memory management ould be simplified future if [https://github.com/golang/go/issues/70224](Proposal 70224) is implemented.


