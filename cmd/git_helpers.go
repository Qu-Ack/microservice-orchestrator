package main 


import (
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
)

func make_clone_path(subdomain string, deployment_id string) string {
	return filepath.Join("/home/quack/hosting", fmt.Sprintf("%v_%v", subdomain, deployment_id))
}

func git_clone(clone_url string, branch string, depth string, subdomain string, clone_path string) error {
	cmd := exec.Command("git", "clone", "--depth", depth, "-b",  branch, "--single-branch", clone_url, clone_path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
