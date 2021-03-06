# Rule 110

[Rule 110](https://en.wikipedia.org/wiki/Rule_110) is a 1-dimensional, turing
complete cellular automata.

Transitions for each bit are determined by its neighbours:

```
- 111 -> 0
- 110 -> 1
- 101 -> 1
- 100 -> 0
- 011 -> 1
- 010 -> 1
- 001 -> 1
- 000 -> 0
```

RHS = 01101110 = 110 (base 10), hence rule 110.


## How

Just `go run main.go` to see a 200 unit automata in action over 5000 periods,
slowed down for viewing convenience (requires go 1.10+). The code just uses the
unicode block and whitespace to represent elements in the cell, so it should
look like this in your terminal.

![hello sierpinski](img/sim.png)


## But Why?

Cellular automata like this are appealing, in theory, because they are turing
complete, also because program and input data exist in the same state space.
This opens the door to learning arbitrary functions over the space of computable
programs, which could be useful.


## Improvements

Implementation is super naive and not very efficient. Notes for future me if this
becomes something that needs to be fast.

1. Use bytes, not bools (which store on bit in eight bytes)
2. Amortise transitions by pre-compiling transitions at byte level
    - since signals propagate at a fixed speed, could also pre-compute multiple step transitions
3. With (2), theoretically only need to allocate one byte per transition thread, by
4. Investigate `Hashlife` for performance techniques