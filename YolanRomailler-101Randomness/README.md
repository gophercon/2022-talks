# Randomness 101

Let us explore together what is randomness, what it means for cryptographic primitives such as TLS, signatures and encryption, and why it is more often than not a challenge to get right.

In this tutorial we will walk through the different _flavours_ of randomness, from secret, local randomness to public, verifiable, distributed randomness. We will see a few bad-randomness examples, see how we can properly sample randomness in our code. We will learn about the infamous "modulo bias" and how to avoid it in Go.

## Speaker Details

- Name: Yolan Romailler
- A Few Handles:
    - GitHub: [@anomalroil](https://github.com/anomalroil)
    - Twitter: [@anomalroil](https://twitter.com/anomalroil)

## Links and ressources
- [The definitive guide to modulo bias](https://research.kudelskisecurity.com/2020/07/28/the-definitive-guide-to-modulo-bias-and-how-to-avoid-it)
- A few useful repos:
    - drand repo: https://github.com/drand/drand/
    - tlock repo: https://github.com/drand/tlock/
    - tlock-js repo: https://github.com/drand/tlock-js/
- Test Timelock Encryption in your browser: https://timevault.drand.love/
- Stay tuned for more info and blog posts around randomness or timelock encryption on drand's blog: https://drand.love/blog/
- Wants to run a drand node and join the League of Entropy? Join us: https://drand.love/partner-with-us/

Be sure to checkout my slides, they contain the Playground links to most of my code snippets!
