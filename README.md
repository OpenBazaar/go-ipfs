# go-ipfs OpenBazaar fork
[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs) [![Build Status](https://travis-ci.org/ipfs/go-ipfs.svg?branch=master)](https://travis-ci.org/ipfs/go-ipfs)

This is the official fork of IPFS used in OpenBazaar. It's comes bundled in the `vendor`
package of openbazaar-go so if you run openbazaar-go you are running this fork.

It is not safe to run the main IPFS codebase in the OpenBazaar network as your node will
not be able to communicate with other OpenBazaar nodes.

## Diff
This fork is currently based on IPFS v0.4.15 with the following changes:

- namesys/publisher.go change DefaultRecordTTL and DefaultPublishLifetime to one week.
- namesys/namesys.go NewNameSystem takes in a database instance for caching records.
- namesys/routing.go resolveOnce stores the resolved record in the database using `cachePrefix` and returns the record when routing returns not found if UsePersistentCache is true.
- namesys/routing.go resolveOnce stores the resolved public key in the database using `keyCachePrefix` and checks the db when fetching public keys.
- namesys/routing.go resolveOnce is modified to optionally accept an alt-root in the format root:suffix.
- namesys/namesys.go NewNameSystem takes in a custom DNSResolver instance which can be nil.

- repo/config/ipns.go Add QuerySize, BackUpAPI, and UsePersistentCache paramters to IPNS config.

- core/core.go startOnlineServices takes in a DNSResolver to initialize the `NameSystem` with.
- core/builder.go add a DNSResolver to the build config to pass into startOnlineServices
- core/bootstrap.go add `DoneChan` to the bootstrap config which is closed when the inital bootstrap finishes. This is in place of blocking for the initial bootstrap.
- core/commands/swarm.go Change the `swarm` `peers` output to []string from a private struct. The access control on the struct made the return unusable otherwise.
- core/commands/ipns.go Initialize NewNameSystem with a NewDNSResolver().
- core/commands/dht.go Swap out DHT import with DHT fork.
- core/coreunix/add.go AddWithContext function modified to use CIDv1.
- core/coreunix/add.go NewAddr updated to use prefix and CIDv1.

Finally, we've had to patch go-multiaddr to handle CIDv1 in the /ipfs/ addr strings. We'll try to get this merged into the main repo.
