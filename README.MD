# Diommy's GoAPI

A little Rest API to practice my Go knowledge

## Objectives

- [] Design a small usecase with All 4 CRUD operations
- [] Implement the underlying structure of the api, mocking the storage
- [] Implement the Get endpoint
- [] Implement Post/Put endpoints
- [] Implement Delete endpoint
- [] Implement the data perstistence layer

## Idea

The Api should be tracking a game record

We should have a number of player, and for each player we should hold a win/loss counter. 

The counter starts at 0 when a player is created, increase by 1 for each win and decrease by one for each loss.

Let's assume that we do not keep a record of all loss/win but only the counter.

let's also assume that if we delete the player, the counter is deleted as well.

We should be able to do 4 things :

- Create a player with a Name
- Register a win or loss for a player 
- Get the counter for the player
- Delete the player


## Nice to have 

- [] Tests
- [] Bootstrap the data persistence layer
- [] Add Auth mech
- [] Add some concurrency use case to the design
