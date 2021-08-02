package test_gomod_replace_dep_a

type OriginalInt int32

func Add(a, b OriginalInt) OriginalInt {
	return a + b + 100
}
