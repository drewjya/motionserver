package schema

import "gorm.io/gorm"

const URL = "https://www.googleapis.com/youtube/v3/search?channelId=UCjmwuAexNpXPzUglE1F11WQ&maxResults=1&key=AIzaSyBYnhuaSpaPXeakbujqOqNIAuKsSNB49Vo&order=date&part=snippet"

type Youtube struct {
	gorm.Model
	VideoId     string
	Thumbnail   string
	PublishedAt string
}
