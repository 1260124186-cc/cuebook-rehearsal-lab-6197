package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/format"
	"example.com/cuebook-rehearsal-lab/internal/model"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, format.Error(err))
		os.Exit(1)
	}
}
func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("expected compose, review, publish, inspect, complete, or serve")
	}
	service := app.NewDemoService()
	switch args[0] {
	case "compose":
		return compose(service, args[1:])
	case "review":
		return review(service, args[1:])
	case "publish":
		return publish(service, args[1:])
	case "inspect":
		return inspect(service, args[1:])
	case "complete":
		return complete(service, args[1:])
	case "serve":
		return serve(service, args[1:])
	default:
		return fmt.Errorf("unknown command %s", args[0])
	}
}
func compose(service *app.Service, args []string) error {
	flags := flag.NewFlagSet("compose", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	show := flags.String("show", "", "show name")
	director := flags.String("director", "", "director name")
	jsonOut := flags.Bool("json", false, "structured output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	response, err := service.Assemble(*show, *director)
	if err != nil {
		return err
	}
	if *jsonOut {
		return printJSON(format.AssembleJSON(response))
	}
	fmt.Println(format.Assemble(response))
	return nil
}
func review(service *app.Service, args []string) error {
	flags := flag.NewFlagSet("review", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	show := flags.String("show", "", "show name")
	department := flags.String("department", "", "department")
	jsonOut := flags.Bool("json", false, "structured output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	response, err := service.Review(*show, model.Department(strings.TrimSpace(*department)), "department-lead")
	if err != nil {
		return err
	}
	if *jsonOut {
		return printJSON(format.ReviewJSON(response))
	}
	fmt.Println(format.Review(response))
	return nil
}
func publish(service *app.Service, args []string) error {
	flags := flag.NewFlagSet("publish", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	show := flags.String("show", "", "show name")
	jsonOut := flags.Bool("json", false, "structured output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	response, err := service.Publish(*show)
	if err != nil {
		return err
	}
	if *jsonOut {
		return printJSON(format.PublishJSON(response))
	}
	fmt.Println(format.Publish(response))
	return nil
}
func inspect(service *app.Service, args []string) error {
	flags := flag.NewFlagSet("inspect", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	show := flags.String("show", "", "show name")
	jsonOut := flags.Bool("json", false, "structured output")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if _, err := service.Assemble(*show, "Mira"); err != nil {
		return err
	}
	response, err := service.Inspect(*show)
	if err != nil {
		return err
	}
	report, err := service.Report(*show)
	if err != nil {
		return err
	}
	history, err := service.History(*show)
	if err != nil {
		return err
	}
	if *jsonOut {
		return printJSON(format.InspectJSON(response))
	}
	fmt.Println(format.Inspect(response))
	fmt.Println(format.Report(report))
	fmt.Println(history.Text())
	return nil
}
func complete(service *app.Service, args []string) error {
	flags := flag.NewFlagSet("complete", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	show := flags.String("show", "", "show name")
	director := flags.String("director", "Mira", "director name")
	if err := flags.Parse(args); err != nil {
		return err
	}
	result, err := service.Complete(*show, *director)
	if err != nil {
		return err
	}
	fmt.Println(app.ScenarioLabel(result))
	return nil
}
func printJSON(payload []byte, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(string(payload))
	return nil
}
