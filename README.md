# Multistring

## Introduction

`multistring` defines a new type, `Wrap`, that has the underlying type `[]string`, with methods defined on it that correspond to<sup>1</sup> the free functions in the standard library's `strings` package. Each method performs the equivalent of applying the corresponding function over the slice, like the `map` higher-order function would.

## Example

```go
w := multistring.Wrap{" apple  ", "banana ", "  celery"}
s := w.TrimSpace().ToUpper().Repeat(3).JoinBy(",")

println(s)
// APPLEAPPLEAPPLE,BANANABANANABANANA,CELERYCELERYCELERY
```

## Method signatures

The methods' signatures are derived from those of the corresponding functions, subject to two transformations:

1. The elements in the `Wrap` will be implicitly passed as the first string argument to the `strings` function 
2. Every return type `T` will be transformed to the slice type `[]T`, unless `T` is `string`, in which case it will be transformed to `Wrap`, to allow chaining.

## Other methods

There are two additional methods on `Wrap` that are not derived from `strings`, `JoinBy` and `MapString`:

* `JoinBy` can be used for cases where, as a final step, you want to join a group of strings (contained in a `Wrap`) with the same separator, rather than join a group of strings (passed to the `Join` method) with a group of separators contained in a `Wrap`.
* `MapString` is the analogue of `strings.Map`, but taking a function that takes a `string` (`func(string) string`), instead of a `rune`.

## Why did I do this? 

For a side project, I needed to apply functions from `strings` over slices of strings. After writing a few `for` loops, I decided to abstract it out into what is now `Wrap`, and then I thought - what if I could generate these method definitions?

A couple of days later...
