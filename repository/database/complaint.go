package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateComplaint(complaint *model.Complaint) error {
	err := config.DB.Create(complaint).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComplaintsByUserID(userID uint) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	err := config.DB.Where("user_id = ?", userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaintByID(id uint) (*model.Complaint, error) {
	var complaint model.Complaint
	err := config.DB.First(&complaint, id).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func GetFeedbacksByComplaintID(complaintID string) ([]*model.Feedback, error) {
	feedbacks := []*model.Feedback{}
	err := config.DB.Where("complaint_id = ?", complaintID).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func GetComplaintsByCategoryAndSort(categoryID uint, sortParam string) ([]*model.Complaint, error) {
	var complaints []*model.Complaint
	query := config.DB.Where("category_id = ?", categoryID)

	switch sortParam {
	case "asc":
		query = query.Order("created_at ASC")
	case "desc":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	err := query.Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetLikes(sortBy string, query string) ([]*model.Like, error) {
	// Buat query sesuai dengan parameter yang diberikan
	queryStr := "SELECT * FROM likes"
	if query != "" {
		queryStr += " WHERE name LIKE '%" + query + "%'"
	}
	if sortBy != "" {
		queryStr += " ORDER BY " + sortBy
	}

	// Eksekusi query ke basis data
	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []*model.Like{}
	for rows.Next() {
		like := &model.Like{}
		err := rows.Scan(&like.ID, &like.UserID, &like.ComplaintID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}
