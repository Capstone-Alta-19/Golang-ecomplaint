package usecase

import (
	"capstone/constant"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"sort"
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

func GetComplaintsByCategoryID(categoryID uint, sortOrder string) ([]*model.Complaint, error) {
	complaints, err := database.GetComplaintsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	publicComplaint := []*model.Complaint{}
	for _, v := range complaints {
		if v.IsPublic == true {
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

	switch sortOrder {
	case constant.Ascending:
		sort.Slice(publicComplaint, func(i, j int) bool {
			return publicComplaint[i].CreatedAt.Before(publicComplaint[j].CreatedAt)
		})
	case constant.Descending:
		sort.Slice(publicComplaint, func(i, j int) bool {
			return publicComplaint[i].CreatedAt.After(publicComplaint[j].CreatedAt)
		})
	default:
		sort.Slice(publicComplaint, func(i, j int) bool {
			return publicComplaint[i].CreatedAt.After(publicComplaint[j].CreatedAt)
		})
	}

	return publicComplaint, nil
}

func GetUserComplaintsByStatus(userID uint, status string) ([]*payload.GetComplaintByStatusResponse, error) {
	complaints, err := database.GetComplaintsByUserID(userID)
	if err != nil {
		return nil, err
	}
	sort.Slice(complaints, func(i, j int) bool {
		return complaints[i].CreatedAt.After(complaints[j].CreatedAt)
	})

	if status == constant.StatusAll {
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

	userComplaints := []*payload.GetComplaintByStatusResponse{}
	for _, v := range complaints {
		if v.Status == status {
			userComplaints = append(userComplaints, &payload.GetComplaintByStatusResponse{
				ID:          v.ID,
				Description: v.Description,
				Status:      v.Status,
			})
		}
	}

	return userComplaints, nil
}

func GetComplaintByID(id uint) (*model.Complaint, error) {
	complaint, err := database.GetComplaintByID(id)
	if err != nil {
		return nil, err
	}
	return complaint, nil
}

func DeleteComplaintByID(userID, complaintID uint) error {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return err
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
