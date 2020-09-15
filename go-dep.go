package main

// details https://tutorialedge.net/golang/an-intro-to-go-dep/
// he dep tool is the “official experiment” dependency management tool for the go programming language.
//It helps you to manage the ever-growing list of dependencies your project needs to maintain without
//a lot of overhead and it can pin you to specific versions of dependencies to ensure stability in your systems.

// - dep init

// The dep command features 5 commands in total:
//
//init - Sets up a new Go project
//status - Reports the status of a project’s dependencies
//ensure - Ensures a dependency is safely vendored in the project
//prune - Prunes your dependencies, this is also done automatically by ensure
//version - Shows the dep version information

// Updating dependencies
// // dry run testing an update
//$ dep ensure -update -n
//// non-dry run
//$ dep ensure -update
//// updates a specific package
//$ dep ensure -update github.com/gorilla/mux
//// updates to a specific version
//$ dep ensure -update github.com/gorilla/mux@1.0.0