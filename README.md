# Pokemon 
Pokemon: apanhá-los todos. Go and Java implementation.

## Author
### João Santos

# Description
Ash walks around an infinite bidimensional space collecting Pokemons. Each spot contains a Pokemon, so every time Ash passes through a non visited spot, he catches the residing Pokemon. However, if Ash goes through a previously visited spot, he cacthes no Pokemon, as it has already been caught. Ash automatically catches the Pokemon where he is spawn.

## Available Commands


* __N:__ Go North
* __S:__ Go South
* __E:__ Go East
* __O:__ Go West

## Example

__Input:__ N

__Output:__ 2

__Input:__ NSNS

__Output:__ 2

__Input:__ ENOS

__Output:__ 4

# Running Instructions
## Java : Maven installation
#### Navigate to Folder [PokemonJava](./PokemonJava), and run the following commands:
* `mvn package`  
* `java -jar target/Pokemon.jar`

#### For testing only:
* `mvn test`

## Go: Manual installation
#### Download and Install GO by following the steps in [The Go WebSite](https://golang.org/doc/install)
#### Navigate to Folder [PokemonGo](./PokemonGo), and run the following commands:
* `go run pokemon.go`  

#### For testing only:
* `go test -v`

## Observations
#### In the GO Implementation there are two functions which (can) deal with move evaluation:
* EvaluateMoveBranching - Performs a switch on a given move and calls the corresponding method
* EvaluateMoveMapped - Makes use of a map which has methods mapped to a given move (E.g map['N'] = goNorth() )
