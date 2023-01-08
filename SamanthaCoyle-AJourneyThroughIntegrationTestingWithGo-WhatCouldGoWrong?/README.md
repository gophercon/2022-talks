# A Journey Through Integration Testing with Go - What Could Go Wrong?

When our team set out to conquer the world and perform integration testing for our project, we arrived at an impasse. The typical flow we've experienced uses the Python Robot framework but being Go developers we chose to take a risk for our org in hopes that it would pay off.

Our team's use of Testcontainers for Go and go-test-report helped us to spin up our services, do some exciting integration tests, and create a nice test report html page for a good UX. Everything was working honky-dory until it was noticed that the tests all passed, but the test report tool showed a red failing status for passing tests. Yes, you read that right - this seemed suspicious. 

This was the first of a few problems we encountered on our journey; from limitations with some open source pkgs we used to disagreements between developers and our validation engineer on when code should be deemed as passing our validation standards. Come along in our journey to find out what our issues were and how we went about alleviating them.

## Samantha Coyle

[Personal Website](http://samcoyle.me/) | 
[Github](https://github.com/sicoyle) | 
[Twitter](https://twitter.com/thesamcoyle) | 
[LinkedIn](https://www.linkedin.com/in/sam-coyle/) |
[Medium](https://medium.com/@samcoyle)

## Resources

[PowerPoint](./SamanthaCoyleGopherCon22.pdf) |
[Testcontainers for Go Github](https://github.com/testcontainers/testcontainers-go) |
[Go Test Report Github](https://github.com/vakenbolt/go-test-report)

See y'all at GopherCon 2023!
