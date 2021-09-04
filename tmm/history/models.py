from django.db import models
from flyadmin.views.charts import bar, pie, line

class History(models.Model):
    title           = models.CharField('标题', max_length=20)#服务
    file            = models.CharField('文件', max_length=200)
    field           = models.CharField('字段', max_length=1000) #请求url
    category      = models.CharField('脱敏类型', max_length=1000) #request response
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    def __str__(self):
        return self.device

    class Meta:
        verbose_name = '脱敏记录'
        verbose_name_plural = '脱敏记录'

    

