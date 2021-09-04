import json
import psutil
from datetime import datetime as dt
from django.shortcuts import render
from django.contrib import admin
from django.conf import settings
from django import forms
from django.apps import apps
from .models import Natural
from django.urls import path
from pyecharts.charts import Bar, Gauge
from pyecharts import options as opts
from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
from flyadmin.views.charts import bar, pie, line

class NaturalAdmin(admin.ModelAdmin):
    model = Natural

    def admin_view_natural(self, request):
        _f1 = lambda x:round(x/1024/1024/1024, 1)
        #将数据写入到里面来到首页数据里(不得修改)
        mem     = psutil.virtual_memory()
        #cpu     = psutil.cpu_percent()
        disk    = psutil.disk_usage('/')
        network = psutil.net_io_counters()
        naturals = Natural.objects.order_by('create_time').all()
        c = line('cpu趋势(G)', [i.create_time.strftime('%Y-%m-%d %H') for i in naturals], {'占用比率':[i.ava for i in naturals]})
        p = pie("内存占比(G)", ('全部','可用'), (_f1(mem.total), _f1(mem.available)))
        d = pie("硬盘占比(G)", ('全部','使用', '可用'), (_f1(disk.total), _f1(disk.used), _f1(disk.free)))
        n = pie("网络情况(G)", ('接收','发送'), (_f1(network.bytes_recv), _f1(network.bytes_sent)))
        html = ['<el-row><div class="grid-content">']
        html.append('<div class="el-col el-col-8">{}</div>'.format(p))
        html.append('<div class="el-col el-col-8">{}</div>'.format(d))
        html.append('<div class="el-col el-col-8">{}</div>'.format(n))
        html.append('</div></el-row>')
        html.append('<el-row><div class="grid-content"><div class="el-col el-col-24">{}</div></div></el-row>'.format(c))
        xplots = ''.join(html)   
        return render(request, "admin/xplots.html", locals()) 
 
    def changelist_view(self, request, extra_content=None):
        return self.admin_view_natural(request)



admin.site.register(Natural, NaturalAdmin)

 
