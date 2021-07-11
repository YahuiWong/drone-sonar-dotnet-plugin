package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	Config struct {
		Key   string
		Name  string
		Host  string
		Token string

		Version         string
		Branch          string
		Sources         string
		Timeout         string
		Inclusions      string
		Exclusions      string
		Level           string
		ShowProfiling   string
		BranchAnalysis  bool
		UsingProperties bool
		buildfile       string
	}
	Plugin struct {
		Config Config
	}
)

func (p Plugin) Exec() error {
	args := []string{
		"/d:sonar.host.url=" + p.Config.Host + "",
		"/d:sonar.login=" + p.Config.Token + "",
	}

	if !p.Config.UsingProperties {
		argsParameter := []string{
			"/k:" + strings.Replace(p.Config.Key, "/", "-", -1) + "",
			"/n:" + p.Config.Name,
			"/version:" + p.Config.Version,
			"/d:sonar.sources=" + p.Config.Sources,
			"/d:sonar.ws.timeout=" + p.Config.Timeout,
			// "/d:sonar.inclusions=" + p.Config.Inclusions,
			// "/d:sonar.exclusions=" + p.Config.Exclusions,
			"/d:sonar.log.level=" + p.Config.Level,
			"/d:sonar.showProfiling=" + p.Config.ShowProfiling,
			"/d:sonar.scm.provider=git",
		}
		args = append(args, argsParameter...)
	}
	if p.Config.Inclusions != "" {
		args = append(args, "/d:sonar.inclusions="+p.Config.Inclusions+"")
	}
	if p.Config.Exclusions != "" {
		args = append(args, "/d:sonar.exclusions="+p.Config.Exclusions+"")
	}
	if p.Config.BranchAnalysis {
		args = append(args, "/d:sonar.branch.name="+p.Config.Branch+"")
	}
	beginargs := append([]string{"sonarscanner", "begin"}, args...)
	begincmd := exec.Command("dotnet", beginargs...)
	// fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
	begincmd.Stdout = os.Stdout
	begincmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Begin:\n")
	begincmderr := begincmd.Run()
	if begincmderr != nil {
		return begincmderr
	}
	buildargs := []string{"build"}
	if p.Config.buildfile != "" {
		buildargs = append(buildargs, p.Config.buildfile)
	}
	buildcmd := exec.Command("dotnet", buildargs...)
	buildcmd.Stdout = os.Stdout
	buildcmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Build:\n")
	builderr := buildcmd.Run()
	if builderr != nil {
		return builderr
	}
	endargs := append([]string{"sonarscanner", "end"}, args...)
	endcmd := exec.Command("dotnet", endargs...)
	endcmd.Stdout = os.Stdout
	endcmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Build:\n")
	enderr := endcmd.Run()
	if enderr != nil {
		return enderr
	}
	return nil
}
