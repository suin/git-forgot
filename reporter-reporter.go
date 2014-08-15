package main

type ReporterReporter struct {
	reporters []Reporter
}

func NewReporterReporter() *ReporterReporter {
	return &ReporterReporter{}
}

func (r *ReporterReporter) Status(path string, status GitStatus) error {
	for _, reporter := range r.reporters {
		err := reporter.Status(path, status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ReporterReporter) Append(reporter Reporter) {
	r.reporters = append(r.reporters, reporter)
}
