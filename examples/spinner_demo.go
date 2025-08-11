package main

import (
	"fmt"
	"time"

	"github.com/bm611/go-ph/internal/spinner"
)

func main() {
	fmt.Println("🎯 Spinner Demo - Showcasing different animation styles\n")

	// Demo different spinner styles
	styles := []struct {
		style   spinner.SpinnerStyle
		name    string
		message string
	}{
		{spinner.DotsStyle, "Dots Style", "Loading with dots animation..."},
		{spinner.LineStyle, "Line Style", "Processing with line animation..."},
		{spinner.ArrowStyle, "Arrow Style", "Fetching data with arrows..."},
		{spinner.PulseStyle, "Pulse Style", "Analyzing with pulse animation..."},
		{spinner.BouncingBar, "Bouncing Bar", "Syncing with bouncing bar..."},
	}

	for i, demo := range styles {
		fmt.Printf("%d. %s:\n", i+1, demo.name)

		s := spinner.New(demo.style, demo.message)
		s.Start()

		// Let it run for 3 seconds
		time.Sleep(3 * time.Second)

		s.StopWithMessage(fmt.Sprintf("✓ %s completed!", demo.name))
		fmt.Println()

		// Brief pause between demos
		time.Sleep(500 * time.Millisecond)
	}

	// Demo error handling
	fmt.Println("6. Error Demo:")
	errorSpinner := spinner.New(spinner.DotsStyle, "Simulating an error scenario...")
	errorSpinner.Start()
	time.Sleep(2 * time.Second)
	errorSpinner.StopWithError("✗ Something went wrong!")
	fmt.Println()

	// Demo message updating
	fmt.Println("7. Dynamic Message Demo:")
	dynamicSpinner := spinner.New(spinner.PulseStyle, "Starting process...")
	dynamicSpinner.Start()

	time.Sleep(1 * time.Second)
	dynamicSpinner.SetMessage("Processing step 1 of 3...")
	time.Sleep(1 * time.Second)
	dynamicSpinner.SetMessage("Processing step 2 of 3...")
	time.Sleep(1 * time.Second)
	dynamicSpinner.SetMessage("Processing step 3 of 3...")
	time.Sleep(1 * time.Second)

	dynamicSpinner.StopWithMessage("✓ All steps completed successfully!")

	fmt.Println("\n🎉 Demo completed! The spinners are ready to use in your application.")
}
