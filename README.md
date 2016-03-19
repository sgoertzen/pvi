#Pom Version Information (pvi)
Compile pom details across many projects

[![Build Status](https://travis-ci.org/sgoertzen/pvi.svg?branch=master)](https://travis-ci.org/sgoertzen/pvi)

Pom Version Information is a command line tool for analyzing multiple maven projects and printing out how they are related.

## Install:
```
go get github.com/sgoertzen/pvi/cmd/pvi
go install github.com/sgoertzen/pvi/cmd/pvi
```

## Usage:
```
usage: pvi [<flags>] [<path>]

Flags:
  -?, --help               Show context-sensitive help (also try --help-long and --help-man).
  -o, --format="text"      Specify the output format. Should be either 'text' or 'json'
  -f, --filename=FILENAME  The file in which the output should be stored. If this is left off the output will be printed to the console
  -n, --nocolor            Do not color the output. Ignored if filename is specified.
  -d, --debug              Output debug information during the run.
  -p, --showpath           Show the path information for each project.
  -v, --version            Show application version.

Args:
  [<path>]  The `directory` that contains subfolders with maven projects. Defaults to current directory. Example: '/user/code/projects/'
```

#### Examples
Run the program in the current directory
```
pvi 
```

Return JSON and store into a file
```
pvi ../../MyCode -o=json -f=myoutput.json
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
