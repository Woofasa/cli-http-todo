package usecase

import (
	"main/internal/domain"
	"slices"
)

func (a *App) Sort(pattern string, list []*domain.Task) []*domain.Task {
	switch pattern {
	case "created_at":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	case "name":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.Title > b.Title {
				return 1
			} else if b.Title > a.Title {
				return -1
			}
			return 0
		})
	case "completed_at":
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CompletedAt != nil && b.CompletedAt == nil {
				return -1
			}
			if b.CompletedAt != nil && a.CompletedAt == nil {
				return -1
			}
			if a.CompletedAt == nil && b.CompletedAt == nil {
				return 0
			}
			if a.CompletedAt.Before(*b.CompletedAt) {
				return -1
			}
			if b.CompletedAt.Before(*a.CompletedAt) {
				return -1
			}
			return 0
		})
	default:
		slices.SortFunc(list, func(a, b *domain.Task) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			} else if b.CreatedAt.Before(a.CreatedAt) {
				return 1
			}
			return 0
		})
	}
	return list
}

func (a *App) Filter(pattern string, list []*domain.Task) []*domain.Task {
	filtered := make([]*domain.Task, 0, len(list))
	switch pattern {
	case "opened":
		for _, v := range list {
			if v.Status {
				filtered = append(filtered, v)
			}
		}
	case "closed":
		for _, v := range list {
			if !v.Status {
				filtered = append(filtered, v)
			}
		}
	default:
		return list
	}
	return filtered
}
