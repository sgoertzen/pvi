#Pom Version Information (pvi)
Compile pom details across many projects

[![Build Status](https://travis-ci.org/sgoertzen/pvi.svg?branch=master)](https://travis-ci.org/sgoertzen/pvi)

Pom Version Information is a command line tool for analyzing multiple maven projects and printing out how they are related.

## Install:
```
go get github.com/sgoertzen/pvi/cmd/pvi
```

## Usage:
```
pvi [-filename (name)] [-format (text or json)] [-path (path)] [-nocolor]
-filename string
    The file in which the output should be stored.  If this is left off the output will be printed to the console
-format 'text' or 'json'
    Specify the output format.  Should be either 'text' or 'json' (default "text")
-path directory
    The directory that contains subfolders with maven projects.  Example: '/user/code/projects/' (default ".")
-nocolor
    Do not color the output.  Ignored if filename is specified.
```

#### Examples
Run the program in the current directory
```
pvi 
```

With options
```
pvi -path /usr/code -format json
```

##Example:
Given the directory structure as follows:

```
/users
    /myname
        /code
            /project1
                /pom.xml
                ...
            /project2
                /pom.xml
                ...
```

You would run the tool with:
```
./pvi /users/myname/code
```

##Example Output:
```
my-parent-project (1.4.2)
--my-dependant-project (2.0.24)

```

If the child project points to an older version of the parent pom then you would get a message like
```
my-parent-project (1.4.2)
--my-dependant-project (2.0.24) ** Warning: looking for parent version: 1.3.1
```

##Development
### Running integration tests
```
go test -tags=integration
```
