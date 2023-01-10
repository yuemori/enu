## enu - Enumerate over slices, maps, channels...

✨ `yuemori/enu` is a Enumerator library based on Go 1.18+ Generics.

This package that provides functions for manipulationg collections like [C# IEnumerable](https://learn.microsoft.com/ja-jp/dotnet/api/system.collections.generic.ienumerable-1) or [Ruby Enumerable](https://docs.ruby-lang.org/en/3.2/Enumerable.html).

We prefer C# IEnumerable and Ruby Enumerable because there modules can be manipurated with a same interfaces and functions for various enumerable collections.

Note: This package inspire and used internally [samber/lo](https://github.com/samber/lo) functions. Thanks.

## Install

```
go get github.com/yuemori/enu
```

## Usage

You can import `enu` using:

```go
import "github.com/yuemori/enu"
```

Then use one of the example below:

```go
count := enu.From([]int{1, 2, 3, 4, 5}).Count()
// 5

r := enu.From([]int{1, 2, 3, 4, 5}).Filter(func(i int, _ int) bool {
  return i%2 != 0
}).ToSlice()
// []int{1, 3, 5}
```

## Concepts

### Basics

`enu` provides `Enumerators` and `Enumerables` .

`Enumerator` is type-specific IEnumerator for each collections.

`Enumerable` provides collection-specific enumerate behaviour.

example below:

```go
// SliceEnumerator provides enumerable behaviour for slice
enumerable := enu.NewSliceEnumerator([]int{1, 2, 3})

// Enumerator provides enumerator interfaces
enumerator := enu.New(enumerable)

r := enumerator.Count()
// 3
```

### Method chain

`Enumeartors` supports mathod chaining.

example below:

```go
r := enu.From([]int{1, 2, 3, 4, 5}).Filter(func(i int, _ int) bool {
  return i%2 != 0
}).Reverse().ToSlice()
// []int{5, 3, 1}
```

### Lazy Enumeration

Importantly, `Enumerator` postpone enumeration and enumerate values only an as-needed basis.

So the `Enumerator` only asks the `Enumerable` for the number it needs.

See below examples:

*Range*

```go
// `enu.FromNumericRange(1, math.MaxInt)` Returns range from 1 to infinity.
r := enu.FromNumericRange(1, math.MaxInt).Take(10).ToSlice()
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

*Generator*

```go
// fibonacci number generator
fibonacci := func() func(int) (int, bool) {
  n, n1 := 0, 1
  return func(i int) (int, bool) {
    fn := func(index int) (int, bool) {
      // Reset variable if index given zero.
      if index == 0 {
        n, n1 = 0, 1
      }
      v := n
      n, n1 = n1, n+n1
      return v, true
    }
    return fn(i)
  }
}

r := enu.FromFunc(fibonacci()).Take(10).ToSlice()
// []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
```

*Channel*

```go
ch := make(chan (int))
done := make(chan struct{})

go func() {
  defer close(ch)
  defer close(done)

  i := 0
  for {
    select {
    case <-done:
      break
    default:
      ch <- i
      i++
    }
  }
}()

r := enu.FromChannelWithDone(ch, done).Take(10).ToSlice()
// []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
<- done // close done channel after finished enumeration
```

and helper functions to use:

```golang
request1 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"https://google.com", nil}
}

request2 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"https://amazon.co.jp", nil}
}

request3 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"", errors.New("some error")}
}

// Aggregate multiple function result by async
r := enu.FromChannel(enu.Parallel(2, request1, request2, request3)).ToSlice()

// []enu.Tuple2[string, error]{
//   {"", errors.New("some error")},
//   {"https://amazon.co.jp", nil},
//   {"https://google.com", nil},
// }
```

*Cache Result*

```golang
ch := make(chan (int))

go func() {
  defer close(ch)

  for i := 1; i < 6; i++ {
    ch <- i
  }
}()

channel := enu.FromChannel(ch)

// `Find` supports lazy iteration.
// stop iteration but channel does not closed here.
r2, ok := channel.Find(func(i, _ int) bool {
  return i == 3
})
// 3, true

