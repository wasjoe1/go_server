# GO Server

## Past experience
This repo is to create a production grade HTTP server in GO without the use of a framework.
In my previous internship, I built a server which proxied FIX messages between Clients & Servers,
and got my hands dirty in socket programming (mainly TCP + SSL layer).

## Why this project?
There were a few things I could improve on this:
- server was written in Python; performance was an issue (high memory & CPU usage)
- APP layer was not covered => handled by the quickfix library to process app layer messages
- Unsure if what i wrote was industry standard

I want to build this GO server for this reasons:
- Learn a language that has better performance (GO/ C++/ Rust)
    => since this course uses GO, I have decided to stick with it as im not familiar in the other 2 languages to try & implement the newly taught concepts
- Try my hand at building an APP layer server
- learn from a course which uses industry standards to create a production grade server
    => previously 1 issue i faced was how do i come up with unit & intergration tests to benchmark my server
    => I also did not know if what i created was the industry standard as i was alrady using python to write it
    (whos not to say that the way i implemented the async coroutines might not have been the best way to)


## Resources
- https://www.youtube.com/watch?v=FknTw9bJsXM <= was inspired to do this because of this youtube video by the 1 and only PrimeAgen
- https://www.boot.dev/courses/learn-http-protocol-golang
- https://www.boot.dev/courses/learn-http-servers-golang

## Project structure
- /textio
    - a mini programming exercise to get myself familiar with the GO language (incomplete but i rather spend time getting familiar with the language via the server project)
- /http_protocol
    - Building a HTTP/1.1 server at the APP layer on top of the already existing TRANSPORT layer (TCP)
- /http_server
