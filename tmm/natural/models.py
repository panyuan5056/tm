import json
import requests
from django.db import models

#策略表
class Natural(models.Model):
    name      = models.CharField("名称", max_length=200) 
    total     = models.FloatField("全部")
    ava       = models.FloatField("剩余")
    create_time   = models.DateTimeField(auto_now=True)  
    update_time   = models.DateTimeField(auto_now_add=True) 

    def __str__(self):
        return self.name
    
    class Meta:
        verbose_name = '系统资源'
        verbose_name_plural = '系统资源'