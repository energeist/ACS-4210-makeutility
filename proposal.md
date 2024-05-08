# ACS-4210 - FINAL PROJECT PROPOSAL

## GOALS

The goal of this project is to create a utility that can quickly calculate the potential outcomes of a Starcraft II fantasy match between two pro players with data available on record.  This utility will use data from [publicly available APIs](http://aligulac.com) and information about the current [professional tournament environment](https://liquipedia.net/starcraft2/Main_Page) to make data-informed predictions.

## AUDIENCE

The audience for this utility is esports fans, particularly those who follow competitive real-time strategy games like Starcraft II.

## PACKAGES USED

| Package | Used For |
| ------- | -------- |
| github.com/joho/godotenv | .env file utility for private keys |
| Gin     | Server & API routing |
| GORM    | Database persistence |

## TASK BREAKDOWN FOR V1.0

| Task Description | Hours | Done By |
| ---------------- | ----- | ------- |
| Establish SQLite DB connection | 3 | Apr 24 |
| Data Research | 3 | Apr 28 |
| Plan basic implementation | 6 | May 1 | 
| Assemble Seed Data | 3 | May 3 |
| Build basic implementation | 12 | May 6 |
| Add basic calculation logic for single match for MVP | 6 | May 8 |

## FUTURE ADDITIONS

- Expand functionality from single 1v1 match to tournament bracket
- - Expand from single game matches to Best-of-N, as would be typical of a real tournament
- Allow user to select players from input instead of random choice from top X
- Allow selection of map pool instead of using GSL 2024 Season 2 map pool for all calculations
- Concurrent API calls to Aligulac to speed up initialization process
- Concurrent Match outcome calculation
- Parallel iteration to make use of threading
- Improvement of calculation model
- Further abstraction of code modules for better separation of concerns
- - Headless client should be as light weight as possible, and could also be supplemented with an actual front end
- - Calculation engine should live outside of the server as its own microservice / job platform
- Implement memcached to prevent overquerying of API & overloading of DB writes