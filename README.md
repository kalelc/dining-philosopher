# Dining Philosopher Problem Solution Using Golang

## Problem

_The Dining Philosopher Problem states that K philosophers seated around a circular table with one chopstick between each pair of philosophers. There is one chopstick between each philosopher. A philosopher may eat if he can pickup the two chopsticks adjacent to him. One chopstick may be picked up by any one of its adjacent followers but not both._ Wikipedia.

## Mutex in Golang

A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.

Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.