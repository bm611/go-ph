package spinner

import (
	"fmt"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// SpinnerStyle defines different types of spinners
type SpinnerStyle int

const (
	DotsStyle SpinnerStyle = iota
	LineStyle
	ArrowStyle
	PulseStyle
	BouncingBar
)

// Spinner represents an animated loading spinner
type Spinner struct {
	style     SpinnerStyle
	message   string
	frames    []string
	interval  time.Duration
	isRunning bool
	mu        sync.RWMutex
	stopCh    chan struct{}
	done      chan struct{}
}

// New creates a new spinner with the specified style and message
func New(style SpinnerStyle, message string) *Spinner {
	s := &Spinner{
		style:    style,
		message:  message,
		interval: 100 * time.Millisecond,
		stopCh:   make(chan struct{}),
		done:     make(chan struct{}),
	}

	s.setFrames()
	return s
}

// setFrames sets the animation frames based on the spinner style
func (s *Spinner) setFrames() {
	switch s.style {
	case DotsStyle:
		s.frames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
		s.interval = 80 * time.Millisecond
	case LineStyle:
		s.frames = []string{"|", "/", "-", "\\"}
		s.interval = 100 * time.Millisecond
	case ArrowStyle:
		s.frames = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
		s.interval = 150 * time.Millisecond
	case PulseStyle:
		s.frames = []string{"◐", "◓", "◑", "◒"}
		s.interval = 200 * time.Millisecond
	case BouncingBar:
		s.frames = []string{
			"[    ]",
			"[=   ]",
			"[==  ]",
			"[=== ]",
			"[ ===]",
			"[  ==]",
			"[   =]",
			"[    ]",
			"[   =]",
			"[  ==]",
			"[ ===]",
			"[=== ]",
			"[==  ]",
			"[=   ]",
		}
		s.interval = 100 * time.Millisecond
	default:
		s.frames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
		s.interval = 80 * time.Millisecond
	}
}

// SetMessage updates the spinner message
func (s *Spinner) SetMessage(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.message = message
}

// SetInterval updates the animation interval
func (s *Spinner) SetInterval(interval time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.interval = interval
}

// Start begins the spinner animation
func (s *Spinner) Start() {
	s.mu.Lock()
	if s.isRunning {
		s.mu.Unlock()
		return
	}
	s.isRunning = true
	s.mu.Unlock()

	go s.animate()
}

// Stop ends the spinner animation
func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.isRunning {
		s.mu.Unlock()
		return
	}
	s.isRunning = false
	s.mu.Unlock()

	close(s.stopCh)
	<-s.done

	// Clear the line
	fmt.Print("\r\033[K")
}

// StopWithMessage stops the spinner and displays a final message
func (s *Spinner) StopWithMessage(message string) {
	s.Stop()

	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true)

	fmt.Println(successStyle.Render(message))
}

// StopWithError stops the spinner and displays an error message
func (s *Spinner) StopWithError(message string) {
	s.Stop()

	errorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true)

	fmt.Println(errorStyle.Render(message))
}

// animate runs the spinner animation loop
func (s *Spinner) animate() {
	defer close(s.done)

	spinnerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("105")).
		Bold(true)

	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("99"))

	frameIndex := 0
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.mu.RLock()
			frame := s.frames[frameIndex]
			message := s.message
			interval := s.interval
			s.mu.RUnlock()

			// Update ticker interval if it changed
			if ticker.C != nil && interval != s.interval {
				ticker.Stop()
				ticker = time.NewTicker(interval)
			}

			// Clear line and print spinner
			spinnerText := spinnerStyle.Render(frame)
			messageText := messageStyle.Render(" " + message)
			fmt.Printf("\r%s%s", spinnerText, messageText)

			frameIndex = (frameIndex + 1) % len(s.frames)
		}
	}
}

// IsRunning returns true if the spinner is currently active
func (s *Spinner) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isRunning
}
