# Parsing Without Code Generators: Parser Combinators in Go with Generics

For reasons too tedious to print here, I needed to write an extensible Markdown parser.  (OK, since you asked, I needed it to implement the software project planning language for [DoomCheck](https://www.doomcheck.com/)). Just as I finished, the first beta release of Go with generics came out, so I did it all again!

Now I’m here to take you on the same journey. Not the Markdown part – Markdown is a parsing nightmare – the parser combinators part.

Parser combinators are a composable way of building parsers in code; they stand in contrast to traditional parser generators like goyacc. They work particularly well in Go now that Go has generics, and they are remarkably easy to implement from scratch. In this tutorial, I’ll do a quick introduction to parser combinators, show you how to use a few primitives to implement a parser for a microlanguage, and dive under the abstraction layer and show you how to implement all the primitives -- each in just a few lines of Go. I aim to give you an appreciation for both parser combinators as an approach to parsing, and for how much better Go generics have made life.

Finally, while I have had a great experience with Go generics overall, there are some frustrating limits around how you can use generics, so in the wrap-up, I will briefly talk about how future versions of Go could increase generics usability for this sort of domain.

### [Presentation](https://docs.google.com/presentation/d/1PfFFXjguakJM13tHWkFxFcy4wee65B26T9Cij2TWC0w/edit#slide=id.g15a27a86b97_0_76)

### [Youtube Video](https://www.youtube.com/watch?v=x5p_SJNRB4U&t=442s)