// Enumerator cache the results of previous,
// and resume iterate to needed.
// and channel closed if lazy function completed.
r3 := channel.ToSlice()
// []int{1, 2, 3, 4, 5} -> 1,2,3 is from cache. 4,5 from source.
```

### Enumerators

This package provides the below Enumerators:

- [Enumerator[T any]](#enumerator)
- [EnumeratorComparable[T comparable]](#enumeratorcomparable)
- [EnumeratorMap[K comparable, V any]](#enumeratormap)
- [EnumeratorNumeric[T constraints.Float | constraints.Integer]](#enumeratornumeric)
- [EnumeratorOrdered[T constraints.Ordered]](#enumeratorordered)

These allow for a common interface and operations specific to each collection.

### Enumerables

This package provides the blow enumerables:

- [RangeEnumerator](#rangeenumerator)
- [NumericRangeEnumerator](#numericrangeenumerator)
- [DateRangeEnumerator](#daterangeenumerator)
- [Generator](#generator)
- [ChannelEnumerator](#channelenumerator)
  - [channel helpers](#channelhelpers)

[Enumerators](#enumerators) require an enumerable interfaces to each collection.

Enumerables provides a useful interface for each collection and usecases.

## Spec: Enumerators

### Enumerator

`Enumerator` is basic enumerator to each collection.

- [New](#new)
- [From](#from)
- [FromChannel](#fromchannel)
- [FromChannelWithDone](#fromchannelwithdone)

### EnumeratorComparable

`Enumerator` is enumerator for `comparable` collections.

- [FromComparable](#fromcomparable)
- [ToComparable](#tocomparable)

### EnumeratorMap

`Enumerator` is enumerator for `map` collections.

- [FromMap](#frommap)
- [ToMap](#tomap)

### EnumeratorNumeric

`Enumerator` is enumerator for `constraints.Integer` and `constraints.Float` collections.

- [FromNumeric](#fromnumeric)
- [ToNumeric](#tonumeric)

### EnumeratorOrdered

`Enumerator` is enumerator for `constraints.Integer` and `constraints.Float` collections.

- [FromOrdered](#fromordered)
- [ToOrdered](#toordered)

## Spec: Enumerator functions

- [Each](#each)
- [ToSlice](#toslice)
- [Count](#count)
- [Filter](#filter)
- [Reject](#reject)
- [Nth](#nth)
- [Find](#find)
- [First](#first)
- [Last](#last)
- [Reverse](#reverse)
- [Sort](#sort) only supported below:
  - [EnumeratorNumeric](#enumeratornumeric)
  - [EnumeratorOrdered](#enumeratorordered)
- [SortBy](#sort_by)
- [IsAll](#isall)
- [IsAny](#isany)
- [Take](#take)
- [ToMap](#tomap)
- [Uniq](#uniq) only supported below:
  - [EnumeratorComparable](#enumeratorcomparable)
  - [EnumeratorNumeric](#enumeratornumeric)
  - [EnumeratorOrdered](#enumeratorordered)
- [Contains](#contains) only supported below:
  - [EnumeratorComparable](#enumeratorcomparable)
  - [EnumeratorNumeric](#enumeratornumeric)
  - [EnumeratorOrdered](#enumeratorordered)
- [IndexOf](#contains) only supported below:
  - [EnumeratorComparable](#enumeratorcomparable)
  - [EnumeratorNumeric](#enumeratornumeric)
  - [EnumeratorOrdered](#enumeratorordered)
- [Min](#min) only supported below:
  - [EnumeratorOrdered](#enumeratorordered)
  - [EnumeratorNumeric](#enumeratornumeric)
- [Max](#max) only supported below:
  - [EnumeratorOrdered](#enumeratorordered)
  - [EnumeratorNumeric](#enumeratornumeric)
- [Sum](#sum) only supported below:
  - [EnumeratorNumeric](#enumeratornumeric)
- [Keys](#keys) only supported below:
  - [EnumeratorMap](#enumeratormap)
- [Values](#values) only supported below:
  - [EnumeratorMap](#enumeratormap)
- [GetEnumerator](#getenumerator)

## Spec: Enumerables

### RangeEnumerator

`RangeEnumerator` is basic range enumerator for

Implements `RangeValuer` interface if you want custom range.

- [FromRange](#fromrange)
- [RangeValuer](#rangevaluer)

### NumericRangeEnumerator

`NumericRangeEnumerator` is range enumerator for `constraints.Integer` and `constraints.Float` .

- [FromNumericRange](#newnumericrange)
- [FromNumericRangeWithStep](#newnumericrangewithstep)

### DateRangeEnumerator

`DateRangeEnumerator` is range enumerator for `time.Time` with `time.Duration` .

- [FromDateRange](#fromdaterange)

### Generator

`Generator` provides generator pattern from func.

- [FromFunc](#fromfunc)

### ChannelEnumerator

`ChannelEnumerator` is enumerator for `chan(T)` .

- [FromChannel](#fromchannel)
- [FromChannelWithDone](#fromchannelwithdone)

And see the [concurrent functions](#cuncurrent--functions).

## Spec: helpers

### Concurrent functions

- [Repeat](#repeat)
- [Parallel](#parallel)
- [Retry](#retry)

### Aggregate functions

- [Map](#map)
- [Reduce](#reduce)

### Helper types

- [Tuple2-9](#tuple)

## Documentation

### New

Returns an `*Enumerator[T]` with `Enumerable` argument.

```go
enumerator := enu.NewSliceEnumerator([]int{1, 2, 3})
r := enu.New(enumerator).Count()
// 3
```

### From

Returns an `Enumerator` with `slice` argument.

```go
r := enu.From([]int{1, 2, 3}).Count()
// 3
```

### FromChannel

Returns an `*Enumerator[T]` with `channel` argument.

```go
ch := make(chan (int))

