/*
@Time : 2022/4/9 22:21
@Author : Yuxue.fan<fanyixuanf@gmail.com>
@File : cron_job.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

func AddCronJob(job model.CronJob) (err error, jobInter model.CronJob) {
	var cjob model.CronJob
	if !errors.Is(global.GVA_DB.Where("job_key = ?", job.JobKey).First(&cjob).Error, gorm.ErrRecordNotFound) { // 判断key是否存在
		return errors.New("job key已经存在"), jobInter
	}
	err = global.GVA_DB.Create(&job).Error
	return err, job
}

func GetCronJobList(p request.CronJobListReq) (err error, list interface{}, total int64) {
	limit := p.PageSize
	offset := p.PageSize * (p.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CronJob{})
	var data []model.CronJob
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&data).Error
	return nil, data, total
}

func ChangeCronJobStatus(id request.ChangeStatusCronJob, s, t int) (err error, data model.CronJob) {
	var job model.CronJob
	db := global.GVA_DB.Model(&model.CronJob{})
	if t == 0 {
		err = db.Where("id = ?", id.Id).First(&job).Updates(map[string]interface{}{
			"status": s,
		}).Error
	} else {
		err = db.Where("id = ?", id.Id).First(&job).Updates(map[string]interface{}{
			"timer_id": t,
		}).Error
	}
	return err, job
}