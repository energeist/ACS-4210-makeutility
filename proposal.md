# ACS-4210 - FINAL PROJECT PROPOSAL

## GOALS

The goal of this project is to create a utility that can quickly calculate the potential outcomes of a Starcraft II fantasy tournament.  This utility will use data from [publicly available APIs](http://aligulac.com) and information about the current [professional tournament environment](https://liquipedia.net/starcraft2/Main_Page) to make data-informed predictions.

## AUDIENCE

The audience for this utility is esports fans, particularly those who follow competitive real-time strategy games like Starcraft II.

## PACKAGES USED

| Package | Used For |
| ------- | -------- |
| github.com/joho/godotenv | .env file utility for private keys |
| Gin     | Server & API routing |
| GORM    | Database persistence |

## TASK BREAKDOWN

| Task Description | Hours | Done By |
| ---------------- | ----- | ------- |
| Establish SQLite DB connection | 3 | Apr 24 |
| Data Research | 3 | Apr 28 |
| Plan basic implementation | 6 | May 1 | 
| Assemble Seed Data | 3 | May 3 |
| Build basic implementation | 12 | May 6 |
| Add calculation logic | 6 | May 8 |
| Add customizable model weights for tuning | 6 | May 10 |