go func() {
  defer close(ch)

  for i := 1; i < 6; i++ {
    ch <- i
  }
}()

r := enu.FromChannel(ch).ToSlice()
// []int{1, 2, 3, 4, 5}
```

### FromChannelWithDone

Returns an `*Enumerator[T]` with two `channel` argument.

Expects T as the first argument and Done channel as the second argument.

```go
ch := make(chan (int))
done := make(chan struct{})

go func() {
  defer close(ch)
  defer close(done)

  i := 0
  for {
    select {
    case <-done:
      break
    default:
      ch <- i
      i++
    }
  }
}()

r := enu.FromChannelWithDone(ch, done).Take(10).ToSlice()
// []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

### FromComparable

Returns an `*EnumeratorComparable[T]` with `comparable` argument.

```go
r := enu.FromComparable([]int{1, 1, 2, 3, 3}).Uniq().ToSlice()
// []int{1, 2, 3}
```

### ToComparable

Returns an `*EnumeratorComparable[T]` with `Enumerable[T comparable]` argument.

```go
// Enumerable[T] does not implement Uniq()
e := enu.From([]int{1, 1, 2, 3, 3})

// Convert to `EnumerableComparable` to be use.
r := enu.ToComparable(e).Uniq().ToSlice()
// []int{1, 2, 3}
```

### FromMap

Returns an `*EnumeratorMap[T]` with `map` argument.

```go
r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Keys()
// []int{1, 2, 3}
```

### ToMap

Returns an `*EnumeratorMap[T]` with `Enumerable[KeyValuePair[K, V]]` argument.

```go
// `Enumerator` does not impelements `Keys()`
e := enu.From([]enu.KeyValuePair[int, string]{
  {Key: 1, Value: "foo"},
  {Key: 2, Value: "bar"},
  {Key: 3, Value: "baz"},
})

// Convert to `EnumerableMap`
r := enu.ToMap(e).Keys()
// []int{1, 2, 3}
```

### FromNumeric

Returns an `*EnumeratorNumeric[T]` with `constraints.Integer` or `constraints.Float` argument.

```go
r := enu.FromNumeric([]int{1, 2, 3}).Sum()
// 6
```

### ToNumeric

Returns an `*EnumeratorNumeric[T]` with `Enumerator[T constraints.Integer | constraints.Float]` argument.

```go
// `Enumrator` does not implement `Sum()`
e := enu.From([]int{1, 2, 3})

// Convert to `EnumeratorNumeric`
r := enu.ToNumeric(e).Sum()
// 6
```

### FromOrdered

