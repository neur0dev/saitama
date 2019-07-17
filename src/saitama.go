/* Saitama app, kill process by name with one punch
Developed by lobocode - lobocode@fedoraproject.org */

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// args holds the commandline args
var args []string

func findAndKillProcessByName(path string, info os.FileInfo, err error) error {

	// We are only interested in files with a path looking like /proc/<pid>/status.
	if strings.Count(path, "/") == 3 {
		if strings.Contains(path, "/status") {

			pid, _ := strconv.Atoi(path[6:strings.LastIndex(path, "/")])

			readIndex, _ := ioutil.ReadFile(path)

			// Extract the process name from within the first line in the buffer
			processName := string(readIndex[6:bytes.IndexByte(readIndex, '\n')])

			oh :=
				`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⣠⣶⡾⠏⠉⠙⠳⢦⡀⠀⠀⠀⢠⠞⠉⠙⠲⡀⠀
		⠀⠀⠀⣴⠿⠏⠀⠀⠀⠀⠀⠀⢳⡀⠀ ⡏⠀⠀⠀⠀ ⢷
		⠀⠀⢠⣟⣋⡀⢀⣀⣀⡀⠀⣀⡀⣧⠀⢸⠀⠀⠀⠀⠀ ⡇
		⠀⠀⢸⣯⡭⠁⠸⣛⣟⠆⡴⣻⡲⣿⠀⣸⠀⠀Oh!⠀⡇
		⠀⠀⣟⣿⡭⠀⠀⠀⠀⠀⢱⠀⠀⣿⠀⢹⠀⠀⠀⠀⠀ ⡇
		⠀⠀⠙⢿⣯⠄⠀⠀⠀⢀⡀⠀⠀⡿⠀⠀⡇⠀⠀⠀⠀⡼
		⠀⠀⠀⠀⠹⣶⠆⠀⠀⠀⠀⠀⡴⠃⠀⠀⠘⠤⣄⣠⠞⠀
		⠀⠀⠀⠀⠀⢸⣷⡦⢤⡤⢤⣞⣁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⢀⣤⣴⣿⣏⠁⠀⠀⠸⣏⢯⣷⣖⣦⡀⠀⠀⠀⠀⠀⠀
		⢀⣾⣽⣿⣿⣿⣿⠛⢲⣶⣾⢉⡷⣿⣿⠵⣿⠀⠀⠀⠀⠀⠀
		⣼⣿⠍⠉⣿⡭⠉⠙⢺⣇⣼⡏⠀⠀⠀⣄⢸⠀⠀⠀⠀⠀⠀
		⣿⣿⣧⣀⣿.........⣀⣰⣏⣘⣆ 
		`

			switch args[1] {
			case "--help", "-h":
				log.Fatalln("\nUsage: saitama <processname>\nKill process with one punch\n\nMandatory arguments\n\n-h  --help	display this help and exit\n-l  --list	list process by name")
			case "-l", "--list":
				fmt.Printf("%s\n", processName)
			default:
				if processName == args[1] {

					proc, _ := os.FindProcess(pid)

					if err = proc.Kill(); err != nil {
						fmt.Printf("Warning: This process owner is 'root'\nPlease use 'sudo'\n")
						os.Exit(1)
						//log.Fatal(err)
					} else {
						fmt.Printf("Killing %s with one punch \n", args[1])
						fmt.Printf("PID: %d %s %s .\n", pid, processName, oh)
						// Execute and exit
						os.Exit(1)
					}
				}

			}
		}
	}

	return nil
}

// main is the entry point of any go application
func main() {
	args = os.Args

	if len(args) == 1 {
		log.Fatalln("\nSaitama: missing operand\nTry 'saitama --help' for more information")
	}

	saitamaExec := filepath.Walk("/proc", findAndKillProcessByName)
	fmt.Printf("%s", saitamaExec)

}
