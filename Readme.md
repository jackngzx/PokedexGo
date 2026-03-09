# Pokedex on Go

## Guided project for Boot.dev backend course

This REPL Pokedex program was built under the guidance of the excellent Boot.dev course. The idea is simple: call upon the [PokeAPI](https://pokeapi.co/docs/v2), parse the information so that users can run a few simple commands to retrieve data at will.

### Commands

1. Help: Provide a helpful message stating all the possible commands.
2. Exit: Exit the program.
3. Map: List out the first 20 maps available using [Location Areas endpoint](https://pokeapi.co/docs/v2#location-areas).
4. Mapb: List out the previous 20 maps, if the user has executed Map before. The program also utilize a cache to optimize performance, i.e if the list of maps were called before, they would be saved into the internal cache so that the next time user called on the previous list of maps, the cache result would be printed out instead of another API call. 
5. Explore: Take the name of a map as an argument and list out all Pokemons available at the location.
6. Catch: The most fun part of the program. Throw a Pokeball and catch a Pokemon! The catch chance is based on the Pokemon [endpoint](https://pokeapi.co/docs/v2#pokemon), specifically base_experience. The higher the base_experience, the harder it would be for the user to catch a Pokemon. Once the Pokemon is caught, it will also be saved in a Pokedex.
7. Pokedex: List of all caught Pokemons. 
8. Inspect: If a Pokemon is caught, user can use the name of the caught Pokemon as an argument alongside with this command to list out details about the Pokemon (name, weight, stats, types etc). 

## Learning

Overall I think this project is a fun one. I had a great time building out the program and learning the process of how to utilize the cache. I did need Boot AI help in some parts of the project, especially around structuring the project with multiple struct types and commands. The hardest part of the project is definitely building the cache and understanding the reaploop method, which I will try to review in a couple of days to make sure that I understand everything completely. 



