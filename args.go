package main

import (
	"fmt"
)

const (
	PATH = "path"
	P = "p" // short version of path
	ACTIVITY = "activity"
	TIME = "time"
	DATE = "date"
)

// Struct RequiredArgs contains values of each argument. RequiredArguments with empty values are
// considered as not passed.
type RequiredArgs struct {
	dataFilePath string
	activity string
	time string
	date string
}

func (ra *RequiredArgs) String() string {
	return ra.date + "," + ra.time + "," + ra.activity
}

// ParseArgs function parses an array of argument strings and returns a new instance
// of RequiredArgs struct type with correct values.
// The function returns an error if the arguments are not passed properly.
func ParseArgs(args []string) (*RequiredArgs, error) {
	var argName string
	a := &RequiredArgs{}

	for i := 1; i < len(args); {
		arg := args[i]

		if len(arg) < 2 {
			return nil, fmt.Errorf("Invalid flag: %s", arg)
		}

		fullArg := true
		if arg[0] == 45 { // 45 is the ascii/UTF-8 code for '-'
			if arg[1] == 45 {
				argName = arg[2:]
			} else {
				fullArg = false
				argName = arg[1:]
			}

			var err error
			i, err = setArgVal(a, argName, i, args, fullArg)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("Command '" + arg + "' not found")
		}
	}

	err := checkRequiredArgsPassed(a)
	if (err != nil) {
		return nil, err
	}

	return a, nil
}

// The function setArgVal assigns the value of an argument specified by a given argName
// to the corresponding field of a given RequiredArgs type instance. 
// The function returns the updated index for the given args array, which is an array of 
// raw argument values similar to os.Args.
// The function returns an error if:
//   - the format of the argument name is incorrect or the arg name doesn't exist.
//   - the argument value is empty or doesn't exist.
func setArgVal(a *RequiredArgs, argName string, argIndex int, args []string, fullArg bool) (int, error) {
	if fullArg {
		return setFullArgVal(a, argName, argIndex, args)
	}

	return setShortArgVal(a, argName, argIndex, args)
}

func setFullArgVal(a *RequiredArgs, argName string, argIndex int, args []string) (int, error) {
	if len(argName) == 0 {
		return argIndex, fmt.Errorf("No argument name specified after --")
	}
	if argIndex >= len(args) - 1 || len(args[argIndex+1]) == 0 {
		return argIndex, fmt.Errorf("No value specified for the argument '" + argName + "'")
	}

	if argName == PATH {
		return argIndex + 2, setDataFilePathVal(a, args[argIndex + 1])
	}
	if argName == ACTIVITY {
		return argIndex + 2, setActivityVal(a, args[argIndex + 1])
	}
	if argName == TIME {
		return argIndex + 2, setTimeVal(a, args[argIndex + 1])
	}
	if argName == DATE {
		return argIndex + 2, setDateVal(a, args[argIndex + 1])
	}

	return argIndex, fmt.Errorf("Argument '" + argName + "' not found")
}

func setShortArgVal(a *RequiredArgs, argName string, argIndex int, args []string) (int, error) {
	if len(argName) == 0 {
		return argIndex, fmt.Errorf("No argument name specified after -")
	}
	if argIndex >= len(args) - 1 || len(args[argIndex+1]) == 0 {
		return argIndex, fmt.Errorf("No value specified for the argument '" + argName + "'")
	}

	if argName == P {
		return argIndex + 2, setDataFilePathVal(a, args[argIndex + 1])
	}

	return argIndex, fmt.Errorf("Argument '" + argName + "' not found")
}

// Returns an error if the filepath is empty
func setDataFilePathVal(a *RequiredArgs, fPath string) error {
	if len(fPath) == 0 {
		return fmt.Errorf("The data filepath can't be empty")
	}

	a.SetDataFilePath(fPath)

	return nil
}

// Returns an error if the activity name is empty.
func setActivityVal(a *RequiredArgs, actName string) error {
	if len(actName) == 0 {
		return fmt.Errorf("The activity name can't be empty")
	}

	a.SetActivity(actName)

	return nil
}

// Returns an error if the time is of incorrect format
func setTimeVal(a *RequiredArgs, t string) error {
	return a.SetTime(t)
}

func setDateVal(a *RequiredArgs, d string) error {
	return a.SetDate(d)
}

func checkRequiredArgsPassed(a *RequiredArgs) error {
	if len(a.dataFilePath) == 0 {
		return fmt.Errorf("Data storage file path is not set. Use --path or -p flags for that.")
	}
	if len(a.activity) == 0 {
		return fmt.Errorf("Activity name is not set. Use --activity flag for that.")
	}
	if len(a.time) == 0 {
		return fmt.Errorf("Time is not set. Use --time flag for that.")
	}
	if len(a.date) == 0 {
		return fmt.Errorf("Date is not set. Use --date flag for that.")
	}

	return nil
}

// SETTERS
func (a *RequiredArgs) SetDataFilePath(fPath string) {
	a.dataFilePath = fPath
}

func (a *RequiredArgs) SetActivity(act string) {
	a.activity = act
}

func (a *RequiredArgs) SetTime(t string) error {
	// TODO: Test the time format
	a.time = t
	return nil
}

func (a *RequiredArgs) SetDate(d string) error {
	// TODO: Test the date format
	a.date = d
	return nil
}
