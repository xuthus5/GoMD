package tools

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

// 热更新 重启服务
func ReloadServer() error {
	cmd := exec.Command("kill", "-HUP", strconv.Itoa(os.Getpid()))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		return err
	}
	log.Println("Server stopped!")
	return nil
}

// 热编译 重启服务
func RebuildServer(servername string) error {
	//关闭服务
	cmd := exec.Command("pkill", servername)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// 运行命令
	if err := cmd.Start(); err != nil {
		return err
	}
	// 保证关闭输出流
	stdout.Close()

	//编译运行
	cmd = exec.Command("bee", "run")
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// 运行命令
	if err := cmd.Start(); err != nil {
		return err
	}
	// 保证关闭输出流
	stdout.Close()
	log.Println("Server Rebuild!")
	return nil
}
