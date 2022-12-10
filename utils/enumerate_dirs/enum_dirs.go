package enumerate_dirs

import (
	"fmt"
	"log"
	"os/exec"
	"quick_passive_recon/pkg/args_proc"
	"quick_passive_recon/pkg/run_settings"
)

func BruteForceDirs(command *args_proc.Command, settings *run_settings.Settings) {
	fmt.Println("Enumerating directories using gobuster on", command.Url)
	word_list := "~/common.txt"

	//Add gobuster to $PATH
	arg1 := "gobuster"
	arg2 := "dir"
	arg3 := fmt.Sprintf("-u %s", command.Url)
	arg4 := fmt.Sprintf("-w %s", word_list)

	fmt.Println("Executing command:", arg1, arg2, arg3, arg4)
	cmd := exec.Command(arg1, arg2, arg3, arg4)

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("FATAL:", err)
	}

	fmt.Println(string(stdout))
}
