# Dining Philosopher Problem Solution Using Golang

## Problem

_The Dining Philosopher Problem states that K philosophers seated around a circular table with one chopstick between each pair of philosophers. There is one chopstick between each philosopher. A philosopher may eat if he can pickup the two chopsticks adjacent to him. One chopstick may be picked up by any one of its adjacent followers but not both._ Wikipedia.


## Contraints

- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

## Mutex in Golang

A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition from happening.

Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.