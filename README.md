# cron-parser-v1

Cron-parser parses a cron expression by applying the rules and rendering the output in a table format.

Required:-
Go1.20 version

Steps to run:-
1. download the source code
2. make sure go1.20 is installed
3. go to project dir cron-parser-v1/
4. Hit "go run main.go"
5. follow the console instructions

Troubleshooting:-
1. check go version
2. "go version"
3. 	with gvm: "gvm list"
4. Use the Go version
5.   "alias go="go1.20" then "go version"
6.   with gvm: "gvm use go1.20"


Rules defined by the application
1. Allowed special characters:- ["/","-",",","*"]
2. Allowed values:-
3.   minutes - 0-59
4.   hour - 0-23
5.   day of month - 0-31
6.   month - 1-12, Jan-Dec English short month name of exactly three characters. Allowed case - upper, lower, and mixed.
7.   day of week - 0-6, Sun-Sat short name of week of exactly three characters. Allowed case - upper, lower, and mixed.
8. Valid input string:- "* * * * * /find"
9. Invalid input string:- "* * /find"
