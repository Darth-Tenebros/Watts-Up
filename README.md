# Eskom Se Poes

![demo image of esp](./img/demo.gif)

Eskom Se Poes is a cli based Load Shedding schedule checker.
Simply install/run the code and you'll be presented with locations. Select the one you care about and you'll be presented with a table showing the schedule for that location.

## Why?
why not?
Also, I just needed a project to get me familiar with the Golang programming language.


## Data
ESP uses data from a [FREE API](https://eskom-calendar-api.shuttleapp.rs/)🐐. You can find the source code for the api [HERE](https://github.com/beyarkay/eskom-calendar-api).

## usage
_you'll need to have go1.21.4 installed to build the app yourself_
1. clone this repo and run the code everytime you want to use the app

2. if on Ubuntu/Ubuntu based OS, find binary in the bin folder and do:
    ```
    cd bin
    sudo mv esp /usr/local/bin/ 
    ```
   then you'll be able to use it system-wide

3. if on a different system (*nix) and you want a binary, clone this repo and do:
    ```
    make
    ```
   then install as in `step 2`


4. windows: i don't know

## TODO
- document code
- add more features
