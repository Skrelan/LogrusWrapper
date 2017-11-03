# LogrusWrapper

## The Problem:
When we define the logging level, logs that are considered lower levels that the requested level are ignored. This can be an issue when you want to push your code to a prod enviroment, where you might want to log only Errors but at that same time you also would like to log out which process is running. 
 

## The Goal
To enable developers to have the ability to have muteable and immuteable logging.

## The Solution
This is a wrapper package for logrus that you can use to have both mutable and immutable logs, based on your preference.
The logging levels still work in the same way, i.e Debug < Info < Warn < Error < Panic. Regular logging happens as it did in the past, and the level can be controlled by the SetLoggingLevel function, but there are two more Logging functions, `AlwaysLog` and `AlwaysLogf` which are the equivalent of `Info` and `Infof` respectivley. 

The above two added functions run on an new instance of logrus and remain un-effected by the change in logging level in the other.
