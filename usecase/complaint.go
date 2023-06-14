package usecase

import (
	"capstone/constant"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"capstone/utils"
	"errors"
	"time"

	"gorm.io/gorm"
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
			CreatedAt:    v.CreatedAt.Format("02/01/2006"),
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
		resp = append(resp, &payload.GetComplaintByStatusResponse{
			ID:          v.ID,
			Description: v.Description,
			Status:      v.Status,
		})
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
		CreatedAt:   complaint.CreatedAt.Format("02 January 2006"),
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

func AdminDeleteComplaintByID(complaintID uint) error {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("complaint not found")
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
			CreatedAt:   v.CreatedAt.Format("02/01/2006"),
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

func UpdateComplaintByID(req *payload.UpdateComplaintRequest, complaintID uint) (*model.Complaint, error) {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return nil, errors.New("complaint not found")
	}

	complaint.Status = req.Status
	complaint.Type = req.Type

	err = database.UpdateComplaint(complaint)
	if err != nil {
		return nil, err
	}
	return complaint, nil
}

func GetUserComplaintID(ComplaintID, userID uint) (*payload.GetUserComplaintIDResponse, error) {
	complaint, err := database.GetComplaintByID(ComplaintID)
	if err != nil {
		return nil, errors.New("complaint not found")
	}
	user, err := database.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	comments := []payload.GetCommentResponse{}
	for _, comment := range complaint.Comments {
		time := time.Since(comment.CreatedAt.Local())
		comments = append(comments, payload.GetCommentResponse{
			ID:           comment.ID,
			PhotoProfile: utils.ConvertToNullString(comment.User.PhotoProfile),
			FullName:     comment.User.FullName,
			Username:     comment.User.Username,
			Description:  comment.Description,
			CreatedAt:    utils.GetTimeAgo(int64(time.Seconds())),
		})
	}

	resp := &payload.GetUserComplaintIDResponse{
		ID:           complaint.ID,
		PhotoProfile: utils.ConvertToNullString(complaint.User.PhotoProfile),
		FullName:     complaint.User.FullName,
		Username:     complaint.User.Username,
		Description:  complaint.Description,
		PhotoURL:     utils.ConvertToNullString(complaint.PhotoURL),
		VideoURL:     utils.ConvertToNullString(complaint.VideoURL),
		IsPublic:     complaint.IsPublic,
		Feedback:     utils.ConvertToNullString(complaint.Feedback.Description),
		CreatedAt:    complaint.CreatedAt.Format("02/01/2006"),
		Comments:     comments,
		UserProfile:  utils.ConvertToNullString(user.PhotoProfile),
	}
	return resp, nil
}

func PinnedComplaint(userID, complaintID uint) error {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("complaint not found")
	}
	pinned, err := database.GetPinnedByComplaintIdAndUserId(userID, complaint.ID)
	if err != nil && err == errors.New("record not found") {
		return err
	}
	if pinned != nil {
		return errors.New("you have pinned this complaint")
	}

	pin := &model.PinnedComplaint{
		UserID:      userID,
		ComplaintID: complaint.ID,
	}

	err = database.AddPinnedComplaint(pin)
	if err != nil {
		return err
	}
	return nil
}

func UnpinnedComplaint(userID, complaintID uint) error {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return errors.New("complaint not found")
	}
	pinned, err := database.GetPinnedByComplaintIdAndUserId(userID, complaint.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if pinned == nil {
		return errors.New("you have not pinned this complaint")
	}

	err = database.DeletePinnedComplaint(pinned)
	if err != nil {
		return err
	}
	return nil
}
