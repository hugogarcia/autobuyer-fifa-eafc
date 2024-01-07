# About
A simple EA FC auto buyer that I created for personal use. But feel free to use it.

It has no interface, and its configuration is manual. The only automated part is the card buyer.

> This auto buyer simulates the EA FC Web App, so if you're playing on you're console you won't be able to run this bot.

# Built with
- Go 1.19

# WARNING!!!!!!!!!!!!!!!!!
YOUR EA FC ACCOUNT ON THE WEB MAY BE BLOCKED FOR TRADING

KEEP THIS RISK IN MIND

# Configure Spreadsheet (CSV)
1. Find the player you want to search for on [Futbin](https://www.futbin.com/).
2. Get the `ID` code, which will be on the left side with the player's information.
3. Add the ID, Name, and Maximum Bid Value for that card to the spreadsheet. The information should be separated by semicolons (;).

# Configure Environment Variables
Firstly, open the `.env` file in the folder. I will show the configuration of each variable below:

### `GAME_SKU`
  - Enter the code for the platform you use to play, the options are:
    - FFA24XBO - Xbox One
    - FFA24XSX - Xbox X
    - FFA24PS5 - PS5
    - FFA24PS4 - PS4
    - FFA24PCC - PC
### `USER_ID`
  - Access the EA FC Web App website and log in.
  - After logged, press F12.
  - Go to the `Console` tab.
  - Type the command `localStorage.getItem('_eadp.identity.pidId')`.
  - Copy this code and paste it into the USER_ID in the file.
### `TOKEN`
  - Access the EA FC Web App website and log in.
  - After logged, press F12.
  - Go to the `Console` tab.
  - Type the command `localStorage.getItem('_eadp.identity.access_token')`.
  - Copy this text and paste it into the TOKEN in the file.

# How to Run
After completing the [configure environment variables](#configure-environment-variables) step, you can run the bot.

## Executable (For windows only)
In the folder, there will be an `exe` file that will start the bot.

## Terminal
- Install Go 1.19
- Run the command `go run main.go`

# My bot is closing
## Cause 1:
  If the bot closes immediately upon execution, there may be an issue with reading the CSV. Check if it is configured correctly.

## Cause 2:
  If the bot closes after running and making bids, it is because EA FC has temporarily blocked access. This is normal, EA FC blocks when there are many web trade attempts. Log out of the account and log in again on the website, and redo the process of configuring the [TOKEN](#token) variable.
## Cause 3:
  If the bot closes after running and making bids, try logging in again on the EA FC Web and check if you can log in. EA FC may block your account on the web; in that case, the bot will never work again 😢 The block will only be on EA FC Web; you can still trade normally within the game.

# Donation
If you liked the bot and can help me:
 
https://www.buymeacoffee.com/hugogarcia