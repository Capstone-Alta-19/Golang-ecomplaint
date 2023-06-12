package usecase

import (
	"capstone/constant"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"capstone/utils"
	"errors"
)

func CreateComplaint(UserID uint, req *payload.CreateComplaintRequest) (*model.Complaint, error) {
	category, err := database.GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	resp := &model.Complaint{
		UserID:      UserID,
		Description: req.Description,
		Type:        req.Type,
		PhotoURL:    req.PhotoURL,
		VideoURL:    req.VideoURL,
		CategoryID:  category.ID,
		IsPublic:    req.IsPublic,
		Status:      constant.StatusPending,
	}
	err = database.CreateComplaint(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetComplaintsByCategoryID(categoryID uint, sortOrder string) ([]*payload.GetComplaintByCategoryIDResponse, error) {
	complaints, err := database.GetComplaintsByCategoryID(categoryID, sortOrder)
	if err != nil {
		return nil, err
	}
	publicComplaint := []*model.Complaint{}
	for _, v := range complaints {
		if v.IsPublic {
			publicComplaint = append(publicComplaint, v)
		}
	}

	for _, complaint := range publicComplaint {
		likesCount, err := database.GetLikesCountByComplaintID(complaint.ID)
		if err != nil {
			return nil, err
		}
		complaint.LikesCount = likesCount
	}

	resp := []*payload.GetComplaintByCategoryIDResponse{}
	for _, v := range publicComplaint {
		resp = append(resp, &payload.GetComplaintByCategoryIDResponse{
			ID:           v.ID,
			PhotoProfile: utils.ConvertToNullString(v.User.PhotoProfile),
			FullName:     v.User.FullName,
			Username:     v.User.Username,
			Category:     v.Category.Name,
			Description:  v.Description,
			PhotoURL:     utils.ConvertToNullString(v.PhotoURL),
			VideoURL:     utils.ConvertToNullString(v.VideoURL),
			IsPublic:     v.IsPublic,
			Feedback:     utils.ConvertToNullString(v.Feedback.Description),
			LikesCount:   v.LikesCount,
			CreatedAt:    v.CreatedAt,
		})
	}
	return resp, nil
}

func GetUserComplaintsByStatus(userID uint, status string) ([]*payload.GetComplaintByStatusResponse, error) {
	complaints, err := database.GetComplaintsByUserID(userID, status)
	if err != nil {
		return nil, err
	}

	resp := []*payload.GetComplaintByStatusResponse{}
	for _, v := range complaints {
		if v.Status == status {
			resp = append(resp, &payload.GetComplaintByStatusResponse{
				ID:          v.ID,
				Description: v.Description,
				Status:      v.Status,
			})
		}
	}

	return resp, nil
}

func GetComplaintByID(id uint) (*payload.GetComplaintByIDResponse, error) {
	complaint, err := database.GetComplaintByID(id)
	if err != nil {
		return nil, errors.New("complaint not found")
	}
	resp := payload.GetComplaintByIDResponse{
		ID:          complaint.ID,
		FullName:    complaint.User.FullName,
		Type:        complaint.Type,
		Category:    complaint.Category.Name,
		Description: complaint.Description,
		PhotoURL:    utils.ConvertToNullString(complaint.PhotoURL),
		VideoURL:    utils.ConvertToNullString(complaint.VideoURL),
		IsPublic:    complaint.IsPublic,
		CreatedAt:   complaint.CreatedAt,
	}
	return &resp, nil
}

func DeleteComplaintByID(userID, complaintID uint) error {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("complaint not found")
	}
	if complaint.UserID != userID {
		return errors.New("you are not the owner of this complaint")
	}

	err = database.DeleteComplaint(complaint)
	if err != nil {
		return err
	}
	return nil
}

func GetTotalComplaints() (*payload.GetTotalComplaintsResponse, error) {
	complaint, err := database.GetTotalComplaints()
	if err != nil {
		return nil, err
	}
	Aspiration, err := database.GetTotalAspirations()
	if err != nil {
		return nil, err
	}
	total := complaint + Aspiration
	resp := &payload.GetTotalComplaintsResponse{
		Total:      total,
		Complaint:  complaint,
		Aspiration: Aspiration,
	}
	return resp, nil
}

func GetAllComplaints(sortBy, typeSort, search string, limit, page int) ([]*payload.GetAllComplaintsResponse, error) {
	offset := utils.GetOffset(limit, page)
	complaints, err := database.GetAllComplaints(sortBy, typeSort, search, limit, offset)
	if err != nil {
		return nil, err
	}

	resp := []*payload.GetAllComplaintsResponse{}
	for _, v := range complaints {
		resp = append(resp, &payload.GetAllComplaintsResponse{
			ID:          v.ID,
			Name:        v.User.FullName,
			Type:        v.Type,
			Category:    v.Category.Name,
			Description: v.Description,
			Status:      v.Status,
			IsPublic:    v.IsPublic,
			CreatedAt:   v.CreatedAt,
		})
	}
	return resp, nil
}

func CreateFeedbackByComplaintID(req *payload.CreateFeedbackRequest, complaintID uint) (*model.Feedback, error) {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return nil, errors.New("complaint not found")
	}
	resp := &model.Feedback{
		ComplaintID: complaint.ID,
		Description: req.Description,
	}

	err = database.CreateFeedback(resp)
	if err != nil {
		return nil, err
	}

	complaint.Status = constant.StatusResolved
	err = database.UpdateComplaint(complaint)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
