package im_mysql_model

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/db"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/utils"
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	//init managers
	for k, v := range config.Config.Manager.AppManagerUid {
		user, err := GetUserByUserID(v)
		if err != nil {
			fmt.Println("GetUserByUserID failed ", err.Error(), v, user)
		} else {
			continue
		}
		var appMgr db.User
		appMgr.UserID = v
		appMgr.Nickname = "AppManager" + utils.IntToString(k+1)
		appMgr.AppMangerLevel = constant.AppAdmin
		err = UserRegister(appMgr)
		if err != nil {
			fmt.Println("AppManager insert error", err.Error(), appMgr, "time: ", appMgr.Birth.Unix())
		}

	}
}

func UserRegister(user db.User) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	user.CreateTime = time.Now()
	if user.AppMangerLevel == 0 {
		user.AppMangerLevel = constant.AppOrdinaryUsers
	}
	if user.Birth.Unix() < 0 {
		user.Birth = utils.UnixSecondToTime(0)
	}
	err = dbConn.Table("users").Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(userID string) (i int64) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return 0
	}
	i = dbConn.Table("users").Where("user_id=?", userID).Delete(db.User{}).RowsAffected
	return i
}

func GetUserByUserID(userID string) (*db.User, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var user db.User
	err = dbConn.Table("users").Where("user_id=?", userID).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserNameByUserID(userID string) (string, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return "", err
	}
	var user db.User
	err = dbConn.Table("users").Select("name").Where("user_id=?", userID).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Nickname, nil
}

func UpdateUserInfo(user db.User) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	dbConn.LogMode(true)
	err = dbConn.Table("users").Where("user_id=?", user.UserID).Update(&user).Error
	return err
}

func SelectAllUserID() ([]string, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var resultArr []string
	err = dbConn.Table("users").Pluck("user_id", &resultArr).Error
	if err != nil {
		return nil, err
	}
	return resultArr, nil
}

func SelectSomeUserID(userIDList []string) ([]string, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	dbConn.LogMode(true)
	if err != nil {
		return nil, err
	}
	var resultArr []string
	err = dbConn.Table("users").Where("user_id IN (?) ", userIDList).Pluck("user_id", &resultArr).Error

	if err != nil {
		return nil, err
	}
	return resultArr, nil
}

func GetUsers(showNumber, pageNumber int32) ([]db.User, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var users []db.User
	if err != nil {
		return users, err
	}
	dbConn.LogMode(true)
	err = dbConn.Table("users").Limit(showNumber).Offset(showNumber * (pageNumber - 1)).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}

func AddUser(userId, phoneNumber, name string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	user := db.User{
		PhoneNumber: phoneNumber,
		Birth:       time.Now(),
		CreateTime:  time.Now(),
		UserID:      userId,
		Nickname:    name,
	}
	result := dbConn.Table("users").Create(&user)
	return result.Error
}

func UserIsBlock(userId string) (bool, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return false, err
	}
	var user db.BlackList
	rows := dbConn.Table("black_lists").Where("uid=?", userId).First(&user).RowsAffected
	if rows >= 1 {
		return true, nil
	}
	return false, nil
}

func BlockUser(userId, endDisableTime string) error {
	user, err := GetUserByUserID(userId)
	if err != nil || user.UserID == "" {
		return err
	}
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	end, err := time.Parse("2006-01-02 15:04:05", endDisableTime)
	if err != nil {
		return err
	}
	if end.Before(time.Now()) {
		return constant.ErrDB
	}
	var blockUser db.BlackList
	dbConn.Table("black_lists").Where("uid=?", userId).First(&blockUser)
	if blockUser.UserId != "" {
		dbConn.Model(&blockUser).Where("uid=?", blockUser.UserId).Update("end_disable_time", end)
		return nil
	}
	blockUser = db.BlackList{
		UserId:           userId,
		BeginDisableTime: time.Now(),
		EndDisableTime:   end,
	}
	result := dbConn.Create(&blockUser)
	return result.Error
}

func UnBlockUser(userId string) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	dbConn.LogMode(true)
	result := dbConn.Where("uid=?", userId).Delete(&db.BlackList{})
	return result.Error
}

type BlockUserInfo struct {
	User             db.User
	BeginDisableTime time.Time
	EndDisableTime   time.Time
}

func GetBlockUserById(userId string) (BlockUserInfo, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var blockUserInfo BlockUserInfo
	blockUser := db.BlackList{
		UserId: userId,
	}
	if err != nil {
		return blockUserInfo, err
	}
	if err = dbConn.Table("black_lists").Where("uid=?", userId).Find(&blockUser).Error; err != nil {
		return blockUserInfo, err
	}
	user := db.User{
		UserID: blockUser.UserId,
	}
	if err := dbConn.Find(&user).Error; err != nil {
		return blockUserInfo, err
	}
	blockUserInfo.User.UserID = user.UserID
	blockUserInfo.User.FaceURL = user.UserID
	blockUserInfo.User.Nickname = user.Nickname
	blockUserInfo.BeginDisableTime = blockUser.BeginDisableTime
	blockUserInfo.EndDisableTime = blockUser.EndDisableTime
	return blockUserInfo, nil
}

