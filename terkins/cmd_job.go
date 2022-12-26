package main

import (
	"context"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"sync"
)

var CmdJob = &cobra.Command{
	Use:   "job",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		ReadSetting()
		getSession()
		goter.Required(host, func(u string) bool { return u != "" }, "run without host", func() { _ = cmd.Help() })
		goter.Required(user, func(u string) bool { return u != "" }, "run without username", func() { _ = cmd.Help() })
		goter.Required(pass, func(u string) bool { return u != "" }, "run without password", func() { _ = cmd.Help() })
		filters := make([]*regexp.Regexp, 0)
		if len(args) > 0 {
			for _, arg := range args {
				compile, err := regexp.Compile(arg)
				if err != nil {
					_, _ = os.Stderr.WriteString(fmt.Sprintf("wrong regExp:%s"))
					os.Exit(1)
				}
				filters = append(filters, compile)
			}
		}
		jobs, _ := session.GetAllJobs(context.Background())
		group := sync.WaitGroup{}
		mutex := sync.Mutex{}
		results := make([]string, 0)
		for _, job := range jobs {
			job := job
			group.Add(1)
			go func() {
				defer group.Done()
				innerJobs, _ := job.GetInnerJobs(context.Background())
				mutex.Lock()
				defer mutex.Unlock()
				for _, innerJob := range innerJobs {
					for _, filter := range filters {
						if !filter.MatchString(innerJob.Base) {
							goto EXIT
						}
					}
					results = append(results, innerJob.Base)
				EXIT:
				}
			}()
		}
		group.Wait()
		for _, result := range results {
			fmt.Println(result)
		}
	},
}
