# Eskom Se Poes

![demo image of esp](./img/demo.gif)

Watts-Up is a cli based Load Shedding schedule checker.
Simply install/run the code and you'll be presented with locations. Select the one you care about and you'll be presented with a table showing the schedule for that location.

## Why?
why not?
Also, I just needed a project to get me familiar with the Golang programming language and some other technologies.


## Data
Watts-Up uses data from a [FREE API](https://eskom-calendar-api.shuttleapp.rs/)🐐. You can find the source code for the api [HERE](https://github.com/beyarkay/eskom-calendar-api).

## usage
_you'll need to have go1.21.4 installed to build the app yourself_
1. option 1: clone this repo and run the code everytime you want to use the app

2. option 2: if on a *nix system, do (from root of project):
    ```
    make
    cd bin
    sudo mv esp /usr/local/bin/
    ```
 - then you'll be able to use it system-wide like:
   ```
    watts
   ```


3. windows: i don't know

## Keys
from the list view:
- `a` : the area name that the selector is currently on will be inserted into your favourites.
- `b` : allows you to go back to the main list view (from favourites or table view).
- `enter` : in list view this will take you to the table view (load schedule) for the current area.
- `q` or `ctrl + c` : quits the program
- `d`: when in the `Favourites` view, press to delete selected area from favourites.


## TODO
- [ ] document code.
- [x] add delete option in the favourites view
- [ ] maybe add aliases for the area names. right now some locations are not searchable by name: e.g Samora,Philippi etc
  fall under `city-of-cape-town-area-16`. Adding aliases would make filtering better.
- [ ] add help menu to all views.
- [ ] Dockerize project
