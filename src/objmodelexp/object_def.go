package objmodelexp

type A struct {
    F1 bool
    F2 int16
    F3 int
}

type B struct {
    F4 bool
    F5 int
    F6 int16
}

type C struct {
    A
    F7 int
}

type D struct {
    F8 int
    C
}

type E struct {
    C
    B
    F9 int
    F10 int
    F11 int
}

type F struct {
    F12 int
    C
    F13 int
    B
    F14 int
}
