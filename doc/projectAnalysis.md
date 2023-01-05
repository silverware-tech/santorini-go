# Project Analysis

List of macro functionalities that the software must offer:

- Representing game objects (board, blocks, workers, players)
- Move workers
- Build blocks
- Detect player loss or win
- Detect end of the game

## Representing game objects

- The **board** is a `NxN` matrix of **cells** called **field**
- A cell contains info about cell **tower** of any level and/or a **worker** *if present on that cell*
- **Players** are represented by a **name** and a pair of **workers** of one **worker group**
- **Workers** are represented by a **worker group** and the **coordinates** in the **board field**
- **Blocks** are stored as tower status in **cells**

## Software components study

Need components for:

- Start/end game
- Handle players interactions (I/O)
- Move workers
- Build blocks
- Handle players turn

Possible components:

- **Game** is the core component that handle game lifecycle
- **I/O** handle interaction with stdin and stdout
- **Players** handle players info and store the turn status
- **Board** handle board objects status
  - **Cell** handle signle cell status

### Q&A

> write here some question or reply to already created ones

# Achitecture

- Game
	- map
		- player
			- ch1
			- ch2

- Player
	- name
	isHuman
	theme

- Character
	- x
	- y
	- skin

- Board
	- field: [][]Cell

- Cell
  - height

# Game Phases

1. Loop aksing players info (name, etc...). Max 4 players
2. Sort starting player
3. for each player:
   1. Check lose
   2. Ask for character
   3. Ask where to move
   4. Check if valid move
   5. Move
   6. Check win
   7. Ask where to build
   8. Check if valid build
   9. Build

# UML
![UML](https://app.diagrams.net/?title=Santorini#R7Vprc9o4FP01zGx3Jjt%2BYEM%2BxhBIW5pNmm0T%2Bk2xha1GtlhbBMiv75UtP8DgmJQATZjJZKwryZLuOffqSKahd%2FxZP0Rj7wtzMG1oijNr6N2Gpp0aTfgvDPPEYBpqYnBD4iQmJTfckCecGNXUOiEOjqQtMXHGKCfjRaPNggDbfMGGwpBNF5uNGHUWDGPk4pLhxka0bL0lDvcSa1tr5fYLTFwvHVk1T5MaH6WN5UoiDzlsWjDp5w29EzLGkyd%2F1sFU%2BC71y%2B3H%2BS0dPJj9T9fR%2F%2Bib9fm%2Fy%2B8nyct6m3TJlhDigL%2F41U8Po97F9%2BbPu%2FFV73qq9M%2FuP53ILsojohPpryuK5jhsaGZD0wFD3ZKL5%2FPUo9GU%2BBQFULJGLOA3sgb8YSFK3ACebZgovES3HnHICYBxJis4G4PV9gh1BmjOJmI5EUf2Q1qyPBaSJ3gtolClggGqQy55BbMqtrgRPcEsphniCNpcpT5SM9MARVy2sRmlaByR%2B3jCoomPQpcEFuOc%2BemL2CRwsCNLGehxgYfsIaORmjnnWWQkgsIbeFbgpUSqj5mPeTiHJrLWaElkZNSphixPcw6rprR5Bf62pQ3JsHGzV2ejfYUwQ4ELPsiG0%2FSl4Zo1hxM0KQ6HKOAeII4t4cWoSEh4KKw0N8U03YCyWomyl8gXa%2FkL4CGB%2B6GKu%2BB%2BXuApxSO%2BlqXRGNnwvkHcptvMLV%2Fl%2BoWJQd8RjRniEcfBQcwgjjhKSCZmMWYk4LGDDAv%2BwI0d5R%2BjYcCEOlBW8zL8ieYh77AAVoNIzCoMDJ5iweIqcnrcp%2FKxzMrK4H%2BelfNFtDdlRZGEC3SowN51%2F9VHT0N8Obzye73P6vTWNU60cr76GF1MfBQI%2FO9hXzmi%2F2roG9qe0VeNEqLYgc1eFsFxHnNZgOh5bl3yWd5mwATMcRL%2FiTmfyx0GTTjL3SlqwYvh%2FE72jwtDUQDIZLE7K1Z252thqNoaIjYJbVwBl3QqbIYuroJVlcJGeKYS1RBTxMnjokz6Hcyqpl0I2H6cretpCuWoKbanKVR1cZPXlBUBfboioFWl%2BRJRoS6JCr19uKKitUYHR2JfkZJYpO%2Fum95fau4cZm3mVeiGVUTbxs4xGJ4O7Ghy3Z4pk45jfiNm80cWKzm8HZEuYJ1f0LghVIVJhQP%2BTk8%2FMIySNkmAhyYurwA9ieMOowxyVDdgSS4jlC6Z%2FnBylECvy5eNZMVrkWPl9NolclgMhY6I%2FPjhw5uL863qyI2zQau9Z8BVfR3iR1Wyc1Wim0uqpF3zmJHmjc1Eid6sIUpahyFK1GaJpz2CaZyZ4v7JxpTqkjeWpGqmnyyYD%2FIWo%2FKItiBHUIhskUmOOWj%2FOaj2besLT0aGsTTeYVy3rr50KSchKZL1MzCngvmdJp9WbartIfmsPAqVr8%2FvEiSFD45nm03PNgkDDuLOdCXcZak7PML9anCvOtrs9Iq8nKyTXkdRsXNRsfwJt%2F7BpoIcf9zBZiVLy%2FpX%2FgQDDjYQrO%2FhLFMZvVtXE%2BlN%2FPY%2FyZklrN7PJzmZ7J%2F%2FJCdx3f0nuZWSQDkkyLR6mC25bosgpuAUUVwfnPvCrHlImO0Qssoj6bOBp70OZFDMf%2FWX7In5Tyf1818%3D)