package response

type BaseInitResponse struct {
    ActivityPeriodicList []string `json:"activityPeriodicList"`
    ActivityRemindAtList []string `json:"activityRemindAtList"`
    ActivityPrivacyList  []string `json:"activityPrivacyList"`
}
