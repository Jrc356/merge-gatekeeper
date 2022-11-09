package status

import (
	"fmt"
	"strings"
)

type status struct {
	totalJobs    []string
	completeJobs []string
	errJobs      []string
	ignoredJobs  []string
	succeeded    bool
}

const (
	jobListSeparatorStr = "\n      -"
)

func formatListToStr(arr []string) string {
	if len(arr) == 0 {
		return "[]"
	}

	str := ""
	for _, a := range arr {
		str = fmt.Sprintf("%s%s %s", str, jobListSeparatorStr, strings.TrimSpace(a))
	}

	return str
}

func (s *status) Detail() string {
	incompleteJobs := s.incompleteJobs()
	result := fmt.Sprintf(
		`%d out of %d

  Total job count: %d
    jobs: %s
  Completed job count: %d
    jobs: %s
  Incomplete job count: %d
    jobs: %s
  Failed job count: %d
    jobs: %s`,
		len(s.completeJobs), len(s.totalJobs),
		len(s.totalJobs), formatListToStr(s.totalJobs),
		len(s.completeJobs), formatListToStr(s.completeJobs),
		len(incompleteJobs), formatListToStr(incompleteJobs),
		len(s.errJobs), formatListToStr(s.errJobs),
	)

	if len(s.ignoredJobs) > 0 {
		result = fmt.Sprintf(
			`%s

  --

  Ignored jobs: %s`, result, formatListToStr(s.ignoredJobs))
	}

	return result
}

func (s *status) IsSuccess() bool {
	// TDOO: Add test case
	return s.succeeded
}

func (s *status) incompleteJobs() []string {
	var incomplete []string
	for _, job := range s.totalJobs {
		isIgnored := false
		for _, ignored := range s.ignoredJobs {
			if job == ignored {
				isIgnored = true
				break
			}
		}
		if isIgnored {
			continue
		}

		isCompleted := false
		for _, completed := range s.completeJobs {
			if job == completed {
				isCompleted = true
				break
			}
		}

		for _, errd := range s.errJobs {
			if job == errd {
				isCompleted = true
				break
			}
		}

		if !isCompleted {
			incomplete = append(incomplete, job)
		}
	}
	return incomplete
}