func GetBlockUsers(showNumber, pageNumber int32) ([]BlockUserInfo, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var blockUserInfos []BlockUserInfo
	var blockUsers []db.BlackList
	if err != nil {
		return blockUserInfos, err
	}
	dbConn.LogMode(true)
	if err = dbConn.Limit(showNumber).Offset(showNumber * (pageNumber - 1)).Find(&blockUsers).Error; err != nil {
		return blockUserInfos, err
	}
	for _, blockUser := range blockUsers {
		var user db.User
		if err := dbConn.Table("users").Where("user_id=?", blockUser.UserId).First(&user).Error; err == nil {
			blockUserInfos = append(blockUserInfos, BlockUserInfo{
				User: db.User{
					UserID:   user.UserID,
					Nickname: user.Nickname,
					FaceURL:  user.FaceURL,
				},
				BeginDisableTime: blockUser.BeginDisableTime,
				EndDisableTime:   blockUser.EndDisableTime,
			})
		}
	}
	return blockUserInfos, nil
}

func GetUserByName(userName string, showNumber, pageNumber int32) ([]db.User, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var users []db.User
	if err != nil {
		return users, err
	}
	dbConn.LogMode(true)
	err = dbConn.Table("users").Where(fmt.Sprintf(" name like '%%%s%%' ", userName)).Limit(showNumber).Offset(showNumber * (pageNumber - 1)).Find(&users).Error
	return users, err
}

func GetUsersCount(user db.User) (int32, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return 0, err
	}
	dbConn.LogMode(true)
	var count int32
	if err := dbConn.Table("users").Where(fmt.Sprintf(" name like '%%%s%%' ", user.Nickname)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetBlockUsersNumCount() (int32, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return 0, err
	}
	dbConn.LogMode(true)
	var count int32
	if err := dbConn.Model(&db.BlackList{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func SetConversation(conversation db.Conversation) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	dbConn.LogMode(true)
	newConversation := conversation
	if dbConn.Model(&db.Conversation{}).Find(&newConversation).RowsAffected == 0 {
		log.NewDebug("", utils.GetSelfFuncName(), "conversation", conversation, "not exist in db, create")
		return dbConn.Model(&db.Conversation{}).Create(conversation).Error
		// if exist, then update record
	} else {
		log.NewDebug("", utils.GetSelfFuncName(), "conversation", conversation, "exist in db, update")
		//force update
		return dbConn.Model(conversation).Where("owner_user_id = ? and conversation_id = ?", conversation.OwnerUserID, conversation.ConversationID).
			Update(map[string]interface{}{"recv_msg_opt": conversation.RecvMsgOpt, "is_pinned": conversation.IsPinned, "is_private_chat": conversation.IsPrivateChat}).Error
	}
}

func SetRecvMsgOpt(conversation db.Conversation) error {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return err
	}
	dbConn.LogMode(true)
	newConversation := conversation
	if dbConn.Model(&db.Conversation{}).Find(&newConversation).RowsAffected == 0 {
		log.NewDebug("", utils.GetSelfFuncName(), "conversation", conversation, "not exist in db, create")
		return dbConn.Model(&db.Conversation{}).Create(conversation).Error
		// if exist, then update record
	} else {
		log.NewDebug("", utils.GetSelfFuncName(), "conversation", conversation, "exist in db, update")
		//force update
		return dbConn.Model(conversation).Where("owner_user_id = ? and conversation_id = ?", conversation.OwnerUserID, conversation.ConversationID).
			Update(map[string]interface{}{"recv_msg_opt": conversation.RecvMsgOpt}).Error
	}
}

func GetUserAllConversations(ownerUserID string) ([]db.Conversation, error) {
	var conversations []db.Conversation
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return conversations, err
	}
	dbConn.LogMode(true)
	err = dbConn.Model(&db.Conversation{}).Where("owner_user_id=?", ownerUserID).Find(&conversations).Error
	return conversations, err
}

func GetConversation(OwnerUserID, conversationID string) (db.Conversation, error) {
	var conversation db.Conversation
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return conversation, err
	}
	err = dbConn.Model(&db.Conversation{
		OwnerUserID:    OwnerUserID,
		ConversationID: conversationID,
	}).Find(&conversation).Error
	return conversation, err
}

func GetConversations(OwnerUserID string, conversationIDs []string) ([]db.Conversation, error) {
	var conversations []db.Conversation
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	if err != nil {
		return conversations, err
	}
	err = dbConn.Model(&db.Conversation{}).Where("conversation_id IN (?) and  owner_user_id=?", conversationIDs, OwnerUserID).Find(&conversations).Error
	return conversations, err
}
