package usecase

import (
	"capstone/model/payload"
	"capstone/repository/database"
	"capstone/utils"
)

func GetNotification() ([]*payload.GetNotificationResponse, error) {
	notifications, err := database.GetNotification()
	if err != nil {
		return nil, err
	}
	resp := []*payload.GetNotificationResponse{}
	for _, v := range notifications {
		resp = append(resp, &payload.GetNotificationResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			PhotoProfile: utils.ConvertToNullString(v.User.PhotoProfile),
			Category:     v.Complaint.Category.Name,
			Description:  v.Complaint.Description,
			ComplaintID:  v.ComplaintID,
			CreatedAt:    v.CreatedAt.Format("02 January 2006"),
		})
	}
	return resp, nil
}
