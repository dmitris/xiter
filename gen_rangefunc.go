// Code generated by internal/gen. DO NOT EDIT.

//go:build rangefunc

package xiter

import (
	"cmp"
	"context"
	"iter"
	"reflect"
)

func All[T any](seq iter.Seq[T], f func(T) bool) bool {
	return _All(seq, f)
}

func Any[T any](seq iter.Seq[T], f func(T) bool) bool {
	return _Any(seq, f)
}

func AppendSplitTo[T1 any, T2 any](seq SplitSeq[T1, T2], s1 []T1, s2 []T2) ([]T1, []T2) {
	return _AppendSplitTo(seq, s1, s2)
}

func AppendTo[T any, S ~[]T](seq iter.Seq[T], s S) S {
	return _AppendTo(seq, s)
}

func Bytes(s string) iter.Seq[byte] {
	return _Bytes(s)
}

func Cache[T any](seq iter.Seq[T]) iter.Seq[T] {
	return _Cache(seq)
}

func Chunks[T any](seq iter.Seq[T], n int) iter.Seq[[]T] {
	return _Chunks(seq, n)
}

func Collect[T any](seq iter.Seq[T]) []T {
	return _Collect(seq)
}

func CollectSize[T any](seq iter.Seq[T], len int) []T {
	return _CollectSize(seq, len)
}

func CollectSplit[T1 any, T2 any](seq SplitSeq[T1, T2]) (y1 []T1, y2 []T2) {
	return _CollectSplit(seq)
}

func Concat[T any](seqs []iter.Seq[T]) iter.Seq[T] {
	return _Concat(seqs)
}

func Contains[T comparable](seq iter.Seq[T], v T) bool {
	return _Contains(seq, v)
}

func Drain[T any](seq iter.Seq[T]) {
	return _Drain(seq)
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return _Enumerate(seq)
}

func Equal[T cmp.Ordered](seq1 iter.Seq[T], seq2 iter.Seq[T]) bool {
	return _Equal(seq1, seq2)
}

func EqualFunc[T1 any, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2], equal func(T1, T2) bool) bool {
	return _EqualFunc(seq1, seq2, equal)
}

func Filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return _Filter(seq, f)
}

func Find[T any](seq iter.Seq[T], f func(T) bool) (r T, ok bool) {
	return _Find(seq, f)
}

func Flatten[T any](seq iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return _Flatten(seq)
}

func Fold[T any](seq iter.Seq[T], reducer func(T, T) T) T {
	return _Fold(seq, reducer)
}

func FromPair[T1 any, T2 any](seq iter.Seq[Pair[T1, T2]]) iter.Seq2[T1, T2] {
	return _FromPair(seq)
}

func Generate[T Addable](start T, step T) iter.Seq[T] {
	return _Generate(start, step)
}

func GoPull[T any](seq iter.Seq[T]) (iter func() (T, bool), stop func()) {
	return _GoPull(seq)
}

func Handle[T any](seq iter.Seq2[T, error], f func(error) bool) iter.Seq[T] {
	return _Handle(seq, f)
}

func IsSorted[T cmp.Ordered](seq iter.Seq[T]) bool {
	return _IsSorted(seq)
}

func IsSortedFunc[T any](seq iter.Seq[T], compare func(T, T) int) bool {
	return _IsSortedFunc(seq, compare)
}

func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return _Limit(seq, n)
}

func Map[T1 any, T2 any](seq iter.Seq[T1], f func(T1) T2) iter.Seq[T2] {
	return _Map(seq, f)
}

func MapKeys[K comparable, V any, M ~map[K]V](m M) iter.Seq[K] {
	return _MapKeys(m)
}

func MapValues[K comparable, V any, M ~map[K]V](m M) iter.Seq[V] {
	return _MapValues(m)
}

func Max[T cmp.Ordered](seq iter.Seq[T]) T {
	return _Max(seq)
}

func Merge[T cmp.Ordered](seq1 iter.Seq[T], seq2 iter.Seq[T]) iter.Seq[T] {
	return _Merge(seq1, seq2)
}

func MergeFunc[T any](seq1 iter.Seq[T], seq2 iter.Seq[T], compare func(T, T) int) iter.Seq[T] {
	return _MergeFunc(seq1, seq2, compare)
}

func Min[T cmp.Ordered](seq iter.Seq[T]) T {
	return _Min(seq)
}

func Of[T any](vals []T) iter.Seq[T] {
	return _Of(vals)
}

func OfChan[T any](c <-chan T) iter.Seq[T] {
	return _OfChan(c)
}

func OfMap[K comparable, V any, M ~map[K]V](m M) iter.Seq2[K, V] {
	return _OfMap(m)
}

func OfSlice[T any, S ~[]T](s S) iter.Seq[T] {
	return _OfSlice(s)
}

func OfSliceIndex[T any, S ~[]T](s S) iter.Seq2[int, T] {
	return _OfSliceIndex(s)
}

func OfValue(v reflect.Value) iter.Seq2[reflect.Value, reflect.Value] {
	return _OfValue(v)
}

func Or[T any](seqs []iter.Seq[T]) iter.Seq[T] {
	return _Or(seqs)
}

func Partition[T any](seq iter.Seq[T], f func(T) bool) (true []T, false []T) {
	return _Partition(seq, f)
}

func PartitionInto[T any](seq iter.Seq[T], f func(T) bool, true []T, false []T) ([]T, []T) {
	return _PartitionInto(seq, f, true, false)
}

func Product[T Multiplyable](seq iter.Seq[T]) T {
	return _Product(seq)
}

func Pull[T any](seq iter.Seq[T]) (iter func() (T, bool), stop func()) {
	return _Pull(seq)
}

func RecvContext[T any](ctx context.Context, c <-chan T) iter.Seq[T] {
	return _RecvContext(ctx, c)
}

func Reduce[T any, R any](seq iter.Seq[T], initial R, reducer func(R, T) R) R {
	return _Reduce(seq, initial, reducer)
}

func Runes[T ~[]byte | ~string](s T) iter.Seq[rune] {
	return _Runes(s)
}

func SendContext[T any](seq iter.Seq[T], ctx context.Context, c chan<- T) {
	return _SendContext(seq, ctx, c)
}

func Skip[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return _Skip(seq, n)
}

func Split[T any](seq iter.Seq[T], f func(T) bool) SplitSeq[T, T] {
	return _Split(seq, f)
}

func Split2[T1 any, T2 any](seq iter.Seq2[T1, T2]) SplitSeq[T1, T2] {
	return _Split2(seq)
}

func StringSplit(s string, sep string) iter.Seq[string] {
	return _StringSplit(s, sep)
}

func Sum[T Addable](seq iter.Seq[T]) T {
	return _Sum(seq)
}

func ToPair[T1 any, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[Pair[T1, T2]] {
	return _ToPair(seq)
}

func V1[T1 any, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[T1] {
	return _V1(seq)
}

func V2[T1 any, T2 any](seq iter.Seq2[T1, T2]) iter.Seq[T2] {
	return _V2(seq)
}

func Windows[T any](seq iter.Seq[T], n int) iter.Seq[[]T] {
	return _Windows(seq, n)
}

func Zip[T1 any, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq[Zipped[T1, T2]] {
	return _Zip(seq1, seq2)
}
