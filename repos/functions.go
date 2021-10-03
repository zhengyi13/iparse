package repos

import (
	"os/exec"
	"strings"
)

func GetRepos() (repos []string, err error) {
	repolist, err := exec.Command("/usr/bin/dnf", "repolist").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(repolist), "\n")
	var rs []string
	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}
		reponame := words[0]
		if reponame == "repo" {
			continue
		}
		rs = append(rs, reponame)
	}
	return rs, nil
}

func GetRepoIni(repo string) (ini []byte, err error) {
	var dumpCmd = "/usr/bin/dnf"
	var dumpArgs = []string{
		"config-manager",
		"--dump",
		repo,
	}

	repodata, err := exec.Command(dumpCmd, dumpArgs...).Output()
	if err != nil {
		return nil, err
	}
	return repodata, nil
}
