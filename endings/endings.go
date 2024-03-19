package endings

type ending struct {
	Strong string
	Weak   string
	Noun   string
}

var Endings = [][]ending{
	{{Strong: "er", Weak: "e"} /*      */, {Strong: "e", Weak: "e"}, {Strong: "es", Weak: "e" /*       */}, {Strong: "e", Weak: "n"}},
	{{Strong: "en", Weak: "n"} /*      */, {Strong: "e", Weak: "e"}, {Strong: "es", Weak: "e" /*       */}, {Strong: "e", Weak: "n"}},
	{{Strong: "em", Weak: "n"} /*      */, {Strong: "er", Weak: "n"}, {Strong: "em", Weak: "n" /*      */}, {Strong: "en", Weak: "n", Noun: "n"}},
	{{Strong: "es", Weak: "n", Noun: "s"}, {Strong: "er", Weak: "n"}, {Strong: "es", Weak: "n", Noun: "s"}, {Strong: "er", Weak: "n"}},
}
