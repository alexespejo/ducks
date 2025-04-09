---
layout: "../../layouts/LayoutSingle.astro"
title: cs179_notes
---

# DESIGNING FINITE AUTOMATA

Whether it be of automaton or artwork, design is a creative process. As such,
it cannot be reduced to a simple recipe or formula. However, you might find
a particular approach helpful when designing various types of automata. That
is, put yourself in the place of the machine you are trying to design and then see
how you would go about performing the machine’s task. Pretending that you are
the machine is a psychological trick that helps engage your whole mind in the
design process.

Let’s design a finite automaton using the “reader as automaton” method just
described. Suppose that you are given some language and want to design a finite
automaton that recognizes it. Pretending to be the automaton, you receive an
input string and must determine whether it is a member of the language the
automaton is supposed to recognize. You get to see the symbols in the string
one by one. After each symbol, you must decide whether the string seen so far is
in the language. The reason is that you, like the machine, don’t know when the
end of the string is coming, so you must always be ready with the answer.

First, in order to make these decisions, you have to figure out what you need
to remember about the string as you are reading it. Why not simply remember
all you have seen? Bear in mind that you are pretending to be a finite automaton
and that this type of machine has only a finite number of states, which means
a finite memory. Imagine that the input is extremely long—say, from here to
the moon—so that you could not possibly remember the entire thing. You have
a finite memory—say, a single sheet of paper—which has a limited storage capacity. Fortunately, for many languages you don’t need to remember the entire
input. You need to remember only certain crucial information. Exactly which
information is crucial depends on the particular language considered.

For example, suppose that the alphabet is {0,1}and that the language consists
of all strings with an odd number of 1s. You want to construct a finite automaton
E1 to recognize this language. Pretending to be the automaton, you start getting an input string of 0s and 1s symbol by symbol. Do you need to remember the
entire string seen so far in order to determine whether the number of 1s is odd?
Of course not. Simply remember whether the number of 1s seen so far is even
or odd and keep track of this information as you read new symbols. If you read
a 1, flip the answer; but if you read a 0, leave the answer as is.
But how does this help you design E1? Once you have determined the neces-
sary information to remember about the string as it is being read, you represent
this information as a finite list of possibilities. In this instance, the possibilities
would be
