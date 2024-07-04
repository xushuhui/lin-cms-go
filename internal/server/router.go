package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/selector"
)

func ListRoutes() map[string]struct{} {
	return map[string]struct{}{
		"/api.Suggest/CreateSuggest":     {},
		"/api.Suggest/UpdateSuggestRead": {},
		"/api.Suggest/Update":            {},
		"/api.Suggest/CreateDraft":       {},
		"/api.Suggest/ListDraft":         {},
		"/api.Suggest/GetDraft":          {},
		"/api.Suggest/ListChannel":       {},
		"/api.Suggest/UpdateRead":        {},
		"/api.Suggest/CreateLike":        {},
		"/api.Suggest/CreateFollow":      {},
		"/api.Suggest/CreateComment":     {},
		"/api.Suggest/CreateCommentLike": {},
		"/api.Suggest/DeleteComment":     {},
		"/api.Suggest/ListFollow":        {},
		"/api.Suggest/ListNotifyLike":    {},
		"/api.Suggest/ListNotifyComment": {},
		"/api.Suggest/ListMyCheck":       {},
		"/api.Suggest/ListMyRelease":     {},
		"/api.Suggest/GetMySuggest":      {},
	}
}

func MustAuthRoutersMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := ListRoutes()[operation]; ok {
			return true
		}
		return false
	}
}
