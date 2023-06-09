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

	resp := &model.Complaint{
		UserID:      UserID,
		Description: req.Description,
		Type:        req.Type,
		PhotoURL:    req.PhotoURL,
		VideoURL:    req.VideoURL,
		CategoryID:  req.CategoryID,
		IsPublic:    req.IsPublic,
		Status:      constant.StatusPending,
	}
	err := database.CreateComplaint(resp)
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
