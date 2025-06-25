package dao

import "mini_douyin/model"

func CreateFavorite(userID, videoID int64) error {
	fav := model.Favorite{UserID: userID, VideoID: videoID}
	return DB.Where(fav).FirstOrCreate(&fav).Error
}

func DeleteFavorite(userID, videoID int64) error {
	return DB.Where("user_id = ? AND video_id = ?", userID, videoID).Delete(&model.Favorite{}).Error
}

func GetFavoriteVideos(userID int64) ([]model.Video, error) {
	var favs []model.Favorite
	if err := DB.Where("user_id = ?", userID).Find(&favs).Error; err != nil {
		return nil, err
	}
	var videos []model.Video
	var ids []int64
	for _, fav := range favs {
		ids = append(ids, fav.VideoID)
	}
	if len(ids) > 0 {
		if err := DB.Where("id IN ?", ids).Find(&videos).Error; err != nil {
			return nil, err
		}
	}
	return videos, nil
}
