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