Returns an `*EnumeratorOrdered[T]` with `constraints.Ordered` argument.

```go
r := enu.FromOrdered([]string{"c", "a", "b"}).Sort()
// //[]string{"a", "b", "c"}
```

### ToOrdered

Returns an `*EnumeratorOrdered[T]` with `Enumerator[T constraints.Ordered]` argument.

```go
// `Enumrator` does not implement `Sum()`
e := enu.From([]string{"c", "a", "b"})

// Convert to `EnumeratorNumeric`
r := enu.ToNumeric(e).Sum()
// 6
```

## Enumerator functions

Enumerators called `ToSlice` function if `lazy` is false.

### Each

- lazy: false
- supported: all

Iterates over elements of a collection and invokes the function over each element.

```go
enu.From([]string{"hello", "world"}).Each(func(x string, _ int) {
  println(x)
})
// prints "hello\nworld\n"
```

### ToSlice

- lazy: false
- supported: all

Transform a enumerable into a slice.

```go
r1 := enu.From([]int{1, 2, 3}).ToSlice()
// []int{1, 2, 3}

r2 := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).ToSlice()
// []enu.KeyValuePair[int, string]{{Key: 1, Value: "foo"}, {Key: 2, Value: "bar"}, {Key: 3, Value: "baz"}}
```

### Count

- lazy: false
- supported: all

Counts the number of elements in the collection.

```go
r1 := enu.From([]int{1, 2, 3, 4, 5}).Count()
// 5

r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Count()
// 3
```

### Filter

- lazy: false
- supported: all

Counts the number of elements in the collection.

```go
r1 := enu.From([]int{1, 2, 3, 4}).Filter(func(x int, index int) bool {
    return x%2 == 0
}).ToSlice()
// []int{2, 4}

r2 := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz", 4: "boo"}).Filter(func(kv enu.KeyValuePair[int, string], _ int) bool {
  return kv.Key%2 == 0
}).ToMap()
// map[int]string{2: "bar", 4: "boo"}
```

### Reject

- lazy: false
- supported: all

The opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.

```go
r1 := enu.From([]int{1, 2, 3, 4}).Reject(func(x int, index int) bool {
    return x%2 == 0
}).ToSlice()
// []int{1, 3}

r2 := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz", 4: "boo"}).Reject(func(kv enu.KeyValuePair[int, string], _ int) bool {
  return kv.Key%2 == 0
}).ToMap()
// map[int]string{1: "foo", 3: "baz"}
```

### Nth

- lazy: false
- supported: all

Returns the element at index nth of collection. If nth is negative, the nth element from the end is returned. An false is returned when nth is out of slice bounds.

```golang
nth, ok := enu.New([]int{0, 1, 2, 3}).Nth(2)
// 2, true

nth, ok := enu.New([]int{0, 1, 2, 3}).Nth(-2)
// 2, true

nth, ok := enu.New([]int{0, 1, 2, 3}).Nth(6)
// 0, false
```

### Find

- lazy: true
- supported: all

Search an element in a slice based on a predicate. It returns element and true if element was found.

This works lazily. If a checked items is found, it stops iteration and does not look into remaining groups.

```golang
slices := enu.From([]int{1, 2, 3, 4, 5})

r1, ok := slices.Find(func(item, _ int) bool {
  return item%2 == 0
})
// 2, true

_, ok = slices.Find(func(item, _ int) bool {
  return item > 10
})
// 0, false

r2, ok := slices.Find(func(_, index int) bool {
  return index == 3
})
// 4, true

ch := make(chan (int))

go func() {
  defer close(ch)

  for i := 1; i < 6; i++ {
    ch <- i
  }
}()

r3, ok := enu.FromChannel(ch).Find(func(i int) bool {
  return i == 3
})
// 3, true
```

### First

- lazy: true
- supported: all

Search a first element in a slice based. It returns element and true if collection was not empty.

This works lazily. If a first item is found, it stops iteration and does not look into remaining groups.

```golang
r1, ok := enu.From([]int{1, 2, 3, 4, 5}).First()
// 1, true

r2, ok := enu.From([]int{}).First()
// 0, false
```

### Last

- lazy: false
- supported: all

