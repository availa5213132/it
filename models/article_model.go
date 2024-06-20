package models

import "gvb/gvb_server/models/ctype"

type ArticleModel struct {
	MODEL
	//ID            string         `json:"id"`                                      // es的id
	//CreatedAt     string         `json:"created_at"`                              // 创建时间
	//UpdatedAt     string         `json:"updated_at"`                              // 更新时间
	Title string `gorm:"size:32" json:"title"` // 文章标题
	//Keyword       string         `json:"keyword,omit(list)"`                      // 关键字
	Abstract      string         `json:"abstract"`                                       // 文章简介
	Content       string         `json:"content"`                                        // 文章内容
	LookCount     int            `json:"look_count"`                                     // 浏览量
	CommentCount  int            `json:"comment_count"`                                  // 评论量
	DiggCount     int            `json:"digg_count"`                                     // 点赞量
	CollectsCount int            `json:"collects_count"`                                 // 收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"` //文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`                  //文章的评论列表
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`                     //文章作者
	UserID        uint           `json:"user_id"`                                        // 用户id
	//UserNickName  string      `json:"user_nick_name" structs:"user_nick_name"` //用户昵称
	//UserAvatar    string      `json:"user_avatar" structs:"user_avatar"`       // 用户头像
	Category   string      `json:"category"`                     // 文章分类
	Source     string      `json:"source"`                       // 文章来源
	Link       string      `json:"link"`                         // 原文链接
	Banner     BannerModel `gorm:"foreignKey:BannerID" json:"-"` //文章封面
	BannerID   uint        `json:"banner_id"`                    // 文章封面id
	NickName   string      `gorm:"sizxe:42" json:"nick_name"`
	BannerPath string      `json:"banner_path"`
	//BannerUrl     string      `json:"banner_url" structs:"banner_url"`         // 文章封面
	Tags ctype.Array `gorm:"type:string;size:64" json:"tags"` // 文章标签
}
