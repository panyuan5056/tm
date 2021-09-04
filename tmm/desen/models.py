from django.db import models
from flyadmin.views.charts import bar, pie, line

class Dense(models.Model):
    title         = models.CharField('标题', max_length=1000) #request response
    desc          = models.TextField('描述')
    file          = models.FileField(upload_to='upload')
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    def __str__(self):
        return self.title

    class Meta:
        verbose_name = '脱敏管理'
        verbose_name_plural = '脱敏管理'

