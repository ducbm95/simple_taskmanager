package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"taskmanager/taskmanager"
)

func main() {
	tm := taskmanager.NewTaskManager()

	for {
		fmt.Print("Please enter command: ADD|REMOVE|START|STOP|IMPORT|EXPORT|EXIT\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := strings.ToLower(scanner.Text())

		if strings.Compare(line, "add") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			fmt.Print("> Cron expression: ")
			scanner.Scan()
			cronExpr := scanner.Text()

			fmt.Print("> Task content: ")
			scanner.Scan()
			taskContent := scanner.Text()

			task := taskmanager.Task{
				Name:        taskName,
				CronExpr:    cronExpr,
				TaskContent: taskContent,
			}

			err := tm.Add(task)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%s added successfully\n", taskName)
			}
		} else if strings.Compare(line, "remove") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			err1 := tm.Remove(taskName)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Printf("%s removed successfully\n", taskName)
			}
		} else if strings.Compare(line, "start") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			err1 := tm.Start(taskName)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Printf("%s started successfully\n", taskName)
			}
		} else if strings.Compare(line, "stop") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			err1 := tm.Stop(taskName)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Printf("%s stopped successfully\n", taskName)
			}
		} else if strings.Compare(line, "export") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			err1 := tm.Export(taskName)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Printf("%s exported successfully\n", taskName)
			}
		} else if strings.Compare(line, "import") == 0 {
			fmt.Print("> Task name: ")
			scanner.Scan()
			taskName := scanner.Text()

			err1 := tm.Import(taskName)
			if err1 != nil {
				fmt.Println(err1)
			} else {
				fmt.Printf("%s imported successfully\n", taskName)
			}
		} else if strings.Compare(line, "exit") == 0 {
			return
		}

	}
}
