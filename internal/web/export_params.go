package web

import "strings"

func normalizeLogsExportWebsiteID(params map[string]string) {
	if len(params) == 0 {
		return
	}
	websiteID := strings.TrimSpace(params["website_id"])
	if websiteID == "" {
		websiteID = strings.TrimSpace(params["websiteId"])
	}
	if websiteID == "" {
		websiteID = strings.TrimSpace(params["id"])
	}
	if websiteID == "" {
		return
	}
	params["id"] = websiteID
}