Search a last element in a slice based. It returns element and true if collection was not empty.

```golang
r1, ok := enu.From([]int{1, 2, 3, 4, 5}).Last()
// 5, true

r2, ok := enu.From([]int{}).Last()
// 0, false
```

### Reverse

- lazy: false
- supported: all

Reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.

```golang
r := enu.From([]int{1, 2, 3, 4, 5}).Reverse().ToSlice()
// []int{5, 4, 3, 2, 1}
```

### Sort

- lazy: false
- supported only:
  - EnumeratorNumeric
  - EnumeratorOrdered

Sorts the slice x using `res[i] < res[j]` , keeping equal elements in their original order.

```golang
r := enu.From([]int{5, 3, 2, 1, 4}).Sort().ToSlice()
// []int{1, 2, 3, 4, 5}
```

### SortBy

- lazy: false
- supported: all

SortBy the slice x using the provided less function, keeping equal elements in their original order.

```golang
r := enu.From([]string{"aa", "bbb", "c"}).SortBy(func(i, j string) bool {
  return len(i) > len(j)
}).ToSlice()
// []string{"bbb", "aa", "c"}
```

### IsAll

- lazy: true
- supported: all

Returns true if all elements meet a specified criterion; false otherwise.

IsAll the slice x using the provided less function, keeping equal elements in their original order.

This works lazily. If a first negative item is found, it stops iteration and does not look into remaining groups.

```golang
r := enu.From([]int{1, 1, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
// true

r2 := enu.From([]int{1, 2, 1, 1, 1}).IsAll(func(item int) bool { return item == 1 })
// false
```

### IsAny

- lazy: true
- supported: all

Returns true if any elements meet a specified criterion; false otherwise.

IsAll the slice x using the provided less function, keeping equal elements in their original order.

This works lazily. If a first positive item is found, it stops iteration and does not look into remaining groups.

```golang
r1 := enu.From([]int{1, 1, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
// false

r2 := enu.From([]int{1, 2, 1, 1, 1}).IsAny(func(item int) bool { return item == 2 })
// true
```

### Take

- lazy: true
- supported: all

Returns a specified number of leading elements.

This works lazily. If a number of leading item is found, it stops iteration and does not look into remaining groups.

```golang
r1 := enu.From([]int{1, 2, 3, 4, 5}).Take(3).ToSlice()
// []int{1, 2, 3}

r2 := enu.FromNumericRange(1, math.MaxInt).Take(10).ToSlice()
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

### ToMap

- lazy: false
- supported: all

Returns a map key-value pairs provided by index and value of slice.

If receiver is `EnumeratorMap`, returns map provided by `KeyValuePair` .

```golang
r1 := enu.From([]int{3, 2, 1}).ToMap()
// map[int]int{0: 3, 1: 2, 2: 1}

r2 := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz", 4: "boo"}).Filter(func(kv enu.KeyValuePair[int, string], _ int) bool {
  return kv.Key%2 == 0
}).ToMap()
// map[int]string{2: "bar", 4: "boo"}
```

### Uniq

- lazy: false
- supported: only
  - EnumeratorComparable
  - EnumeratorNumeric
  - EnumeratorOrdered

Returns elements that are not duplicates.

```golang
r := enu.FromComparable([]int{1, 1, 2, 3, 3}).Uniq().ToSlice()
// []int{1, 2, 3}
```

### Contains

- lazy: false
- supported: only
  - EnumeratorComparable
  - EnumeratorNumeric
  - EnumeratorOrdered

Returns true if an element is present in a collection.

```golang
r1 := enu.FromComparable([]int{1, 1, 2, 3, 3}).Contains(1)
// true

r2 := enu.FromComparable([]int{1, 1, 2, 3, 3}).Contains(4)
// false
```

### IndexOf

- lazy: false
- supported: only
  - EnumeratorComparable
  - EnumeratorNumeric
  - EnumeratorOrdered

Returns the index at which the first occurrence of a value is found in an array or return -1 if the value cannot be found.

```golang
r1 := enu.FromComparable([]int{1, 1, 2, 3, 3}).IndexOf(2)
// 2

