# Pokemon 
Pokemon: apanhá-los todos. Implementação em Go e Java.

## Author
### João Santos

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
