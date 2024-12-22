# Silent Map Access in Go

This example demonstrates a potential issue in Go where accessing a non-existent key in a map returns the zero value for the map's value type instead of raising an error or exception. This can make debugging difficult, as the code might run without any obvious errors, but produce unexpected results.

## The Bug

The `bug.go` file contains a simple Go program that demonstrates this behavior.  A map is created, a key is added and then deleted. Accessing the deleted key will return 0 instead of signaling an error. This behavior can lead to subtle and hard-to-find bugs if not handled carefully.

## The Solution

The `bugSolution.go` shows one approach to mitigating this by checking whether a key exists before accessing it using the `value, ok := m["a"]` pattern.