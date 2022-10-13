# Performance in a High-throughput SQL Database: Real-world Profiler Tips

In this tutorial, Zach will show you how to profile your codebase to
make it faster and share some performance-killing anti-patterns to
avoid, all in the context of a novel SQL database written in golang.

Zach and the team at DoltHub built Dolt, the world’s first SQL
database that you can branch and merge, fork and clone, push and pull
like a git repository. They wrote the storage and query engine from
the ground up, then made it as fast as MySQL by using golang’s
profiler. Join Zach as he shares what they learned on this journey.
He will demo how to use golang’s profiler tools to uncover slowdowns
in your code. He will also detail many examples of slow code from
DoltHub’s own code base and show you how the team fixed them. Finally,
Zach will discuss the architectural choices for their database’s tuple
storage, and how they contribute to its performance.

Watch slides online [here](https://docs.google.com/presentation/d/e/2PACX-1vR7HrSEoNYJuE9cS2Df_sPPli2vs2ZfzB-np2dtrTNZn-yLCq4ZXUjkEWm0CxtCLLnctr1KL2D2kmub/pub?start=false&loop=false&delayms=3000), or browse the PDF file in this directory.

Additional info:

* [Dolt on GitHub](https://github.com/dolthub/dolt)
* [dolthub.com](https://dolthub.com)
* [Dolt Discord server](https://discord.com/invite/RFwfYpu)
