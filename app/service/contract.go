package service

import (
	"errors"
	"fastIM/app/model"
	"time"
)

type ContactService struct {}

//添加好友
func (c *ContactService) AddFriend(userid int64, dstid int64) error {
	if dstid == userid {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经添加了好友
	friend := model.Contact{}
	model.DbEngine.Where("ownerid = ?", userid).
		And("dstobj = ?", dstid).
		And("cate = ?", model.ConcatCateUser).
		Get(&friend)
	//如果好友已经存在，则不添加
	if friend.Id > 0 {
		return errors.New("该好友已经添加过了")
	}
	//开启事物
	session := model.DbEngine.NewSession()
	session.Begin()
	//插入两条好友关系数据
	_, s1 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})
	_, s2 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})
	if s1 == nil && s2 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if s1 != nil {
			return s1
		}
		return s2
	}
}

//搜索社群
func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	model.DbEngine.Where("ownerid = ? and cate = ?", userId, model.ConcatCateComunity).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	model.DbEngine.In("id", comIds).Find(&coms)
	return coms
}

//根据名字搜索社群
func (c *ContactService) SearchCommunityByName(cname string) model.Community {
	com := model.Community{}
	model.DbEngine.Where("name = ?", cname).Get(&com)
	return com
}

//搜索好友
func (c *ContactService) SearchFriend(userId int64) []model.User {
	friends := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	model.DbEngine.Where("ownerid = ? and cate = ?", userId, model.ConcatCateUser).Find(&friends)
	for _, v := range friends {
		objIds = append(objIds, v.Dstobj)
	}
	users := make([]model.User, 0)
	if len(objIds) == 0 {
		return users
	}
	model.DbEngine.In("id", objIds).Find(&users)
	return users
}

//根据姓名搜索用户
func (c *ContactService) SearchFriendByName(mobile string) model.User {
	user := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&user)
	return user
}

func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	// 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	model.DbEngine.Where("ownerid = ? and cate = ?", userId, model.ConcatCateComunity).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}

//添加群
func (c *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := model.DbEngine.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := model.DbEngine.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.ConcatCateComunity,
				Createat: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return com, err
	}
}

//用户加群
func (c *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.ConcatCateComunity,
	}
	model.DbEngine.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := model.DbEngine.InsertOne(cot)
		return err
	} else {
		return nil
	}
}
