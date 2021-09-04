from django.db import models

#策略表
class Online(models.Model):
    name      = models.CharField("名称", max_length=200)
    server    = models.CharField("探针地址", max_length=200)
    token     = models.TextField("探针token", max_length=200, blank=True, null=True)
    message   = models.TextField("备注", max_length=200, blank=True, null=True)
    cpu       = models.CharField("cpu", max_length=200,default="100%", editable=False)
    memery    = models.CharField("内存", max_length=200,default="100%", editable=False)
    disk      = models.CharField("总计硬盘空间", max_length=200,default="100%", editable=False)
    disk2     = models.CharField("剩余硬盘空间", max_length=200,default="100%", editable=False)
    status    = models.IntegerField("状态", default=1, choices=((1, "未连接"),(2, "正常"),(3, "异常")), editable=False)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True)  

    class Meta:
        verbose_name = '探针配置'
        verbose_name_plural = '探针配置'
    
    def __str__(self):
        return self.name