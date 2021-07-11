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
	}
	Plugin struct {
		Config Config
	}
)

func (p Plugin) Exec() error {
	args := []string{
		"/d:sonar.host.url=" + p.Config.Host,
		"/d:sonar.login" + p.Config.Token,
	}

	if !p.Config.UsingProperties {
		argsParameter := []string{
			"/k:" + strings.Replace(p.Config.Key, "/", ":", -1),
			"/d:sonar.projectName=" + p.Config.Name,
			"/d:sonar.projectVersion=" + p.Config.Version,
			"/d:sonar.sources=" + p.Config.Sources,
			"/d:sonar.ws.timeout=" + p.Config.Timeout,
			"/d:sonar.inclusions=" + p.Config.Inclusions,
			"/d:sonar.exclusions=" + p.Config.Exclusions,
			"/d:sonar.log.level=" + p.Config.Level,
			"/d:sonar.showProfiling=" + p.Config.ShowProfiling,
			"/d:sonar.scm.provider=git",
		}
		args = append(args, argsParameter...)
	}

	if p.Config.BranchAnalysis {
		args = append(args, "/d:sonar.branch.name="+p.Config.Branch)
	}

	begincmd := exec.Command("dotnet sonarscanner begin", args...)
	// fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
	begincmd.Stdout = os.Stdout
	begincmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Begin:\n")
	begincmderr := begincmd.Run()
	if begincmderr != nil {
		return begincmderr
	}
	buildcmd := exec.Command("dotnet build ")
	buildcmd.Stdout = os.Stdout
	buildcmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Build:\n")
	builderr := buildcmd.Run()
	if builderr != nil {
		return builderr
	}
	endcmd := exec.Command("dotnet sonarscanner end", args...)
	endcmd.Stdout = os.Stdout
	endcmd.Stderr = os.Stderr
	fmt.Printf("==> Code Analysis Build:\n")
	enderr := endcmd.Run()
	if enderr != nil {
		return enderr
	}
	return nil
}
