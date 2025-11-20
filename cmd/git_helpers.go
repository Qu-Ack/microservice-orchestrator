package main 


import (
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
)

func make_clone_path(subdomain string) string {
	return filepath.Join("/home/quack/hosting", fmt.Sprintf("%v_%v", subdomain, random_string_from_charset(6)))
}

func git_clone(clone_url string, branch string, depth string, subdomain string) error {
	cmd := exec.Command("git", "clone", "--depth", depth, "-b", "--single-branch" ,branch, clone_url, make_clone_path(subdomain))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
