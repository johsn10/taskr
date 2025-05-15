# TaskR - Productivity IDE tailored for rScript
Manage, write and run your scripts all from your terminal. Automate your tasks with a flexible, yet simple all-in-one. 

## How to "rScript" your way around
rScript is a scripting language mainly to automate tedious terminal work. It's an interpreted, statically typed language with intuitive syntax. 

## Example of syntax
````bash
// This is a comment

-> touch example.txt

let fileContents <- cat example.txt
let siteContents <- curl google.com

let isEqual = false
if fileContents == siteContents 
|true let isEqual = true
|false -> echo "isEqual was false""

!output isEqual
````