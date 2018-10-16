/*
package test implements a testing suite spanning multiple packages.

Using github.com/stretchr/testify/suite requires each subpackage to reimplement the toplevel testing suite
and register itself again.

test takes a different approach as it does not use a godlevel construct, but instead has utility getter functions
returning "stateful" and providing a cleanup function.

Implement test requires the machine to be connected to a network and a recent version of the docker runtime.

TODO: some services are implemented as docker containers, which take a few seconds to start.
Currently when two tests require a service, two docker containers are deployed, instead of cleaning
the first and reusing the connection. Once the testsize of carapace grows, we might need to look into this.
*/
package test
