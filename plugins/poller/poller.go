package poller

import "time"

// Poller ...
type Poller struct {
	Job    func()
	done   chan struct{}
	ticker *time.Ticker
}

// SyncJob ...
type SyncJob func()

// NewPoller ...
func NewPoller(job func(), duration time.Duration) *Poller {
	poller := &Poller{
		Job:    job,
		done:   make(chan struct{}),
		ticker: time.NewTicker(duration),
	}

	return poller
}

// Stop ...
func (p *Poller) Stop() {
	close(p.done)
}

// Run ...
func (p *Poller) Run() {
	for {
		select {
		case <-p.ticker.C:
			p.Job()
		case <-p.done:
			p.ticker.Stop()
			return
		}
	}
}
