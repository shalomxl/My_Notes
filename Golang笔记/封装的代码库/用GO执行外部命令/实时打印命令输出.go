func asyncLog(reader io.ReadCloser) error {
	cache := "" //缓存不足一行的日志信息
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
			fmt.Printf("%s%s\n", cache, line)
			cache = s[len(s)-1]
		}
		if err == io.EOF{
			break
		}
	}
	return nil
}

func execute(cmd *exec.Cmd) error {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Starting command: %s Error: %s ......", err.Error(), cmd.String())
		return err
	}

	go asyncLog(stdout)
	go asyncLog(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Waiting for command execution: %s Error: %s ......", cmd.String(), err.Error())
		return err
	}

	return nil
}

//	exec包相关文档链接🔗：<https://pkg.go.dev/os/exec@go1.17.7>		TODO：其中涉及到unix操作系统相关概念，需要学习一哈