This is my solution for Exercise 1 of Gophercises by Jon Calhoun (https://gophercises.com/)

The aim was to create a quiz game that reads a CSV file with questions/answers and ask each question one by one until all have been answered or until a timer runs out. It then gives the score at the end. 

Instructions
============

Assuming you have Go installed already, you can simply run:

```
go run main.go
```

Though I would recommend building it first so that you can run it with the additional flags:

```
go build main.go
```

Then run it with the optional flags:

```
./main -filename:csvfile.csv -timer:20 -random=true
```

By default the program will load problems.csv but you can create your own csv file in the format "question,answer" and use the filename flag to tell the program to use that instead of the default.

The timer is 30 seconds by default but you can use the timer flat to change that to whatever you want

And by using the random flag you can tell the program to randomize the order of the questions, but by default it will always give them in the same order as they are in the CSV file

Lessons
=======

I have used/learned the following functions of Go in this program:
* Flags
* Basic file operations and parsing CSV files
* Randomizing a slice
* Getting user input 
* Go routines and WaitGroups
* Simple time functions
* Using panic to handle errors
* Using for to iterate through a slice

Official Solution
=================

I watched Jon's official solution to this exerice and made the following observations.

Jon included better error handling wheras I only relied on panic. Going forward I'd like to put more effort into error handling and giving meaningful error messages.

I used os.ReadFile where Jon used os.Open so I feel I need a better understanding about file operations and the best way to do them.

In Jon's solution he created a custom type to store the question/answers in. Using struct is something I had read about but not yet tried using so its something I'd like to do in future as it seems to be a better, object-orientated way of doing things in Go.

For the timer function I used go routines as that seemed like the simplest option. In Jon's solution he used the Timer function on seperate channels with switch statements used to check for the response from that other channel. In some ways the timer function seems like the most obvious choice but I feel like my solution was more elegant. It was a good opportunity for me to use go functions in my solution but I think I need to better understand the timer function and the best times to use them.

I'm glad that I put more of my code in to functions than Jon did but there's still room for further refactoring of my code.

I also put a little more effort into the  user experience by telling the user how many questions there are and how much time they'll have then waiting for the user to press enter before it all starts. Also, the solution didn't go through the "extras" such as randomizing the questions but I'm glad my solution seems to work as intended.
