/*
	PLEASE NOTE

	This simple app is done for the backend developer test for the PT. Telekomunikasi Indonesia. If you have any queries, please contact me at the below contact.

	Name	: Syafrial Azis
	Email	: syafrialazis13@gmail.com

	This is an automated editor signature template

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

/*

	Declaring variables

*/
var targetDir string
var targetFormat string
var originFile string
var fileName string
var fileNameWithExtension []string
var dirInputFlag int

/*

	Processing the arguments.
	Looping through the parameters that is sent to the command.

*/

func processArgs(args []string) string {

	originFile = args[1]
	fileNameWithExtension := strings.Split(path.Base(originFile), ".")
	fileName = fileNameWithExtension[0]
	dirInputFlag = 0
	targetFormat = "txt"
	targetDir, dirErr := os.Getwd()
	fmt.Println("Dir selected", targetDir)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	for i := range args {

		if strings.ToUpper(args[i]) == "-O" {
			dirInputFlag = 1
			targetDir = args[i+1]
		}

		if strings.ToUpper(args[i]) == "-T" {
			targetFormat = args[i+1]
		}
	}

	return targetDir
}

/*

	Main executable block

*/

func main() {
	/*

		Help parameter comes first

	*/

	if len(os.Args) == 2 && strings.ToUpper(os.Args[1]) == "-H" {
		fmt.Println("Please use the utility with the format ==>   logCopy [-h] [-t [FileType]] [-o [DestinationDirectory]]")
	} else {

		targetDir = processArgs(os.Args)

		/*

			Checking the OS compatibility with the path.Base() function
			Due to its limitation that only allows it to handle file names and directories from UNIX based systems only (that uses "/" instead of "\")

		*/
		if dirInputFlag != 1 {
			if runtime.GOOS == "windows" {
				fmt.Println("Please provide the destination directory [-o] when invoking this app from windows based operating systems")
				return
			} else {
				targetDir += "/" + fileName + "." + targetFormat
				fmt.Println("Test path", path.Dir(targetDir))
				fmt.Println("File Name", fileName)
			}
		}

		/*

			Copying the file

		*/
		input, err := ioutil.ReadFile(originFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(targetDir)

		err = ioutil.WriteFile(targetDir, input, 0644)
		if err != nil {
			fmt.Println("Error creating file", targetDir)
			fmt.Println(err)
			return
		}

		fmt.Println(fileName)

	}
}