r2 := enu.FromComparable([]int{1, 1, 2, 3, 3}).IndexOf(4)
// -1
```

### Min

- lazy: false
- supported: only
  - EnumeratorOrdered
  - EnumeratorNumeric

Search the minimum value of a collection.

Returns zero value when collection is empty.

```golang
r1 := enu.FromOrdered([]int{2, 1, 3}).Min()
// 1

r2 := enu.FromOrdered([]int{}).Min()
// 0
```

### Max

- lazy: false
- supported: only
  - EnumeratorOrdered
  - EnumeratorNumeric

Search the maximum value of a collection.

Returns zero value when collection is empty.

```golang
r1 := enu.FromOrdered([]int{2, 1, 3}).Max()
// 3

r2 := enu.FromOrdered([]int{}).Max()
// 0
```

### Sum

- lazy: false
- supported: only
  - EnumeratorNumeric

Sums the values in a collection.

If collection is empty 0 is returned.

```golang
r1 := enu.FromNumeric([]int{2, 1, 3}).Sum()
// 6

r2 := enu.FromNumeric([]int{}).Sum()
// 0
```

### Keys

- lazy: false
- supported: only
  - EnumeratorMap

Returns an array of the map keys.

```golang
r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Keys()
// []int{1, 2, 3}
```

### Values

- lazy: false
- supported: only
  - EnumeratorMap

Returns an array of the map values.

```golang
r := enu.FromMap(map[int]string{1: "foo", 2: "bar", 3: "baz"}).Values()
// []string{"foo", "bar", "baz"}
```

### GetEnumerator

- lazy: true
- supported: all

Returns `*Enumerator[T any]` .

```golang
// Example: `ToNumeric` only supports `*Enumerator[T any]` , but enu.FromOrdered returns EnumeratorOrderd
// Get `*Enumreator[T any]` if GetEnumerator to use.
enumerator := enu.FromOrdered([]int{1, 2, 3}).GetEnumerator()
r := enu.ToNumeric(enumerator).Sum()
// 6
```

## Enumrables

### FromRange

Returns `*Enumerator[T any]` with `*RangeEnumerator[T1, T2]` argument.

Expects range values impelements [RangeValuer](#rangevaluer) interface.

```go
type Month time.Month

// Next returns next month.
func (m Month) Next(step int) enu.RangeValuer[time.Month, int] {
	val := int(m) + step
	if val > 12 {
		val -= 12
	}

	return Month(val)
}

// Value returns current month.
func (m Month) Value() time.Month {
	return time.Month(m)
}

// Compare current and other elements
func (m Month) Compare(other time.Month) int {
	if int(m) < int(other) {
		return -1
	}
	if int(m) > int(other) {
		return 1
	}
	return 0
}

r1 := enu.FromRange[time.Month, int](Month(time.January), Month(time.July), 1).ToSlice()
// []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July}

r2 := enu.FromRange[time.Month, int](Month(time.January), Month(math.MaxInt), 1).Take(13).ToSlice()
// []time.Month{
//   time.January,
//   time.February,
//   time.March,
//   time.April,
//   time.May,
//   time.June,
//   time.July,
//   time.August,
//   time.September,
//   time.October,
//   time.November,
//   time.December,
//   time.January,
// }
```

### RangeValuer

Interface to `RangeEnumerator`.

```go
type RangeValuer[T1, T2 any] interface {
  // Compare self and other
  Compare(T1) int

  // Returns current value.
  Value() T1

  // Returns Next RangeValuer
  Next(step T2) RangeValuer[T1, T2]
}
```

### FromNumericRange

Returns an `*EnumeratorNumeric[T]` with `min` and `max` argument.

If `min` greater than `max` empty range is returned.

```go
r1 := enu.FromNumericRange(1, 5).ToSlice()
//  []int{1, 2, 3, 4, 5}

r2 := enu.FromNumericRange(-5, -1).ToSlice()
// []int{-5, -4, -3, -2, -1}

r3 := enu.FromNumericRange(-1, -5).ToSlice()
// []int{}

r4 := enu.FromNumericRange(1, math.MaxInt).Take(10).ToSlice()
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

