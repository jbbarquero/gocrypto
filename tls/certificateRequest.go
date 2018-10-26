package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var env = flag.String("env", "DEV", "Environment: DEV, TST, ACC, or PRD")
var app = flag.String("app", "", "Application name")
var machines = flag.String("machines", "", "Machine name")

func main() {

	flag.Parse()

	validateFlags(*env, *app, *machines)

	for _, machine := range strings.Split(*machines, ",") {
		invokeCSR(*env, *app, machine)
	}

}

func validateFlags(env string, app string, machines string) {
	switch env {
	case "DEV":
	case "TST":
	case "ACC":
	case "PRD":
	default:
		log.Fatalf("Error, unexpected env: [%s]. Usage generate_csr ENV APP MACHINE\n", env)
	}

	if app == "" {
		log.Fatal("Error, empty app. Usage generate_csr ENV APP MACHINE")
	}

	if machines == "" {
		log.Fatal("Error, empty machine. Usage generate_csr ENV APP MACHINE")
	}
}

func invokeCSR(env string, app string, machine string) {

	var cmdOut []byte
	var err error

	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--verify", "HEAD"}
	log.Printf("%s %s\n", cmdName, strings.Join(cmdArgs, " "))
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	firstSix := sha[:6]
	fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)

	lsCmd := "/bin/bash"
	lsArgs := []string{"-c", "ls", "-las", "localhost.*"}
	log.Printf("%s %s\n", lsCmd, strings.Join(lsArgs, " "))
	var lsOut []byte
	if lsOut, err = exec.Command(lsCmd, lsArgs...).CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, "ls error: "+err.Error(), err)
		os.Exit(2)
	}
	fmt.Println("ls: ", string(lsOut))

	command := "openssl"
	args := []string{"req", "-utf8", "-sha256", "-newkey rsa:4096", "-keyout " + machine + ".key", "-nodes",
		"-out " + machine + ".csr",
		"-subj \"/O=Malsolo/OU=Unit1/OU=Certificate Authorities/OU=" + env + "/OU=" + app + "/CN=" + machine + ".com\"",
	}

	log.Printf("%s %s\n", command, strings.Join(args, " "))

	cmd := exec.Command(command, args...)
	var out []byte
	out, error := cmd.CombinedOutput()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("openssl: ", string(out))
}