r5 := enu.FromNumericRange(1, math.Inf(0)).Take(3).ToSlice()
// []float64{1, 2, 3}
```

### FromNumericRangeWithStep

Returns an `*EnumeratorNumeric[T]` with `min` , `max` and `step` argument.

If `min` greater than `max` empty range is returned.

```go
r1 := enu.FromNumericRangeWithStep(1, 5, 1).ToSlice()
// []int{1, 2, 3, 4, 5}

r2 := enu.FromNumericRangeWithStep(1.1, 5.1, 1.0).ToSlice()
// []float64{1.1, 2.1, 3.1, 4.1, 5.1}
```

### FromDateRange

Returns an `*Enumerator[T]` with `start` , `end` and `duration` argument.

If `start` greater than `end` empty range is returned.

```go
r1 := enu.FromDateRange(
  time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
  time.Date(2023, 5, 20, 23, 59, 59, 0, time.UTC),
  time.Hour,
).ToSlice()
// []time.Time{
//   time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 20, 22, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 20, 23, 59, 59, 0, time.UTC),
// }

r2 := enu.FromDateRange(
  time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
  time.Date(2099, 5, 23, 21, 59, 59, 0, time.UTC),
  time.Hour*24,
).Take(5).ToSlice()
// []time.Time{
//   time.Date(2023, 5, 20, 21, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 21, 21, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 22, 21, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 23, 21, 59, 59, 0, time.UTC),
//   time.Date(2023, 5, 24, 21, 59, 59, 0, time.UTC),
// }
```

### FromFunc

Returns an `*Enumerator[T]` with `func(index int) (T, bool)` func.

```go
fibonacci := func() func(int) (int, bool) {
  n, n1 := 0, 1
  return func(i int) (int, bool) {
    fn := func(index int) (int, bool) {
      // Reset variable if index given zero.
      if index == 0 {
        n, n1 = 0, 1
      }
      v := n
      n, n1 = n1, n+n1
      return v, true
    }
    return fn(i)
  }
}

g := enu.FromFunc(fibonacci())

r1, ok := g.First()
// 0, true

r2 := g.Take(10).ToSlice()
// []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
```

## ChannelHelpers

### Repeat

Executes a number of repeated functions in a goroutine and returns the result in channel.

```golang
mockHttpRequest := func(n int) enu.Tuple2[string, error] {
  if n%2 == 0 {
    return enu.Tuple2[string, error]{"OK", nil}
  }
  return enu.Tuple2[string, error]{"", errors.New("some error")}
}

// Aggregate function result by repeat and async
r := enu.FromChannel(enu.Repeat(2, 5, mockHttpRequest)).ToSlice()
// []enu.Tuple2[string, error]{
//   {"OK", nil},
//   {"OK", nil},
//   {"OK", nil},
//   {"", errors.New("some error")},
//   {"", errors.New("some error")},
// }
```

### Parallel

Executes a multiple functions in a goroutine and returns the result in channel.

```golang
request1 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"https://google.com", nil}
}

request2 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"https://amazon.co.jp", nil}
}

request3 := func() enu.Tuple2[string, error] {
  return enu.Tuple2[string, error]{"", errors.New("some error")}
}

// Aggregate multiple function result by async
r := enu.FromChannel(enu.Parallel(2, request1, request2, request3)).ToSlice()
// []enu.Tuple2[string, error]{
//   {"", errors.New("some error")},
//   {"https://amazon.co.jp", nil},
//   {"https://google.com", nil},
// }
```

### Retry

Executes a retryable functions in a goroutine while to succeeded, and returns the result in channel.

```golang
mockHttpRequest := func(n int) (string, error) {
  if n < 5 {
    n += 1
    return "", errors.New("some error")
  }
  return "OK", nil
}

// Returns [result[T], true] if function was succeed.
r1, ok := enu.FromChannel(enu.Retry(3, 5, mockHttpRequest)).First()
// {"ok", true}, true

// Returns [empty[T], false] if function was failed with reach to maxRetry.
r2, ok := enu.FromChannel(enu.Retry(3, 3, mockHttpRequest)).First()
// {"", false}, true
```

## Aggregate functions

### Map

Manipulates a slice of one type and transforms it into a slice of another type:

### Reduce

Reduces a collection to a single value.
